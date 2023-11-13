package ac;

import java.util.Scanner;

/**
 * 快速排序 https://www.acwing.com/problem/content/787/
 * 1. (l + r) >> 1，中间值取左边，最后一步取到的是j，所以分区是(l, j)和(j + 1, r)
 * 2. (l + r + 1) >> 1，中间值取右边，最后一步取到的是i，所以分区是(l, i - 1)和(i, r)
 */
class T785 {

    static void swap(int[] q, int i, int j){
        int t = q[i];
        q[i] = q[j];
        q[j] = t;
    }
    
    static void quick_sort(int[] q, int l, int r) {
        if (l >= r) return;
        
        int i = l - 1;
        int j = r + 1;
        int x = q[(l + r) >> 1];
        
        while (i < j) {
            while (q[++i] < x);
            while (q[--j] > x);
            if (i < j) swap(q, i, j);
        }
        
        quick_sort(q, l, j);
        quick_sort(q, j + 1, r);
    }
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] q = new int[n];
        for (int i = 0; i < n; i++) {
            q[i] = sc.nextInt();
        }
        quick_sort(q, 0, n - 1);
        for (int i = 0; i < n; i++) {
            System.out.print(q[i] + " ");
        }
        sc.close();
    }
}
