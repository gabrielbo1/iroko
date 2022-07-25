package pkg

import "testing"

func TestInCreate(t *testing.T) {
	inVal := InCreate([]int{1})
	if inVal != "(1)" {
		t.Errorf("Error InCreate function, expected (1) returned %v", inVal)
	}
	inVal = InCreate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if inVal != "(1,2,3,4,5,6,7,8,9,10)" {
		t.Errorf("Error InCreate function, expected (1) returned %v", inVal)
	}
	inVal = InCreate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if inVal != "(1,2,3,4,5,6,7,8,9)" {
		t.Errorf("Error InCreate function, expected (1) returned %v", inVal)
	}
	inVal = InCreate([]int{})
	if inVal != "" {
		t.Errorf("Error InCreate function, expected (1) returned %v", inVal)
	}
}
