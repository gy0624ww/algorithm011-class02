学习笔记

### 二分查找

二分查找算法是典型的「减治思想」的应用，我们使用二分查找将待搜索的区间逐渐缩小，以达到「缩减问题规模」的目的；

掌握二分查找的两种思路：

思路 1：在循环体内部查找元素：while (left <= right)；
思路 2：在循环体内部排除元素：while (left < right)。
全部使用左闭右闭区间，不建议使用左闭右开区间，反而使得问题变得复杂；



最简单的二分查找思路：在一个有序数组里查找目标元素。特别像以前电视「猜价格」上的猜价格游戏：运气好，一下子猜中，如果主持人说猜高了，下一步就应该往低了猜，如果主持人说猜低了，下一步就应该就往高了猜。这个思路把待搜索区间 [left, right] 分为 3 个部分：

mid 位置（只有 1 个元素）；
[left, mid - 1] 里的所有元素；
[mid + 1, right] 里的所有元素；
于是，二分查找就是不断地在区间 [left, right] 里根据 left 和 right 的中间位置 mid = (left + right) / 2 的元素大小，也就是看 nums[mid] 与 target 的大小关系：

如果 nums[mid] == target ，返回 mid；
如果 nums[mid] < target ，由于数组有序，mid 以及 mid 左边的所有元素都小于 target，目标元素可能在区间 [mid + 1, right] 里，因此设置 left = mid + 1；
如果 nums[mid] > target ，由于数组有序，mid 以及 mid 右边的所有元素都大于 target，目标元素可能在区间 [left, mid - 1] 里，因此设置 right = mid - 1。
循环体内一定有 3 个分支，并且第 1 个分支一定用于退出循环，或者直接返回目标元素；

退出循环以后，left 和 right 的位置关系为 [right, left] ，返回 left 或者 right 需考虑清楚

循环可以继续的条件是 while (left <= right)，特别地，当 left == right 即当待搜索区间里只有一个元素的时候，查找也必须进行下去；
int mid = (left + right) / 2; 在 left + right 整形溢出的时候，mid 会变成负数，回避这个问题的办法是写成 int mid = left + (right - left) / 2;