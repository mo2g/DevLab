# DevLab

[English Overview](./README.md)

![DevLab AI Lab overview](./docs/assets/hero-ai-lab.svg)

DevLab 是我对外展示 AI 开源工作的入口仓库，围绕三条主线组织内容：`AI coding workflow`、`token analytics`、`agent security`。

## Start Here

- [文档总览](./docs/README.zh-CN.md)
- [AI 项目索引](./docs/ai/README.zh-CN.md)
- [AI Cookbook](./docs/cookbook/README.zh-CN.md)
- [Token Insight 项目页](./docs/ai/token-insight.zh-CN.md)
- [Codex Composer 项目页](./docs/ai/codex-composer.zh-CN.md)
- [AgentScan 项目页](./docs/ai/agentscan.zh-CN.md)

## About

这个仓库不再只是一个零散的多语言实验集合。我把它重新整理成一个文档优先的 AI 开源入口，用来回答几个最实际的问题：

- 我当前在做什么 AI 相关项目
- 我在这些项目里扮演什么角色
- 我关注哪些工程问题，以及我是怎么做取舍的

这里既会放当前重点维护的 AI 项目，也会保留过去的多语言实验代码，方便快速了解我的技术背景和工作方式。

## Featured Projects

### Token Insight

- Role: 发起者 / 维护者
- Domain: 面向 AI coding tools 的本地优先 token 可观测性
- Stack: Rust, React, SQLite
- Status: Active build

这是一个面向 AI 编码工具的 token usage analytics 项目。它的重点不是做一个新的云平台，而是把本地散落的 usage artifacts 统一采集、标准化、查询和可视化，让成本、趋势、来源健康度这些信息变得可观察、可导出、可分享。

查看详情：

- [项目主页](./docs/ai/token-insight.zh-CN.md)
- [英文文档](./docs/ai/token-insight.md)
- [GitHub 仓库](https://github.com/mo2g/token-insight)

### Codex Composer

- Role: 发起者 / 维护者
- Domain: 面向仓库的可复制 Codex 工作流
- Stack: JavaScript, Shell, Markdown
- Status: Active iteration

这个项目关注的不是单次 prompt，而是如何把 Codex 的使用方式沉淀成一个仓库级别可复制、可验证、可交接的 workflow。核心价值在于把 `AGENTS.md`、验证命令、技能、人工 review 边界这些隐性规则显式化。

查看详情：

- [项目主页](./docs/ai/codex-composer.zh-CN.md)
- [英文文档](./docs/ai/codex-composer.md)
- [GitHub 仓库](https://github.com/mo2g/codex-composer)

### AgentScan

- Role: 参与者 / 维护协作者
- Domain: AI Agent 暴露面发现与安全审计
- Stack: Go, React, SQLite
- Status: Active security tooling

这个项目面向的是一个很现实的问题：越来越多 AI Agent 被直接暴露在内网或公网环境中，而部署者往往没有传统安全经验。AgentScan 把发现、指纹识别、漏洞验证、任务调度和报告输出放在同一个平台里，让 AI Agent 资产能够被真正纳入安全治理。

查看详情：

- [项目主页](./docs/ai/agentscan.zh-CN.md)
- [英文文档](./docs/ai/agentscan.md)
- [GitHub 仓库](https://github.com/AutoScan/agentscan)

## Related AI Work

除了上面的重点项目，这里也记录我参与或维护过的相关 AI 工作，但不会把它们包装成我主导构建的项目：

- [one-api](https://github.com/mo2g/one-api)：围绕 LLM 网关、接口分发和多模型接入场景的维护型 fork 工作。
- [MetaGPT](https://github.com/mo2g/MetaGPT)：基于 fork 的多 Agent 软件流程探索与实践入口。

## Cookbook

我希望 DevLab 不只是“项目链接列表”，还要能沉淀方法和经验。因此这里新增了一组 AI Cookbook，专门记录我在项目推进中踩过的坑、做过的工程取舍，以及哪些经验可以被别的团队直接复用。

- [Cookbook 索引](./docs/cookbook/README.zh-CN.md)
- [Codex Composer 开发经验](./docs/cookbook/codex-composer-workflow.zh-CN.md)
- [Token Insight 开发经验](./docs/cookbook/token-insight-local-first-analytics.zh-CN.md)
- [AgentScan 工程思考](./docs/cookbook/agentscan-agent-security-notes.zh-CN.md)

## Legacy Lab

仓库里的旧目录仍然保留，它们代表了更早期的多语言实验和具体场景解决方案。这部分不再作为首页主角，但仍然保留可访问入口：

- [golang](./golang)
- [java](./java)
- [js](./js)
- [php](./php)
- [python](./python)

## Collaboration

如果你关注以下方向，这个仓库就是一个适合快速判断我工作风格的入口：

- AI developer tooling
- 内部 AI workflow enablement
- local-first analytics
- agent security

我更偏好务实合作：先看项目、再看取舍、最后再谈落地方式。DevLab 的目标就是把这些判断信息提前公开出来。
