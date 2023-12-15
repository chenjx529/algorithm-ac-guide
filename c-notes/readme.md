## 1. 基础算法

### 快排

```java
static void quick_sort(int[] q, int l, int r) {
    if (l >= r) return;
    
    int i = l - 1;
    int j = r + 1;
    int x = q[(l + r) >> 1];
    
    while (i < j) {
        while (q[++i] < x);
        while (q[--j] > x);
        if (i < j) swap(q, i, j);
    }
    
    quick_sort(q, l, j);
    quick_sort(q, j + 1, r);
}
```

### 归并

```java
static void merge_sort(int[] q, int l, int r) {
    if (l >= r) return;
    
    int mid = (r + l) >> 1;
    merge_sort(q, l, mid);
    merge_sort(q, mid + 1, r);

    int k = 0, i = l, j = mid + 1;
    int[] t = new int[r - l + 1];
    while (i <= mid && j <= r) {
        if (q[i] <= q[j]) t[k++] = q[i++];
        else t[k++] = q[j++];
    }
    
    while (i <= mid) t[k++] = q[i++];
    while (j <= r) t[k++] = q[j++];
    
    for (i = l, j = 0; i <= r; i++, j++) {
        q[i] = t[j];
    }
}
```

### 二分

```java
// 二分左边界，条件是 q[mid] >= x，当 q[mid] == x，mid就是左边界，最大值最小的问题
static int bsearch_left(int[] q, int x, int l, int r) {
    while (l < r) {
        int mid = (l + r) >> 1; // 取左边的数
        if (q[mid] >= x) r = mid;
        else l = mid + 1;
    }
    return l;
}

// 二分右边界，条件是 q[mid] <= x, 当 q[mid] == x，mid就是右边界，最小值最大的问题
static int bsearch_right(int[] q, int x, int l, int r) {
    while (l < r) {
        int mid = (l + r + 1) >> 1;  // 取右边的数
        if (q[mid] <= x) l = mid;
        else r = mid - 1;
    }
    return r;
}
```
### 前缀和

```java
// 构造一维前缀和
s[i] = a[1] + a[2] + ... a[i]
// 计算数组部分和 
a[l] + ... + a[r] = s[r] - s[l - 1]

// 构造二维前缀和
s[i][j] = a[i][j] + s[i - 1][j] + s[i][j - 1] - s[i - 1][j - 1]
// 计算数组部分和 
s[x2, y2] - s[x1 - 1, y2] - s[x2, y1 - 1] + s[x1 - 1, y1 - 1]
```

### 差分

```java
// 构造一维差分    
b[l] += c
b[r + 1] -= c
// 一维差分求和，显示变化后的数组
for (int i = 1; i <= n; i++) {
    b[i] += b[i - 1];
    printf("%d ", B[i]);
}

// 构造二维差分
b[x1][y1] += c
b[x2 + 1][y1] -= c
b[x1][y2 + 1] -= c
b[x2 + 1][y2 + 1] += c
// 二维差分求和，显示变化后的数组
for (int i = 1; i <= n; i++)
    for (int j = 1; j <= m; j++)
        b[i][j] += b[i - 1][j] + b[i][j - 1] - b[i - 1][j - 1];
```

### 位运算

```java
// 求n的第k位数字
n >> k & 1 
// 一般分开写
n & 1 == 1
n >>= 1

// 返回n的最后一位1
n & -n  
```

### 双指针

```c++
int l = 0, r = n - 1;

while (l < r) {
    func();
    if (check()) r--;
    else l++;
} 

for (int i = 0, j = 0; i < n; i++) {
    while (j < i && check(i, j)) j++ ;
    func();
}
```

### 区间合并

```c++
Integer[][] q = new Integer[n][2];
for (int i = 0; i < n; i++) {
    q[i][0] = sc.nextInt();
    q[i][1] = sc.nextInt();
}
Arrays.sort(q, (a, b) -> a[0] - b[0]);

int st = Integer.MIN_VALUE, ed = Integer.MIN_VALUE;
int res = 0;
for (int i = 0; i < n; i++) {
    if (ed < q[i][0]) {
        if (st != Integer.MIN_VALUE) {
            res++;
        }
        st = q[i][0];
        ed = q[i][1];
    } else {
        ed = Math.max(ed, q[i][1]);
    }
}
if (st != Integer.MIN_VALUE) {
    res++;
} 
```

