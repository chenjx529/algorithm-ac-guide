package ac

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

// 二维前缀和 https://www.acwing.com/problem/content/798/
// S[i][j] = S[i - 1][j] + S[i][j - 1] - S[i - 1][j - 1] + a[i][j]
// S[x2][y2] - S[x2][y1 - 1] - S[x1 - 1][y2] + S[x1 - 1][y1 - 1]
// 使用fmt会超时，建议使用bufio，使用ScanWords作为间隔符号，一个个读取
func T796() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    m, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    t, _ := strconv.Atoi(scanner.Text())

    var s [1005][1005]int

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            scanner.Scan()
            x, _ := strconv.Atoi(scanner.Text())
            s[i][j] = s[i - 1][j] + s[i][j - 1] - s[i - 1][j - 1] + x
        }
    }
    for i := 1; i <= t; i++ {
        scanner.Scan()
        x1, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        y1, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        x2, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        y2, _ := strconv.Atoi(scanner.Text())
        fmt.Println(s[x2][y2] - s[x2][y1 - 1] - s[x1 - 1][y2] + s[x1 - 1][y1 - 1])
    }
}