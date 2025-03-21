package min

import "fmt"

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
	Swap(i, j int)
}

func Min(data Miner) interface{} {
	min := data.ElemIx(0)
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			min = data.ElemIx(i)
		}
		// else {
		// 	data.Swap(i, i-1)
		// }
	}
	return min
}

func (p IntArray) List() string {
	s := "{ "
	for i := 0; i < p.Len(); i++ {
		// if p[i] == 0 {
		// 	continue
		// }
		s += fmt.Sprintf("%v ", p[i])
	}
	s += " }"
	return s
}

func (p IntArray) String() string {
	return p.List()
}

type IntArray []int

func (p IntArray) Len() int                  { return len(p) }
func (p IntArray) ElemIx(ix int) interface{} { return p[ix] }
func (p IntArray) Less(i, j int) bool        { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }

type StringArray []string

func (p StringArray) Len() int                  { return len(p) }
func (p StringArray) ElemIx(ix int) interface{} { return p[ix] }
func (p StringArray) Less(i, j int) bool        { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }
