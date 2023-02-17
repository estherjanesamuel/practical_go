package hammingweight


func hammingWeight(num uint32) int {
	result := 0
	for i := 0; i < 32; i++ {
		if (num & (1 << i)) >> i == 1 {
			result++
		}
	}
	return result
}

func hammingWeight2(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	return (n & 1) + hammingWeight2(n >> 1)
}

func HammingWeight2(n uint32) uint32 {
	return hammingWeight2(n)
}

func HammingWeight(n uint32) int {
	return hammingWeight(n)
}