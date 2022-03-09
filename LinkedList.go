package main

import (
	"fmt"
)

type LinkedList struct {
	head *LinkedListNode
}

func (l *LinkedList) InsertAthead(data int) {
	newNode := &LinkedListNode{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *LinkedList) InsertNodeAthead(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *LinkedList) InsertNodeAtTail(node *LinkedListNode) {
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

func (l *LinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		l.InsertAthead(lst[i])
	}
}

func (l *LinkedList) DisplayLinkedList() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
}
