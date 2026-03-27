# Token Insight

[中文](./README.zh-CN.md)

- Repository: [mo2g/token-insight](https://github.com/mo2g/token-insight)
- Related notes: [Token Insight local-first analytics notes](../cookbook/token-insight-local-first-analytics/README.md)

## My Role

Builder and maintainer of a local-first token analytics stack for AI coding tools.

## Problem It Solves

AI coding usage data is usually fragmented across local artifacts, vendor-specific logs, and ephemeral pricing assumptions. Token Insight is meant to turn that scattered data into a usable system for inspection, filtering, trend analysis, export, and sharing without making cloud upload the default.

## Stack

Rust backend and CLI, React plus Vite frontend, SQLite storage, and helper scripts for refresh, export, and share image generation.

## Current Stage

Active build. The current shape already covers ingestion, normalization, dashboard views, and operational scripts. The next layer is better adapter coverage, stronger fixtures, and more polished share and release workflows.

## Why It Is Worth Watching

Most AI tooling discussions focus on model quality or prompt quality. Token Insight focuses on observability instead: where token usage comes from, how it changes over time, and how to keep that visibility local-first. That makes it useful both as a product idea and as an engineering pattern for privacy-aware analytics.
