package ac;

import java.util.Scanner;

/**
 * 快速排序 https://www.acwing.com/problem/content/787/
 * 1. (l + r) >> 1，中间值取左边，最后一步取到的是j，所以分区是(l, j)和(j + 1, r)
 * 2. (l + r + 1) >> 1，中间值取右边，最后一步取到的是i，所以分区是(l, i - 1)和(i, r)
 */
public class T785 {

    static int[] q = new int[100005];
    
    public static void swap(int i, int j) {
        int t = q[i];
        q[i] = q[j];
        q[j] = t;
    }
    
    public static void quick_sort(int l, int r) {
        if (l >= r) return;
        
        int i = l - 1;
        int j = r + 1;
        int x = q[l + r >> 1];  // 如果是Go语言的话，需要写成(l + r) >> 1
        
        while (i < j) {
            while (q[++i] < x);  // 这里必须严格小于
            while (q[--j] > x);  // 这里必须严格大于
            if (i < j) swap(i, j);
        }
        
        quick_sort(l, j);
        quick_sort(j + 1, r);
    }
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        for (int i = 0; i < n; i++) 
            q[i] = sc.nextInt();
        
        quick_sort(0, n - 1);
        
        for (int i = 0; i < n; i++) 
            System.out.print(q[i] + " ");
            sc.close();
    }
}
