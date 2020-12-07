package main
// 两数字相加，思路：从链表最左边开始，同为相加，要考虑到进位的问题。注意真假怕判断。
// 熟悉go中链表结构体，以及if语句的语法
import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
  type ListNode struct {
	     Val int
	     Next *ListNode
	}
	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {	 
		var carry = 0
		var head *ListNode
		var tail *ListNode
	   for l1 != nil || l2 != nil { // 有一个不为0才可以继续，否则返回零
		   var i, j = 0,0 // 声明两个变量进行存值
		   if l1 != nil {
			   i = l1.Val
			   l1 = l1.Next
		   }
		   if l2 != nil {
			   j = l2.Val
			   l2 = l2.Next
		   }
		   var sum = i + j + carry
		   sum = (i + j + carry) % 10
		   carry = (i + j + carry) /10
		   if head == nil { // 如果为空，说明为首位
			   head = &ListNode{Val:sum}
			   tail = head   // 指向相同的地址
		   } else {
			   tail.Next = &ListNode{Val:sum}
			   tail = tail.Next
		   }
	   }
	   if carry != 0 { // 注意最后要判断进位与否
		   tail.Next = &ListNode{Val:carry}
	   }
	   return head
   }
func main()  {
	// 刚开始学go，为了熟悉语法和验证结果，构造了两个链表，并带入到上述方法中，看控制台的结果如下
	var l1 *ListNode = new(ListNode)
	l1.Val = 2
	var l2 *ListNode = new(ListNode)
	l2.Val = 5
	var l3 *ListNode = new(ListNode)
	l3.Val = 3
	l1.Next = l2
	l2.Next = l3
	var l4 *ListNode = new(ListNode)
	l4.Val = 2
	var l5 *ListNode = new(ListNode)
	l5.Val = 5
	var l6 *ListNode = new(ListNode)
	l6.Val = 3
	l4.Next = l5
	l5.Next = l6
	var x = addTwoNumbers(l1,l4)
	fmt.Println(1, x.Next.Next.Val)
	fmt.Println(2, x.Next.Val)
	fmt.Println(3, x.Val)
}