// Package fmr provides filter, map and reduce functions to use with Go slices
package fmr

// FilterFunction defines the filter function prototype
type FilterFunction func(elem interface{}) bool

// MapFunction defines the map function prototype
type MapFunction func(elem interface{}) interface{}

// ReduceFunction defines the reduce function prototype
type ReduceFunction func(elem1 interface{}, elem2 interface{}) interface{}

// FunctionChain contains the slice data and the functions to be applied
type FunctionChain struct {
	data      interface{}
	functions []interface{}
}

// SetSlice sets the slice to be iterated
func SetSlice(data interface{}) *FunctionChain {
	return &FunctionChain{
		data:      data,
		functions: []interface{}{},
	}
}

// Filter adds a filter function to the function chain
func (fc *FunctionChain) Filter(f FilterFunction) *FunctionChain {
	fc.functions = append(fc.functions, f)
	return fc
}

// Map adds a map function to the function chain
func (fc *FunctionChain) Map(f MapFunction) *FunctionChain {
	fc.functions = append(fc.functions, f)
	return fc
}

// Reduce adds a reduce function to the function chain
func (fc *FunctionChain) Reduce(f ReduceFunction) *FunctionChain {
	fc.functions = append(fc.functions, f)
	return fc
}

// Get runs the filter-map-reduce chain and gets a result
func (fc *FunctionChain) Get() (interface{}, error) {
	// Creates a new context to apply the function chain on the data
	ctx := createContext(fc.data)
	for _, fn := range fc.functions {
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

func (fc *FunctionChain) Channel() <-chan interface{} {
	channel := make(chan interface{})

	go func() {
		defer close(channel)

		// Creates a new context to apply the function chain on the data
		ctx := createContext(fc.data)
		for _, fn := range fc.functions {
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
				channel <- ctx.err
				break
			}
		}
		channel <- ctx.data
	}()

	return channel
}
