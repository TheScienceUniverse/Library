package des

import(
	"fmt"
)

var IP = [64]byte {
0x39, 0x31, 0x29, 0x21, 0x19, 0x11, 0x09, 0x01,
0x3b, 0x33, 0x2b, 0x23, 0x1b, 0x13, 0x0b, 0x03,
0x3d, 0x35, 0x2d, 0x25, 0x1d, 0x15, 0x0d, 0x05,
0x3f, 0x37, 0x2f, 0x27, 0x1f, 0x17, 0x0f, 0x07,
0x38, 0x30, 0x28, 0x20, 0x18, 0x10, 0x08, 0x00,
0x3a, 0x32, 0x2a, 0x22, 0x1a, 0x12, 0x0a, 0x02,
0x3c, 0x34, 0x2c, 0x24, 0x1c, 0x14, 0x0c, 0x04,
0x3e, 0x36, 0x2e, 0x26, 0x1e, 0x16, 0x0e, 0x06 }

var FP = [64]byte {
0x27, 0x07, 0x2f, 0x0f, 0x37, 0x17, 0x3f, 0x1f,
0x26, 0x06, 0x2e, 0x0e, 0x36, 0x16, 0x3e, 0x1e,
0x25, 0x05, 0x2d, 0x0d, 0x35, 0x15, 0x3d, 0x1d,
0x24, 0x04, 0x2c, 0x0c, 0x34, 0x14, 0x3c, 0x1c,
0x23, 0x03, 0x2b, 0x0b, 0x33, 0x13, 0x3b, 0x1b,
0x22, 0x02, 0x2a, 0x09, 0x32, 0x12, 0x3a, 0x1a,
0x21, 0x01, 0x29, 0x08, 0x31, 0x11, 0x39, 0x19,
0x20, 0x00, 0x28, 0x07, 0x30, 0x10, 0x38, 0x18 }

var EB = [48]byte {
0x03, 0x00, 0x01, 0x02, 0x03, 0x04,
0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c,
0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
0x0f, 0x10, 0x11, 0x12, 0x13, 0x14,
0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c,
0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x00 }

var SB = [8][4][16]byte {
{{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13}},
{{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9}},
{{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12}},
{{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14}},
{{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3}},
{{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13}},
{{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12}},
{{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11}}}

func B2b(n uint, B byte) []byte {
	b := make([]byte, n)
	var i uint
	for i = 0; i < n; i++ {
		b[n-i-1] = (B >> i) & 0x01
	}
	return b
}
func b2B(n int, b []byte) byte {
	var B byte = 0
	var i int
	for i = 0; i < n; i++ {
		B = (B << 1) | b[i]
	}
	return B
}

func expandBytes(X []byte) []byte {
	Y := make([]byte, 64)
	var i, j uint
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			Y[8*i + j] = ((X[i] >> (7-j)) & 0x01)
		}
	}
	return Y
}
func shrinkBits(X []byte) []byte {
	Y := make([]byte, 8)
	var i, j uint
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			Y[i] = (Y[i] << 1) | X[8*i + j]
		}
	}
	return Y
}

func iniPermute(X []byte) []byte {
	Y := make([]byte, 64)
	var i int
	for i = 0; i < 64; i++ {
		Y[i] = X[IP[i]]
	}
	return Y
}
func finPermute(X []byte) []byte {
	Y := make([]byte, 64)
	var i int
	for i = 0; i < 64; i++ {
		Y[i] = X[FP[i]]
	}
	return Y
}

func getKey(r int) []byte {
	K := make([]byte, 48)
	var i int
	for i = 0; i < 48; i++ {
		K[i] = 1
	}
	return K
}

func coreDES(r int, R []byte) []byte {
	Y := make([]byte, 48)
	var i int
	var rs, cs, x byte
	for i = 0; i < 48; i++ {
		Y[i] = R[EB[i]]
	}
//fmt.Println(Y)
	for i = 0; i < 8; i++ {
		rs = (Y[6*i] << 1) | Y[6*i + 5]
		cs = 0
		for x = 1; x < 5; x++ {
			cs = (cs << 1) | Y[6*i + int(x)]
		}
		x = SB[i][rs][cs]
//Z := B2b(4, x)
//fmt.Println(Z)
copy(R[4*i : 4*i+4], B2b(4,x))
	}
//fmt.Println(R)
	K := getKey(r)
	for i = 0; i < 48; i++ {
		Y[i] &= K[i]
	}
	return R
}

func FeistalRound(r int, X []byte) []byte {
	var i int
	L := make([]byte, 32)
	R := make([]byte, 32)
	for i = 0; i < 32; i++ {
		L[i] = X[i]; R[i] = X[32+i];
	}
	Y := coreDES(r, R)

	for i = 0; i < 32; i++ {
		L[i] ^= Y[i]
	}

	copy(X[0:32], R[:])
	copy(X[32:64], L[:])
	return X
}

func DES(X []byte) []byte {

	var i int
	X = expandBytes(X)
	X = iniPermute(X)

	for i = 0; i < 16; i++ {
		X = FeistalRound(i, X)
	}

	X = finPermute(X)
	X = shrinkBits(X)
	fmt.Println(X)

	return X
}
