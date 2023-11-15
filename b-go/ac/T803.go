package ac

import (
    "sort"
    "fmt"
    "math"
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
    
    st := math.MinInt32
    ed := math.MinInt32
    res := 0
    for i := 0; i < n; i++ {
        if ed < q[i].Left {
            if st != math.MinInt32 {
                res++
            }
            st = q[i].Left
            ed = q[i].Right
        } else {
            ed = max(ed, q[i].Right)
        }
    }
    
    if st != math.MinInt32 {
        res++
    }
    
    fmt.Print(res)
}