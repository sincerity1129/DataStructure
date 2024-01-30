package main

import "fmt"

// LinkedList 구조를 go는 가지고 있지 않음
// 이와 같은 방식을 활용해서 만들 수 있음
type Node[T string] struct {
	UserName *Node[T]
	PreUser  *Node[T]
	val      T
}

func AutoGetList(users []string) (*Node[string], *Node[string]) {
	var getPointer *Node[string]
	var backGetPointer *Node[string]

	for idx, user := range users {
		if idx == 0 {
			basic := &Node[string]{nil, nil, user}
			backGetPointer = basic
		} else {
			backGetPointer = &Node[string]{nil, backGetPointer, user}
		}
	}
	for n := backGetPointer; n != nil; n = n.PreUser {
		getPointer = &Node[string]{getPointer, nil, n.val}
	}
	return getPointer, backGetPointer
}
func main() {
	// 구조 설명 -> {UserName, PreUser, val}
	// UserName은 지속적으로 nil를 통해 계속 다음 포인터를 잡아서 넣어줌
	// basic에서 UserName의 nil은 basic.UserName을 통해 값을 채워주는 방식

	// //기본적인 LinkedList 알아보기 위한 basic 버전
	// basic := &Node[string]{nil, nil, "민수"}
	// basic.UserName = &Node[string]{nil, basic, "철수"}
	// basic.UserName.UserName = &Node[string]{nil, basic.UserName, "영희"}

	// backBasic := basic.UserName.UserName

	// for n := backBasic; n != nil; n = n.PreUser {
	// 	fmt.Printf("node val: %d\n", n.val)

	// 함수를 활용해 자동 만들기
	users := []string{"민수", "철수", "영희"}
	next, pre := AutoGetList(users)
	fmt.Print(next, pre)
}
