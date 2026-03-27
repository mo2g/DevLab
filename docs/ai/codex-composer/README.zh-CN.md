# Codex Composer

[English](./README.md)

- Repository: [mo2g/codex-composer](https://github.com/mo2g/codex-composer)
- 相关笔记：
  - [Codex Composer 开发经验](../../cookbook/codex-composer/codex-composer-workflow/README.zh-CN.md)
  - [我如何思考 Codex 工作流自动化](../../cookbook/codex-composer/codex-workflow-automation-tradeoffs/README.zh-CN.md)
  - [Codex Composer 轻量标准工作流收敛](codex-composer/codex-composer-lightweight-workflow/README.zh-CN.md)

## 我的角色

一个轻量级 starter kit 的发起者与维护者，用于把 Codex 工作流基础能力接入真实仓库。

## 它解决什么问题

大多数团队开始使用 AI coding 时，依赖的是重复的 prompt 习惯，而不是仓库契约。这会导致规划规则、验证命令、merge 边界和交接预期都停留在隐含状态。Codex Composer 把这些关键部分封装进一个轻量 bootstrap 中，让真实仓库可以在不引入重流程的前提下获得稳定工作方式。

## 技术栈

JavaScript、shell 脚本、Markdown 契约文件、模板资源、安装逻辑，以及 quickstart 与 merge guidance 等仓库级资产。

## 当前阶段

Active iteration。当前已经支持仓库 bootstrap 和最小任务流，但最核心的价值仍然在于继续优化 ergonomics、示例质量，以及安装后仓库契约的清晰度。

## 为什么值得关注

Codex Composer 不只是一个模板仓库，它在尝试把 AI coding 行为沉淀为仓库级基础设施。如果这层做稳，团队就可以不再把“好 prompt”当作个人经验，而是把 AI workflow 视为工程环境的一部分。
