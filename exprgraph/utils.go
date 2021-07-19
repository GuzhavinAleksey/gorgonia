package exprgraph

import (
	"fmt"
	"strconv"

	"gorgonia.org/gorgonia/values/dual"
	"gorgonia.org/tensor"
)

// T2T tries to find a `tensor.Tensor` from a Tensor
// it returns nil if no tensor is found
func T2T(a Tensor) tensor.Tensor {
	switch t := a.(type) {
	case *Node:
		if t.Tensor == nil {
			return nil
		}
		return T2T(t.Tensor.(Tensor))
	case *dual.Dual:
		panic("Not handled yet")
	case tensor.Tensor:
		return t
	default:
		return nil
	}
}

func consFmtStr(a fmt.State, c rune) string {
	retVal := "%"
	acceptable := []rune{'+', '-', ' ', '#', '0'}
	for _, f := range acceptable {
		if a.Flag(int(f)) {
			retVal = retVal + string(f)
		}
	}
	width, wok := a.Width()
	prec, pok := a.Precision()
	if wok {
		retVal = retVal + strconv.Itoa(width)
	}
	if pok {
		retVal = retVal + "." + strconv.Itoa(prec)
	}
	retVal = retVal + string(c)
	return retVal
}
