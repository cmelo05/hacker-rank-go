package migratorybirds

import "testing"

func TestMigratoryBirds(t *testing.T){
	input := []int32{1,2,3,4,5,4,3,2,1,3,4};
	var expectedResult int32 = 3

	actualResult := migratoryBirds(input)

	if(actualResult != expectedResult){
		t.Errorf("Wanted %d got %d",expectedResult,actualResult)
	}
}