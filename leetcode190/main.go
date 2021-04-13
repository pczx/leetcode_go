package main

func reverseBits(num uint32) uint32 {
	bits := make([]uint32, 32)
	for i := 0; i < 32; i++ {
		bits[31-i] = (num & uint32(1))
		num >>= 1
	}

	var ans uint32
	for i, v := range bits {
		if v == 1 {
			ans = 1 << i
		}
	}

	return ans
}
