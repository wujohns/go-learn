# go 相关学习笔记
先说结论，`go` 就是一坨shit，从18/3/20后的两年内都改变不了这个事实。谷歌的开源项目若没有被开源社区反逼（比如 `io.js` 最终的结局都是垃圾堆）。

对 `go` 的开坑与填坑只是工作的需要，个人对这块毫无兴趣，推荐使用 `rust`。

## 初始开发环境配置
win下的安装：
在 google code 的下载页下载msi安装包：https://golang.org/dl/

环境变量配置：  
GOROOT：go安装的位置（ex：D:\Go）  
执行路径：安装位置的bin目录（ex：win下为 D:\Go\bin）  
GOPATH：包加载路径，可设置多个用分号分隔，使用import时会在这些路径中寻找

一些依赖的包的安装（配合调试）：
`gopkgs`: go get github.com/uudashr/gopkgs/cmd/gopkgs  
`godef`: go get github.com/rogpeppe/godef  
`dlv`: go get github.com/derekparker/delve/cmd/dlv  
`gocode`: go get github.com/nsf/gocode  
`goreturns`: go get github.com/sqs/goreturns  
等，可以在 vscode 的 `GO` 插件安装后执行 `install all` 安装所有调试时依赖的包

需要注意的是为了规避 go 在包管理方面的坑，这里的工程不会任何的第三方包，只依赖内置库

## vscode 配置相关
这里使用 `vscode` 配合 `go` 的开发，为了便于开发需要对 `vscode` 做以下配置：  
1. 安装 `Go` 插件以便于对代码编写提供方便的支持  
1. 配置 `./vscode/launch.json` 作为调试启动的配置

`launch.json` 样板：

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceRoot}/main.go",
            "env": {},
            "args": [],
            "showLog": true
        }
    ]
}
```

然后 `F5` 或点击启动按钮既可开始调试项目（支持断点调试等）

## 简易案例
在 [demo](/demo) 目录中是一个 `go` 工程的简单案例，主要为在工程内的模块化案例。可以配合 [demo说明文档](/demo.md) 进行查看