# Codex Composer Development Notes

[中文](README.zh-CN.md)

Related project: [mo2g/codex-composer](https://github.com/mo2g/codex-composer)

## Problem

When teams start using Codex, the real leverage rarely comes from one prompt. It comes from a growing set of implicit repository habits: when to plan first, when to implement directly, where verification commands live, what must stay human-reviewed, and which knowledge should persist in the repo. If that only lives in chat history, it is hard to reuse and hard to align across people.

The core problem behind Codex Composer is not "how to write a smarter prompt" but "how to make AI coding workflow portable at repository scope."

## Design Decisions

- Put the repository contract in `AGENTS.md` so rules exist before tasks.
- Centralize verification commands in `.codex/config.toml` so test entrypoints do not need to be re-explained.
- Keep only a few high-value capabilities such as planning, implementation, and merge readiness instead of creating a large protocol surface.
- Use templates and an install script for bootstrap rather than turning every target repo into a new platform.
- Keep human review and human merge explicit so delivery responsibility stays clear.

## Pitfalls

- It is easy to mistake "reusable" for "more configuration," which makes the template harder to adopt than the target repo itself.
- If the repository contract is too abstract, AI appears to follow process while skipping verifiable work.
- Over-relying on a single-thread workflow raises collaboration cost when tasks are naturally separable.
- If installation rewrites too much of the README or directory structure, users read it as takeover instead of enablement.

## Tradeoffs

- Between lightweight templates and heavily constrained process, I prefer the former because repositories differ more than people expect.
- Between automatic merge and manual merge, I choose manual merge because responsibility boundaries need to stay explicit.
- Between many mediocre skills and a few high-quality skills, I prefer a smaller set that solves frequent problems well.
- Between telling AI what to do and shaping a repository AI can understand, the second approach is more durable.

## Reusable Lessons

- Reusable AI workflow starts with repository contracts, not task prompts.
- Verification commands should always have a stable entrypoint instead of living in memory.
- Template value comes from removing repeated decisions, not from adding new learning cost.
- If an AI workflow cannot explain where humans should step in, it is not ready for serious team use.

## Next Steps

- Improve bootstrap ergonomics so adoption friction is lower in existing repositories.
- Add stronger example repositories and migration notes to reduce first-use cost.
- Clarify the reasoning behind the design, not only the file inventory.
- Keep distilling cross-repository lessons instead of expanding too early into a broad framework.
