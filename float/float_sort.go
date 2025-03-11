package float

import (
	"fmt"
	"math/rand"
	"time"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for i := data.Len() - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type Float64Array []float64

func (p Float64Array) Len() int           { return len(p) }
func (p Float64Array) Less(i, j int) bool { return p[i] < p[j] }
func (p Float64Array) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NewFloat64Array() Float64Array {
	return make([]float64, 25)
}

func (p Float64Array) Fill(n int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < n; i++ {
		p[i] = 100 * (rand.Float64())
	}
}

func (p Float64Array) List() string {
	s := "{ "
	for i := 0; i < p.Len(); i++ {
		if p[i] == 0 {
			continue
		}
		s += fmt.Sprintf("%3.1f ", p[i])
	}
	s += " }"
	return s
}

func (p Float64Array) String() string {
	return p.List()
}

// // float64 is necessary as input to math.Sqrt()
// package main

// import (
// 	"fmt"
// 	float64 "the-way-to-go/float"
// )

// func main() {
// 	f1 := float64.NewFloat64Array()
// 	f1.Fill(10)
// 	fmt.Printf("Before sorting %s\n", f1)
// 	float64.Sort(f1)
// 	fmt.Printf("After sorting %s\n", f1)
// 	if float64.IsSorted(f1) {
// 		fmt.Println("The float64 array is sorted!")
// 	} else {
// 		fmt.Println("The float64 array is NOT sorted!")
// 	}
// }

// /* Output:
// The length of the vector p1 is: 5.000000
// The length of the vector p2 is: 6.403124
// The length of the vector p1 after scaling is: 25.000000
// Point p1 after scaling has the following coordinates: X 15.000000 - Y 20.000000
// */
