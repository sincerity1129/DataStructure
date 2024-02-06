# shell sorted
## process
    첫번째: gap의 간격을 기준으로 반복문을 수행하여 진행
    둘째: gap 값과 gap보다 뒷 열에 있는 값들을 순차적으로 gap까지 비교
    셋째: gap의 간격을 순차적으로 줄여나가면서 모든 값을 비교하여 정렬
## 장단점
    장점: 삽입정렬을 개선하여 효율적으로 동작
    -> 삽입정렬의 경우 인접한 사이에서만 정렬하지만 이를 효율적으로
        간격을 도입하여 멀리에 있는 대/소를 비교할 수 있음
    
    단점: 간격 선택에 따라 효율성이 달라짐

## 코드 구성
#### python
    python 기준 코드
#### golang
    python to golang
    golang의 경우 별도의 while문이 없기 때문에 for문을 활용하여 코드 작성
#### java
    python to java