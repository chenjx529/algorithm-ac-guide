package ac;
import java.util.Scanner;

/**
 * 归并排序
 */
public class T787 {
    
    static int[] q = new int[100005];
    static int[] t = new int[100005];
    
    public static void merge_sort(int l, int r) {
        if (l >= r) return;
        
        int mid = (r + l) >> 1;
        merge_sort(l, mid);
        merge_sort(mid + 1, r);

        int k = l, i = l, j = mid + 1;
        while (i <= mid && j <= r) {
            if (q[i] <= q[j]) t[k++] = q[i++];
            else t[k++] = q[j++];
        }
        
        while (i <= mid) t[k++] = q[i++];
        while (j <= r) t[k++] = q[j++];
        
        for (i = l, j = l; i <= r; i++, j++)
            q[i] = t[j];
    }
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        for (int i = 1; i <= n; i++)
            q[i] = sc.nextInt();

        merge_sort(0, n);

        for (int i = 1; i <= n; i++)
            System.out.print(q[i] + " ");
        sc.close();
    }
}
