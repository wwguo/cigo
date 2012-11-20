package base

import "testing"


// not finished 
func (A Vector) TestInsertionSort (t *testing.T) {
	A = Vector{5, 2, 4, 6, 1, 3}
	expect := Vector{1, 2, 3, 4, 5, 6}
	A.InsertionSort(true)
    if !EqualVector(A, expect) {
		t.Errorf("Sorting result isn't right.\n   Expected: %v;\n   Actual: %v\n", expect, A)
	}
}

func (A Vector) TestMergeSort (t *testing.T) {
	A = Vector{2, 5, 4, 7, 1, 2, 3, 6}
	expect := Vector{1, 2, 2, 3, 4, 5, 6, 7}
	A.MergeSort(0, len(A), true)
    if !EqualVector(A, expect) {
		t.Errorf("Sorting result isn't right.\n   Expected: %v;\n   Actual: %v\n", expect, A)
	}
}

