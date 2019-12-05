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
  Val int
  Next *ListNode
}


func reverseList(head *ListNode) (result *ListNode) {
    
    var prev *ListNode
    curr := head

    for curr != nil {
      if curr.Next == nil  {
        result = curr
      }

      nCurr := curr.Next
      curr.Next = prev
      prev = curr
      curr = nCurr
    }

    return
}

func main() {
  fmt.Println("Hello World")
}