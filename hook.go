package hook

var (
	IgnoreHook bool
)

type Option[T interface{}] func(T)

type Stopper[T interface{}] interface {
	Stop(...Option[T])
}

type stopper[T interface{}] struct {
	value T
}

func (s stopper[T]) Stop(opts ...Option[T]) {
	for ix := range opts {
		opts[ix](s.value)
	}
}

type empty[T interface{}] struct{}

func (s empty[T]) Stop(opts ...Option[T]) {}

func Start[T interface{}](value T, opts ...Option[T]) Stopper[T] {
	if IgnoreHook {
		return empty[T]{}
	}
	for ix := range opts {
		opts[ix](value)
	}
	return stopper[T]{value: value}
}
