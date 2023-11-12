package ac

import "fmt"

// 前缀和 https://www.acwing.com/problem/content/797/
// S[i] = S[i - 1] + a[i]，i要从1开始算
// a[l] + ... + a[r] = S[r] - S[l - 1]
func T795() {
    var n, t, x, l, r int
    fmt.Scan(&n, &t)
    q := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        fmt.Scan(&x)
        q[i] = q[i - 1] + x
    }
    for i := 1; i <= t; i++ {
        fmt.Scan(&l, &r)
        fmt.Println(q[r] - q[l - 1])
    }
}