package ac

import "fmt"

// 快速排序 https://www.acwing.com/problem/content/787/
// 1. (l + r) >> 1，中间值取左边，最后一步取到的是j，所以分区是(l, j)和(j + 1, r)
// 2. (l + r + 1) >> 1，中间值取右边，最后一步取到的是i，所以分区是(l, i - 1)和(i, r)
func quickSort(q []int, l, r int) {
	if l >= r { // 终止条件
		return
	}
	x := q[(l + r) >> 1] // 确定分界点
	i, j := l - 1, r + 1  // 两个指针，因为do while要先自增/自减
	for i < j {       // 每次迭代
		for { // do while 语法
			i++ // 交换后指针要移动，避免没必要的交换
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
		if i < j { // swap 两个元素
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j) // 递归处理左右两段
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
