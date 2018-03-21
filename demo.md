# demo 的相关说明
这里对附带的 `demo` 做说明，作为 `go` 工程样板的相关参考。

## vscode 配置相关
### 工作区配置
参考 `.vscode/settings.json` 中的配置，其中：
`go.gopath`: 设定当前工作区的 `GOPATH`，建议设定为系统环境变量中的GOPATH再加上当前工程自己独有的GOPATH，方便将自己代码逻辑的拆分。  
`search.exclude` 与 `files.exclude`： 设定搜索屏蔽与文件显示屏蔽，避免对项目开发不必要的文件干扰

```json
// settings.json
{
    "go.gopath": "E:/gopath/;${workspaceRoot}/libs",
    "search.exclude": {
        "libs/pkg": true,
        "debug": true
    },
    "files.exclude": {
        "**/.git": true,
        "libs/pkg": true,
        "debug": true
    }
}
```

### 调试配置
参考 `.vscode/launch.json` 中的配置，需要注意的是以下配置项：
`program`: 该程序的入口文件，启动调试时将会从此处编译调试  
`env.GOPATH`： 建议设定为系统环境变量中的GOPATH再加上当前工程自己独有的GOPATH  

```json
// launch.json
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
            "env": {
                "GOPATH": "E:/gopath/;${workspaceRoot}/libs"
            },
            "args": [],
            "showLog": true
        }
    ]
}
```

## 模块化
通过上述在编辑器内的准备工作，就可以开始这一步操作了，上述中的demo中将 `${workspaceRoot}/libs` 设定为了该工程独有的 `GOPATH`, 按照 `go` 的规范我们可以将自己写的模块放置在 `libs/src` 目录下，参考对照 `libs/src/tryone` 目录下的内容以及 `main.go` 中对该模块的调用，即可习得相关的细节。