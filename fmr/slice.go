package fmr

type FilterFunction func(elem interface{}) bool
type MapFunction func(elem interface{}) interface{}
type ReduceFunction func(elem1 interface{}, elem2 interface{}) interface{}

type Slice struct {
	data      interface{}
	functions []interface{}
}

func NewSlice(data interface{}) *Slice {
	return &Slice{
		data:      data,
		functions: []interface{}{},
	}
}

func (s *Slice) Filter(f FilterFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

func (s *Slice) Map(f MapFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

func (s *Slice) Reduce(f ReduceFunction) *Slice {
	s.functions = append(s.functions, f)
	return s
}

func (s *Slice) Get() (interface{}, error) {
	ctx := CreateContext(s.data)
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
