package interfac_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	F() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) F() string {
	return "F()"
}

func TestClient(t *testing.T) {
	var p Programmer // 接口变量
	p = &GoProgrammer{}
	fmt.Println(p.F())
}
