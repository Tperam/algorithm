package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := ProduceRandomData(3000)
	QuickSort(data, 0, len(data)-1)
	fmt.Println(VerifyDataSort(data))
}

func TestQuickSort1(t *testing.T) {
	data := ProduceRandomData(3000)
	QuickSort1(data, 0, len(data)-1)
	fmt.Println(VerifyDataSort(data))
}

func BenchmarkQuickSort(b *testing.B) {
	data := ProduceRandomData(300000)
	QuickSort(data, 0, len(data)-1)
}

func BenchmarkQuickSort1(b *testing.B) {
	data := ProduceRandomData(300000)
	QuickSort1(data, 0, len(data)-1)
}

func ProduceRandomData(length int) []int {
	data := make([]int, length, length)

	for i := 0; i < length; i++ {
		data[i] = rand.Intn(10000)
	}

	return data
}

func VerifyDataSort(data []int) bool {
	length := len(data)
	for i := 0; i < length-1; i++ {
		if data[i] > data[i+1] {
			return false
		}
	}
	return true
}
