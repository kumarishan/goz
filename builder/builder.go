package builder

type Builder[A any, C any] interface {
	AddOne(ele A) Builder[A, C]
	Clear() Builder[A, C]
	Result() C
}
