package go_nd

import "github.com/c2akula/go.nd/nd/iter"

// Scale performs the element-wise operation y = a*x.
func Scale(a float64, x, y iter.Iterator) {
	if x.Len() != y.Len() {
		panic("input iterators must have same size.")
	}

	xd := x.Data()
	if x == y {
		for _, k := range x.Ind() {
			xd[k] *= a
		}
		return
	}

	yd := y.Data()
	xi := x.Ind()
	yi := y.Ind()
	for i := 0; i < x.Len(); i++ {
		yd[yi[i]] = a * xd[xi[i]]
	}
}
