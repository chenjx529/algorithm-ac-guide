## 1. 导论

出现映射的问题，想到哈希表。

出现数组分块的问题，考虑单调栈、dp。

区间最大值问题，单调队列。

左右两边最大最小值问题，单调栈。

单调的，考虑二分，双指针。

二维数组求和问题，考虑前缀和。

数字都在 1 到 n 之间，查找一个有范围的整数，考虑二分查找。

Integer和int的区别，一个是对象，默认值是null；一个是基本类型，默认值是0。

## 2. 输入

通用Scanner：

```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] q = new int[n];
        for (int i = 0; i < n; i++) {
            q[i] = sc.nextInt();
        }
    }
}
```

缓冲流BufferedReader：

```java
import java.io.*;
public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(reader.readLine());
        int[][] a = new int[n][2];
        String[] line;
        for (int i = 0; i < n; i++) {
            line = reader.readLine().split(" ");
            a[i][0] = Integer.parseInt(line[0]);
            a[i][1] = Integer.parseInt(line[1]);
        }
    }
}
```

## 3. 数组

数组是new出来的，数组是对象，数组作为参数传入函数后经过操作底层数据会发生改变的。

初始化：

```java
// 一维数组
int[] arr1 = new int[3]; 
int[] arr2 = new int[]{1,2,3,4,5};
int[] arr3 = {1,2,3,4,5};

// 二维数组
int[][] arr4 = new int[3][4]; 
int[][] arr5 = {{1,2,3,4},{4,5,6,7},{7,8,9,10}};
```

长度：array.length

判空：

```java
// 一维数组
if (arr == null || arr.length == 0) return false;

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
public void reverse(int[] array) {
    int n = array.length;
    for (int i = 0; i < n / 2; i++) {
        int tmp = array[i];
        array[i] = array[n - 1 - i];
        array[n - 1 - i] = tmp;
    }
}
```

数组排序：

```java
// 对数组a从fromIndex到toIndex-1排序，toIndex不参与排序
Arrays.sort(Object[] a, fromIndex, toIndex)

// 想要使用自定义的Comparator，必须使用包装类
// 原理：返回负数，表示a<b；返回0，表示a==b；返回正数，表示a>b。
Arrays.sort(Integer[] a, (a, b) -> a - b) // 顺序
```

数组去重：使用ArrayList

二分查找（排序好的数组，返回左边界的值，找不到返回负数）：

```java
Array.binarySearch(Object[] a, Object key)
```

## 4. ArrayList

+ add(obj)
+ get(index)
+ remove(index)
+ size()
+ contains(obj)
+ indexOf(obj) 
+ clear()

排序（和Arrays.sort()类似）：

```java
Collections.sort(list)
```

二维数组：

```java
List<List<Integer>> resultList = new ArrayList<>();
resultList.add(flagList); // 二维数组赋值，flagList为一维数组
```

去重：

```java
HashSet<Integer> tmp = new HashSet<Integer>(arraylist);
arraylist = new ArrayList<Integer>(tmp);
```

## 5. LinkedList

LinkedList 可以作为堆栈、队列来使用。

+ 数组：和ArrayList的api一样
+ 队列：addLast(E e)、addFirst(E e)、removeFirst()、removeLast()、getFirst()、getLast()
+ 堆栈：push()、pop()

## 6. HashMap

+ put(K key, V value)
+ get(Object key)
+ getOrDefault(key, defaultValue)
+ remove(Object key)
+ containsKey(Object key)
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

## 7. HashSet

+ add(obj)
+ remove(objectx)
+ size()
+ contains(obj)
+ clear()

## 8. String

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

字符串排序：

```java
char[] array = str.toCharArray();// 转换为字符数组
Arrays.sort(array);// 根据字典排序
String key = new String(array);
```

