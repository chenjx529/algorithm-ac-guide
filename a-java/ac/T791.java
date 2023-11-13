package ac;

import java.util.Scanner;

/**
 * 二分 https://www.acwing.com/problem/content/791/
 */
class T791 {

    // 二分左边界，条件是 q[mid] >= x，当 q[mid] == x，mid就是左边界，最大值最小的问题
    static int bsearch_left(int[] q, int x, int l, int r) {
        while (l < r) {
            int mid = (l + r) >> 1; // 取左边的数
            if (q[mid] >= x) r = mid;
            else l = mid + 1;
        }
        return l;
    }

    // 二分右边界，条件是 q[mid] <= x, 当 q[mid] == x，mid就是右边界，最小值最大的问题
    static int bsearch_right(int[] q, int x, int l, int r) {
        while (l < r) {
            int mid = (l + r + 1) >> 1;  // 取右边的数
            if (q[mid] <= x) l = mid;
            else r = mid - 1;
        }
        return r;
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int t = sc.nextInt();
        int[] q = new int[n];
        for (int i = 0; i < n; i++)
            q[i] = sc.nextInt();

        for (int i = 1; i <= t; i++){
            int x = sc.nextInt();
            int l = bsearch_left(q, x, 0, n - 1);
            int r = bsearch_right(q, x, 0, n - 1);
            if (q[l] == x) System.out.println(l + " " + r);
            else System.out.println("-1 -1");
        }
        sc.close();
    }
}
