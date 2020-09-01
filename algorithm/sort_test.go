package algorithm

import (
	"reflect"
	"testing"
)

type numberTest struct {
	in[] int
	out []int
}

func TestNormalSort(t *testing.T) {

	numberTests := []numberTest{
		{[]int{5,2,5,7,1,0,8}, []int{0,1,2,5,5,7,8}},
		{[]int{6,5,4,1,2,3}, []int{1,2,3,4,5,6}},
	}

	for _, expected := range numberTests {
		bubbleSort(expected.in)
		if !reflect.DeepEqual(expected.in, expected.out) {
			t.Error("Unit test execute failed")
		}
	}

	numberTests = []numberTest{
		{[]int{5,2,5,7,1,0,8}, []int{0,1,2,5,5,7,8}},
		{[]int{6,5,4,1,2,3}, []int{1,2,3,4,5,6}},
	}

	for _, expected := range numberTests {
		insertSort(expected.in)
		if !reflect.DeepEqual(expected.in, expected.out) {
			t.Error("Unit test execute failed")
		}
	}

	numberTests = []numberTest{
		{[]int{5,2,5,7,1,0,8}, []int{0,1,2,5,5,7,8}},
		{[]int{6,5,4,1,2,3}, []int{1,2,3,4,5,6}},
	}

	for _, expected := range numberTests {
		selectSort(expected.in)
		if !reflect.DeepEqual(expected.in, expected.out) {
			t.Error("Unit test execute failed")
		}
	}
}

func TestMergeSort(t *testing.T)  {
	numberTests := []numberTest{
		{[]int{5,2,5,7,1,0,8}, []int{0,1,2,5,5,7,8}},
		{[]int{6,5,4,1,2,3}, []int{1,2,3,4,5,6}},
	}

	for _, expected := range numberTests {
		result := mergeSort(expected.in)
		if !reflect.DeepEqual(result, expected.out) {
			t.Errorf("Unit test execute failed, intput %v, output %v", expected.in, result)
		}
	}
}

func TestQuickSort(t *testing.T)  {

}