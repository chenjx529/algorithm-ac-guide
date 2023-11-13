package ac;

import java.util.*;

/**
 * 前缀和 https://www.acwing.com/problem/content/797/
 * S[i] = S[i - 1] + a[i]，i要从1开始算
 * a[l] + ... + a[r] = S[r] - S[l - 1]
 */
class T795 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int t = sc.nextInt();
        int[] s = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            int a = sc.nextInt();
            s[i] = s[i - 1] + a;
        }
        
        for (int i = 0; i < t; i++) {
            int l = sc.nextInt();
            int r = sc.nextInt();
            System.out.println(s[r] - s[l - 1]);
        }
        sc.close();
    }
}
