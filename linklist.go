package linklist

import "fmt"

type Node struct {
	x         int
	pre, next *Node
}

type LinkList struct {
	hptr *Node
	pre  *Node
	cnt  int
}

func (ll *LinkList) add(num ...int) {

	// when head is null
	if len(num) == 1 {
		(*ll).push(num[0])
		return
	}

	if (*ll).cnt < num[1] {
		return
	}

	node := Node{x: num[0], pre: nil, next: nil}
	(*ll).cnt++

	if num[1] == 0 {
		node.next = (*ll).hptr
		(*ll).hptr.pre = &node
		(*ll).hptr = &node
		return
	}

	nextPtr := (*ll).hptr
	for i := 1; i < num[1]; i++ {
		nextPtr = (*nextPtr).next
	}

	node.pre = nextPtr
	node.next = nextPtr.next
	nextPtr.next = &node

	if node.next == nil {
		(*ll).pre = &node
	}
}

func (ll *LinkList) remove(index int) {

	if index == 0 {
		(*ll).pop()
		return
	}

	if index == (*ll).cnt {
		(*ll).dequeue()
		return
	}

	if (*ll).cnt < index {
		return
	}

	ptr := (*ll).hptr
	for i := 1; i < index; i++ {
		ptr = (*ptr).next
	}

	(*ptr).next.pre = (*ptr).pre
	(*ptr).pre.next = (*ptr).next
	(*ll).cnt--
}

func (ll *LinkList) push(num int) {

	if (*ll).hptr == nil {
		(*ll).hptr = &Node{x: num, pre: nil, next: nil}
		(*ll).pre = (*ll).hptr
		(*ll).cnt++
		return
	}

	node := Node{x: num, pre: (*ll).pre, next: nil}
	(*ll).pre.next = &node
	(*ll).pre = &node
	(*ll).cnt++
}

func (ll *LinkList) pop() {

	if (*ll).pre == nil {
		return
	}

	if (*ll).pre.pre == nil {
		(*ll).pre = nil
		(*ll).hptr = nil
	} else {
		(*ll).pre = (*ll).pre.pre
		(*ll).pre.next = nil
	}
	(*ll).cnt--
}

func (ll *LinkList) enqueue(num int) {
	(*ll).push(num)
}

func (ll *LinkList) dequeue() {

	if (*ll).hptr == nil {
		return
	}

	(*ll).hptr = (*ll).hptr.next
	(*ll).cnt--
}

func (ll *LinkList) show() {
	node := (*ll).hptr

	if node == nil {
		fmt.Println(nil)
		return
	}

	for i := 0; i <= (*ll).cnt; i++ {
		fmt.Println(node.x)

		if node.next == nil {
			break
		}

		node = (*node).next
	}
}

func (ll *LinkList) getCount() int {
	return (*ll).cnt
}

func (ll *LinkList) get(index int) (int, string) {

	if (*ll).cnt <= index {
		return 0, "error"
	}

	ptr := (*ll).hptr
	for i := 1; i <= index; i++ {
		ptr = (*ptr).next
	}

	return ptr.x, "success"
}

func (ll *LinkList) getAll() ([]int, string) {

	if (*ll).cnt == 0 {
		return []int{}, "success"
	}

	num := make([]int, (*ll).cnt)
	node := (*ll).hptr

	for i := 0; i <= (*ll).cnt; i++ {
		num[i] = node.x

		if node.next == nil {
			break
		}

		node = (*node).next
	}

	return num, "success"
}
