import random

def set_merge_sort(arr):
    arr_size = len(arr)
    if arr_size > 1:
        mid = arr_size // 2  # 중간 인덱스 계산
        left_arr = arr[:mid]  # 왼쪽 부분 배열
        right_arr = arr[mid:]  # 오른쪽 부분 배열

        # 왼쪽 부분 배열과 오른쪽 부분 배열을 재귀적으로 정렬
        set_merge_sort(left_arr)
        set_merge_sort(right_arr)

        i = j = k = 0   
            
        # 두 부분 배열을 합치기
        while i < len(left_arr) and j < len(right_arr):
            if left_arr[i] < right_arr[j]:
                arr[k] = left_arr[i]
                i += 1
            else:
                arr[k] = right_arr[j]
                j += 1
            k += 1

        # 남은 요소들 처리
        while i < len(left_arr):
            arr[k] = left_arr[i]
            i += 1
            k += 1

        while j < len(right_arr):
            arr[k] = right_arr[j]
            j += 1
            k += 1

if __name__ == "__main__":
    original_list = [random.randint(0,100) for _ in range(10)]

    # 병합 정렬 수행
    set_merge_sort(original_list)

    # 정렬된 결과 출력
    print("정렬된 리스트:", original_list)