package stdops

import (
	"context"
	"testing"

	"github.com/chewxy/hm"
	"github.com/stretchr/testify/assert"
	"gorgonia.org/gorgonia/internal/datatypes"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/shapes"
	"gorgonia.org/tensor"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

func TestlteVV(t *testing.T) {
	op := lteVV{}
	// basic test
	assert.Equal(t, 2, op.Arity())

	/* Do (using tensor-tensor) */

	// set up
	var a, b, c values.Value
	var expectedType hm.Type
	var expectedShape shapes.Shape
	var err error

	a = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6}))
	b = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 40, 50, 60}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass shape checking. Err: %v", err)
	}

	// actually doing and testing
	if c, err = op.Do(context.Background(), a, b); err != nil {
		t.Fatalf("Expected lteVV{} to work correctly. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correct := []bool{true, true, true, true, true, true}
	assert.Equal(t, correct, c.Data())

	/* PreallocDo (using scalar-scalar to make sure things don't go wrong) */

	// set up
	a = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{1}))
	b = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{2}))
	c = tensor.New(tensor.WithShape(), tensor.WithBacking([]bool{false}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass shape checking. Err: %v", err)
	}

	// actually PreallocDo-ing and testing
	c, err = op.PreallocDo(context.Background(), c, a, b)
	if err != nil {
		t.Fatalf("Expected lteVV{}'s Prealloc to work. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correctScalar := true
	assert.Equal(t, correctScalar, c.Data())

	// bad cases: fails  typecheck and shapecheck
	a = tensor.New(tensor.WithShape(2, 3), tensor.Of(tensor.Float64))
	b = tensor.New(tensor.WithShape(), tensor.Of(tensor.Float64))
	if expectedType, err = typecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteVV{} to NOT pass type checking. Got ~(%v %v) =  %v ", datatypes.TypeOf(a), datatypes.TypeOf(b), expectedType)
	}
	if expectedShape, err = shapecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteVV{} to NOT pass shape checking. Got expectedShape = %v", expectedShape)
	}

}

func TestlteVS(t *testing.T) {
	op := lteVS{}
	// basic test
	assert.Equal(t, 2, op.Arity())

	/* Do */

	// set up
	var a, b, c values.Value
	var expectedType hm.Type
	var expectedShape shapes.Shape
	var err error

	a = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6}))
	b = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{100}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVS{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVS{} to pass shape checking. Err: %v", err)
	}

	// actually doing and test
	if c, err = op.Do(context.Background(), a, b); err != nil {
		t.Fatalf("Expected lteVS{} to work correctly. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correct := []bool{true, true, true, true, true, true}
	assert.Equal(t, correct, c.Data())

	/* PreallocDo */

	// set up - create a new preallocated result
	c = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]bool{false, false, false, false, false, false}))

	// actually PreallocDo-ing and checking
	c, err = op.PreallocDo(context.Background(), c, a, b)
	if err != nil {
		t.Fatalf("Expected lteVS{}'s Prealloc to work. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	assert.Equal(t, correct, c.Data())

	/* bad cases: lteVS{} on tensor-tensor */

	b = tensor.New(tensor.WithShape(2, 3), tensor.Of(tensor.Float64))
	// we won't type check because the type system is not a dependent type system, thus
	// lteVS : (a → b → a) will always type check without errors
	if expectedShape, err = shapecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteVS{} to NOT pass shape checking. Got %v ~ (%v, %v) = %v", op.ShapeExpr(), a.Shape(), b.Shape(), expectedShape)
	}
}

func TestlteSV(t *testing.T) {
	op := lteSV{}
	// basic test
	assert.Equal(t, 2, op.Arity())

	/* Do */

	// set up
	var a, b, c values.Value
	var expectedType hm.Type
	var expectedShape shapes.Shape
	var err error

	a = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{100}))
	b = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteSV{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteSV{} to pass shape checking. Err: %v", err)
	}

	// actually doing and test
	if c, err = op.Do(context.Background(), a, b); err != nil {
		t.Fatalf("Expected lteSV{} to work correctly. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correct := []bool{false, false, false, false, false, false}
	assert.Equal(t, correct, c.Data())

	/* PreallocDo */

	// set up - create a new preallocated result
	c = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]bool{false, false, false, false, false, false}))

	// actually PreallocDo-ing and checking
	c, err = op.PreallocDo(context.Background(), c, a, b)
	if err != nil {
		t.Fatalf("Expected lteVS{}'s Prealloc to work. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	assert.Equal(t, correct, c.Data())

	/* bad cases: lteSV{} on tensor-tensor */

	a = tensor.New(tensor.WithShape(2, 3), tensor.Of(tensor.Float64))
	// we won't type check because the type system is not a dependent type system, thus
	// lteSV : (a → b → b) will always type check without errors
	if expectedShape, err = shapecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteSV{} to NOT pass shape checking. Got %v ~ (%v, %v) = %v", op.ShapeExpr(), a.Shape(), b.Shape(), expectedShape)
	}
}

