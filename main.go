package main

import (
	"fmt"
)

func testfor() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("Last i=%d", i)
}

func testfor2() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("i=%d\n", i)
	}
}

func testfor3() {
	i := 1
	for i <= 10 {
		i += 2
		if i%2 == 0 {
			continue

		}
		fmt.Printf("i=%d\n", i)
	}
}

func testfor5() {
	for n, i := 10, 1; n <= 19 && i <= 10; n, i = n+1, i+1 {
		fmt.Printf("n*i=%d%d\n", n, i)
	}
}

func main() {
	// testfor()
	testfor2()
	testfor3()
	testfor5()
}
