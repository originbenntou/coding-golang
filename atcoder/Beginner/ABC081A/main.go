package main

import (
	"fmt"
	"unsafe"
)

// 5ms
func main() {
	var data []byte
	_, _ = fmt.Scanf("%s", &data)
	p := unsafe.Pointer(&data)

	var s = '1'
	checker := *(*byte)(unsafe.Pointer(&s))

	var b int
	for i := 0; i < len(data); i++ {
		if (*(*string)(p))[i] == checker {
			b++
		}
	}

	fmt.Printf("%d", b)
}

// 5ms
func main() {
	var data string
	_, _ = fmt.Scanf("%s", &data)

	var b int
	for i := 0; i < len(data); i++ {
		if rune(data[i]) == '1' {
			b++
		}
	}

	fmt.Printf("%d", b)
}

// 8ms
func main() {
	var data string
	_, _ = fmt.Scanf("%s", &data)

	var b int
	for i := 0; i < len(data); i++ {
		if string(data[i]) == string('1') {
			b++
		}
	}

	fmt.Printf("%d", b)
}
