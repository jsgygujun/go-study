package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	sum, carry := 0, 0
	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum += v1 + v2 + carry
		sum, carry = sum%10, sum/10
		next := new(ListNode)
		next.Val = sum
		curr.Next = next
		curr = next
	}
	return dummy.Next
}

func main() {

}
