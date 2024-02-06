import random

def set_bubble_sort(arr):
    arr_size = len(arr)
    # 모든 요소에 대해 반복
    for i in range(arr_size):
        # 마지막 i개의 요소는 이미 정렬되었으므로 제외
        for j in range(0, arr_size-i-1):
            # 인접한 두 요소를 비교하여 정렬
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]


if __name__ == "__main__":
    # 테스트용 리스트
    original_list = [random.randint(0,100) for _ in range(100)]
    # 버블 정렬 수행
    set_bubble_sort(original_list)

    # 정렬된 결과 출력
    print("정렬된 배열:", original_list)