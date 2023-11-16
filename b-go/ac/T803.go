package ac

import (
    "sort"
    "fmt"
    "os"
    "bufio"
)

type Pair struct {
    Left int
    Right int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 区间合并 https://www.acwing.com/problem/content/805/
func T803() {
    in := bufio.NewReader(os.Stdin)
    var n, l, r int
    fmt.Fscan(in, &n)
    
    q := make([]Pair, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &l, &r)
        q[i] = Pair{l, r}
    }
    
    sort.Slice(q, func(i, j int) bool {  
        return q[i].Left < q[j].Left
    })
    
    ed := q[0].Right
    res := 1
    for i := 1; i < n; i++ {
        if ed < q[i].Left {
            res++
            ed = q[i].Right
        } else {
            ed = max(ed, q[i].Right)
        }
    }
    fmt.Print(res)
}