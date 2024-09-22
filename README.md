# TelMsg

## 项目简介
TelMsg 是一个基于 **Telegram Bot** 和 **Gin** 框架的非常简单通知系统。用户可以通过 HTTP 请求向 Telegram Bot 发送通知。

## 准备工作

1. Telegram Bot Token
2. 一台公网服务器
3. 服务器部署go环境

## 安装与使用

### 1. 克隆项目
首先，克隆项目到本地机器上：
```sh
git clone <repository-url>
cd <repository-directory>
```

### 2.安装依赖

```sh
go mod tidy
```

### 3.运行项目

```sh
go run main.go <your-telegram-bot-token>
```

### 4.使用API

项目启动后，会在 http://localhost:6001 上运行一个 HTTP 服务器。你可以通过 POST 请求向 /notice 端点发送通知。请求参数包括：  
token: 用户的 token
title: 消息标题
content: 消息内容
示例请求：

```sh
#!/bin/bash
usertoken="token"
req="test"
curl --location --request POST 'http://localhost:6001/notice' \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "token=$usertoken&title=title&content=$req" \
-s
```

![image-20240922170909283](https://img.xlasm.com/i/2024/09/22/66efdeb63395d.png)

## 用户Token的获取

1. 机器人添加命令 `start`

2. 运行项目
3. 点击`start` ,返回用户token

每次重新运行项目.都需要重新`start`获取一次