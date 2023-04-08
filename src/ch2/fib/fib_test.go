package fib

import "testing"

//func TestFibList(t *testing.T) {
//	var (
//		a int = 1
//		b int = 1
//	)
//	for i := 0; i < 5; i++ {
//		t.Log(b)
//		var tmp int = b
//		b = a + b
//		a = tmp
//	}
//}

func TestEchange(t *testing.T) {
	var (
		a = 1
		b = 2
	)
	//tmp := a
	//a = b
	//b = tmp
	a, b = b, a
	t.Log(a, b)
}
