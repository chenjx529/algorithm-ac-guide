package ac

import (
    "fmt"
    "bufio"
    "os"
)

// 二维差分 https://www.acwing.com/problem/content/800/
func insert798(b [][]int, x1, y1, x2, y2, c int) {
    b[x1][y1] += c
    b[x2 + 1][y1] -= c // 超出范围的地方减去
    b[x1][y2 + 1] -= c // 超出范围的地方减去
    b[x2 + 1][y2 + 1] += c // 减多了就加上
}

func T798() {
    in := bufio.NewReader(os.Stdin)
    
    // n m q
    var n, m, q int
    fmt.Fscan(in, &n, &m, &q)
    
    
    // b
    b := make([][]int, n + 2)
    for i := 0; i < n + 2; i++ {
        b[i] = make([]int, m + 2)
    }
    
    // 差分初始化
    var c int
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            fmt.Fscan(in, &c)
            insert798(b, i, j, i, j, c)
        }
    }
    
    // 差分
    var x1, y1, x2, y2 int
    for i := 1; i <= q; i++ {
        fmt.Fscan(in, &x1, &y1, &x2, &y2, &c)
        insert798(b, x1, y1, x2, y2, c)
    }
    
    // 前缀和
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            b[i][j] += b[i - 1][j] + b[i][j - 1] - b[i - 1][j - 1]
            fmt.Printf("%d ", b[i][j])
        }
        fmt.Printf("\n")
    }
}