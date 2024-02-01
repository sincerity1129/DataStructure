import random
from collections import defaultdict
def quick_sort(arr, standard):
    if len(arr) <= 1:
        return arr
    
    if standard !=0 or len(arr)>standard:
        standard=0
        
    pivot = arr[standard]
    arr_sorted =defaultdict(list)
    for x in arr:
        if x < pivot:
            arr_sorted["left"].append(x)
        elif x == pivot:
            arr_sorted["middle"].append(x)
        elif x> pivot:
            arr_sorted["right"].append(x)
    return quick_sort(arr_sorted["left"],standard) + arr_sorted["middle"] + quick_sort(arr_sorted["right"],standard)

        
if __name__=="__main__":
    # 테스트
    original_list = [random.randint(1, 100) for _ in range(100)]
    sorted_list = quick_sort(original_list, 10)

    if sorted(original_list)==sorted_list:
        print("쿽 정렬이 정상적으로 완료되었습니다.")
    else:
        print("알고리즘을 확인해 주세요.")