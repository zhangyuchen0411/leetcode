- [315. 计算右侧小于当前元素的个数](#315-计算右侧小于当前元素的个数)
  - [题目](#题目)
  - [题解](#题解)
    - [方法一：离散化树状数组](#方法一离散化树状数组)

------------------------------

# 315. 计算右侧小于当前元素的个数

## 题目

给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： `counts[i]` 的值是 `nums[i]` 右侧小于 `nums[i]` 的元素的数量。

示例:

```
输入: [5,2,6,1]
输出: [2,1,1,0] 
解释:
5 的右侧有 2 个更小的元素 (2 和 1).
2 的右侧仅有 1 个更小的元素 (1).
6 的右侧有 1 个更小的元素 (1).
1 的右侧有 0 个更小的元素.
```

- 来源：力扣（LeetCode）
- 链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
- 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 题解

这题严重超纲。直接上答案吧。


> 链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/solution/ji-suan-you-ce-xiao-yu-dang-qian-yuan-su-de-ge-s-7/

### 方法一：离散化树状数组

**预备知识**

「**树状数组**」是一种可以动态维护序列**前缀和**的数据结构，它的功能是：

- **单点更新** update(i, v)： 把序列 i 位置的数加上一个值 v，在该题中 v = 1
- **区间查询** query(i)： 查询序列 $[1 \cdots i]$ 区间的区间和，即 i 位置的前缀和

修改和查询的时间代价都是 $O(\log n)$，其中 n 为需要维护前缀和的序列的长度。

**思路与算法**

记题目给定的序列为 a，我们规定 $a_i$ 的取值集合为 a 的「值域」。我们**用桶来表示值域中的每一个数**，桶中记录这些数字出现的次数。假设 $a = \{5, 5, 2, 3, 6\}$，那么遍历这个序列得到的桶是这样的：

```
index  ->  1 2 3 4 5 6 7 8 9
value  ->  0 1 1 0 2 1 0 0 0
```

**转化为动态维护前缀和问题**

记 value 序列为 v，我们可以看出它第 i - 1 位的前缀和表示「有多少个数比 i 小」(上面1-4的和为2，表示比5小的有2个)。那么我们可以从后往前遍历序列 a，记当前遍历到的元素为 $a_i$，我们把 $a_i$ 对应的桶的值自增 1，记 $a_i = p$，**把 v 序列 p - 1 位置的前缀和加入到答案中算贡献**。为什么这么做是对的呢，因为我们在循环的过程中，我们把原序列分成了两部分，后半部部分已经遍历过（已入桶），前半部分是待遍历的（未入桶），那么我们求到的 p - 1 位置的前缀和就是「已入桶」的元素中比 p 小的元素的个数总和。这种动态维护前缀和的问题我们可以用「树状数组」来解决。

**用离散化优化空间**

我们显然可以用数组来实现这个桶，可问题是如果 $a_i$ 中有很大的元素，比如 $10^9$，我们就要开一个大小为 $10^9$ 的桶，内存中是存不下的。这个桶数组中很多位置是 0，有效位置是稀疏的，我们要想一个办法让有效的位置全聚集到一起，减少无效位置的出现，这个时候我们就需要用到一个方法——**离散化**。离散化的方法有很多，但是目的是一样的，即**把原序列的值域映射到一个连续的整数区间，并保证它们的偏序关系不变**。 这里我们将原数组**去重后排序**，原数组每个数**映射**到去重排序后这个数对应位置的**下标**，我们称这个下标为这个对应数字的 id。已知数字获取 id 可以在去重排序后的数组里面做**二分查找**，已知 id 获取数字可以直接把 id 作为下标访问去重排序数组的对应位置。大家可以参考代码和图来理解这个过程。

![](assets/no_0315_count_of_smaller_numbers_after_self.gif)

其实，计算每个数字右侧小于当前元素的个数，这个问题我们在「剑指 Offer 51. 数组中的逆序对」题解的「方法二：离散化树状数组」中遇到过，在统计逆序对的时候，只需要统计每个位置右侧小于当前元素的个数，再对它们求和，就可以得到逆序对的总数。这道逆序对的题可以作为本题的补充练习。

```go
var a, c []int

func countSmaller(nums []int) []int {
    resultList := []int{}
    // 去重并且排序，生成 a 数组
    discretization(nums)
    // 为什么 +5 ？其实加1就够了吧，0没有使用到，而且是 a.len()+1
    c = make([]int, len(nums) + 5)
    for i := len(nums) - 1; i >= 0; i-- {
        // id 是 nums[i] 在 a 数组中的下标
        id := getId(nums[i])
        // c[0] + ... + c[id-1]
        resultList = append(resultList, query(id - 1))
        // 把 c[id] 到最后的位置都自增1.
        update(id)
    }
    // 倒序
    for i := 0; i < len(resultList)/2; i++ {
        resultList[i], resultList[len(resultList)-1-i] = resultList[len(resultList)-1-i], resultList[i]
    }
    return resultList
}

func lowBit(x int) int {
    return x & (-x)
}

func update(pos int) {
    for pos < len(c) {
        c[pos]++
        pos += lowBit(pos)
    }
}

func query(pos int) int {
    ret := 0
    for pos > 0 { // 0 索引没有使用到
        ret += c[pos]
        pos -= lowBit(pos)
    }
    return ret
}

func discretization(nums []int) {
    set := map[int]struct{}{}
    for _, num := range nums {
        set[num] = struct{}{} 
    }
    a = make([]int, 0, len(nums))
    for num := range set {
        a = append(a, num)
    }
    sort.Ints(a)
}

func getId(x int) int {
    return sort.SearchInts(a, x) + 1
}
```

**复杂度分析**

假设题目给出的序列长度为 n。

- 时间复杂度：我们梳理一下这个算法的流程，这里离散化使用哈希表去重，然后再对去重的数组进行排序，时间代价为 $O(n \log n)$；初始化树状数组的时间代价是 $O(n)$；通过值获取离散化 id 的操作单次时间代价为 $O(\log n)$；对于每个序列中的每个元素，都会做一次查询 id、单点修改和前缀和查询，总的时间代价为 $O(n \log n)$。故渐进时间复杂度为 $O(n \log n)$。
- 空间复杂度：这里用到的离散化数组、树状数组、哈希表的空间代价都是 $O(n)$，故渐进空间复杂度为 $O(n)$。
