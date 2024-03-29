# Count sorted
## process
    첫번째: 배열 내에서 max 값을 기준으로 빈 배열 생성
    둘째: 생성된 배열을 기준으로 0~max까지 개수 카운팅
    셋째: 배열의 개수만큼 간격을 가진 배열로 다시 구조화
    넷째: 배열을 위해 정렬하고자 하는 배열 길이 만큼 빈 배열 생성
    다섯째: 간격을 가진 배열과 정렬 배열을 조합하여 정렬
## 장단점
    장점: 범위가 제한되어 있는 정수형 데이터에서 효율성이 높음
    -> O(n)+k의 속도를 가지기 때문에 제한된 상황에선 빠름, 단 k가 무한대 일 경우에는 속도가 한정 없이 늘어나서 비효율적
    
    단점: 메모리 사용량이 증가하며, 정수 이외의 실수나 문자열에서 사용하기 어려움
    -> 세번째 과정에서 max 값 기준으로 구조화하는 과정에서 max 값이 크면 메모리 사용량이 증가

## 코드 구성
#### python
    python 기준 코드
#### golang
    python to golang
#### java
    python to java
    쿽정렬과 다르게 array 써도 값이 제한되어 있어 사용 가능(max 값 무한대 배제)