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
n >> k & 1 // 求n的第k位数字
n & -n  // 返回n的最后一位1
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

### 离散化

```java
all = new ArrayList<>(new HashSet<Integer>(all)); // 去重
Collections.sort(all); // 排序

// 左边界二分查找
public static int find(int x) {
    int l = 0, r = all.size() - 1;
    while (l < r) {
        int mid = l + r >> 1;
        if (x <= all.get(mid)) r = mid;
        else l = mid + 1;
    }
    return l + 1;
}

// 新的下标
find(x)
```

### 区间合并

```c++
vector<PII> merge(vector<PII> &segs) {
    vector<PII> res;
    sort(segs.begin(), segs.end());
    
    int st = -2e9, ed = -2e9;
    for (auto seg : segs)
        if (ed < seg.first) {
            if (st != -2e9) res.push_back({st, ed});
            st = seg.first, ed = seg.second;
        }
        else ed = max(ed, seg.second);
    
    if (st != -2e9) res.push_back({st, ed});
    return res;
}
```

## 2. 数据结构

### kmp

```c++
const int N = 10010, M = 100010;
int n, m, ne[N];
char p[N], s[M];

cin >> n >> p + 1 >> m >> s + 1; 

// 求模式串p的 next 数组
for (int i = 2, j = 0; i <= n; i ++ )
{
    while (j && p[i] != p[j + 1]) j = ne[j];
    if (p[i] == p[j + 1]) j ++ ;
    ne[i] = j;
}

// kmp匹配
for (int i = 1, j = 0; i <= m; i ++ )
{
    while (j && s[i] != p[j + 1]) j = ne[j];
    if (s[i] == p[j + 1]) j ++ ;
    if (j == n)
    {
        j = ne[j];
        // 成功后的逻辑
    }
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

### Dijstra

```c++
int h[N], e[N], ne[N], w[N], idx;
int n, m;
typedef pair<int, int> pa;
int d[N];
bool st[N]; // 和spfa的st数组不一样，这个表示加入集合

int dijstra()
{
    memset(d, 0x3f, sizeof d);
    d[1] = 0;
    priority_queue<pa, vector<pa>, greater<pa> > q;
    q.push({0, 1});
    
    while (q.size())
    {
        pa t = q.top();
        q.pop();
        int v = t.second, dist = t.first;
        
        if (st[v]) continue;
        st[v] = true;
        
        for (int i = h[v]; i != -1; i = ne[i])
        {
            int j = e[i];
            if (d[j] > dist + w[i])
            {
                d[j] = dist + w[i];
                q.push({d[j], j});
            }
        }
    }
    
    if (d[n] == 0x3f3f3f3f) return -1;
    else return d[n];
}

void add(int u, int v, int c)
{
    e[idx] = v, ne[idx] = h[u], w[idx] = c, h[u] = idx++;    
}

int main()
{
    cin >> n >> m;
    memset(h, -1, sizeof h);
    
    for (int i = 0; i < m; i++)
    {
        int a, b, c;
        cin >> a >> b >> c;
        add(a, b, c);
    }
    cout << dijstra() << endl;
    return 0;
}
```

### BF算法

```c++
int n, m, k;
vector<vector<int> > e(10005, vector<int>(3, 0));
int d[N], backup[N];

int main()
{
    cin >> n >> m >> k;
    for (int i = 0; i < m; i++) cin >> e[i][0] >> e[i][1] >> e[i][2];
    
    memset(d, 0x3f, sizeof d);
    d[1] = 0;
    
    for (int i = 0; i < k; i++)
    {
        memcpy(backup, d, sizeof d);
        for (int j = 0; j < m; j++)
        {
            int u = e[j][0], v = e[j][1], w = e[j][2];
            d[v] = min(d[v], backup[u] + w);
        }
    }
    
    if (d[n] > 0x3f3f3f3f / 2) cout << "impossible" << endl;
    else cout << d[n] << endl;
    
    return 0;
}
```

### spfa

判断负环就加一个 `cnt` 数组

```c++
int n, m;
int idx, e[N], ne[N], h[N], w[N];
bool st[N];
int d[N];

void add(int u, int v, int c)
{
    e[idx] = v, ne[idx] = h[u], w[idx] = c, h[u] = idx++; 
}

int main()
{
    cin >> n >> m;
    memset(h, -1, sizeof h);
    for (int i = 0; i < m; i++)
    {
        int a, b, c;
        cin >> a >> b >> c;
        add(a, b, c);
    }
    
    memset(d, 0x3f, sizeof d);
    d[1] = 0;
    queue<int> q;
    q.push(1);
    st[1] = true;
    
    while (q.size())
    {
        int t = q.front();
        q.pop();
        st[t] = false;
        
        for (int i = h[t]; i != -1; i = ne[i])
        {
            int j = e[i];
            if (d[j] > d[t] + w[i])
            {
                d[j] = d[t] + w[i];
                if (!st[j])
                {
                    q.push(j);
                    st[j] = true;
                }
            }
        }
    }
    
    if (d[n] == 0x3f3f3f3f) cout << "impossible" << endl;
    else cout << d[n] << endl;
    
    return 0;
}
```

### prim

无向图，2条边

```c++
const int inf = 0x3f3f3f3f;
int g[N][N], d[N];
int n, m;
bool st[N];

int prim()
{
    int res = 0;
    memset(d, 0x3f, sizeof d);
    // d[1] = 0 加不加无所谓，反正没有算入res
    
    for (int i = 0; i < n; i++)
    {
        int t = -1;
        for (int j = 1; j <= n; j++)
            if (!st[j] && (t == -1 || d[t] > d[j])) t = j;
                
        if (i && d[t] == inf) return inf;
        if (i) res += d[t];
        st[t] = true;
        
        for (int j = 1; j <= n; j++) d[j] = min(d[j], g[t][j]);
    }
    return res;
}

int main()
{
    cin >> n >> m;
    memset(g, 0x3f, sizeof g);
    for (int i = 0; i < m; i++)
    {
        int a, b, c;
        cin >> a >> b >> c;
        g[a][b] = min(g[a][b], c);
        g[b][a] = g[a][b]; // 无向图
    }
    int res = prim();
    if (res == inf) cout << "impossible" << endl;
    else cout << res << endl;
    return 0;
}
```

### Kruskal

结构体排序要快很多

```c++
struct Edge
{
    int u, v, w;
    bool operator<(Edge& W) const
    {
        return w < W.w;
    }
}e[M];
int n, m, p[N];

int find(int x)
{
    if (p[x] != x) p[x] = find(p[x]);
    return p[x];
}

int main()
{
    cin >> n >> m;
    for (int i = 1; i <= n; i++) p[i] = i;
    int a, b, c;
    for (int i = 0; i < m; i++)
    {
        scanf("%d%d%d", &a, &b, &c);
        e[i] = {a, b, c};
    }
    
    sort(e, e + m);
    int res = 0, cnt = 0;
    for (int i = 0; i < m; i++)
    {
       
        int u = e[i].u, v = e[i].v;
        int a = find(u), b = find(v);
        if (a != b)
        {
            p[a] = b;
            cnt++;
            res += e[i].w;
        }
    }
    
    if (cnt < n - 1) cout << "impossible" << endl;
    else cout << res << endl;
    
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

## 5. 动态规划

















