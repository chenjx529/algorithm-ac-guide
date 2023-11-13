package ac;

import java.util.*;


/**
 * 差分 https://www.acwing.com/problem/content/799/
 * a[i] = b[1] + b[2] + ... + b[i]
 * a[i] ~ a[j] 所有的数都加上 x：b[i] += x, b[j + 1] -= x
 */
class T797{
    static void insert(int[] b, int l, int r, int x) {
        b[l] += x;
        b[r + 1] -= x;
    }
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int t = sc.nextInt();
        int[] b = new int[n + 2];
        for (int i = 1; i <= n; i++) {
            int x = sc.nextInt();
            insert(b, i, i, x);
        }
        for (int i = 1; i <= t; i++) {
            int l = sc.nextInt();
            int r = sc.nextInt();
            int x = sc.nextInt();
            insert(b, l, r, x);
        }
        for (int i = 1; i <= n; i++) {
            b[i] += b[i - 1];
            System.out.print(b[i] + " ");
        }
        sc.close();
    }
}