package base

import "testing"

// b := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
// A := matrix.MakeMatrix(4,4,b)
// Z := matrix.MakeZero(5,3)
// U := matrix.MakeUnit(4,5)
// I := matrix.MakeIdentity(4)
	
	// fmt.Printf("%v\n", A)
	// fmt.Printf("%v\n", Z)
	// fmt.Printf("%v\n", U)
	// fmt.Printf("%v\n", I)

// not finished 
func TestMakeZero(t *testing.T) {
	// Z := matrix.MakeZero(5,3)
	// result := Matrix{[]flaot64{0 0 0 0 0 0 0 0 0 0 0 0 0 0 0}, 5, 3}
    // if Z != result {
	// 	t.Errorf("net = %g, expect = 20.\n", net )
	// }
}

// func TestParse(t *testing.T) {
// 	s := `[1 2 3;4 5 6]`
// 	A, err := ParseMatlab(s)
	
// 	if err != nil {
// 		t.Fatal(err)
// 	}
	
// 	Ar := MakeDenseMatrix([]float64{1,2,3,4,5,6}, 2, 3)
// 	if !Equals(A, Ar) {
// 		t.Error()
// 	}
// }