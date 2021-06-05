package fmr

import (
	"fmt"
	"reflect"
)

type Context struct {
	data interface{}
	err  error
}

func CreateContext(data interface{}) *Context {
	return &Context{
		data: data,
		err:  nil,
	}
}

func (c *Context) filterer(l reflect.Value, f FilterFunction) []interface{} {
	var result []interface{}
	for i := 0; i < l.Len(); i++ {
		item := l.Index(i).Interface()
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func (c *Context) mapper(l reflect.Value, f MapFunction) []interface{} {
	var result []interface{}
	for i := 0; i < l.Len(); i++ {
		item := l.Index(i).Interface()
		mapped := f(item)
		result = append(result, mapped)
	}
	return result
}

func (c *Context) reducer(l reflect.Value, f ReduceFunction) interface{} {
	if l.Len() == 0 {
		return nil
	}
	result := l.Index(0).Interface()
	for i := 1; i < l.Len(); i++ {
		item := l.Index(i).Interface()
		result = f(result, item)
	}
	return result
}

func (c *Context) executeFilter(f FilterFunction) {
	t := reflect.ValueOf(c.data)
	if t.Kind() == reflect.Slice {
		c.data = c.filterer(t, f)
	} else {
		c.err = fmt.Errorf("filter: need a slice to iterate, but found %o", c.data)
	}
}

func (c *Context) executeMap(f MapFunction) {
	t := reflect.ValueOf(c.data)
	if t.Kind() == reflect.Slice {
		c.data = c.mapper(t, f)
	} else {
		c.err = fmt.Errorf("map: need a slice to iterate, but found %o", c.data)
	}
}

func (c *Context) executeReduce(f ReduceFunction) {
	t := reflect.ValueOf(c.data)
	if t.Kind() == reflect.Slice {
		c.data = c.reducer(t, f)
	} else {
		c.err = fmt.Errorf("reduce: need a slice to iterate, but found %o", c.data)
	}
}
