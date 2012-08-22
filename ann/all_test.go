package ann

import "testing"

var zz = []float64 {1, 2, 3, 4}
var	vv = []float64 {4, 3, 2, 1}

func TestSu(t *testing.T) {
    if net := su(4, zz, vv); net != 20 {
		t.Errorf("net = %g, expect = 20.\n", net )
	}
}

func TestPu(t *testing.T) {
    if net := pu(4, zz, vv); net != 288 {
		t.Errorf("net = %g, expect = 288.\n", net )
	}
}

// TODO TestNeuron
// neuron AND and OR
// z1 := []float64 {0, 0, 1, 1}
// z2 := []float64 {0, 1, 0, 1}
// vv := []float64 {1, 2}
// var oo []float64
// n := ann.Neuron{Aggregate: ann.SU(2), Activate: ann.Linear(1), Weights: vv}
// for i := 0; i < 4; i++ {
// 	signal := []float64{z1[i],z2[i]}
// 	n.GenerateOutput(signal) 
// 	if n.Output - 2 < 0 {
// 		oo = append(oo, 0)
// 	} else {
// 		oo = append(oo, 1)
// 	}
// }
// fmt.Println(n,oo)
// n1 := ann.Neuron{Aggregate: ann.SU(2), Activate: ann.Linear(1), Weights: vv, Output: 23}
// n2 := ann.Neuron{Aggregate: ann.SU(2), Activate: ann.Linear(1), Weights: vv, Output: 12}
// fmt.Println(n1,n2)
// n := ann.Neuron{Aggregate: ann.SU(2), Activate: ann.Linear(1), Weights: vv, Inputs: []*float64{&n1.Output,&n2.Output}}
// n.GenerateOutput(nil)
// fmt.Println(n)
