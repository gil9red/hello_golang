// hello_world project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	// Ожидание закрытия консоли
	var input string
	fmt.Scanln(&input)
}
