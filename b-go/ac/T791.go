package ac

import "fmt"

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

func T791() {
	var n, t, x int
	fmt.Scan(&n, &t)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&q[i])
	}

	for i := 0; i < t; i++ {
		fmt.Scan(&x)
		l := bsearch_l(q, x, 0, n - 1)
		r := bsearch_r(q, x, 0, n - 1)
		if q[l] == x {
			fmt.Printf("%d %d\n", l, r)
		} else {
			fmt.Println("-1 -1")
		}
	}
}