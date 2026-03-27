# Token Insight Local-First Analytics Notes

[中文](./README.zh-CN.md)

Related project: [mo2g/token-insight](https://github.com/mo2g/token-insight)

## Problem

Usage data from AI coding tools is naturally fragmented. Different tools scatter token, model, source, time, and cost information across local files, caches, logs, or exported artifacts. When teams want usage analytics, the first instinct is often to plug into an external SaaS, but that immediately creates two problems: privacy boundaries and incompatible data structures across tools.

Token Insight starts from a stronger position: local-first is a product stance, not just a deployment mode. If local usage events are not normalized first, later analysis, sharing, or synchronization never gets a stable foundation.

## Design Decisions

- Build ingestion and normalization before the dashboard so the data model does not get rewritten around early UI decisions.
- Use SQLite by default to keep the system inspectable and runnable on one machine.
- Separate backend collection from frontend visualization so schema evolution is not constrained by UI pace.
- Ship refresh, export, and share-image actions as scripts to lower the barrier for routine use.
- Support multi-source input from the start so the product does not become trapped inside one tool ecosystem.

## Pitfalls

- Usage artifacts from AI tools are not stable, and field names or timestamps can change easily.
- Cost estimation tied directly to live pricing makes historical data harder to explain.
- Once data is visualized, expectations for correctness rise fast, and normalization drift becomes highly visible.
- Local-first improves privacy trust but also means installation, refresh, and debugging need to stay simple or adoption drops.

## Tradeoffs

- Between local-first and multi-user collaboration, I preserve local truth first and consider optional sync later.
- Between SQLite and a more complex analytical store, I prefer the simpler, more portable, more debuggable option.
- Between a universal schema and source-specific detail, I prefer a stable event backbone that can expand over time.
- Between flashy charts and explainable data models, I choose explainability because trust in analytics comes from interpretation.

## Reusable Lessons

- If data sources are messy by nature, define the normalized event model before deciding what the pages should look like.
- Local-first should shape product narrative, deployment, and data boundaries instead of staying a slogan.
- Export is not a side feature; it often determines whether the system enters a real workflow.
- Shareable views matter because analytics products often spread through screenshots and exports before formal adoption.

## Next Steps

- Expand source coverage so more AI coding tools can feed the same view.
- Strengthen parser fixtures and automated verification to reduce regression risk in the adapter layer.
- Keep improving dashboard expression so trends, distribution, and anomalies read faster.
- Explore safer sync and sharing mechanisms without weakening the local-first stance.
