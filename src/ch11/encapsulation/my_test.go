package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Name string
	Age  int
}

//func (e Employee) F() string {
//	fmt.Println(unsafe.Pointer(&e))
//	return e.Name
//}

func (e *Employee) F() string {
	fmt.Println(unsafe.Pointer(e))
	return e.Name
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"Bob", 20}
	fmt.Println(unsafe.Pointer(&e))
	fmt.Println(e.F())
}
