package pp

import "fmt"

type Option struct {
	Name string
	Age int
	Class int
}



type OptionFunc func(*Option)


func WithName(name string) OptionFunc {
	return func(option *Option) {
		option.Name = name
	}
}

func WithAge(age int ) OptionFunc {
	return func(option *Option) {
		option.Age = age
	}
}

func WithClass(class int) OptionFunc {
	return func(option *Option) {
		option.Class = class
	}
}

func Apply(op *Option, opts * []OptionFunc) {
	for _, optfunc := range *opts {
		optfunc(op)
	}
}

func Test() {

	op := &Option{}

	opt := make([]OptionFunc, 0)
	opt = append(opt, WithName("huxiaoyu"))
	opt = append(opt, WithAge(25))
	opt = append(opt, WithClass(15))
	Apply(op, &opt)

	fmt.Println(op)
}


// ==============================

type DialOption interface {
	apply(option *Option)
}


type DialOptionFunc func(option *Option)

func (d DialOptionFunc) apply(option *Option) {
	d(option)
}

func newOptionFunc(f func(op *Option)) DialOption {
	return DialOptionFunc(f)
}

func WithName2(name string) DialOption {
	return newOptionFunc(func(op *Option){
		op.Name = name
	})
}

func WithAge2(age int) DialOption {
	return newOptionFunc(func(op *Option){
		op.Age = age
	})
}

func WithClass2(class int) DialOption {
	return newOptionFunc(func(op *Option){
		op.Class = class
	})
}

func Apply2(op *Option, opts *[]DialOption) {
	for _, optfunc := range *opts {
		optfunc.apply(op)
	}
}

func Test2() {

	op := &Option{}

	opt := make([]DialOption, 0)
	opt = append(opt, WithName2("huxiaoyu"))
	opt = append(opt, WithAge2(25))
	opt = append(opt, WithClass2(15))
	Apply2(op, &opt)

	fmt.Println(op)
}