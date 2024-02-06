package algorithms.sort.mergeSort;

import java.util.Random;
import java.util.Arrays;

public class Merge {
    public static void setMerge(int[] arr) {
        int arrLength = arr.length;
        if (arrLength > 1) {
            
            int mid = arrLength / 2;
            
            int[] leftArr = new int[mid];
            int[] rightArr = new int[arrLength-mid];
                
            System.arraycopy(arr, 0, leftArr, 0, leftArr.length);
            System.arraycopy(arr, mid, rightArr, 0, rightArr.length);

            setMerge(leftArr);
            setMerge(rightArr);

            int i = 0;
            int j = 0;
            int k = 0;

            while (i<leftArr.length && j<rightArr.length) {
                if (leftArr[i] < rightArr[j]) {
                    arr[k] = leftArr[i];
                    i++;
                }else{
                    arr[k] = rightArr[j];
                    j++;
                }
                k++;
            }

            while (i<leftArr.length) {
                arr[k] = leftArr[i];
                i++;
                k++;
            }

            while (j<rightArr.length) {
                arr[k] = rightArr[j];
                j++;
                k++;
            }
        }
    }

    public static void main(String[] args) {
        int listLength = 27;
        int[] originalList = new int[listLength];
        Random random = new Random();

        for (int i=0; i<originalList.length; i++) {
            int randNum = random.nextInt(100); 
            originalList[i] = randNum;
        }
        setMerge(originalList);
        System.out.println(Arrays.toString(originalList));
    }
}
