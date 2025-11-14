package main

import "testing"

func TestMakeList(t *testing.T) {
	output := makeList(1, 2, 3)
	expectedValues := []int{1, 2, 3}

	for _, expected := range expectedValues {
		if output.Val != expected {
			t.Errorf("Expected %v but got %v", expected, output.Val)
		}

		output = output.Next
	}

	if output != nil {
		t.Errorf("Expected list to be entirely consumed")
	}
}

// Input: head = [0,3,1,0,4,5,2,0]
// Output: [4,11]
func TestCase1(t *testing.T) {

	head := makeList(0, 3, 1, 0, 4, 5, 2, 0)

	newHead := mergeNodes(head)

	expectedValues := []int{4, 11}
	for _, expected := range expectedValues {
		if newHead.Val != expected {
			t.Errorf("Expected %v got %v", expected, newHead.Val)
		}
		newHead = newHead.Next
	}
}

// Input: head = [0,1,0,3,0,2,2,0]
// Output: [1,3,4]
func TestCase2(t *testing.T) {
	head := makeList(0, 1, 0, 3, 0, 2, 2, 0)

	newHead := mergeNodes(head)

	expectedValues := []int{1, 3, 4}
	for _, expected := range expectedValues {
		if newHead.Val != expected {
			t.Errorf("Expected %v got %v", expected, newHead.Val)
		}
		newHead = newHead.Next
	}
}
