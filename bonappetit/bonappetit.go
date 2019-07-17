package bonappetit

import "fmt"

// Complete the bonAppetit function below.
func bonAppetit(bill []int32, k int32, b int32) {
	var sum int32 = 0
	for i, v := range bill {
		if int32(i) != k {
			sum += v
		}
	}

	split := sum / 2

	if split == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - split)
	}
}
