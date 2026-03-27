# AgentScan Engineering Notes

[中文](./README.zh-CN.md)

Related project: [AutoScan/agentscan](https://github.com/AutoScan/agentscan)

## Problem

AI agent deployment is spreading quickly, but many operators do not have a traditional security engineering background. The result is predictable: the service starts, ports are open, default configs stay in place, access control is weak, and operations teams may not even know the instances are exposed. Traditional scanners can say "there is a port here," but they usually cannot answer "is this an AI agent, what can it do, and how serious is the risk?"

AgentScan matters because it treats AI agents as a distinct asset class instead of lumping them in with generic web services or middleware.

## Design Decisions

- Split the scan flow into L1, L2, and L3 so discovery, identification, and verification stay separable.
- Go beyond asset discovery and include task scheduling, alerting, reporting, and frontend visibility so results can be operated.
- Shape the backend around the scan pipeline while the frontend is organized around risk interpretation and task state.
- Keep configuration precedence explicit so the path from development to more formal deployment is predictable.
- Maintain a platform mindset for security checks instead of writing one-off scripts for a single agent product.

## Pitfalls

- Over-aggressive fingerprinting creates false positives, which then contaminate downstream vulnerability judgments.
- Vulnerability validation without strong boundaries can drift from "inspection" into "target disturbance."
- Once tasks, alerts, and reporting are added, complexity grows far beyond a simple CLI scanner.
- Security users rarely stop at detection; they also need to know why the finding matters and what to do next.

## Tradeoffs

- Between scan speed and identification accuracy, some restraint is necessary because not every target tolerates aggressive concurrency.
- Between broader agent coverage and deeper support for fewer types, the balance has to be revisited continuously.
- Between a strong CLI experience and a platform product, both automation use and visual operations need support.
- The depth of verification cannot be judged only by technical possibility; default behavior also needs to remain conservative.

## Reusable Lessons

- If a new system spreads quickly, it will likely become a new asset surface and attack surface just as fast.
- Layered scan architecture is easier to maintain and easier to explain than a single all-in-one probe.
- Security tools earn trust through conservative defaults, traceable outcomes, and clear remediation language.
- In AI security products, technical detection and operator-facing communication matter equally.

## Next Steps

- Improve the maintainability of fingerprint and vulnerability data so rule updates can keep pace with the ecosystem.
- Keep strengthening permissions, rate limits, observability, and health checks as platform capabilities.
- Optimize concurrency and scheduling strategies for larger-scale scan environments.
- Make reports and alerts align more closely with real remediation workflows instead of acting as plain exports.
