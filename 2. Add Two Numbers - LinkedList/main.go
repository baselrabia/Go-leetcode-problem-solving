package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
 

    count := 0
    var x, y int
    carry := 0
    sum := 0
    last := 0

    for l1!=nil || l2!=nil {
        x, y = 0, 0 
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

        sum = x + y + carry
        last = 0
        if( sum >= 10 ){
            last = sum % 10
            carry = 1
        } else {
            last = sum
            carry = 0
        }

        fmt.Println("sum:", sum, "last:", last, "carry:", carry)

        newNode := &ListNode{Val: last}

        if count == 0 {
            result = newNode
        } else {
            current := result
            for current.Next != nil {
                current = current.Next
            }
            current.Next = newNode
        }

        count++  
	}

    if(carry == 1){
        last = 1
        
        current := result
        for current.Next != nil {
            current = current.Next
        }
        current.Next = &ListNode{Val: last}
    }

    fmt.Println("sum:", sum, "last:", last, "carry:", carry)
 
	return result

}


  
func (l *ListNode) prepend(n *ListNode) *ListNode {
  // the head node gonna be the second node
  second := l
  // the new node gonna be the head and it's next is the previous head 
  l = n
  l.Next = second;

  return l
}


func (l *ListNode) print(){
    for l != nil {
        fmt.Printf("%d  ", l.Val)
        l = l.Next
    }
    fmt.Printf("\n")
}

func main() {
    //[9,9,9,9,9,9,9]
    l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9,Next:&ListNode{Val: 9,Next:&ListNode{Val: 9,Next:
          &ListNode{Val: 9,Next:&ListNode{Val: 9,Next:&ListNode{Val: 9}}}}}}}

    l2 := &ListNode{Val: 9,Next: &ListNode{Val: 9,Next:&ListNode{Val: 9,Next:&ListNode{Val: 9}}}}


    l1.print()
    l2.print()

    // node2 :=  &ListNode{Val: 22}
    
    // l1.prepend(node2).print()

    // fmt.Println("result:")
	
    addTwoNumbers(l1, l2).print()
}
