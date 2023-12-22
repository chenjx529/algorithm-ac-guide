package ac;

import java.util.Scanner;

// 并查集  https://www.acwing.com/problem/content/838/
public class T836 {
    static int find(int[] p, int x) {
        if (p[x] != x) {
            p[x] = find(p, p[x]);
        }
        return p[x];
    }
    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int m = sc.nextInt();
        int[] p = new int[n + 1];
        for (int i = 1; i <= n; i++) p[i] = i;
        
        for (int i = 0; i < m; i++) {
            char c = sc.next().charAt(0); // 读取一个字符
            if (c == 'M') {
                int a = sc.nextInt();
                int b = sc.nextInt(); 
                p[find(p, a)] = find(p, b);
            } else {
                int a = sc.nextInt();
                int b = sc.nextInt(); 
                if (find(p, a) == find(p, b)) {
                    System.out.println("Yes");
                } else {
                    System.out.println("No");
                }
            }
        }
        sc.close();
    }
}
