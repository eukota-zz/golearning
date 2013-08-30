package main

import(
  "fmt"
  "bytes"
  "strconv"
)

func main() {
  myList := NewSLList()
  myList = myList.pop_front()
  fmt.Printf(myList.print())
  fmt.Printf("\n")

}

type sllist struct{
  data int
  next *sllist
}

func NewSLList() sllist{
  fmt.Println("Building linked list of ints. Enter each number on its own line:")
  l := sllist{}
  first := true
  b := true
  for b==true{
    var i int
    fmt.Scanf("%d", &i)
    if i==666 {
      b = false
    }
    if first==true {
      l.data=i
      first=false 
    } else {
      l.push_back(i) 
    }
  } 
  l.pop_back() 
  return l
}

func (l *sllist) pop_front() sllist{
  return *l.next
}

func (l *sllist) pop_back(){
  node := l
  nextnode := node.next
  for nextnode.next != nil {
    node = nextnode
    nextnode = nextnode.next
  }
  node.next = nil
}
func (l *sllist) push_back(i int){
  newNode := &sllist{data:i}
  node := l
  for node.next != nil {
    node = node.next
  }
  node.next = newNode
}

func (l *sllist) print() string{
  node := l
  var buffer bytes.Buffer
  for node != nil {
    buffer.WriteString(strconv.Itoa(node.data))
    buffer.WriteString("\n")
    node = node.next   
  }
  return buffer.String()
}

