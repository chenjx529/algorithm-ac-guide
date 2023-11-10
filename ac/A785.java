package ac;
import java.util.Scanner;

public class A785 {
    static int[] q = new int[100005];
    
    public static void swap(int i, int j) {
        int t = q[i];
        q[i] = q[j];
        q[j] = t;
    }
    
    public static void quick(int l, int r) {
        if (l >= r) return;
        
        int i = l - 1, j = r + 1, x = q[l + r >> 1];
        
        
        while (i < j) {
            while (q[++i] < x);
            while (q[--j] > x);
            if (i < j) swap(i, j);
        }
        
        quick(l, j);
        quick(j + 1, r);
    }
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        for (int i = 0; i < n; i++) 
            q[i] = sc.nextInt();
        
        quick(0, n - 1);
        
        for (int i = 0; i < n; i++) 
            System.out.print(q[i] + " ");
    }
}
