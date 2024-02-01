package sort.linkedList;

import java.util.Arrays;
import java.util.LinkedList;
/*
 자바는 기본적으로 linkedList를 지원하고 있음
 따라서 별도의 포인터를 사용하여 구성 할 필요 없이 아래와 같이 간단하게 예제 실행하면 가능

 */
public class Linked {
    public static void main(String[] args) {
        LinkedList<String> linkedList = new LinkedList<>();

        // 전체 요소 LinkedList에 추가
        String[] datas = {"민수", "철수", "영희"};
        linkedList.addAll(Arrays.asList(datas));
        System.out.println("초기 LinkedList: "+linkedList);

        // first insert LinkedList에 추가
        linkedList.addFirst("first");
        // last insert LinkedList에 추가
        linkedList.addLast("end");
        // 특정 인덱스에 LinkedList에 추가
        linkedList.add(1, "firstNext");
        System.out.println("LinkedList 값 추가: "+linkedList);

        // 특정 인덱스 값 가져오기
        String linkedIndexParsing = linkedList.get(2);
        System.out.println("index 2: " + linkedIndexParsing);

        // 맨 앞의 요소 제거
        linkedList.removeFirst();
        System.out.println("first 값 삭제: " + linkedList);
        // 맨 뒤의 요소 제거
        linkedList.removeLast();
        System.out.println("last 값 삭제: " + linkedList);

    }
    
}