func TestlteVV_RetSame(t *testing.T) {
	op := lteVV{retSame: true}
	// basic test
	assert.Equal(t, 2, op.Arity())

	/* Do (using tensor-tensor) */

	// set up
	var a, b, c values.Value
	var expectedType hm.Type
	var expectedShape shapes.Shape
	var err error

	a = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6}))
	b = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 40, 50, 60}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass shape checking. Err: %v", err)
	}

	// actually doing and testing
	if c, err = op.Do(context.Background(), a, b); err != nil {
		t.Fatalf("Expected lteVV{} to work correctly. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correct := []float64{1, 1, 1, 1, 1, 1}
	assert.Equal(t, correct, c.Data())

	/* PreallocDo (using scalar-scalar to make sure things don't go wrong) */

	// set up
	a = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{1}))
	b = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{2}))
	c = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{0}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVV{} to pass shape checking. Err: %v", err)
	}

	// actually PreallocDo-ing and testing
	c, err = op.PreallocDo(context.Background(), c, a, b)
	if err != nil {
		t.Fatalf("Expected lteVV{}'s Prealloc to work. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correctScalar := 1.0
	assert.Equal(t, correctScalar, c.Data())

	// bad cases: fails  typecheck and shapecheck
	a = tensor.New(tensor.WithShape(2, 3), tensor.Of(tensor.Float64))
	b = tensor.New(tensor.WithShape(), tensor.Of(tensor.Float64))
	if expectedType, err = typecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteVV{} to NOT pass type checking. Got ~(%v %v) =  %v ", datatypes.TypeOf(a), datatypes.TypeOf(b), expectedType)
	}
	if expectedShape, err = shapecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteVV{} to NOT pass shape checking. Got expectedShape = %v", expectedShape)
	}

}

func TestlteVS_RetSame(t *testing.T) {
	op := lteVS{retSame: true}
	// basic test
	assert.Equal(t, 2, op.Arity())

	/* Do */

	// set up
	var a, b, c values.Value
	var expectedType hm.Type
	var expectedShape shapes.Shape
	var err error

	a = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6}))
	b = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{100}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVS{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteVS{} to pass shape checking. Err: %v", err)
	}

	// actually doing and test
	if c, err = op.Do(context.Background(), a, b); err != nil {
		t.Fatalf("Expected lteVS{} to work correctly. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correct := []float64{1, 1, 1, 1, 1, 1}
	assert.Equal(t, correct, c.Data())

	/* PreallocDo */

	// set up - create a new preallocated result
	c = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{0, 0, 0, 0, 0, 0}))

	// actually PreallocDo-ing and checking
	c, err = op.PreallocDo(context.Background(), c, a, b)
	if err != nil {
		t.Fatalf("Expected lteVS{}'s Prealloc to work. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	assert.Equal(t, correct, c.Data())

	/* bad cases: lteVS{} on tensor-tensor */

	b = tensor.New(tensor.WithShape(2, 3), tensor.Of(tensor.Float64))
	// we won't type check because the type system is not a dependent type system, thus
	// lteVS : (a → b → a) will always type check without errors
	if expectedShape, err = shapecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteVS{} to NOT pass shape checking. Got %v ~ (%v, %v) = %v", op.ShapeExpr(), a.Shape(), b.Shape(), expectedShape)
	}
}

func TestlteSV_RetSame(t *testing.T) {
	op := lteSV{retSame: true}
	// basic test
	assert.Equal(t, 2, op.Arity())

	/* Do */

	// set up
	var a, b, c values.Value
	var expectedType hm.Type
	var expectedShape shapes.Shape
	var err error

	a = tensor.New(tensor.WithShape(), tensor.WithBacking([]float64{100}))
	b = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{1, 2, 3, 4, 5, 6}))

	// type and shape checks
	if expectedType, err = typecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteSV{} to pass type checking. Err: %v", err)
	}
	if expectedShape, err = shapecheck(op, a, b); err != nil {
		t.Fatalf("Expected lteSV{} to pass shape checking. Err: %v", err)
	}

	// actually doing and test
	if c, err = op.Do(context.Background(), a, b); err != nil {
		t.Fatalf("Expected lteSV{} to work correctly. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	correct := []float64{0, 0, 0, 0, 0, 0}
	assert.Equal(t, correct, c.Data())

	/* PreallocDo */

	// set up - create a new preallocated result
	c = tensor.New(tensor.WithShape(2, 3), tensor.WithBacking([]float64{0, 0, 0, 0, 0, 0}))

	// actually PreallocDo-ing and checking
	c, err = op.PreallocDo(context.Background(), c, a, b)
	if err != nil {
		t.Fatalf("Expected lteVS{}'s Prealloc to work. Err: %v", err)
	}
	assert.Equal(t, expectedType, datatypes.TypeOf(c))
	assert.True(t, expectedShape.Eq(c.Shape()))
	assert.Equal(t, correct, c.Data())

	/* bad cases: lteSV{} on tensor-tensor */

	a = tensor.New(tensor.WithShape(2, 3), tensor.Of(tensor.Float64))
	// we won't type check because the type system is not a dependent type system, thus
	// lteSV : (a → b → b) will always type check without errors
	if expectedShape, err = shapecheck(op, a, b); err == nil {
		t.Fatalf("Expected lteSV{} to NOT pass shape checking. Got %v ~ (%v, %v) = %v", op.ShapeExpr(), a.Shape(), b.Shape(), expectedShape)
	}
}
