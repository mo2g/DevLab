# AgentScan

[English](./README.md)

- Repository: [AutoScan/agentscan](https://github.com/AutoScan/agentscan)
- 相关笔记：[AgentScan 工程思考](../cookbook/agentscan-agent-security-notes/README.zh-CN.md)

## 我的角色

AI Agent 资产发现与安全审计平台的参与者与维护协作者。

## 它解决什么问题

越来越多 AI Agent 系统以弱默认配置、宽暴露面和低运维意识的方式被部署出来。通用资产扫描器可以告诉你某个服务开着，但很难继续解释它是不是 AI Agent、具有什么特征、对应什么风险。AgentScan 的目标是把暴露 AI Agent 视作一个需要被发现、验证、报告和跟进的独立安全资产面。

## 技术栈

Go 后端与 CLI、Gin API、GORM 持久层、React 前端、WebSocket 实时更新，以及围绕扫描管线的任务、告警和报告能力。

## 当前阶段

Active platform build。当前方向已经覆盖分层扫描、dashboard、定时任务、告警和报告，下一步重点是进一步加固、扩展 signatures，并提升真实扫描环境下的规模适应能力。

## 为什么值得关注

AgentScan 处在 AI 系统与安全运营的交叉点上。它有意思的不只是产品表面，更在于它展示了一件事：当 AI-native 基础设施被当作资产类别来治理，而不是当作演示工程来对待时，很多安全问题才会真正进入可管理状态。
