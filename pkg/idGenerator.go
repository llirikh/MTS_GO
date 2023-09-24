package pkg

type IdGenerator func(title string) uint64

func ReversePolynomialHash() IdGenerator {
	const (
		base uint64 = 1331
		mod  uint64 = 179
	)
	return func(str string) uint64 {
		var hash uint64 = 0
		for i := 0; i < len(str); i++ {
			hash = (hash*base + uint64(str[i])) % mod
		}
		return hash
	}
}

func PolynomialHash() IdGenerator {
	const (
		base uint64 = 137
		mod  uint64 = 327
	)
	return func(str string) uint64 {
		var hash uint64 = 0
		for i := len(str) - 1; i >= 0; i-- {
			hash = (hash*base + uint64(str[i])) % mod
		}
		return hash
	}
}
