学习笔记

抵制人肉递归
找最近重复性，找到最近最简方法，将其拆解成可重复解决的问题（重复子问题）
数学归纳法思维（枚举，n=1怎么样 n=2怎么样，不要过多，前几层即可要不就会变成人肉递归，最后泛化成n=n怎么样）

代码模板：

```
public void recur(int level, int param) {
    // 终止条件
    if (level > MAX_LEVEL) {
        // 处理结果
        return;
    }
    // 处理当前层逻辑
    process(level, param);

    // 进入下一层
    recur(level:level+1, newParam);
    // 一些当前层需要清理的东西
}
```