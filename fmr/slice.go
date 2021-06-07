// Provides filter, map and reduce functions to use with Go slices
package fmr

// Filter function
type FilterFunction func(elem interface{}) bool

// Map function
type MapFunction func(elem interface{}) interface{}

// Reduce function
type ReduceFunction func(elem1 interface{}, elem2 interface{}) interface{}

// Contains the slice data and the function chain to be applied
type Slice struct {
	data      interface{}
	functions []interface{}
}

// Sets the slice to be iterated
func SetSlice(data interface{}) *Slice {
	return &Slice{
		data:      data,
		functions: []interface{}{},
	}
}

// Adds a filter function to the chain
func (s *Slice) Filter(f FilterFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

// Adds a map function to the chain
func (s *Slice) Map(f MapFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

// Adds a reduce function to the chain
func (s *Slice) Reduce(f ReduceFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

// Runs the filter-map-reduce chain and gets a result
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
