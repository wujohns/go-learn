# go 相关学习笔记
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

## GFW 相关
直接使用 `vpn` 的话由于特征过于明显容易被封杀，比较不容易被封杀的 `shadowsocks` 由于采用的是 `socks5` 协议无法直接供 `go` 使用。

### 代理方案
这里使用 [cow](https://github.com/cyfdecyf/cow/) 在本地创建 `http` 代理，并将 `shadowsocks` 作为二级代理使用。从而提供可以被 `go` 使用的 `http` 代理。

在 [release 页面](https://github.com/cyfdecyf/cow/releases) 下载自己系统需要用到的版本。

windows 版本下载解压后，通过配置其中的 `rc.txt` 即可完成对 `cow` 的配置，若想通过本地的 `shadowsocks` 客户端进行连接，以下是作为参考的配置，可从通过 `socks5` 连接与直连 `shadowsocks` 任选其一：

```txt
# socks5
proxy = socks5://127.0.0.1:1080

# shadowsocks
proxy = ss://<加密方式>:<密码>@<服务器地址>:<端口>
```

同时配置 `cow` 开启的 `http` 代理端口，如下采用默认的 `7777` 端口:

```txt
listen = http://127.0.0.1:7777
```

配置完成后点击 `cow-hide.exe` 或 `cow-taskbar` 即可启动服务。
综上所述，配置可以如下：

```txt
# rc.txt
proxy = socks5://127.0.0.1:1080
listen = http://127.0.0.1:7777
```

其他详细操作可以参考 `cow` 的文档。

### 配置 go get 使用代理
很多教程中为设置 `http_proxy` 与 `https_proxy` 环境变量来让 `go get` 使用代理，但该环境变量的配置会导致所有的 `http` 请求都走该代理令人困扰。

在 `win` 中采用设置临时环境变量的方式让在用到 `go get` 时走代理：  
```
set http_proxy=http://127.0.0.1:7777
set https_proxy=http://127.0.0.1:7777
``` 

在 `vscode` 中则可以设置工作区的 `http.proxy` 参数为 `http://127.0.0.1:7777` 来实现内置命令运行时走 `http` 代理。

总之这块是比较坑（说一个笑话，go的包管理）