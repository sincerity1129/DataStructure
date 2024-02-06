package algorithms.sort.quickSort;

import java.util.Collections;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Random;

public class Quick {

    public static ArrayList<Integer> setQuickSort(ArrayList<Integer> arr, int standard) {
        if (arr.size() <= 1) {
            return arr;
        }

        if (standard !=0 || arr.size()>standard){
            standard=0;
        }

        int pivotIndex = arr.size() / 2;
        int pivot = arr.get(pivotIndex);

        Map<String, ArrayList<Integer>> arrSorted = new HashMap<>();
        arrSorted.put("left", new ArrayList<>());
        arrSorted.put("middle", new ArrayList<>());
        arrSorted.put("right", new ArrayList<>());

        for (int x : arr) {
            if (x < pivot) {
                arrSorted.get("left").add(x);
            } else if (x == pivot) {
                arrSorted.get("middle").add(x);
            } else if (x> pivot) {
                arrSorted.get("right").add(x);
            }
        }

        ArrayList<Integer> result = new ArrayList<>();
        result.addAll(setQuickSort(arrSorted.get("left"), standard));
        result.addAll(arrSorted.get("middle"));
        result.addAll(setQuickSort(arrSorted.get("right"), standard));

        return result;
    }
    
    public static void main(String[] args) {
        int arrayLength = 100;
        ArrayList<Integer> originalList = new ArrayList<>();
        Random random = new Random();

        for (int i = 0; i < arrayLength; i++) {
            int randomNumber = random.nextInt(100) + 1; // 1부터 100까지의 난수 생성
            originalList.add(randomNumber);
        }
        ArrayList<Integer> resultList = setQuickSort(originalList, 10);
        Collections.sort(originalList);
        if (originalList.equals(resultList)){
            System.out.println("쿽 정렬이 정상적으로 완료되었습니다.");
        }else{
            System.out.println("알고리즘을 확인해 주세요.");
        }
    }
}