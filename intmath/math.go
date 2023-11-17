package intmath

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	float32 | float64
}

func GCD[I Integer](a, b I) I {
	for b > 0 {
		t := a % b
		a = b
		b = t
	}
	return a
}

func CoPrimes[I Integer](a, b I) bool {
	if a&1 == 0 && b&1 == 0 {
		return false
	}
	return GCD(a, b) == 1
}

func Pow[I1 Integer, I2 Integer](base I1, pow I2) I1 {
	acc := I1(1)
	if pow < 0 {
		panic("fuck golang's type system")
	}
	for i := I2(0); i < pow; i++ {
		acc *= base
	}
	return acc
}

func Abs[S Signed](num S) S {
	if num < 0 {
		num *= -1
	}
	return num
}
