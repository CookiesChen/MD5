package md5

import "fmt"

var(
	X1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,15}
	X2 = []int{1, 6,11, 0, 5,10,15, 4, 9,14, 3, 8,13, 2, 7,12}
	X3 = []int{5, 8,11,14, 1, 4, 7,10,13, 0, 3, 6, 9,12,15, 2}
	X4 = []int{0, 7,14, 5,12, 3,10, 1, 8,15, 6,13, 4,11, 2, 9}

	T = []uint32{
		0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
		0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
		0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
		0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
		0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
		0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
		0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
		0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
		0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
	}

	S1 = []uint{ 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22 }
	S2 = []uint{ 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20}
	S3 = []uint{ 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23 }
	S4 = []uint{ 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21 }

)

func compression(Y []uint32)  {
	start := make([]uint32, 4)
	for i := 0; i < 4; i++ {
		start[i] = CV[i]
	}
	for i := 0; i < 4; i++ {
		g(Y, i)
	}
	CV[0] += start[0]
	CV[1] += start[1]
	CV[2] += start[2]
	CV[3] += start[3]
}

func g(Y []uint32, times int)  {
	var X []int
	var function func(uint32, uint32, uint32) uint32
	var S []uint
	switch times {
	case 0:
		X = X1
		function = F
		S = S1
		break
	case 1:
		X = X2
		function = G
		S = S2
		break
	case 2:
		X = X3
		function = H
		S = S3
		break
	case 3:
		X = X4
		function = I
		S = S4
		break
	}
	for i := 0; i < 16; i++ {
		a, b, c, d := CV[0], CV[1], CV[2], CV[3]
		CV[1] = b + CLS(a + function(b, c, d) +  Y[X[i]] + T[times*16 + i], S[i])
		CV[0], CV[2], CV[3] = d, b, c
		fmt.Printf("%x %x %x %x\n", CV[0],CV[1],CV[2],CV[3])
	}

}

func F(b, c, d uint32) uint32 {
	return (b&c)|((^b)&d)
}

func G(b, c, d uint32) uint32 {
	return (b&d)|(c&(^d))
}

func H(b, c, d uint32) uint32 {
	return b^c^d
}

func I(b, c, d uint32) uint32 {
	return c^(b|(^d))
}

func CLS(num uint32, bitNum uint) uint32 {
	high, low := num, num
	high = high << bitNum
	low = low >> (32 - bitNum)
	return high + low
}