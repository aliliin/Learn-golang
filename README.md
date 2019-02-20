
<h1 align="center">golang learn</h1>

<p align="center"> macOS 下搭建 Go 语言开发环境。</p>

## 安装

1. 直接下载 [.pkg](https://golang.org/dl/) 文件双击进行安装 <br/>
    在命令行输入  ` go version ` 进行确认是否安装成功 <br/>
    成功提示 ` go version go1.11.5 darwin/amd64 `

## 配置

2.  配置环境变量

    ### Bash 

    <hr> 

    编辑你的 ` ~/.bash_profile ` 文件加入如下行

    ``` 
        export GOPATH=$HOME/go  
    ``` 
    
    退出并保存你的 ` ~/.bash_profile ` 文件

    ``` 
        source ~/.bash_profile
    ```

    ### Zsh
    <hr> 
    编辑你的 ` ~/.zshrc ` 文件加入如下行

    ``` 
        export GOPATH=$HOME/go
    ```
    退出并保存你的 ` ~/.zshrc ` 文件

    ``` 
        source ~/.zshrc
    ```

    #### 注意：`GOPATH:` 是你日常开发的文件根目录


## 使用

1. 在你的 `GOPATH:` 开发根目录下创建文件夹 hello

```
     mkdir hello 
     cd  hello
     touch hello.go
 ```
 2. 打开新建的文件 ` hello.go ` 输入如下代码，并保存
 ```
     package main
     
    import (
      "fmt"
    )

    func main() {
      fmt.Println("hello");
    }
 ```
 3. 在命令行输入 `go build hello.go ` 生成 ` exec` 文件，进入文件夹中 直接双击这个 exec 文件，就能看到成功的运行结果 ` hello `  了。
 