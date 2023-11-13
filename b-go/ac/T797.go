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
func insert(b []int, l, r, x int) {
    b[l] += x
    b[r + 1] -= x
}

func T797() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    t, _ := strconv.Atoi(scanner.Text())
    b := make([]int, n + 2)
    
    for i := 1; i <= n; i++ {
        scanner.Scan()
        x, _ := strconv.Atoi(scanner.Text())
        insert(b, i, i, x)
    }

    for i := 1; i <= t; i++ {
        scanner.Scan()
        l, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        r, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        x, _ := strconv.Atoi(scanner.Text())
        insert(b, l, r, x)
    }
    for i := 1; i <= n; i++ {
        b[i] += b[i - 1]
        fmt.Printf("%d ", b[i])
    }
}