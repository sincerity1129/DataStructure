package main

// LinkedList 구조를 go는 가지고 있지 않음
// 이와 같은 방식을 활용해서 만들 수 있음
type Node[T any] struct {
	Pointer *Node[T]
	data    T
}

func GetLinkedList(basicNode *Node[any], nowNode *Node[any]) *Node[any] {
	basicNode.Pointer = nowNode
	return basicNode.Pointer
}

func main() {
	// 구조 설명 -> {UserName, PreUser, val}
	// Pointer 지속적으로 nil를 통해 계속 다음 포인터를 잡아서 넣어줌
	// basic에서 Pointer nil은 basic.Pointer 통해 값을 채워주는 방식

	users := []string{"민수", "철수", "영희"}
	// LinkedList 자동 함수 만들기
	basic := &Node[any]{nil, users[0]}
	// 포인터로 만들기 때문에 주소는 그대로 있기 때문에 변수가 바뀌더라도 저장 위치는 그대로 남겨 있음을 이용
	basicList := basic
	for idx, user := range users {
		if idx == 0 {
			continue
		}
		basicList = GetLinkedList(basicList, &Node[any]{nil, user})
	}
}
