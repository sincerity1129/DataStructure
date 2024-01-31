import numpy as np

class Quicksort:
    @classmethod
    def _partition(cls, arr, standard):
        if len(arr)<=1:
            return 0
        
        # 기준점
        pivot = arr[standard]
        prev=1
        last=len(arr)-1
                
        while True:
            while prev < len(arr) and arr[prev] <= pivot:
                prev += 1
            while last > 0 and arr[last] > pivot:
                last -= 1
            if prev >= last:
                arr[standard], arr[last] = arr[last], arr[standard]
                return last
            arr[prev], arr[last] = arr[last], arr[prev]
            
            
    @classmethod
    def set_quick_sort(cls, arr, standard=0):
        if len(arr)<=1:
            return
        
        idx = cls._partition(arr, standard)
        cls.set_quick_sort(arr[:idx])
        cls.set_quick_sort(arr[idx+1:])
        
        return arr
        
if __name__=="__main__":
    arr = np.random.randint(1, 101, 100)
    print(f"before: {arr}")
    arr = Quicksort.set_quick_sort(arr, 10)
    print(f"after: {arr}")