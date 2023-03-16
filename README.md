![](https://socialify.git.ci/nu1r/GolangCallbackLoader/image?font=Raleway&language=1&logo=https%3A%2F%2Fs1.ax1x.com%2F2022%2F09%2F12%2FvXqOUI.jpg&name=1&pattern=Signal&theme=Light)

重写[AlternativeShellcodeExec](https://github.com/aahmad097/AlternativeShellcodeExec)成GOLang版本的
# BUG

- SysEnumSourceFiles: 调用正常, 但是没上线。
- LdrEnumerateLoadedModules：函数地址无效
- LdrpCallInitRoutine：函数地址无效
- RtlUserFiberStart：待写
  
除了上述几个加载器，其他都是可以正常使用的