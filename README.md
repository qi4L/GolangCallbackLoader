![](https://socialify.git.ci/nu1r/GolangCallbackLoader/image?font=Raleway&language=1&logo=https%3A%2F%2Fs1.ax1x.com%2F2022%2F09%2F12%2FvXqOUI.jpg&name=1&pattern=Signal&theme=Light)

重写[AlternativeShellcodeExec](https://github.com/aahmad097/AlternativeShellcodeExec)成GOLang版本的
# BUG

其中几个实现中的BUG

- EnumICMProfiles: EnumICMProfilesW API无法检索上下文
- EnumPropsEx： EnumPropsExW API没持久回调
- EnumPropsW：EnumPropsW API没持久回调
- InitOnceExecuteOnce: 成功回调但是无法加载地址
- SysEnumSourceFiles: 没持久回调
- 函数地址问题, 目前我不知道怎么吧函数地址转换成函数来调用, Google的答案是:
  - functionGo 中的类型不可寻址且不可比较，因为：函数指针表示函数的代码。而函数字面量创建的匿名函数的代码在内存中只存储一次，无论返回匿名函数值的代码运行了多少次。
  - 存在此问题的有:
  - FiberContextEdit
  - LdrEnumerateLoadedModules
  - LdrpCallInitRoutine
  - RtlUserFiberStart
  
除了上述几个加载器，其他都是可以正常使用的