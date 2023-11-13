package ac

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

// 二维差分 https://www.acwing.com/problem/content/800/
func insert798(b [][]int, x1, y1, x2, y2, c int) {
    b[x1][y1] += c
    b[x2 + 1][y1] -= c // 超出范围的地方减去
    b[x1][y2 + 1] -= c // 超出范围的地方减去
    b[x2 + 1][y2 + 1] += c // 减多了就加上
}

func T798() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    
    // n m q
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    m, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    q, _ := strconv.Atoi(scanner.Text())
    
    // b
    b := make([][]int, n + 2)
    for i := 0; i < n + 2; i++ {
        b[i] = make([]int, m + 2)
    }
    
    // 差分初始化
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            scanner.Scan()
            c, _ := strconv.Atoi(scanner.Text())
            insert798(b, i, j, i, j, c)
        }
    }
    
    // 差分
    for i := 1; i <= q; i++ {
        scanner.Scan()
        x1, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        y1, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        x2, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        y2, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        c, _ := strconv.Atoi(scanner.Text())
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