// Package fmr provides filter, map and reduce functions to use with Go slices
package fmr

// FilterFunction defines the filter function prototype
type FilterFunction func(elem interface{}) bool

// MapFunction defines the map function prototype
type MapFunction func(elem interface{}) interface{}

// ReduceFunction defines the reduce function prototype
type ReduceFunction func(elem1 interface{}, elem2 interface{}) interface{}

// Slice contains the slice data and the function chain to be applied
type Slice struct {
	data      interface{}
	functions []interface{}
}

// SetSlice sets the slice to be iterated
func SetSlice(data interface{}) *Slice {
	return &Slice{
		data:      data,
		functions: []interface{}{},
	}
}

// Filter adds a filter function to the function chain
func (s *Slice) Filter(f FilterFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

// Map adds a map function to the function chain
func (s *Slice) Map(f MapFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

// Reduce adds a reduce function to the function chain
func (s *Slice) Reduce(f ReduceFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

// Get runs the filter-map-reduce chain and gets a result
func (s *Slice) Get() (interface{}, error) {
	// Creates a new context to apply the function chain on the data
	ctx := createContext(s.data)
	for _, fn := range s.functions {
		switch f := fn.(type) {
		case FilterFunction:
			ctx.executeFilter(f)
		case MapFunction:
			ctx.executeMap(f)
		case ReduceFunction:
			ctx.executeReduce(f)
		default:
		}

		if ctx.err != nil {
			return nil, ctx.err
		}
	}
	return ctx.data, nil
}
