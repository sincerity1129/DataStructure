# SQuick sorted
### process
    첫째: pivot를 기준으로 하여 왼쪽에 pivot보다 작은 수 오른쪽에 큰수를 가진 배열 생성
    둘째: 재귀함수를 이용해서 작은 수만 모여 있는 배열을 동일한 방식으로 배열 값이 1일이 될 때까지 진행
    셋째: 큰 수도 둘째와 동일한 방식으로 진행
### 장단점
    장점: sort 정렬 중에 보편적으로 사용되며, 정렬 속도가 빠름
    -> O(nlogn)의 속도를 가짐
    
    단점: pivot 기준에 따라 최악의 속도를 가질 수 있음
    -> pivot 기준을 index=0으로 두었을 때 나머지 배열이 역순으로 정렬되어 있을 경우
        O(n^2)만큼 실행해야되기 때문에 속도 느려질 수 있음
### 코드 구성
#### python
    python 기준 코드
#### golang
    python to golang convert
#### java
    python to java convert
    array를 쓰지 않고 List 사용한 이유는 pivot 기준으로 값이 가변적이어서 고정된 array 쓸 수 없음