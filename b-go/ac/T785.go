package ac

import "fmt"

// 快速排序 https://www.acwing.com/problem/content/787/
// 1. (l + r) >> 1，中间值取左边，最后一步取到的是j，所以分区是(l, j)和(j + 1, r)
// 2. (l + r + 1) >> 1，中间值取右边，最后一步取到的是i，所以分区是(l, i - 1)和(i, r)
func quickSort(q []int, l, r int) {
    if l >= r {
        return
    }
    
    i, j, x := l - 1, r + 1, q[(l + r) >> 1]
    for i < j {
        for {
            i++
            if q[i] >= x {
                break
            }
        }
        for {
            j--
            if q[j] <= x {
                break
            }
        }
        if i < j{
            q[i], q[j] = q[j], q[i]
        }
    }
    
    quickSort(q, l, j)
    quickSort(q, j + 1, r)
}

func T785() {
	var n int
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	quickSort(q, 0, n - 1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}
