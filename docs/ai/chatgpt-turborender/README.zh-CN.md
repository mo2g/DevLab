# ChatGPT TurboRender

[English](./README.md)

- 仓库： [mo2g/ChatGPT-TurboRender](https://github.com/mo2g/ChatGPT-TurboRender)
- 相关文档：
  - [架构说明](https://github.com/mo2g/ChatGPT-TurboRender/blob/main/docs/architecture.zh-CN.md)
  - [受控 Chrome Cookbook](https://github.com/mo2g/ChatGPT-TurboRender/blob/main/docs/cookbook-controlled-chrome.zh-CN.md)

## 我的角色

Chromium 优先的 ChatGPT 长对话响应性能扩展的发起者 / 维护者。

## 它解决的问题

超长 ChatGPT 会话很容易把浏览器拖慢：DOM 节点越来越多，流式输出不断触碰已经很重的节点树，滚动开始变卡，输入延迟也会升高。TurboRender 通过裁剪冷历史、保留热区、按需恢复旧消息，把这类渲染压力降下来。

## 技术栈

TypeScript、WXT、Manifest V3、Playwright、Vitest，以及 Chromium / Edge 扩展打包流程。

## 当前阶段

Active build。仓库里已经有核心的 parking / restore 逻辑、架构说明、浏览器打包流程和扩展级测试覆盖，接下来主要是扩展更多 DOM 变化场景下的鲁棒性，以及补充更真实的长对话样本。

## 为什么值得关注

TurboRender 把 ChatGPT 卡顿当作渲染压力问题，而不是 prompt 问题。这使它适合作为浏览器侧 AI 工具的参考案例：本地优先、可逆、并且对宿主 DOM 保持保守。
