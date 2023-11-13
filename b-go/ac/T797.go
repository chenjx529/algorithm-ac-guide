package ac

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

// 差分 https://www.acwing.com/problem/content/799/
// a[i] = b[1] + b[2] + ... + b[i]
// a[i] ~ a[j] 所有的数都加上 x：b[i] += x, b[j + 1] -= x
func insert797(b []int, l, r, x int) {
    b[l] += x
    b[r + 1] -= x
}

func T797() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    
    // n t
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    t, _ := strconv.Atoi(scanner.Text())

    // b
    b := make([]int, n + 2)
    
    // 初始化差分
    for i := 1; i <= n; i++ {
        scanner.Scan()
        x, _ := strconv.Atoi(scanner.Text())
        insert797(b, i, i, x)
    }

    // 差分
    for i := 1; i <= t; i++ {
        scanner.Scan()
        l, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        r, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        x, _ := strconv.Atoi(scanner.Text())
        insert797(b, l, r, x)
    }

    // 前缀和
    for i := 1; i <= n; i++ {
        b[i] += b[i - 1]
        fmt.Printf("%d ", b[i])
    }
}