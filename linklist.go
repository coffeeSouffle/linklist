package linklist

import "fmt"

type Node struct {
	x         int
	Pre, Next *Node
}

type LinkList struct {
	Hptr *Node
	Pre  *Node
	Cnt  int
}

func (ll *LinkList) Add(num ...int) {

	// when head is null
	if len(num) == 1 {
		(*ll).Push(num[0])
		return
	}

	if (*ll).Cnt < num[1] {
		return
	}

	node := Node{x: num[0], Pre: nil, Next: nil}
	(*ll).Cnt++

	if num[1] == 0 {
		node.Next = (*ll).Hptr
		(*ll).Hptr.Pre = &node
		(*ll).Hptr = &node
		return
	}

	NextPtr := (*ll).Hptr
	for i := 1; i < num[1]; i++ {
		NextPtr = (*NextPtr).Next
	}

	node.Pre = NextPtr
	node.Next = NextPtr.Next
	NextPtr.Next = &node

	if node.Next == nil {
		(*ll).Pre = &node
	}
}

func (ll *LinkList) Remove(index int) {

	if index == 0 {
		(*ll).Pop()
		return
	}

	if index == (*ll).Cnt {
		(*ll).Dequeue()
		return
	}

	if (*ll).Cnt < index {
		return
	}

	ptr := (*ll).Hptr
	for i := 1; i < index; i++ {
		ptr = (*ptr).Next
	}

	(*ptr).Next.Pre = (*ptr).Pre
	(*ptr).Pre.Next = (*ptr).Next
	(*ll).Cnt--
}

func (ll *LinkList) Push(num int) {

	if (*ll).Hptr == nil {
		(*ll).Hptr = &Node{x: num, Pre: nil, Next: nil}
		(*ll).Pre = (*ll).Hptr
		(*ll).Cnt++
		return
	}

	node := Node{x: num, Pre: (*ll).Pre, Next: nil}
	(*ll).Pre.Next = &node
	(*ll).Pre = &node
	(*ll).Cnt++
}

func (ll *LinkList) Pop() {

	if (*ll).Pre == nil {
		return
	}

	if (*ll).Pre.Pre == nil {
		(*ll).Pre = nil
		(*ll).Hptr = nil
	} else {
		(*ll).Pre = (*ll).Pre.Pre
		(*ll).Pre.Next = nil
	}
	(*ll).Cnt--
}

func (ll *LinkList) Enqueue(num int) {
	(*ll).Push(num)
}

func (ll *LinkList) Dequeue() {

	if (*ll).Hptr == nil {
		return
	}

	(*ll).Hptr = (*ll).Hptr.Next
	(*ll).Cnt--
}

func (ll *LinkList) Show() {
	node := (*ll).Hptr

	if node == nil {
		fmt.Println(nil)
		return
	}

	for i := 0; i <= (*ll).Cnt; i++ {
		fmt.Println(node.x)

		if node.Next == nil {
			break
		}

		node = (*node).Next
	}
}

func (ll *LinkList) GetCount() int {
	return (*ll).Cnt
}

func (ll *LinkList) Get(index int) (int, string) {

	if (*ll).Cnt <= index {
		return 0, "error"
	}

	ptr := (*ll).Hptr
	for i := 1; i <= index; i++ {
		ptr = (*ptr).Next
	}

	return ptr.x, "success"
}

func (ll *LinkList) GetAll() ([]int, string) {

	if (*ll).Cnt == 0 {
		return []int{}, "success"
	}

	num := make([]int, (*ll).Cnt)
	node := (*ll).Hptr

	for i := 0; i <= (*ll).Cnt; i++ {
		num[i] = node.x

		if node.Next == nil {
			break
		}

		node = (*node).Next
	}

	return num, "success"
}

// construct
func LL() LinkList {
	return LinkList{}
}
