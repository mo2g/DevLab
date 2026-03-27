# Token Insight

[English](./README.md)

- Repository: [mo2g/token-insight](https://github.com/mo2g/token-insight)
- 相关笔记：[Token Insight 开发经验](../cookbook/token-insight-local-first-analytics/README.zh-CN.md)

## 我的角色

本地优先 token analytics 项目的发起者与维护者，目标用户是 AI coding tools 的使用者和团队。

## 它解决什么问题

AI coding usage data 往往分散在本地 artifacts、供应商特定日志和不断变化的价格假设里。Token Insight 的目标是把这些分散信息整理成一个真正可用的系统，用于检查、过滤、趋势分析、导出和分享，而且不把云端上传当成默认前提。

## 技术栈

Rust 后端与 CLI，React + Vite 前端，SQLite 存储，以及用于刷新、导出、分享图生成的辅助脚本。

## 当前阶段

Active build。当前版本已经覆盖 ingestion、normalization、dashboard 视图和一组运维脚本，下一步重点是扩大 adapter 覆盖、补强 fixtures，以及打磨分享与发布体验。

## 为什么值得关注

很多 AI 工具讨论聚焦在模型质量或 prompt 质量上，而 Token Insight 关注的是可观测性：token usage 从哪里来、如何随时间变化、以及怎样在保持 local-first 的前提下获得持续可见性。这既是一个产品方向，也是一种偏重隐私边界的 analytics 工程模式。
