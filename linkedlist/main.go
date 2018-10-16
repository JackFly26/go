package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func main() {
	list := *newList("10 20 500 30 -1")
	list.Print()
}

type Node struct {
	next, prev *Node
	value      int
}

type LinkedList struct {
	head   Node
	length int
}

func newList(str string) *LinkedList {
	ll := new(LinkedList).Init()
	list := strings.Split(str, " ")
	nums, _ := sliceAtoi(list)
	ll.head = Node{&ll.head, &ll.head, nums[0]}
	nums = nums[1:]
	for _, num := range nums {
		if num == -1 {
			return ll
		}
		ll.Append(num)
	}
	return ll
}

func (l LinkedList) At(index int) *Node {
	node := &l.head
	switch index < 0 {
	case true:
		for i := 0.0; i < math.Abs(float64(index)); i++ {
			node = node.prev
		}
		return node
	default:
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	}
}

func (l *LinkedList) InsertAt(index int, value int) {
	node := l.At(index - 1)
	next := node.next
	new := Node{next, node, value}
	node.next = &new
	next.prev = &new
	l.head = *l.At(0)
	l.length++
}

func (l LinkedList) Init() *LinkedList {
	l.head = *new(Node)
	l.length = 1
	return &l
}

func (l *LinkedList) Append(num int) {
	new := Node{&l.head, l.head.prev, num}
	l.head.prev.next = &new
	l.head.prev = &new
	l.length++
}

func (l LinkedList) Print() {
	node := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(node.value)
		node = *node.next
	}
}
func (l LinkedList) PrintRev() {
	node := l.head
	for i := 0; i < l.length; i++ {
		node = *node.prev
		fmt.Println(node.value)
	}
}
