package sort.shell;

import java.util.Arrays;
import java.util.Random;

public class Shell {

    public static void setShell(int[] arr, int gap) {
        int arrLength = arr.length;
        int divide = gap;

        while (gap>0) {
            for(int i=gap; i<arrLength; i++) {
                int tmp = arr[i];
                int j = i;
                while (gap<=j && arr[j-gap] > tmp) {
                    arr[j] = arr[j-gap];
                    j -= gap;
                } 
                arr[j] = tmp;
            }
            gap /=divide;
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

        setShell(originalList,2);
        System.out.println(Arrays.toString(originalList));
    }
}
