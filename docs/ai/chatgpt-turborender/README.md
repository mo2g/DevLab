# ChatGPT TurboRender

[中文](./README.zh-CN.md)

- Repository: [mo2g/ChatGPT-TurboRender](https://github.com/mo2g/ChatGPT-TurboRender)
- Related notes:
  - [Architecture Notes](https://github.com/mo2g/ChatGPT-TurboRender/blob/main/docs/architecture.md)
  - [Controlled Chrome Cookbook](https://github.com/mo2g/ChatGPT-TurboRender/blob/main/docs/cookbook-controlled-chrome.md)

## My Role

Builder and maintainer of a Chromium-first extension that keeps long ChatGPT threads responsive without replacing the native UI.

## Problem It Solves

Very long ChatGPT sessions can turn the browser sluggish: DOM trees grow large, streamed output keeps touching already heavy nodes, scrolling gets sticky, and input latency climbs. TurboRender reduces that pressure by trimming cold history before first render, preserving a hot interaction window, and restoring older turns on demand.

## Stack

TypeScript, WXT, Manifest V3, Playwright, Vitest, and Chromium / Edge extension packaging.

## Current Stage

Active build. The repo already covers the core parking and restore logic, architecture notes, browser packaging, and extension-level test coverage. The next layer is broader DOM-variant hardening and more real-world long-thread cases.

## Why It Is Worth Watching

TurboRender treats ChatGPT slowdown as a rendering-pressure problem, not a prompt problem. That makes it a useful reference for browser-side AI tooling that needs to stay local-first, reversible, and conservative about the host DOM.
