- [941. 有效的山脉数组](#941-有效的山脉数组)
  - [官方题解](#官方题解)
    - [方法一：线性扫描](#方法一线性扫描)

------------------------------

# 941. 有效的山脉数组

给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

- A.length >= 3
- 在 0 < i < A.length - 1 条件下，存在 i 使得：
    - `A[0] < A[1] < ... A[i-1] < A[i]`
    - `A[i] > A[i+1] > ... > A[A.length - 1]`


![](assets/no_0941_valid_mountain_array.png)

示例 1：

```
输入：[2,1]
输出：false
```

示例 2：

```
输入：[3,5,5]
输出：false
```

示例 3：

```
输入：[0,3,2,1]
输出：true
```

提示：

- 0 <= A.length <= 10000
- 0 <= `A[i]` <= 10000 

--------------------

- 来源：力扣（LeetCode）
- 链接：https://leetcode-cn.com/problems/valid-mountain-array
- 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 官方题解

> 链接：https://leetcode-cn.com/problems/valid-mountain-array/solution/you-xiao-de-shan-mai-shu-zu-by-leetcode-solution/

### 方法一：线性扫描

按题意模拟即可。我们从数组的最左侧开始向右扫描，直到找到第一个不满足 `A[i] < A[i + 1]` 的下标 i，那么 i 就是这个数组的最高点的下标。如果 i = 0 或者不存在这样的 i（即整个数组都是单调递增的），那么就返回 false。否则从 i 开始继续向右扫描，判断接下来的的下标 j 是否都满足 `A[j] > A[j + 1]`，若都满足就返回 true，否则返回 false。

```go
func validMountainArray(a []int) bool {
    i, n := 0, len(a)

    // 递增扫描
    for ; i+1 < n && a[i] < a[i+1]; i++ {
    }

    // 最高点不能是数组的第一个位置或最后一个位置
    if i == 0 || i == n-1 {
        return false
    }

    // 递减扫描
    for ; i+1 < n && a[i] > a[i+1]; i++ {
    }

    return i == n-1
}
```

复杂度分析

- 时间复杂度：$O(N)$，其中 N 是数组 A 的长度。
- 空间复杂度：$O(1)$。
