# 1984 Newspeak 聊天系统

基于乔治·奥威尔《1984》中Newspeak概念设计的聊天室内，具有思想审查和语言净化功能。

## 功能特性
- 实时消息传递
- 基于 go langchain 接入大模型自动思想审查和语言净化
- 基于Web的前端界面

## 项目结构

```
newspeak-chat/
├── etc/                - 配置文件
├── frontend/           - 前端资源(HTML, CSS, JS)
├── internal/           - 核心应用逻辑
│   ├── config/         - 配置处理
│   ├── handler/        - HTTP处理器
│   ├── logic/          - 业务逻辑
│   ├── svc/            - 服务上下文
│   ├── types/          - 数据类型
│   └── ws/             - WebSocket处理器
├── main.go             - 应用入口
└── newspeak.api        - API定义
```

## 配置说明

编辑`.env`文件设置环境相关配置

## API文档

API文档详见`newspeak.api`和`etc/1984.openapi.json`文件

## 服务器部署

推荐使用 caddy 进行反向代理

## TODO

- [ ] 修复聊天功能