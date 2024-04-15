package main

import (
	"dataStructure/linkedList/doubleLinkedList"
	"dataStructure/linkedList/singleLinkedList"
	"testing"
)

func BenchmarkSingleLinkedList(b *testing.B) {
	var l singleLinkedList.LinkedList[int]
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}
func BenchmarkSingleLinkedList2(b *testing.B) {
	var l singleLinkedList.LinkedList[int]
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse2()
}

func BenchmarkDoubleLinkedList(b *testing.B) {
	var l doubleLinkedList.LinkedList[int]
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}
