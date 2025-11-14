package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {

	result := head
	current := head

	for {
		next := current.Next

		if next.Val != 0 {
			current.Val += next.Val
			current.Next = next.Next
		} else {
			if next.Next == nil { // end of list
				current.Next = nil
				return result
			}
			current = current.Next // move the pointer
		}

	}

}

func makeList(values ...int) *ListNode {
	start := &ListNode{}
	current := start

	for _, value := range values {
		node := &ListNode{Val: value}
		current.Next = node
		current = current.Next
	}

	return start.Next
}

func main() {
	input := makeList(0, 1, 2, 3, 0, 4, 5, 6, 0)
	output := mergeNodes(input)

	for output != nil {
		println("VALUE: %v", output.Val)
		output = output.Next
	}
}
