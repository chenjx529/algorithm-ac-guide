我的评价是：面试的时候还是用java吧，Go的标准库不是很全，比如heap要自己实现才可以用。

## 1. 输入输出

循环读入多行数据，不知道有几行：

```go
func main() {
    var a, b int
    for { 
        if _, err := fmt.Scan(&a, &b); err != io.EOF {
            fmt.Println(a + b)
        } else {
            break
        }
    }
}
```

第一行只有一个数字n，表示后面的行数，其他每行固定个数：

```go
func main() {
    var n, a, b int
    fmt.Scan(&n) // 
    for i := 0; i < n; i++ {
        fmt.Scan(&a, &b)
        fmt.Println(a + b)
    }
}

// 如果是正常的业务开发的话，是需要获取err并进行处理的。我们一般不用，因为笔试有时间限制的
func main() {
    var t, a, b int
    if _, err := fmt.Scan(&t); err != nil {
        // 处理错误
        return
    } 
    for t > 0 {
        if _, err := fmt.Scan(&a, &b); err != nil {
        	// 处理错误
        	return
    	} 
        fmt.Println(a + b)
        t--
    }
}
```

输入数据有多行，每行固定个数，读取到特殊数据(如0,0)时停止：

```go
func main() {
    var a, b int
    for {
        fmt.Scan(&a, &b)
        if a == 0 && b == 0 {
            break
        }
        fmt.Println(a + b)
    }
}
```

可以使用bufio包，读入一行数据再进行处理：

```go
func main() {
    input := bufio.NewScanner(os.Stdin) //创建并返回一个从os.Stdin读取数据的Scanner
    for input.Scan(){ 
        nums := strings.Split(input.Text(), " ") //分割字符串
        if nums[0] == "0" { // 判断是否结束
            break
        }
        
        sum := 0
        for _, i := range data {
			val, _ := strconv.Atoi(i) //将字符串转换为int
			sum += val
		}
        fmt.Println(sum)
    }
}
```

字符串排序：

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := strings.Split(input.Text(), ",") // 字符串分割
		sort.StringSlice.Sort(str)              // 字符串排序，这里借助了sort.StringSlice方法来实现
		fmt.Println(strings.Join(str, ","))     // 通过Join来连接， 也可以手动连接
	}
}
```

## 2. 排序

数字：

```go
num := [] int{1, 7, 5, 2, 6}
sort.Ints(num) //升序
fmt.Println(num)
sort.Sort(sort.Reverse(sort.IntSlice(num))) //降序
fmt.Println(num)
```

## 3. 数据结构

### 3.1 map

```go
func main() {
    // 创建一个空的map，键是字符串类型，值是整数类型
    myMap := make(map[string]int)
    // 添加键值对到map中
    myMap["apple"] = 10
    myMap["banana"] = 20
    myMap["orange"] = 15
    // 获取map中的值
    fmt.Println("apple:", myMap["apple"])     // 输出：apple: 10
    fmt.Println("banana:", myMap["banana"])   // 输出：banana: 20
    fmt.Println("orange:", myMap["orange"])   // 输出：orange: 15
    // 修改map中的值
    myMap["apple"] = 5
    fmt.Println("Modified apple:", myMap["apple"])  // 输出：Modified apple: 5
    // 删除map中的键值对
    delete(myMap, "orange")
    fmt.Println("After deleting orange:", myMap)   // 输出：After deleting orange: map[apple:5 banana:20]
}
```

### 3.2 set

```go
func main() {
    type exists struct{}
    // 创建一个空的map，键是字符串类型，值是空
    myMap := make(map[string]exists)
    // 添加值到set中
    myMap["apple"] = exists{}
    myMap["banana"] = exists{}
    myMap["orange"] = exists{}
    // 获取map中的值
    _, ok := myMap["apple"]
    fmt.Println(ok)
    // 删除set中的值
    delete(myMap, "orange")
    fmt.Println("After deleting orange:", myMap)
}
```

### 3.3 list

`container/list`包提供了标准库中的双向链表（Doubly Linked List）实现。

```go
func main() {
	// 创建一个双向链表
	myList := list.New()
	// 在链表尾部添加元素
	myList.PushBack(10)
	myList.PushBack(20)
	myList.PushBack(30)
	// 在链表头部添加元素
	myList.PushFront(5)
	// 遍历链表并打印元素
	fmt.Println("Elements in the list:")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	// 删除链表头部和尾部的元素
	myList.Remove(myList.Front())
	myList.Remove(myList.Back())
	// 遍历链表并打印元素
	fmt.Println("Elements after removing:")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
```

### 3.4 heap

`container/heap`包提供了堆的实现，用于构建优先队列。

```go
// 定义一个整数切片类型
type IntHeap []int
// 实现heap.Interface接口的Len方法
func (h IntHeap) Len() int {
	return len(h)
}
// 实现heap.Interface接口的Less方法
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
// 实现heap.Interface接口的Swap方法
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
// 实现heap.Interface接口的Push方法
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
// 实现heap.Interface接口的Pop方法
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	// 创建一个空的最小堆
	minHeap := &IntHeap{}
	// 将元素添加到最小堆
	heap.Push(minHeap, 10)
	heap.Push(minHeap, 20)
	heap.Push(minHeap, 5)
	// 获取最小堆的最小值（优先级最高的元素）
	fmt.Println("Minimum value:", (*minHeap)[0])  // 输出：Minimum value: 5
	// 弹出最小堆的最小值
	fmt.Println("Popped value:", heap.Pop(minHeap))  // 输出：Popped value: 5
}
```

