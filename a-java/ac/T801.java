package ac;

import java.util.*;

// 二进制中1的个数 https://www.acwing.com/problem/content/803/
class T801 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        for (int i = 0; i < n; i++) {
            int x = sc.nextInt();
            int res = 0;
            while (x != 0) {
                if ((x & 1) == 1) {
                    res++;
                }
                x >>= 1;
            }
            System.out.print(res + " ");
        }
        sc.close();
    }
}