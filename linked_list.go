package main

// import (
//   "fmt"
// )

type List struct {
  root Node
}

type Node struct {
  next *Node
  value interface{}
}

func (e *Element) Next() *Element{
  if node := e.next; e.list != nil && node != &e.list.root{
    return n
  }
  return nil
}


func (n *Node) init() *List{
  l.root.next = &l.root
  l.len = 0
  return l
}

func (n *Node) insert(e, target *Element) *Element {
  node := target.Next()
  target.next = e
  e.next = node
  l.len = l.len + 1
  return e
}

func (n *Node) insertValue(val interface{}, at *Element) *Element {
  
}

func (n *Node) insertAtFront(val interface{}) *Element{

}

func (n *Node) insertAtEnd(val interface{}) *Element {

}

func New() *List {

}

func main() {
  list := New()
  node1 := list.insertValue(1, l.root.next)
  node2 := list.insertValue(2, node1)
  node3 := list.insertValue(3, node2)
  node4 := list.insertValue(4, node3)
  node5 := list.insertValue(5, node4)
  node6 := list.insertValue(6, node5)
}

