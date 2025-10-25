package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		b      byte   = math.MaxUint8  // 2^8 - 1
		smallI int32  = math.MaxInt32  // 2^31 - 1
		bigI   uint64 = math.MaxUint64 // 2^64 - 1
	)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
