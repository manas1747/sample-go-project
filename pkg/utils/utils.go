package utils

import "fmt"

func init() {
	fmt.Println("utils::init()")
}

var (
	Name = "Manas Khandeshe"
	Age  = 37
)

func GetNumber() int {
	return 15 * 5
}
