package main

import "fmt"

type Node struct {
    Val string
    Next *Node
}

type LinkedList struct {
    Head *Node
    Length int
}
    
func (l *LinkedList) prepend(n *Node){
  // the head node gonna be the second node
  second := l.Head
  // the new node gonna be the head and it's next is the previous head 
  l.Head = n
  l.Head.Next = second
  l.Length++
}

func (l *LinkedList) append(n *Node){
  // the head node gonna be the second node
//   first := l.Head
    if l.Head == nil {
        l.Head = n
    } else {
        current := l.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = n
    }
    l.Length++

    // if(l.Head== nil){
    //     l.Head = n
    // }else{
    //     for l.Head.Next != nil {
    //         l.Head = l.Head.Next
    //     }
    //     l.Head.Next = n
    // }
//   l.Length++
}

func (l *LinkedList) print(){
    node := l.Head
    for l.Length != 0 {
        fmt.Printf("%s  ", node.Val)
        node = node.Next
        l.Length--
    }
    fmt.Printf("\n")
}

func (l *LinkedList) delete(value string){
    if(l.Length == 0){return }

    node := l.Head
    if (node.Val == value){
        l.Head = node.Next
        l.Length--
        return 
    }
    
    for (node.Next.Val != value){
        if(node.Next.Next != nil){return }
        node = node.Next
    }
    node.Next  = node.Next.Next
    l.Length--
}

func main() {
   myList := LinkedList{}
   node1 := &Node{Val:"1"}
   node2 := &Node{Val:"2"}
   node3 := &Node{Val:"3"}
//    node4 := &Node{Val:"4"}
//    node5 := &Node{Val:"5"}
//    node6 := &Node{Val:"6"}

   myList.append(node1)
   myList.prepend(node2)
   myList.append(node3)
//    myList.prepend(node4)
//    myList.append(node5)
//    myList.append(node6)

    fmt.Println(myList)
    myList.print()
    // myList.delete("15")

    // myList.print()
//    fmt.Printf("Size: %d\n", ListNode.Size())
}
