package structure.queue_stack;

import java.util.Stack;
import java.util.Queue;
import java.util.LinkedList;

public class Basic {
    public static void main(String[] args) {

        Stack<String> stack = new Stack<>();

        stack.push("a");
        stack.push("b");
        stack.push("c");
        
        while (!stack.isEmpty()) {
            System.out.println("stack: "+ stack.pop()); 
        }

        System.out.println("---------------------------------");
        Queue<String> queue = new LinkedList<>();
        queue.offer("a");
        queue.offer("b");
        queue.offer("c");

        while (!queue.isEmpty()) {
            System.out.println("queue: " + queue.poll());
        }
    }
}