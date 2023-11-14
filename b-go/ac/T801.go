package ac

import "fmt"

// 二进制中1的个数 https://www.acwing.com/problem/content/803/
func T801() {
    var n, x int
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        fmt.Scan(&x)
        res := 0
        
        for x != 0 {
            if (x & 1) == 1 {
                res++
            }
            x >>= 1
        }
        
        fmt.Printf("%d ", res)
    }
}