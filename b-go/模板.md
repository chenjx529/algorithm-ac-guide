```go
package main

import (
    "fmt"
)

func main() {
	fmt.Println("hello word")
}
```

## 1. 基础算法

### 排序

```go
// 快排
func quickSort(q []int, l int, r int) {
	if l >= r {
		return
	}

	i, j, x := l-1, r+1, q[(l+r)>>1] // 这里必须加上括号，否则出现计算优先级的问题

	for i < j {
		for i++; q[i] < x; i++ {}
		for j--; q[j] > x; j-- {}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

// 归并
func merge_sort(q []int, tmp []int, l int, r int) {
    if l >= r {
        return
    }
    
    mid := (l + r) >> 1
    merge_sort(q, tmp, l, mid)
    merge_sort(q, tmp, mid + 1, r)
    
    k, i, j := 0, l, mid + 1
    for i <= mid && j <= r {
        if q[i] <= q[j] {
            tmp[k] = q[i]; k++; i++
        } else {
            tmp[k] = q[j]; k++; j++
            // res += mid - i + 1  // 一个有(mid - i + 1)个数和q[j]是逆序对关系
        }
    }
    
    for i <= mid {
        tmp[k] = q[i]; k++; i++
    }
    
    for j <= r {
        tmp[k] = q[j]; k++; j++
    }
    
    for i, j = l, 0; i <= r; i, j = i+1, j+1 {
        q[i] = tmp[j]
    }
}
```

### 二分

```go
// 左边界
func binary_l(q []int, k, l, r int) int {
    for l < r {
        mid := (l + r) >> 1
        if (k <= q[mid]) {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}

// 右边界
func binary_r(q []int, k, l, r int) int {
    for l < r {
        mid := (l + r + 1) >> 1
        if (k >= q[mid]) {
            l = mid
        } else {
            r = mid - 1
        }
    }
    return r
}
```

