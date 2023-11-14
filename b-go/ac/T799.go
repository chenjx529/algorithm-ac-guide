package ac

import "fmt"

// 双指针 最长连续不重复子序列 https://www.acwing.com/problem/content/801/
func T790() {
    var n int
    fmt.Scan(&n)
    q := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&q[i])
    }
    m := make(map[int]int)
    res := 0
    for i, j := 0, 0; i < n; i++ {
        m[q[i]]++
        for j < i && m[q[i]] > 1 {
            m[q[j]]--
            j++
        }
        // math.Max()方法只能用于float64类型，需要自己实现
        if res < i - j + 1 {
            res = i - j + 1
        }
    }
    fmt.Print(res)
}