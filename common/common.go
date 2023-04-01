package common

type Options struct {
	Query map[string]interface{}
}

type Option interface {
	Apply(*Options)
}

type OptionFunc func(*Options)

func (f OptionFunc) Apply(o *Options) {
	f(o)
}
