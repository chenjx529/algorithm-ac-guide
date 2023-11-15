## 1. 导论

出现映射的问题，想到哈希表。

出现数组分块的问题，考虑单调栈、dp。

区间最大值问题，单调队列。

左右两边最大最小值问题，单调栈。

单调的，考虑二分，双指针。

二维数组求和问题，考虑前缀和。

数字都在 1 到 n 之间，查找一个有范围的整数，考虑二分查找。

Integer和int的区别，一个是对象，默认值是null；一个是基本类型，默认值是0。

```java
import java.util.Scanner;

public class Main {
    static int[] q = new int[100005];

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        for (int i = 0; i < n; i++)
            q[i] = sc.nextInt();

        System.out.println();

        // 循环输入
        while (sc.hasNext()) {
            int x = sc.nextInt();
        }
    }
}
```

## 2. 数组

初始化：

```java
int[] arr1 = new int[3]; 
int[] arr2 = new int[]{1,2,3,4,5};
int[] arr3 = {1,2,3,4,5};
```

长度：array.length

数组判断：array.getClass().isArray()

传参：

```java
public static int[] printArray(int[] arr) {
    for (int j : arr) {
        System.out.println(j);
    }
    return new int[]{1, 3, 5, 7, 9};
}
```

二维数组：

```java
int[][] datas = new int[3][4]; 
int[][] m = {{1,2,3,4},{4,5,6,7},{7,8,9,10}};
```

使用clone()：

+ 一维数组：深克隆，重新分配空间，并将元素复制过去
+ 二维数组：浅克隆，传递引用。

判空：

```java
// 一维数组
if (array == null || array.length == 0) return false;

// 二维数组
if (matrix == null) return false;
if (matrix.length == 0) return false;
if (matrix[0].length == 0) return false;
```

转化为ArrayList：

```
List list = Arrays.asList(arr);
```

翻转：

```java
public int[] reverse(int[] array) {
    int n = array.length;
    for (int i = 0; i < n / 2; i++) {
        int tmp = array[i];
        array[i] = array[n - 1 - i];
        array[n - 1 - i] = tmp;
    }
    return array;
}
```

数组排序：

```java
// 对数组a的下标从fromIndex到toIndex-1的元素排序，注意：下标为toIndex的元素不参与排序
Arrays.sort(Object[] a, fromIndex, toIndex)

// 想要使用自定义的Comparator，必须使用包装类
// Comparator的原理：对于(a, b)，返回负数，表示a<b；返回0，表示a==b；返回整数，表示a>b。
Arrays.sort(Integer[] a, (a, b) -> a - b) // 顺序
```

数组去重：推荐使用ArrayList

二分查找（排序好的数组，返回左边界的值，找不到返回小于0的数）：

```java
Array.binarySearch(Object[] a, Object key)
```

交换两个数：

```java
static int[] q = new int[100005];
public static void swap(int i, int j) {
    int t = q[i];
    q[i] = q[j];
    q[j] = t;
}
```

## 3. ArrayList

+ add(obj)
+ get(index)
+ remove(index)
+ size()
+ contains(obj)
+ indexOf(obj) 
+ clear()

排序：

```java
Collections.sort(list)
```

二维数组：

```java
List<List<Integer>> resultList = new ArrayList<>();
resultList.add(flagList); // 二维数组赋值，flagList为一维数组
```

转换数组：

```java
List<Integer> temp = new ArrayList<>();
int[] result = new int[temp.size()]; 
for (int i = 0; i < temp.size(); i++) 
	result[i] = temp.get(i);
```

去重：

```java
HashSet<Integer> tmp = new HashSet<Integer>(arraylist);
arraylist = new ArrayList<Integer>(tmp);
```

## 4. LinkedList

LinkedList 也可以作为堆栈、队列来使用。

```java
public class LinkedList<E> extends AbstractSequentialList<E> implements List<E>, Deque<E>, Cloneable, java.io.Serializable
```

+ 数组：和ArrayList的api一样
+ 队列：addLast(E e)、addFirst(E e)、removeFirst()、removeLast()、getFirst()、getLast()
+ 堆栈：push()、pop()

排序：

```java
Collections.sort(list)
```

## 5. HashMap

+ put(K key, V value)
+ get(Object key)
+ remove(Object key)
+ containsKey(Object key)
+ getOrDefault(key, defaultValue)
+ clear()

遍历，Map不能直接使用迭代器进行遍历，需要先转成Set：

```java
// 方式一
for (String key : map.keySet()) {
    String value = map.get(key);
    System.out.println(key + ": " + value);
}

// 方式二：
for (Entry<String, String> entry : map.entrySet()) {
    String key = entry.getKey();
    String value = entry.getValue();  
    System.out.println(key + ": " + value);
}
```

## 6. HashSet

+ add(obj)
+ remove(objectx)
+ size()
+ contains(obj)
+ clear()

## 7. String

```java
// 字符串的创建
String str = "chen";

// 获取字符串的长度
str.length();

// 获取目标字符
str.charAt(index);

// 字符串比较，推荐使用equals
str.equals(str2);

// 字符串连接
"chen.".concat("jia-xin");

// 字符串分割，如果按照英文句点“.”进行切分，必须写"\\."（两个反斜杠）
String[] str = "chen.jia.xin".split("\\.");

// 字符串替换，不改变原来的字符串
"1234".replace("12", "346"); // 34634

// 返回子串，左闭右开 substring (int beginIndex, int endIndex)
"chen".substring(0, 2);// ch

// 字符串比较，按字典序比较两个字符串，返回的是第一个不一样字符的ASCII的差值
"chenc".compareTo("chena");// 2    

// 字符串查找，从坐标fromIndex开始查找，找到返回索，没有找到的话，返回-1  indexOf(str, fromIndex)
"chenjiaxinchen".indexOf("chen", 1);// 10

// 包含 contains(CharSequence chars)
"chenjiaxinchen".contains("chen");

// 消除前面和尾部的空格
string.trim();

// 空字符串 isEmpty()
"".isEmpty();
```
