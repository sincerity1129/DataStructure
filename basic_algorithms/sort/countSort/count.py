import random

def set_counting_sort(arr):   
    # 카운트 배열 초기화
    count = [0] * (max(arr) + 1)
    
    # 카운트 배열 채우기
    for num in arr:
        count[num] += 1
    
    # 누적 합 계산
    for i in range(1, len(count)):
        count[i] += count[i - 1]
    
    # 정렬된 결과를 저장할 배열 초기화
    sorted_arr = [0] * len(arr)
    
    # 입력 배열을 순회하면서 정렬된 배열에 요소 배치
    for num in arr:
        sorted_arr[count[num] - 1] = num
        count[num] -= 1
    
    return sorted_arr

if __name__ == "__main__":
    # 예제 사용
    original_list = [random.randint(0, 100) for _ in range(100)]
    sorted_list = set_counting_sort(original_list)

    print("입력 배열:", original_list)
    print("정렬된 배열:", sorted_list)