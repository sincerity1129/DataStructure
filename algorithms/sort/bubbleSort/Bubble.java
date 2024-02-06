package algorithms.sort.bubbleSort;

import java.util.Random;
import java.util.Arrays;
public class Bubble {

    public static void setBubble(int[] arr) {
        int arrLength = arr.length;
        
        for (int i=0; i<arrLength; i++){
            for (int j=0; j<arrLength-i-1; j++){
                if (arr[j]>arr[j+1]){
                    int tmp = arr[j];
                    arr[j] = arr[j+1];
                    arr[j+1] = tmp;
                }
            }
        }
    }

 public static void main(String[] args) {
    int listLength = 30; 
    int[] orginalList = new int[listLength];
    Random random = new Random();

    for (int i=0; i<orginalList.length; i++){
        int randNum = random.nextInt(100);
        orginalList[i] = randNum;
    }
    setBubble(orginalList);
    System.out.println(Arrays.toString(orginalList));

 }   
}
