package basic_algorithms.sort.selectSort;

import java.util.Random;
import java.util.Arrays;

public class Select {

public static void setSelect(int[] arr) {
    int arrLength = arr.length;

    for (int i=0; i<arrLength; i++) {
        int minIndex = i;

        for (int j=i+1; j<arrLength; j++) {
            if ( arr[j]<arr[minIndex]){
                minIndex = j;
            }
        }
        int tmp = arr[minIndex];
        arr[minIndex] = arr[i];
        arr[i] = tmp;
    } 
}

    public static void main(String[] args) {
        int listLength = 30;
        int[] originalList = new int[listLength];
        Random random = new Random();


        for (int i=0; i<originalList.length; i++) {
            int randNum = random.nextInt(100);
            originalList[i] = randNum;
        }

        setSelect(originalList);
        System.out.println(Arrays.toString(originalList));
    }
    
}
