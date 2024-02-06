import random

def heapify(arr, list_size, idx):
    largest = idx  # 현재 노드를 가장 큰 값으로 설정
    left_child = 2 * idx + 1
    right_child = 2 * idx + 2

    # 왼쪽 자식이 루트보다 크면 largest 갱신
    if left_child < list_size and arr[left_child] > arr[largest]:
        largest = left_child

    # 오른쪽 자식이 루트보다 크면 largest 갱신
    if right_child < list_size and arr[right_child] > arr[largest]:
        largest = right_child

    # largest가 루트가 아니라면 교환
    if largest != idx:
        arr[idx], arr[largest] = arr[largest], arr[idx]

        # 교환 후 재귀적으로 heapify 수행
        heapify(arr, list_size, largest)

def heap_sort(arr):
    list_size = len(arr)

    # 힙을 만듭니다. (최대 힙)
    for i in range(list_size // 2 - 1, -1, -1):
        heapify(arr, list_size, i)
        
    # 루트 노드를 제거하고 힙을 다시 정렬합니다.
    for i in range(list_size - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]  # 루트 노드와 마지막 노드를 교환
        heapify(arr, i, 0)

if __name__ == "__main__":
    # 예시
    original_list = [random.randint(0,100) for _ in range(100)]
    heap_sort(original_list)
    print("힙 정렬 결과:", original_list)