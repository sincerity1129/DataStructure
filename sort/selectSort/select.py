import random

def set_selection_sort(arr):
    arr_size = len(arr)

    # 배열을 순회하며 최소값을 선택하여 정렬
    for i in range(arr_size):
        # 현재 위치부터 배열 끝까지 최소값을 찾음
        min_index = i
        for j in range(i+1, arr_size):
            if arr[j] < arr[min_index]:
                min_index = j

        # 현재 위치와 최소값 위치를 교환
        arr[i], arr[min_index] = arr[min_index], arr[i]

if __name__ == "__main__":
    original_list = [random.randint(0,100) for _ in range(100)]

    # 선택 정렬 수행
    set_selection_sort(original_list)

    # 정렬된 결과 출력
    print("정렬된 리스트:", original_list)