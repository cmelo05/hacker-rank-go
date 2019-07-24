package utopiantree

import "testing"

func TestUtopianTree(t *testing.T) {
	firstInput := []int32{0, 1, 4}
	firstExpectedResult := []int32{1, 2, 7}

	for i, val := range firstInput {
		result := utopianTree(val)
		expectedResult := firstExpectedResult[i]

		if result != expectedResult {
			t.Errorf("Wanted %d got %d", expectedResult, result)
		}
	}
}
