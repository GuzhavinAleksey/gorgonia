package stdops

import (
	"context"
	"runtime/trace"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia/types"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// ElEq is a tensor-tensor elementwise equal-to.
type ElEq struct {
	binop
	retSame bool
}

// String implements fmt.Stringer.
func (op ElEq) String() string { return "=" }

// Type returns the type: (·) : a → a → a or (·) :  a → a → b
func (op ElEq) Type() hm.Type {
	a := hm.TypeVariable('a') // (T U) or U
	if op.retSame {
		return hm.NewFnType(a, a, a)
	}
	b := types.MakeDependent(a, tensor.Bool) // (T Bool) or Bool
	return hm.NewFnType(a, a, b)
}

// Do performs elementwise equal-to.
func (op ElEq) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise equal-to but with a preallocated return value.
// PreallocDo allows ElEq to implement ops.PreallocOp.
func (op ElEq) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// ElEqVS is a tensor-scalar elementwise equal-to.
type ElEqVS struct {
	binopVS
	retSame bool
}

// String implements fmt.Stringer.
func (op ElEqVS) String() string { return "=·" }

// Type returns the type: (·) : a → b → a or (·) :  a → b → c
func (op ElEqVS) Type() hm.Type {
	a := hm.TypeVariable('a') // (T U) or U
	b := hm.TypeVariable('b') // U
	if op.retSame {
		return hm.NewFnType(a, b, a)
	}
	c := types.MakeDependent(a, tensor.Bool) // (T Bool) or Bool
	return hm.NewFnType(a, b, c)
}

// Do performs elementwise equal-to.
func (op ElEqVS) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise equal-to but with a preallocated return value.
// PreallocDo allows ElEqVS to implement ops.PreallocOp.
func (op ElEqVS) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// ElEqSV is a scalar-tensor elementwise equal-to.
type ElEqSV struct {
	binopSV
	retSame bool
}

// String implements fmt.Stringer.
func (op ElEqSV) String() string { return "·=" }

// Type returns the type: (·) : a → b → b or (·) :  a → b → c
func (op ElEqSV) Type() hm.Type {
	a := hm.TypeVariable('a') // U
	b := hm.TypeVariable('b') // (T U) or U
	if op.retSame {
		return hm.NewFnType(a, b, b)
	}
	c := types.MakeDependent(b, tensor.Bool) // (T Bool) or Bool
	return hm.NewFnType(a, b, c)
}

// Do performs elementwise equal-to.
func (op ElEqSV) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise equal-to but with a preallocated return value.
// PreallocDo allows ElEqSV to implement ops.PreallocOp.
func (op ElEqSV) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}
