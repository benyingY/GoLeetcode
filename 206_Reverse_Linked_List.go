/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isValid(s string) bool {
	size := len(s)
	stack = make([]byte,size)
	top := 0

	for i:=0;i<size;i++{
		ch = s[i]

		switch ch {
		case '(':
			stack[top]=ch+1
			top++
		case '[','{':
			stack[top]=ch+2
			top++
		case ')',']','}':
			if top >0 && stack[top-1]==ch {
				top--
			}
			else{
				return false
			}
		}
	}
	return top==0
}