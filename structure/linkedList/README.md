# LinkedList
     데이터 요소의 집합을 저장하는 자료 구조로, 각 요소는 다음 요소를 가리키는 링크(참조)를 포함하는 집합ㄴ
## Array vs LinkedList
### Array 장단점
    장점: 랜덤 액서스(특정 인덱스 값 변경) 시에 O(1)으로 속도가 빠름
    -> 연속 메모리로 특정 배열을 만들면 한꺼번에 생성되어 그 공간 만큼 메모리가 생성
        따라서 첫 메모리에서 특정 주소까지 바로 연결하면 되는 구조라서 빠름
    
    단점: 삽입 삭제 시 느림 
    -> 연속 메모리 이기 때문에 삽입이나 삭제 시에 다시 메모리를 연속적으로 할당 해야 하므로
        이 경우에 다시 메모리 할당하는 과정에서 속도가 느려짐
### LinkedList
    장점: 삽입 삭제 시 빠름
    -> 불연속 메모리로 그때 그때 메모리를 할당하기 때문에 특정 액서스까지 접근 후
        삽입 삭제 영역의 포인터 부분만 변경해서 연결하면 되는 구조이기 때문에 빠름
    
    단점: 랜덤 액서스 시에 O(n)로 속도가 느림
    -> 불연속 메모리 이기 때문에 정확한 주소를 찾기 위해서 포인터를 통해 지점까지
        계속 이동을 통해서 접근해야함
## 코드 구성
### python
    python의 경우 별도 LinkedList를 지원하지 않기 때문에 별도의 code 만듦
### golang
    python 기반 알고리즘 만든 후 golang에 동일하게 convert
### java
    LinkedList를 지원하기 때문에 이를 활용하여 간략하게 쓰임 확인