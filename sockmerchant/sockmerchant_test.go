package sockmerchant

import "testing"

func TestSockMerchant(t *testing.T) {
	firstInput := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	firstExpectedResult := int32(3)

	secondInput := []int32{4, 5, 5, 5, 6, 6, 4, 1, 4, 4, 3, 6, 6, 3, 6, 1, 4, 5, 5, 5}
	secondExpectedResult := int32(9)

	firstResult := sockMerchant(int32(len(firstInput)), firstInput)
	secondResult := sockMerchant(int32(len(secondInput)), secondInput)

	if firstExpectedResult != firstResult {
		t.Errorf("Expected %d got %d", firstExpectedResult, firstResult)
	}

	if secondExpectedResult != secondResult {
		t.Errorf("Expected %d got %d", secondExpectedResult, secondResult)
	}
}
