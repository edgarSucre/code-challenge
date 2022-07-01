package linkedlist

type Key interface {
	string | int
}

type Node[T Key] struct {
	Val  T
	Next *Node[T]
}
