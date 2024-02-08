package linkedlist_test

import (
	"testing"
)

// LinkedList 구조를 go는 가지고 있지 않음
// 이와 같은 방식을 활용해서 만들 수 있음
type Node[T any] struct {
	Pointer *Node[T]
	data    T
}

func _append(nowNode *Node[any], nextNode *Node[any], final bool) *Node[any] {
	if final && nowNode.Pointer != nil {
		for n := nowNode; n != nil; n = n.Pointer {
			nowNode = n
		}
	}
	nowNode.Pointer = nextNode
	return nextNode
}

func SetLinkedList(basicNode *Node[any], datas []string, index *int, exist bool, final bool) *Node[any] {
	nowNode := basicNode
	var existNode *Node[any]

	if index != nil {
		for i := 0; i < *index; i++ {
			if nowNode == nil {
				return nil
			}
			nowNode = nowNode.Pointer
		}
		existNode = nowNode.Pointer
	}

	for _, data := range datas {
		nowNode = _append(nowNode, &Node[any]{nil, data}, final)
	}
	if exist {
		nowNode.Pointer = existNode
	}
	return basicNode
}

func _delete(deleteNode *Node[any], existNode *Node[any]) {
	deleteNode.Pointer = existNode
}

func DelteLinkedList(basicNode *Node[any], index *int) *Node[any] {
	nodePage := 0
	var nowNode *Node[any]
	var existNode *Node[any]
	var deleteNode *Node[any]

	nowNode = basicNode
	for n := nowNode; n != nil; n = n.Pointer {
		if nodePage == (*index - 1) {
			deleteNode = nowNode.Pointer
		} else if nodePage == (*index + 1) {
			existNode = nowNode.Pointer
			break
		}
		nowNode = nowNode.Pointer
		nodePage += 1
	}
	_delete(deleteNode, existNode)
	return basicNode
}

func TestLinked(t *testing.T) {
	// 구조 설명 -> {UserName, PreUser, val}
	// Pointer 지속적으로 nil를 통해 계속 다음 포인터를 잡아서 넣어줌
	// basic에서 Pointer nil은 basic.Pointer 통해 값을 채워주는 방식
	var result *Node[any]
	users := []string{"민수", "철수", "영희"}
	// LinkedList 자동 함수 만들기
	basic := &Node[any]{nil, users[0]}

	// 기본 LinkedList 생성 방법
	SetLinkedList(basic, users[1:], nil, false, true)
	// 마지막에 노드 추가
	SetLinkedList(basic, []string{"b"}, nil, false, true)

	// 중간에 노드 추가
	index := 2
	SetLinkedList(basic, []string{"1", "2"}, &index, true, false)

	// 중간에 노드 삭제
	result = DelteLinkedList(basic, &index)
	t.Log(result)
}
