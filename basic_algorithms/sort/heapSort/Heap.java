package basic_algorithms.sort.heapSort;
import java.util.Random;
import java.util.Arrays;

public class Heap {

    public static void heapProcess(int[] arr, int arrLength, int idx) {
        int largest = idx;

        int leftNode = idx*2+1;
        int rightNode = idx*2+2;

        if (leftNode<arrLength && arr[leftNode]>arr[largest]) {
            largest = leftNode;
        }
        if (rightNode<arrLength && arr[rightNode]>arr[largest]) {
            largest = rightNode;
        }

        if (largest!=idx) {
            int tmp = arr[largest];
            arr[largest] = arr[idx];
            arr[idx] = tmp;

            heapProcess(arr, arrLength, largest);
        }
    }

    public static void setHeap(int[] arr) {
        int arrLength = arr.length;

        for (int i=(arrLength/2)-1; 0<=i; i--) {
            heapProcess(arr, arrLength, i);
        }

        for (int i=arrLength-1; 0<i; i--) {
            int tmp = arr[0];
            arr[0] = arr[i];
            arr[i] = tmp;

            heapProcess(arr, i, 0);
        }
    }
    public static void main(String[] args) {
        int listLength = 100;
        int[] originalList = new int[listLength];
        Random random = new Random();

        for (int i=0; i<originalList.length; i++) {
            int randNum = random.nextInt(100);
            originalList[i] = randNum;
        }

        setHeap(originalList);
        System.out.println(Arrays.toString(originalList));
    }
}
