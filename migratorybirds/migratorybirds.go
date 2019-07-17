package migratorybirds

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {
	firstMax, _ := getMax(arr)

	arr2 := make([]int32, firstMax)

	for _, v := range arr {
		arr2[v-1]++
	}

	_, index := getMax(arr2)

	return index
}

func getMax(arr []int32) (max int32, index int32) {
	max = arr[0]
	index = 0
	for i, v := range arr {
		if v > max {
			max = v
			index = int32(i)
		}
	}
	index++
	return
}
