package ac;

import java.util.*;

// 二维差分 https://www.acwing.com/problem/content/800/
class T798 {
    static void insert(int[][] b, int x1, int y1, int x2, int y2, int c) {
        b[x1][y1] += c;
        b[x1][y2 + 1] -= c;
        b[x2 + 1][y1] -= c;
        b[x2 + 1][y2 + 1] += c;
    }
    
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int m = scanner.nextInt();
        int q = scanner.nextInt();
        
        // 初始化差分数组
        int[][] b = new int[1005][1005];
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                int c = scanner.nextInt();
                insert(b, i, j, i, j, c);
            }
        }
        
        // 差分
        for (int i = 1; i <= q; i++) {
            int x1 = scanner.nextInt();
            int y1 = scanner.nextInt();
            int x2 = scanner.nextInt();
            int y2 = scanner.nextInt();
            int c = scanner.nextInt();
            insert(b, x1, y1, x2, y2, c);
        }
        
        // 前缀和
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                b[i][j] += b[i - 1][j] + b[i][j - 1] - b[i - 1][j - 1];
                System.out.print(b[i][j] + " ");
            }
            System.out.print("\n");
        }

        scanner.close();
    }
}