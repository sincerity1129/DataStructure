package structure.tree;

import java.util.Queue;
import java.util.LinkedList;

public class basic {

        String value;
        basic left, right;


    public basic(String value) {
        this.value = value;
        this.left = null;
        this.right = null;
    }

    public static void dfs(basic root){
        if (root == null) {
            return;
        }

        System.out.println(root.value + " ");
        dfs(root.left);
        dfs(root.right);
    }

    public static void bfs(basic root) {
        if (root == null) {
            return;
        }

        Queue<basic> queue = new LinkedList<>();
        queue.offer(root);

        while (!queue.isEmpty()) {
            basic node = queue.poll();
            System.out.print(node.value + " ");

            if (node.left != null) {
                queue.offer(node.left);
            }
            if (node.right != null) {
                queue.offer(node.right);
            }
        }

    }

    public static void main(String[] args) {
        basic root = new basic("1");
        root.left = new basic("2");
        root.right = new basic("3");
        root.left.left = new basic("4");
        root.left.right = new basic("5");
        root.right.left = new basic("6");
        root.right.right = new basic("7");

        dfs(root);
        System.out.println("-------------------");
        bfs(root);
    }    
}
