package ac

import "fmt"

// 归并排序 https://www.acwing.com/problem/content/789/
func mergeSort(q []int, l, r int) {
    if l >= r {
        return
    }
    
    mid := (l + r) >> 1
    mergeSort(q, l, mid)
    mergeSort(q, mid + 1, r)
    
    i, j, t:= l, mid + 1, make([]int, 0)
    for i <= mid && j <= r {
        if q[i] <= q[j] {
            t = append(t, q[i])
            i++
        } else {
            t = append(t, q[j])
            j++
            // res += mid - i + 1  // 有(mid - i + 1)个数和q[j]是逆序对关系
        }
    }
    
    for i <= mid {
        t = append(t, q[i])
        i++
    }
    
    for j <= r {
        t = append(t, q[j])
        j++
    }
    
    copy(q[l: r + 1], t)
}

func T787() {
    var n int
    fmt.Scan(&n)
    q := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&q[i])
    }
    mergeSort(q, 0, n - 1)
	for _, v := range q {
	    fmt.Printf("%d ", v)
	}
}