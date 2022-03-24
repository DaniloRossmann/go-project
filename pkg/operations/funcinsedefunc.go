package operations

func Result(x ...int) func() int {
	res := 0
	for _, v := range x {
		res += v
	}

	return func() int {
		return res * res
	}

}
