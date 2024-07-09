package algo

func FibRecurse(n uint64) uint64 {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return FibRecurse(n-1) + FibRecurse(n-2)
}

func FibItr(n uint64) []uint64 {
	arr := make([]uint64, n+1)

	arr[0] = 0

	if n > 0 {
		arr[1] = 1
	}

	for i := uint64(2); i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr
}
