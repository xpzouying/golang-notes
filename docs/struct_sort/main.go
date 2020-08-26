package main

import (
	"log"
	"math"
	"sort"
)

type Value int

// 按照value大小排序
func SortValues(values []Value) {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
}

// 按照与v的差值进行排序，差值越小，优先级越高
func SortValuesByDist(values []Value, v Value) {

	sort.Slice(values, func(i, j int) bool {
		dist1 := int(math.Abs(float64(values[i]) - float64(v)))
		dist2 := int(math.Abs(float64(values[j]) - float64(v)))

		return dist1 < dist2
	})

}

func f1() {
	v1 := Value(3)
	v2 := Value(1)
	v3 := Value(6)

	values := []Value{v1, v2, v3}
	SortValues(values)
	log.Printf("value sorted: %+v", values)
}

func f2() {
	v1 := Value(3)
	v2 := Value(1)
	v3 := Value(6)

	values := []Value{v1, v2, v3}
	SortValuesByDist(values, 3)
	log.Printf("value sorted: %+v", values)
}

func main() {
	// --- 按照大小排序 ---
	f1()

	// --- 按照距离排序 ---
	f2()
}
