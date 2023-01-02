package optional

//该用法为golang的optional设计模式，旨在通用option

type Opt[T any] func(t *T)

// BindOpt 为结构体绑定定义好的options
func BindOpt[T any](t *T, opts ...Opt[T]) {
	for _, opt := range opts {
		opt(t)
	}
}

type OptWithExcept[T any] func(t *T) error

// BindOptWithExcept 为结构体绑定定义好的options，且支持options异常中止
func BindOptWithExcept[T any](t *T, opts ...OptWithExcept[T]) error {
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}
