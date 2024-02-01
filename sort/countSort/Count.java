package sort.countSort;

import java.util.Arrays;
import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class Count {

    public static int[] SortCount(int[] arr) {
        int maxValue = arr[0];
        for (int value : arr) {
            if (maxValue<value) {
                maxValue = value;
            }
        }

        int[] defaultCount = new int[maxValue+1];
        for (int value : arr) {
            defaultCount[value] += 1;
        }

        for (int i=1; i<defaultCount.length; i++){
            defaultCount[i] += defaultCount[i-1];
        }

        int[] sortedCount = new int[arr.length];
        for (int value: arr) {
            sortedCount[defaultCount[value]-1] = value;
            defaultCount[value] -=1;
        }
        return sortedCount;
    }

    public static void main(String[] args) {
        int listLength = 30;
        int[] originalList = new int[listLength];
        Random random = new Random();

        for (int i=0; i<listLength; i++) {
            int randomNumber = random.nextInt(100) + 1;
            originalList[i] = randomNumber;
        }

        int[] result = SortCount(originalList);
        System.out.println(Arrays.toString(originalList));
        System.out.println(Arrays.toString(result));
    }
    
}
