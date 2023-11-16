面试的时候用java，Go的标准库不是很全，比如list和heap要自己实现才可以用。

最小值: math.MinInt32

## 1. 输入输出

普通读入：

```go
var n int
fmt.Scan(&n)
```

缓冲Reader读入一个数据:

```go
in := bufio.NewReader(os.Stdin)
var n int
fmt.Fscan(in, &n)
```

缓冲Scanner读入一个数据:

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(bufio.ScanWords)
scanner.Scan()
n, _ := strconv.Atoi(scanner.Text())
```

缓冲Scanner读入一行数据：

```go
func main() {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan(){ 
        line := strings.Split(input.Text(), " ")
        if nums[0] == "0" { 
            break
        }
        for _, i := range line {
			val, _ := strconv.Atoi(i) //将字符串转换为int
		}
    }
}
```

循环读入多行数据，未知行数：

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

## 2. 排序

```go
tmp := []int{1, 7, 5, 2, 6}
sort.Slice(tmp, func(i, j int) bool {  
    return tmp[i] < tmp[j]
})
```

## 3. 数据结构

### 3.1 slice

```go
切片：var arr []int 或者 arr := make([]int, n)

数组：var arr [n]int

数组变切片：arr[:]

有长度的就是数组，没有长度的就是切片。

二维数组: var arr [n][m]int

二维切片: 
arr := make([][]int, n)
for i := 0; i < n; i++ {
    arr[i] = make([]int, m)
}
```

### 3.2 map

```go
func main() {
    myMap := make(map[string]int)

    // 添加key-value
    myMap["apple"]++
    myMap["banana"] = 20
    myMap["orange"] = 15

    // 获取value，两种形式都行
    v := myMap["apple"]
    v, ok := myMap["apple"]

    // 修改value
    myMap["apple"] = 5
    myMap["apple"]--
    fmt.Println(myMap["apple"]) // 4
    
    // 删除key-value
    delete(myMap, "orange")
    fmt.Println(myMap)   // 输出：map[apple:4 banana:20]
}
```

### 3.3 set

```go
func main() {
    type exists struct{}
    myMap := make(map[string]exists) // 空的map

    // 添加
    myMap["apple"] = exists{}
    myMap["banana"] = exists{}
    myMap["orange"] = exists{}
    
    // 查询
    _, ok := myMap["apple"]
    
    // 删除
    delete(myMap, "orange")
    fmt.Println("After deleting orange:", myMap)
}
```
