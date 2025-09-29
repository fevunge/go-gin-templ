package loop

import "fmt"

const count = 10

func Loop() {
	fmt.Println("Count to until 10")

	for i := range count {
		fmt.Println(i + 2)
	}
}
