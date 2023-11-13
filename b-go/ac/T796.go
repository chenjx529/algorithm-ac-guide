package ac

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

// 二维前缀和 https://www.acwing.com/problem/content/798/
// S[i][j] = S[i - 1][j] + S[i][j - 1] - S[i - 1][j - 1] + a[i][j]
// S[x2][y2] - S[x2][y1 - 1] - S[x1 - 1][y2] + S[x1 - 1][y1 - 1]
// 使用fmt会超时，建议使用bufio
func T796() {
    var n, m, t, x, x1, y1, x2, y2 int
    fmt.Scan(&n, &m, &t)
    var s [1005][1005]int
    // s := make([][]int, n + 1)
    // for i := 1; i <= n; i++ {
    //     s[i] = make([]int, m + 1)
    // }
    scanner := bufio.NewScanner(os.Stdin)
    for i := 1; i <= n; i++ {
        scanner.Scan() // Scan 方法 该方法好比 iterator 中的 Next 方法
        line := strings.Split(scanner.Text(), " ")
        for j := 1; j <= m; j++ {
            // fmt.Scan(&x)
            x, _ = strconv.Atoi(line[j - 1])
            s[i][j] = s[i - 1][j] + s[i][j - 1] - s[i - 1][j - 1] + x
        }
    }
    for i := 1; i <= t; i++ {
        scanner.Scan()
        line := strings.Split(scanner.Text(), " ")
        // fmt.Scan(&x1, &y1, &x2, &y2)
        x1, _ = strconv.Atoi(line[0])
        y1, _ = strconv.Atoi(line[1])
        x2, _ = strconv.Atoi(line[2])
        y2, _ = strconv.Atoi(line[3])
        fmt.Println(s[x2][y2] - s[x2][y1 - 1] - s[x1 - 1][y2] + s[x1 - 1][y1 - 1])
    }
}