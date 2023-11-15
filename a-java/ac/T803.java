package ac;

import java.util.*;

// 区间合并 https://www.acwing.com/problem/content/805/
class T803 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        Integer[][] q = new Integer[n][2];
        for (int i = 0; i < n; i++) {
            q[i][0] = sc.nextInt();
            q[i][1] = sc.nextInt();
        }
        Arrays.sort(q, (a, b) -> a[0] - b[0]);

        int st = Integer.MIN_VALUE, ed = Integer.MIN_VALUE;
        int res = 0;
        for (int i = 0; i < n; i++) {
            if (ed < q[i][0]) {
                if (st != Integer.MIN_VALUE) {
                    res++;
                }
                st = q[i][0];
                ed = q[i][1];
            } else {
                ed = Math.max(ed, q[i][1]);
            }
        }

        if (st != Integer.MIN_VALUE) {
            res++;
        } 

        System.out.println(res);
        sc.close();
    }
}