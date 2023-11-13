package ac;

import java.util.Scanner;

/**
 * 归并排序 https://www.acwing.com/problem/content/789/
 */
class T787 {
    public static void merge_sort(int[] q, int l, int r) {
        if (l >= r) return;
        
        int mid = (r + l) >> 1;
        merge_sort(q, l, mid);
        merge_sort(q, mid + 1, r);

        int k = 0, i = l, j = mid + 1;
        int[] t = new int[r - l + 1];
        while (i <= mid && j <= r) {
            if (q[i] <= q[j]) t[k++] = q[i++];
            else t[k++] = q[j++];
        }
        
        while (i <= mid) t[k++] = q[i++];
        while (j <= r) t[k++] = q[j++];
        
        for (i = l, j = 0; i <= r; i++, j++) {
            q[i] = t[j];
        }
    }
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        
        int n = sc.nextInt();
        int[] q = new int[n];
        for (int i = 0; i < n; i++) {
            q[i] = sc.nextInt();
        }

        merge_sort(q, 0, n - 1);

        for (int i = 0; i < n; i++) {
            System.out.print(q[i] + " ");
        }
        sc.close();
    }
}
