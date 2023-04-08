package map_ext

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	//值为函数的map
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	t.Log(m[1](2), m[2](2))
}

func TestMapForSet(t *testing.T) {
	//值为bool的map，实现set
	set := map[int]bool{}
	set[1] = true
	t.Log(set)

	n := 3
	if set[n] {
		t.Log(set[n])
	} else {
		t.Log(set[n])
	}

}
