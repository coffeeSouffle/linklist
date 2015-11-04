package linklist

import "fmt"

type Node struct {
	Pre, Next *Node
	Value     interface{}
}

type LinkList struct {
	Hptr *Node
	Tptr *Node
	Cnt  int
}

func (ll *LinkList) Add(v interface{}, num int) {

	// when head is null
	if ll.Cnt == num {
		ll.Push(v)
		return
	}

	if ll.Cnt < num {
		return
	}

	node := Node{Value: v, Pre: nil, Next: nil}
	ll.Cnt++

	if num == 0 {
		node.Next = ll.Hptr
		ll.Hptr.Pre = &node
		ll.Hptr = &node
		return
	}

	NextPtr := ll.Hptr
	for i := 1; i < num; i++ {
		NextPtr = (*NextPtr).Next
	}

	node.Pre = NextPtr
	node.Next = NextPtr.Next
	NextPtr.Next.Pre = &node
	NextPtr.Next = &node

	if node.Next == nil {
		ll.Tptr = &node
	}
}

func (ll *LinkList) Remove(index int) {

	if index == 0 {
		ll.Pop()
		return
	}

	if index == ll.Cnt {
		ll.Dequeue()
		return
	}

	if ll.Cnt < index {
		return
	}

	ptr := ll.Hptr
	for i := 1; i < index; i++ {
		ptr = (*ptr).Next
	}

	(*ptr).Next.Pre = (*ptr).Pre
	(*ptr).Pre.Next = (*ptr).Next
	ll.Cnt--
}

func (ll *LinkList) Push(v interface{}) {

	if ll.Hptr == nil {
		ll.Hptr = &Node{Value: v, Pre: nil, Next: nil}
		ll.Tptr = ll.Hptr
		ll.Cnt++
		return
	}

	node := Node{Value: v, Pre: ll.Tptr, Next: nil}
	ll.Tptr.Next = &node
	ll.Tptr = &node
	ll.Cnt++
}

func (ll *LinkList) Pop() {

	if ll.Tptr == nil {
		return
	}

	if ll.Tptr.Pre == nil {
		ll.Tptr = nil
		ll.Hptr = nil
	} else {
		ll.Tptr = ll.Tptr.Pre
		ll.Tptr.Next = nil
	}
	ll.Cnt--
}

func (ll *LinkList) Enqueue(v interface{}) {
	ll.Push(v)
}

func (ll *LinkList) Dequeue() {

	if ll.Hptr == nil {
		return
	}

	ll.Hptr = ll.Hptr.Next
	ll.Cnt--
}

func (ll *LinkList) Show() {
	node := ll.Hptr

	if node == nil {
		fmt.Println(nil)
		return
	}

	for i := 0; i <= ll.Cnt; i++ {
		fmt.Println(node.Value)

		if node.Next == nil {
			break
		}

		node = (*node).Next
	}
}

func (ll *LinkList) GetCount() int {
	return ll.Cnt
}

func (ll *LinkList) Get(index int) (interface{}, string) {

	if ll.Cnt <= index {
		return 0, "error"
	}

	ptr := ll.Hptr
	for i := 1; i <= index; i++ {
		ptr = (*ptr).Next
	}

	return ptr.Value, "success"
}

func (ll *LinkList) GetAll() ([]interface{}, string) {

	if ll.Cnt == 0 {
		return make([]interface{}, 0), "success"
	}

	num := make([]interface{}, ll.Cnt)
	node := ll.Hptr

	for i := 0; i <= ll.Cnt; i++ {
		num[i] = node.Value

		if node.Next == nil {
			break
		}

		node = (*node).Next
	}

	return num, "success"
}

func (ll *LinkList) GetReAll() ([]interface{}, string) {

	if ll.Cnt == 0 {
		return make([]interface{}, 0), "success"
	}

	num := make([]interface{}, ll.Cnt)
	node := ll.Tptr

	for i := 0; i <= ll.Cnt; i++ {
		num[i] = node.Value

		if node.Pre == nil {
			break
		}

		node = node.Pre
	}

	return num, "success"
}

// construct
func LL() LinkList {
	return LinkList{}
}
