package primitive

import (
	"math"
	_ "math"
	"testing"
	_ "unsafe"
)

func TestMaxInt(t *testing.T) {
	//fmt.Printf("Type: %T, Value: %v\n", 1<<64, 1<<64)	// overflows int

	if math.MaxInt8 != 1<<7-1 {
		panic("[WRONG] math.MaxInt8 == 1 << 7 - 1")
	}
	if math.MinInt8 != -1<<7 {
		panic("[WRONG] math.MinInt8 == -1 << 7")
	}
	if math.MaxInt16 != 1<<15-1 {
		panic("[WRONG] math.MaxInt16 == 1 << 15 - 1")
	}
	if math.MinInt16 != -1<<15 {
		panic("[WRONG] math.MinInt16 == -1 << 15")
	}
	if math.MaxInt32 != 1<<31-1 {
		panic("[WRONG] math.MaxInt32 == 1 << 31 - 1")
	}
	if math.MinInt32 != -1<<31 {
		panic("[WRONG] math.MinInt32 == -1 << 31")
	}
	if math.MaxInt64 != 1<<63-1 {
		panic("[WRONG] math.MaxInt64 == 1 << 63 - 1")
	}
	if math.MinInt64 != -1<<63 {
		panic("[WRONG] math.MinInt64 == -1 << 63")
	}
	if math.MaxUint8 != 1<<8-1 {
		panic("[WRONG] math.MaxUint8 == 1 << 8 - 1")
	}
	if math.MaxUint16 != 1<<16-1 {
		panic("[WRONG] math.MaxUint16 == 1 << 16 - 1")
	}
	if math.MaxUint32 != 1<<32-1 {
		panic("[WRONG] math.MaxUint32 == 1 << 32 - 1")
	}
	if math.MaxUint64 != 1<<64-1 {
		panic("[WRONG] math.MaxUint64 == 1 << 64 - 1")
	}

}
