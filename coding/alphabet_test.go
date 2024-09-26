package coding

import "fmt"

func Example() {
	alpha := Alphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	char := alpha.Char(15, 11)
	fmt.Println(string(char), "=", alpha.Index(15, char))

	// Output: TODO
}
