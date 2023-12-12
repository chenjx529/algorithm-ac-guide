package ac;

import java.io.*;

// 单调队列 https://www.acwing.com/problem/content/156/
class T154{
    
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String[] line = reader.readLine().split(" ");
        int n = Integer.parseInt(line[0]);
        int k = Integer.parseInt(line[1]);
        int[] a = new int[n + 1];
        int[] q = new int[n + 1];
        line = reader.readLine().split(" ");
        for (int i = 0; i < n; i++) {
            a[i] = Integer.parseInt(line[i]);
        }
        
        int hh = 0, tt = -1;
        for (int i = 0; i < n; i++) {
            if (hh <= tt && q[hh] < i - k + 1) hh ++; // 移动对头，q ∈ [i - k + 1, i]
            while (hh <= tt && a[q[tt]] >= a[i]) tt--; // 移动队尾，注意等于号
            q[++tt] = i;
            if (i >= k - 1) {
                System.out.print(a[q[hh]] + " ");
            }
        }
        System.out.print("\n");
        
        hh = 0;
        tt = -1;
        for (int i = 0; i < n; i++) {
            if (hh <= tt && q[hh] < i - k + 1) hh ++; // 移动对头，q ∈ [i - k + 1, i]
            while (hh <= tt && a[q[tt]] <= a[i]) tt--; // 移动队尾，注意等于号
            q[++tt] = i;
            if (i >= k - 1) {
                System.out.print(a[q[hh]] + " ");
            }
        }
        
    }
}