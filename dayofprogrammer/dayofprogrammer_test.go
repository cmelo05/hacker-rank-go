package dayofprogrammer

import "testing"

func TestDayOfProgrammer(t *testing.T) {
	firstExpectedResult := "12.09.2016"
	var firstInput int32 = 2016

	secondExpectedResult := "13.09.2017"
	var secondInput int32 = 2017

	thirdExpectedResult := "12.09.1800"
	var thirdInput int32 = 1800

	fourthExpectedResult := "13.09.2100"
	var fourthInput int32 = 2100

	fifthExpectedResult := "26.09.1918"
	var fifthInput int32 = 1918

	firstResult := dayOfProgrammer(firstInput)
	secondResult := dayOfProgrammer(secondInput)
	thirdResult := dayOfProgrammer(thirdInput)
	fourthResult := dayOfProgrammer(fourthInput)
	fifthResult := dayOfProgrammer(fifthInput)

	if firstExpectedResult != firstResult {
		t.Errorf("Expected " + firstExpectedResult + " got " + firstResult)
	}

	if secondExpectedResult != secondResult {
		t.Errorf("Expected " + secondExpectedResult + " got " + secondResult)
	}

	if thirdExpectedResult != thirdResult {
		t.Errorf("Expected " + thirdExpectedResult + " got " + thirdResult)
	}

	if fourthExpectedResult != fourthResult {
		t.Errorf("Expected " + fourthExpectedResult + " got " + fourthResult)
	}

	if fifthExpectedResult != fifthResult {
		t.Errorf("Expected " + fifthExpectedResult + " got " + fifthResult)
	}

}
