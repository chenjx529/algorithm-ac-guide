package ac;

import java.util.PriorityQueue;
import java.util.Scanner;

// 堆排序 https://www.acwing.com/problem/content/840/
public class T838 {
        public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int m = sc.nextInt();
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();
        for (int i = 0; i < n; i++) {
            minHeap.add(sc.nextInt());
        }
        for (int i = 0; i < m; i++) {
            System.out.print(minHeap.poll() + " ");
        }
        sc.close();
    }
}
