package my_map_test

import "testing"

func TestInitMap(t *testing.T) {
	// 创建map
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1, m1[2])
	t.Logf("len m1=%d", len(m1))
	// 创建空map
	m2 := map[int]int{}
	t.Log(m2)
	// 添加元素
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	// 通过make创建map
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	// 定义空map
	m1 := map[int]int{}
	// 取值，有初始值
	t.Log(m1[0])
	// 判断值是否存在
	if v, ok := m1[1]; ok {
		t.Log("v is", v)
	} else {
		t.Log("v is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	// 遍历map
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
