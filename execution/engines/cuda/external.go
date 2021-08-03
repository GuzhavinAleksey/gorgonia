package cuda

import (
	"runtime"

	"github.com/pkg/errors"
	"gorgonia.org/cu"
	cublas "gorgonia.org/cu/blas"
	cudnn "gorgonia.org/cu/dnn"
	"gorgonia.org/gorgonia/internal/allocator"
)

// this file implements all the methods that are required to fulil the External interface

//var _ External = &Engine{}

const (
	// Any address of a variable residing in global memory or returned by one of the
	// memory allocation routines from the driver or runtime API is always aligned to at
	// least 256 bytes.
	//
	memalign    = 32
	scalarAlign = 8

	minAllocBits = 8
	minAllocSize = 1 << minAllocBits
)

// HasFunc returns true if the execution is external (cgo/cuda/openCL) AND the external device contains the function with the given name
func (e *Engine) HasFunc(name string) bool { _, ok := e.f[name]; return ok }

// Sync returns a channel of sync signals
func (e *Engine) Sync() chan struct{} { return e.syncChan }

// Signal signals the machine to do work
func (e *Engine) Signal() {
	e.workAvailable <- true
}

// Context returns the BatchedContext
func (e *Engine) Context() *cu.BatchedContext { return &e.c }

// CUDNNContext returns the cuDNN context
func (e *Engine) CUDNNContext() *cudnn.Context { return &e.n }

// BLASContext returns the cuBLAS context
func (e *Engine) BLASContext() *cublas.Standard { return &e.b }

// Modules returns the loaded modules indexed by name
func (e *Engine) Modules() map[string]cu.Module { return e.m }

// Functions returns the loaded functions indexed by name
func (e *Engine) Functions() map[string]cu.Function { return e.f }

// ElemGridSize calculates the gridsize for elementwise operations. n is the number of elements
func (e *Engine) ElemGridSize(n int) (gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ int) {
	maxThreads := e.mtpb
	maxGridX := e.mgdx
	maxGridY := e.mgdy
	maxGridZ := e.mgdz

	blockDimX = 1
	blockDimY = 1
	blockDimZ = 1
	gridDimX = 1
	gridDimY = 1
	gridDimZ = 1

	blocks := calcBlocks(n, maxThreads)
	switch {
	case blocks == 1:
		blockDimX = n
	case blocks >= maxGridX*maxGridY*maxGridZ:
		// what kind of monstrosity is this??!
	case blocks >= maxGridX*maxGridY:
		gridDimX = maxGridX
		gridDimY = maxGridY
		gridDimZ = calcBlocks(blocks%(maxGridX*maxGridY), maxGridZ)
		blockDimX = maxThreads
	case blocks >= maxGridX:
		gridDimX = maxGridX
		gridDimY = calcBlocks(blocks%(maxGridX), maxGridY)
		blockDimX = maxThreads
	default:
		gridDimX = blocks
		blockDimX = maxThreads
	}
	return
}

// Init creates a CUDA engine with the given size for the given device
func (e *Engine) Init(device cu.Device, size int64) (err error) {
	e.Lock()
	initialized := e.initialized
	e.Unlock()

	if initialized {
		return nil
	}

	e.Lock()
	defer e.Unlock()
	e.d = device
	if err = e.doInit(size); err != nil {
		e.Unlock()
		err2 := e.Close()
		if err2 != nil {
			return errors.Wrapf(err, "Failed to initialize CUDA Engine with size %d for device %v. Additionally, there were errors that occurred when cleaning up %v", size, device, err)
		}
		return errors.Wrapf(err, "Failed to initialize CUDA Engine with size %d for device %v", size, device)
	}
	e.initialized = true
	return
}

