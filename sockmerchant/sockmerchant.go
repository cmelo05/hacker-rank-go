package sockmerchant

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	i := 0
	values := make(map[int]int)

	for i < len(ar)-1 {
		_, present := values[i]
		if present {
			i++
			continue
		}

		j := i + 1
		for j < len(ar) {
			if ar[i] == ar[j] {
				values[j] = i
				break
			}
			j++
		}
		i++
	}

	return int32(len(values))
}
