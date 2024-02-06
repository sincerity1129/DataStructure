import random

def shell_sort(arr, gap):
    arr_size = len(arr)
    divide = gap
    
    while gap > 0:
        for i in range(gap, arr_size):
            temp = arr[i]
            j = i
            while j >= gap and arr[j - gap] > temp:
                arr[j] = arr[j - gap]
                j -= gap
            arr[j] = temp
        gap //= divide

if __name__ == "__main__":
    original_list = [random.randint(0,100) for _ in range(100)]
    shell_sort(original_list, 13)
    print("정렬된 배열:", original_list)