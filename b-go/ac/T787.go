package ac

import "fmt"

func merge_sort(q []int, tmp []int, l int, r int) {
    if l >= r {
        return
    }
    
    mid := (l + r) >> 1
    merge_sort(q, tmp, l, mid)
    merge_sort(q, tmp, mid + 1, r)
    
    k, i, j := l, l, mid + 1
    for i <= mid && j <= r {
        if q[i] <= q[j] {
            tmp[k] = q[i]; k++; i++
        } else {
            tmp[k] = q[j]; k++; j++
        }
    }
    
    for i <= mid {
        tmp[k] = q[i]; k++; i++
    }
    
    for j <= r {
        tmp[k] = q[j]; k++; j++
    }
    
    for i, j = l, l; i <= r; i, j = i+1, j+1 {
        q[i] = tmp[j]
    }
}

func T787() {
    var n int
    fmt.Scan(&n)
    
    q := make([]int, n)
    tmp := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&q[i])
    }
    
    merge_sort(q, tmp, 0, n - 1)
    
	for _, v := range q {
	    fmt.Printf("%d ", v)
	}
}