package main

/*
动态规划
通过把原问题分解为相对简单的子问题的方式求解复杂问题的方法。动态规划常常适用于有重叠子问题和最优子结构性质的问题。

基本思想
若要解一个给定问题，我们需要解其不同部分（即子问题），再合并子问题的解以得出原问题的解。 通常许多子问题非常相似，
为此动态规划法试图仅仅解决每个子问题一次，从而减少计算量： 一旦某个给定子问题的解已经算出，则将其记忆化存储，以便
下次需要同一个子问题解之时直接查表。 这种做法在重复子问题的数目关于输入的规模呈指数增长时特别有用。

### 分治与动态规划
共同点：二者都要求原问题具有最优子结构性质,都是将原问题分而治之,分解成若干个规模较小(小到很容易解决的程序)的子问题.
然后将子问题的解合并,形成原问题的解.

不同点：分治法将分解后的子问题看成相互独立的，通过用递归来做。
　　　　动态规划将分解后的子问题理解为相互间有联系,有重叠部分，需要记忆，通常用迭代来做。

一、最优子结构
如果一个问题的解结构包含了其子问题的最有解，那么称此问题具有最优子结构性质。
使用动态规划算法是，用子问题的最优解来构造问题的最优解，必须考察最优解中用到的所有子问题。

二、重叠子问题
如果递归算法反复的计算相同的子问题，而不是简单的生成新的子问题，就称问题具有重复子问题性质。
例如斐波拉契数列的计算，递归算法中就会有很多的重复计算的子问题。
在动态规划中，通常使用数组保存中间子问题的结果，拒绝重复计算同一个子问题。
 */
