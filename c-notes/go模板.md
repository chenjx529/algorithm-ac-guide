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

### 快排

```go
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
```

### 归并

```go
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
```

### 二分

```go
// 二分左边界，条件是 q[mid] >= x，当 q[mid] == x，mid就是左边界，最大值最小的问题
func bsearch_l(q []int, x, l, r int) int {
	for l < r {
		mid := (l + r) >> 1
		if q[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// 二分右边界，条件是 q[mid] <= x, 当 q[mid] == x，mid就是右边界，最小值最大的问题
func bsearch_r(q []int, x, l, r int) int {
	for l < r {
		mid := (l + r + 1) >> 1
		if q[mid] <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return r
}
```

