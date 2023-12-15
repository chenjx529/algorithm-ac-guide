package ac;

import java.util.*;

// 单调栈 https://www.acwing.com/problem/content/832/
class T830 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        
        int[] stk = new int[n];
        int tt = -1;
        for (int i = 0; i < n; i++) {
            int x = sc.nextInt();
            while (tt >= 0 && stk[tt] >= x) tt--;
            if (tt >= 0) System.out.print(stk[tt] + " ");
            else System.out.print(-1 + " ");
            stk[++tt] = x;
        }
        sc.close();
    }
}