## 2. 数据结构

### 单调队列

```java
int[] a, q;
int hh = 0, tt = -1;
for (int i = 0; i < n; i++) {
    // 移动对头，q ∈ [i - k + 1, i]，这个区间是滑动窗口
    if (hh <= tt && q[hh] < i - k + 1) hh ++; 
    
    // 移动队尾，注意有等号。如果是单调递减的话，只需要改成<=，其他地方不用修改
    while (hh <= tt && a[q[tt]] >= a[i]) tt--; 
    
    // 入队，此时前面的数都比a[i]小
    q[++tt] = i; 
    
    // 找出滑动窗口中的最大值/最小值
    if (i >= k - 1) {
        System.out.print(a[q[hh]] + " ");
    }
}
```

### 单调栈

```java
int[] stk = new int[n];
int tt = -1;
for (int i = 0; i < n; i++) {
    int x = sc.nextInt();
    
    // 单调递增的栈
    while (tt >= 0 && stk[tt] >= x) tt--; 
    
    // 输出每个数左边离它最近的比它大 / 小的数
    if (tt >= 0) System.out.print(stk[tt] + " ");
    else System.out.print(-1 + " ");
    
    // 入栈
    stk[++tt] = x;
}
```

### 并查集

```c++
int p[N];

int find(int x)
{
    if (p[x] != x) p[x] = find(p[x]);
    return p[x];
}

for (int i = 1; i <= n; i ++ ) p[i] = i;


p[find(a)] = find(b);
```

## 3. 图论

### 拓扑排序

```c++
int d[N], q[N];

bool topsort()
{
    int hh = 0, tt = -1;
	
    for (int i = 1; i <= n; i ++ )
        if (!d[i]) q[ ++ tt] = i;

    while (hh <= tt)
    {
        int t = q[hh ++ ];

        for (int i = h[t]; i != -1; i = ne[i])
        {
            int j = e[i];
            if (-- d[j] == 0) q[ ++ tt] = j;
        }
    }

    return tt == n - 1;
}
```

### 染色法判断二分图

无向图，2个add，

```c++
bool dfs(int u, int c)
{
    color[u] = c;
    for (int i = h[u]; i != -1; i = ne[i])
    {
        int j = e[i];
        if (color[j] == -1)
        {
            if (!dfs(j, !c)) return false; // 染成相反的颜色
        }
        else if (color[j] == c) return false;
    }
    return true;
}
```

### 匈牙利算法求最大匹配

```c++
int h[505], e[N], ne[N], idx;
int n1, n2, m;
int match[505];
bool st[505];

void add(int u, int v)
{
    e[idx] = v, ne[idx] = h[u], h[u] = idx++;
}

bool find(int x)
{
    for (int i = h[x]; i != -1; i = ne[i])
    {
        int j = e[i];
        if (!st[j])
        {
            st[j] = true;
            if (!match[j] || find(match[j]))
            {
                match[j] = x;
                return true;
            }
        }
    }
    return false;
}

int main()
{
    cin >> n1 >> n2 >> m;
    memset(h, -1, sizeof h);
    for (int i = 0; i <= m; i++)
    {
        int a, b;
        cin >> a >> b;
        add(a, b);
    }
    
    int res = 0;
    for (int i = 1; i <= n1; i++)
    {
        memset(st, false, sizeof st);
        if (find(i)) res++;
    }
    cout << res << endl;
    return 0;
}
```

## 4. 数学

### 线性筛质数

```c++
int primes[N], cnt;
bool st[N];

void get_primes(int n)
{
    for (int i = 2; i <= n; i ++ ) 
    {
        if (!st[i]) primes[cnt ++ ] = i;
        for (int j = 0; primes[j] <= n / i; j ++ ) 
        {
            st[primes[j] * i] = true; 
            if (i % primes[j] == 0) break;// 核心步骤，我们的目的是使用最小质因子筛出合数，避免重复，保证了时间复杂度为O(n)
        }
    }
}
```

### 快速幂

```c++
long long qmi(int m, int k, int p)
{
    long long res = 1, t = m; 
    while (k)
    {
        if (k & 1)  res = res * t % p;
        t = t * t % p; 
        k >>= 1;
    }
    return res;
}
```



















