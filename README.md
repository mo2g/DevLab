# DevLab

[中文主页](./README.zh-CN.md)

![DevLab AI Lab overview](./docs/assets/devlab-preview.jpg)

DevLab is the public entry point for AI work I build, maintain, or actively contribute to around four themes: AI coding workflow, token analytics, agent security, and ChatGPT workflow performance.

## Start Here

- [Chinese overview](./README.zh-CN.md)
- [Documentation](./docs/README.md)
- [AI project index](./docs/ai/README.md)
- [Cookbook](./docs/cookbook/README.md)

## About

I use this repository as a documentation-first front page for current AI projects and as a durable home for earlier multi-language experiments. The goal is practical: show what I am building, what I am maintaining, and how I approach AI developer tooling in real repositories.

## Featured Projects

| Project | Role | Domain | Stack | Status |
| --- | --- | --- | --- | --- |
| [Token Insight](./docs/ai/token-insight/README.md) / [Repo](https://github.com/mo2g/token-insight) | Builder / Maintainer | Local-first token observability for AI coding tools | Rust, React, SQLite | Active build |
| [Codex Composer](./docs/ai/codex-composer/README.md) / [Repo](https://github.com/mo2g/codex-composer) | Builder / Maintainer | Reproducible Codex workflow bootstrap for repositories | JavaScript, Shell, Markdown | Active iteration |
| [ChatGPT TurboRender](./docs/ai/chatgpt-turborender/README.md) / [Repo](https://github.com/mo2g/ChatGPT-TurboRender) | Builder / Maintainer | Keep long ChatGPT conversations responsive without replacing the native UI | TypeScript, WXT, Manifest V3, Playwright, Vitest | Active build |
| [AgentScan](./docs/ai/agentscan/README.md) / [Repo](https://github.com/AutoScan/agentscan) | Contributor / Maintainer | Exposed AI agent discovery and security audit | Go, React, SQLite | Active security tooling |

### ChatGPT TurboRender

- Role: Builder / Maintainer
- Domain: Browser-side ChatGPT long-thread responsiveness
- Stack: TypeScript, WXT, Manifest V3, Playwright, Vitest
- Status: Active build

This project keeps long ChatGPT sessions responsive by reducing browser render pressure instead of replacing the native UI. It trims cold history before first render, preserves a hot interaction window, and restores old turns on demand so long conversations stay usable.

View details:

- [Project page](./docs/ai/chatgpt-turborender/README.md)
- [Chinese documentation](./docs/ai/chatgpt-turborender/README.zh-CN.md)
- [GitHub repository](https://github.com/mo2g/ChatGPT-TurboRender)

## Related AI Work

- [one-api](https://github.com/mo2g/one-api): maintenance-oriented fork work around LLM gateway and API distribution scenarios.
- [MetaGPT](https://github.com/mo2g/MetaGPT): participation via fork-based exploration of multi-agent software workflows.

## Cookbook

- [Cookbook Index](./docs/cookbook/README.md)
- [How I Think About Codex Workflow Automation](./docs/cookbook/codex-workflow-automation-tradeoffs/README.md)
- [Codex Composer development notes](./docs/cookbook/codex-composer-workflow/README.md)
- [Token Insight local-first analytics notes](./docs/cookbook/token-insight-local-first-analytics/README.md)
- [AgentScan engineering notes](./docs/cookbook/agentscan-agent-security-notes/README.md)

## Legacy Lab

Earlier experiments remain available as reference code:

- [golang](./golang)
- [java](./java)
- [js](./js)
- [php](./php)
- [python](./python)

## Collaboration

I am interested in collaborations around AI developer tools, internal AI workflow enablement, local-first analytics, and agent security. DevLab is structured to make that evaluation quick: projects, notes, and older experiments are all linked from one place.
