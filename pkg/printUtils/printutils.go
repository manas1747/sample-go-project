package printUtils

import "fmt"

func init() {
	fmt.Println("printUtils::init()")
}

func PrintAstericks() {
	maxColumns := 10

	for i := 0; i < maxColumns; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := maxColumns; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