func (e *Engine) doInit(size int64) (err error) {
	e.workAvailable = make(chan bool)
	e.syncChan = make(chan struct{})
	e.finishChan = make(chan struct{})
	e.a = allocator.Make(memalign)

	// create and set context
	var cuctx cu.CUContext
	ctxFlag := cu.SchedAuto
	if cuctx, err = e.d.MakeContext(ctxFlag); err != nil {
		if err == cu.OutOfMemory {
			free, total, err2 := cu.MemInfo()
			if err2 != nil {
				return errors.Wrapf(err, "Out of memory. Additionally errors were found while retrieving mem info %v", err2)
			}
			return errors.Wrapf(err, "Out of memory. Free: %v, total %v | %v", free, total, cuctx)
		}
		return errors.Wrapf(err, "Failed to make context for device %d", e.d)
	}
	e.c = *(cu.NewBatchedContext(cu.CtxFromCUContext(e.d, cuctx, ctxFlag), e.d))

	var attrs []int
	if attrs, err = e.d.Attributes(cu.WarpSize, cu.MaxThreadsPerBlock, cu.MaxGridDimX, cu.MaxGridDimY, cu.MaxGridDimZ, cu.MaxBlockDimX, cu.MaxBlockDimY, cu.MaxBlockDimZ); err != nil {
		return errors.Wrapf(err, "Failed to get attributes for device %v.", e.d)
	}

	e.warp = attrs[0]
	e.mtpb = attrs[1]
	e.mgdx = attrs[2]
	e.mgdy = attrs[3]
	e.mgdz = attrs[4]
	e.mbdx = attrs[5]
	e.mbdy = attrs[6]
	e.mbdz = attrs[7]

	e.m = make(map[string]cu.Module)
	e.f = make(map[string]cu.Function)

	// actual work to allocate from graphics card

	if e.freeMem, e.totalMem, err = cu.MemInfo(); err != nil {
		return errors.Wrapf(err, "Failed to get free and total mem for device %v", e.d)
	}

	// actually reserve memory for the allocator
	var allocsize int64 = 2*size + (size / 2) + minAllocSize
	if size <= 0 {
		// if the hint given is undefined, we just reserve 80% off free memory
		allocsize = e.freeMem * 8 / 10
	}

	if allocsize >= e.freeMem {
		return errors.Errorf("Unable to get %v bytes. Free memory available %v", allocsize, e.freeMem)
	}
	ptr, err := cu.MemAllocManaged(allocsize, cu.AttachGlobal)
	if err != nil {
		return errors.Wrapf(err, "Failed to allocate %v bytes of managed memory for %v", allocsize, e.d)
	}
	e.a.Reserve(uintptr(ptr), allocsize)
	e.n = *(cudnn.NewContext())
	go e.Run()
	return nil
}

// Close cleans up the machine, and closes all available resources
func (e *Engine) Close() error {
	e.Lock()
	defer e.Unlock()
	e.c.Cleanup() // frees all ancillary allocations in C land
	if e.c.Context == nil {
		return nil
	}
	cu.SetCurrentContext(e.c.Context.CUDAContext())

	// Unload all modules (and consequently all functions)
	for name, mod := range e.m {
		if err := mod.Unload(); err != nil {
			return errors.Wrapf(err, "Failed to unload module %v", name)
		}
	}

	// Free all CUDA memory
	if e.a.Start() != 0 {
		cu.MemFree(cu.DevicePtr(e.a.Start()))
	}
	e.a.Reset()

	closeB := func() error { return e.b.Close() }

	if err := e.c.Do(closeB); err != nil {
		return errors.Wrap(e.err, "Failed to close cuBLAS context")
	}

	closeN := func() error { return e.n.Close() }

	if err := e.c.Do(closeN); err != nil {
		return errors.Wrap(e.err, "Failed to close cuDNN context")
	}

	if e.workAvailable != nil {
		close(e.workAvailable)
	}

	if err := e.c.Close(); err != nil {
		return errors.Wrapf(err, "Failed to cloes CUDA Context ")
	}

	runtime.Gosched()   // make sure everyone has a fair play
	close(e.finishChan) // tell Run() to finish up
	e.wg.Wait()         // wait for Run() to finish
	e.initialized = false
	e.running = false
	return nil
}

// DoWork sends a signal to the batched CUDA Context to actually do work
func (e *Engine) DoWork() error {
	e.c.DoWork()
	return e.c.Errors()
}

// Run initialises and starts the engine
func (e *Engine) Run() { e.once.Do(e.run) }

func (e *Engine) run() {
	e.Lock()
	if e.running {
		e.Unlock()
		return
	}
	e.Unlock()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	e.wg.Add(1)
	defer e.wg.Done()

	// finish initialization
	if err := e.b.Init(cublas.WithContext(&e.c)); err != nil {
		panic(err)
	}

loop:
	for {
		select {
		case <-e.c.WorkAvailable():
			e.c.DoWork()
			if err := e.c.Errors(); err != nil {
				e.err = err
				break loop
			}
		case w := <-e.c.Work():
			if w != nil {
				err := w()
				e.c.ErrChan() <- err

				if err != nil {
					e.err = err
					break loop
				}
			} else {
				// if nil, it means the channel hasa been closed
				break loop
			}
		case <-e.finishChan:
			break loop
		}
	}
	return
}

// blockThread is an easier version of calculating <<threads, blocks>> for CUDA. Useful for debugging
func (e *Engine) blockThread(n, dev int) (blocks, threads int) {
	switch {
	case n <= 32:
		threads = 32
	case n <= 64:
		threads = 64
	case n <= 128:
		threads = 128
	case n <= 256:
		threads = 256
	case n <= 512:
		threads = 512
	default:
		threads = 1024
	}

	blocks = (n + threads - 1) / threads
	if blocks < 0 || blocks > 128 {
		blocks = 128
	}
	return
}

// it's just a generic ceiling function. Added here to avoid mixing with any potential ceilInt operation
func calcBlocks(n, maxThreads int) int { return (n + maxThreads - 1) / maxThreads }
