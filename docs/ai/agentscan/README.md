# AgentScan

[中文](./README.zh-CN.md)

- Repository: [AutoScan/agentscan](https://github.com/AutoScan/agentscan)
- Related notes: [AgentScan engineering notes](../cookbook/agentscan-agent-security-notes/README.md)

## My Role

Contributor and maintenance collaborator on an AI agent asset discovery and security audit platform.

## Problem It Solves

AI agent systems are increasingly deployed with weak defaults, broad host exposure, and limited operator awareness. Generic asset scanners can find open services, but they do not explain agent-specific fingerprints, workflows, or risk. AgentScan treats exposed AI agents as a concrete security surface that needs inventory, verification, reporting, and operational follow-up.

## Stack

Go backend and CLI, Gin API, GORM persistence, React frontend, live updates through WebSocket, and a task, alert, and reporting layer around the scanning pipeline.

## Current Stage

Active platform build. The current direction already spans layered scanning, dashboarding, scheduled tasks, alerts, and reporting. The next step is deeper hardening, broader signatures, and better scale characteristics for real-world scanning environments.

## Why It Is Worth Watching

AgentScan sits at the intersection of AI systems and practical security operations. That makes it interesting beyond the immediate product surface: it is also a case study in how AI-native infrastructure becomes governable once it is treated as an asset class instead of a demo stack.
