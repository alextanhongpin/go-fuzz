package main

import (
	"fmt"
	"strconv"
)

func main() {
	// NOTE: When panic occurs in golang, the arguments will be printed out. If it's an integer, they will be formatted as hexadecimal (base 16) instead of decimal (base 10) numbers.
	// bug.FaintHeart(1000001)
	var (
		h = "0xf4241" // hex
		d = 1000001   // decimal
	)
	toHex := fmt.Sprintf("0x%x", d)
	fmt.Println("decimal to hex conversion:", toHex)

	// Ignore the 0x
	toDec, _ := strconv.ParseUint(h[2:], 16, 32)
	fmt.Println("hex to decimal conversion:", toDec)
}

// panic: wohooo:1000001
//
// goroutine 1 [running]:
// github.com/alextanhongpin/go-fuzz.FaintHeart(0xf4241)
//         /Users/alextanhongpin/Documents/golang/src/github.com/alextanhongpin/go-fuzz/gofuzz.go:8 +0xb
// d
// main.main()
//         /Users/alextanhongpin/Documents/golang/src/github.com/alextanhongpin/go-fuzz/cmd/main.go:11 +
// 0x2e
// exit status 2
//
// shell returned 1
