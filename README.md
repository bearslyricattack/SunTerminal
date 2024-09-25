# SunTerminal

# 介绍

命令行运维小工具。
如果你时常看不懂命令行的输出，时常要拿它问ai，可以考虑使用此工具。
运行start命令，将会创建一个交互式命令行工具。在这里运行的每个命令的返回值都会询问大模型后返回，模型的信息使用配置文件配置。运行效果如下：

![](/test.png)

# 安装

## macos

1.clone项目

2.找到deployment/sun

3.sudo cp sun /usr/local/bin/ 复制到bin目录下

4.授予权限

5.运行命令

## windows

与运行exe文件相同.

也可以自行编译运行.

# 命令

共有两个命令，config与start。

## config

指定配置文件位置 生成配置文件模版。

一个参数 path 含义为配置文件位置。

```
model:
    - name: "chat" 
      path: "https://api.openai.com/v1/chat/completions"
      key: ""
      type: "gpt-3.5-turbo"
context: "chat"

```

参数类型为 ：模型 地址 key 类型 目前仅支持chatgpt官方api.

使用与kubeconfig相同的context机制。

## start

开启交互式客户端

一个参数 path 含义同上。

```
aterminal start --path /Users/wpy
```

# 总结

希望大家用的开心！

```
我去问穿过山谷的风你在哪 你去了哪啊
我去问淌过天空的水那是你的泪吗
我去问落下的星 天上的哪颗是你啊
--《云海里的帆》赵仰睿
```

