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
	even, err := fmr.SetSlice(INT_LIST).
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
	empty, err := fmr.SetSlice(INT_LIST).
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
	empty, err := fmr.SetSlice([]int{}).
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
	_, err := fmr.SetSlice(0).
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
	twice, err := fmr.SetSlice(INT_LIST).
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
	empty, err := fmr.SetSlice([]int{}).
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
	_, err := fmr.SetSlice(0).
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
	sum, err := fmr.SetSlice(INT_LIST).
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
	sum, err := fmr.SetSlice([]int{}).
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
	_, err := fmr.SetSlice(0).
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
	sum, err := fmr.SetSlice(INT_LIST).
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

func TestFilterChannelOk(t *testing.T) {
	ch := fmr.SetSlice(INT_LIST).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n%2 == 0
		}).
		Channel()

	value := <-ch
	switch v := value.(type) {
	case error:
		t.Error(v)
	case []interface{}:
		if !equals(v, []interface{}{2, 4, 6}) {
			t.Error("error filtering even numbers")
		}
	default:
		t.Error("wrong result type")
	}
}

func TestFilterChannelNoArray(t *testing.T) {
	ch := fmr.SetSlice(0).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n%2 == 0
		}).
		Channel()

	value := <-ch
	switch value.(type) {
	case error:
		//
	default:
		t.Error("should raise an error")
	}
}

func TestChannelClosed(t *testing.T) {
	ch := fmr.SetSlice(INT_LIST).
		Filter(func(it interface{}) bool {
			n := it.(int)
			return n%2 == 0
		}).
		Channel()

	i := 0
	for value := range ch {
		if i > 0 {
			t.Error("channel should close after result")
		}

		switch v := value.(type) {
		case error:
			t.Error(v)
		case []interface{}:
			if !equals(v, []interface{}{2, 4, 6}) {
				t.Error("error filtering even numbers")
			}
		default:
			t.Error("wrong result type")
		}
		i += 1
	}

}
