package utopiantree

// Complete the utopianTree function below.
func utopianTree(n int32) int32 {
	var height int32 = 1

	for i := 1; int32(i) <= n; i++ {
		if i%2 == 0 {
			height++
		} else {
			height *= 2
		}
	}
	return height
}
