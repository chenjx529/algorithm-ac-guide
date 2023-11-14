package ac;

import java.util.*;

// 双指针 最长连续不重复子序列 https://www.acwing.com/problem/content/801/
class T799 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] q = new int[n];
        for (int i = 0; i < n; i++) {
            q[i] = sc.nextInt();
        }
            
        HashMap<Integer, Integer> m = new HashMap<>();
        int res = 0;
        for (int i = 0, j = 0; i < n; i++) {
            m.put(q[i], m.getOrDefault(q[i], 0) + 1);
            while (j < i && m.get(q[i]) > 1) {
                m.put(q[j], m.get(q[j]) - 1);
                j++;
            }
            res = Math.max(res, i - j + 1);
        }
        System.out.print(res);
        sc.close();
    }
}