![](https://socialify.git.ci/nu1r/GolangCallbackLoader/image?font=Raleway&language=1&logo=https%3A%2F%2Fs1.ax1x.com%2F2022%2F09%2F12%2FvXqOUI.jpg&name=1&pattern=Signal&theme=Light)

重写[AlternativeShellcodeExec](https://github.com/aahmad097/AlternativeShellcodeExec)成GOLang版本的
# BUG

其中几个实现中的BUG

- EnumICMProfiles: EnumICMProfilesW API无法检索上下文
- EnumPropsEx： EnumPropsExW API没持久回调
- EnumPropsW：EnumPropsW API没持久回调
- InitOnceExecuteOnce: 成功回调但是无法加载地址
- SysEnumSourceFiles: 没持久回调
- FiberContextEdit：还在改
- LdrEnumerateLoadedModules：获取到的函数地址无效
- LdrpCallInitRoutine：申请到的函数地址无效
- RtlUserFiberStart：待写
  
除了上述几个加载器，其他都是可以正常使用的