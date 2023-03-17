package fn

func Map[IN any, OUT any](slice []IN, f func(IN) OUT) []OUT {
	var ret []OUT
	for _, item := range slice {
		ret = append(ret, f(item))
	}
	return ret
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var ret []T
	for _, item := range slice {
		if f(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func Reduce[IN any, OUT any](slice []IN, init OUT, f func(OUT, IN) OUT) OUT {
	ret := init
	for _, item := range slice {
		ret = f(ret, item)
	}
	return ret
}

func For[T any](slice []T, f func(T)) {
	for _, item := range slice {
		f(item)
	}
}
