package test

import (
	"testing"

	"github.com/garciaguimeras/go-fmr/fmr"
)

var INT_LIST = []int{1, 2, 3, 4, 5, 6}

func equals(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFilterNumbers(t *testing.T) {
	even, err := fmr.NewSlice(INT_LIST).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n%2 == 0
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if !equals(even.([]interface{}), []interface{}{2, 4, 6}) {
		t.Error("error filtering even numbers")
	}
}

func TestFilterNumbersWithEmptyResult(t *testing.T) {
	empty, err := fmr.NewSlice(INT_LIST).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n > 20
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if !equals(empty.([]interface{}), []interface{}{}) {
		t.Error("error getting empty result")
	}
}

func TestFilterEmptyArray(t *testing.T) {
	empty, err := fmr.NewSlice([]int{}).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n > 20
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if !equals(empty.([]interface{}), []interface{}{}) {
		t.Error("error getting empty result")
	}
}

func TestFilterNoArray(t *testing.T) {
	_, err := fmr.NewSlice(0).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n > 20
		}).
		Get()

	if err == nil {
		t.Error("should raise an error")
	}
}

func TestMapNumbers(t *testing.T) {
	twice, err := fmr.NewSlice(INT_LIST).
		Map(func(it interface{}) interface{} {
			n := it.(int)
			return n * 2
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if !equals(twice.([]interface{}), []interface{}{2, 4, 6, 8, 10, 12}) {
		t.Error("error mapping numbers")
	}
}

func TestMapEmptyArray(t *testing.T) {
	empty, err := fmr.NewSlice([]int{}).
		Map(func(it interface{}) interface{} {
			n := it.(int)
			return n * 2
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if !equals(empty.([]interface{}), []interface{}{}) {
		t.Error("error getting empty result")
	}
}

func TestMapNoArray(t *testing.T) {
	_, err := fmr.NewSlice(0).
		Map(func(it interface{}) interface{} {
			n := it.(int)
			return n * 2
		}).
		Get()

	if err == nil {
		t.Error("should raise an error")
	}
}

func TestReduceNumbers(t *testing.T) {
	sum, err := fmr.NewSlice(INT_LIST).
		Reduce(func(res interface{}, it interface{}) interface{} {
			r := res.(int)
			n := it.(int)
			return r + n
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if sum.(int) != 21 {
		t.Error("error reducing numbers")
	}
}

func TestReduceEmptyArray(t *testing.T) {
	sum, err := fmr.NewSlice([]int{}).
		Reduce(func(res interface{}, it interface{}) interface{} {
			r := res.(int)
			n := it.(int)
			return r + n
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if sum != nil {
		t.Error("error getting null result")
	}
}

func TestReduceNoArray(t *testing.T) {
	_, err := fmr.NewSlice(0).
		Reduce(func(res interface{}, it interface{}) interface{} {
			r := res.(int)
			n := it.(int)
			return r + n
		}).
		Get()

	if err == nil {
		t.Error("should raise an error")
	}
}

func TestCombined(t *testing.T) {
	sum, err := fmr.NewSlice(INT_LIST).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n%2 == 0
		}).
		Map(func(it interface{}) interface{} {
			n := it.(int)
			return n * 2
		}).
		Reduce(func(res interface{}, it interface{}) interface{} {
			r := res.(int)
			n := it.(int)
			return r + n
		}).
		Get()

	if err != nil {
		t.Error(err)
	}
	if sum.(int) != 24 {
		t.Error("error in combined function")
	}
}
