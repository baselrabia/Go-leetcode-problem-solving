package main 


// 1 


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res := new(ListNode)
    temp := res

    for list1 != nil && list2 != nil {
       
        if (list1.Val < list2.Val){
            temp.Next = list1
            temp = temp.Next
            list1 = list1.Next
        }else{
            temp.Next = list2
            temp = temp.Next
            list2 = list2.Next
        }
    }

    for list1 != nil {  
        temp.Next = list1
        temp = temp.Next
        list1 = list1.Next
    }

    for list2 != nil {
        temp.Next = list2
        temp = temp.Next
        list2 = list2.Next
    }

    res = res.Next 
    return res
}
