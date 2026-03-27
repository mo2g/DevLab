# Codex Composer Lightweight Workflow Convergence

[中文](README.zh-CN.md)

Related project: [mo2g/codex-composer](https://github.com/mo2g/codex-composer)

## Problem

The useful part of this iteration was not "adding more automation." It was narrowing the default Codex App workflow until it stayed light enough to use every day and explicit enough to review.

The target questions were:

1. what should stay in the default path
2. what should become optional instead of mandatory
3. where should human intent be clarified before planning starts
4. how should verification happen when the repo stack is not known in advance

## Design Decisions

- Keep the default path thread-first.
- Keep split and worktree as optional tools, not as the mental default.
- Keep `planner`, `implementer`, and `change-check` as the only skill surface.
- Let `planner` clarify intent first, then write the bounded implementation plan.
- Let `change-check` inspect the diff, add direct tests when needed, and choose verification dynamically from the real repo.
- Keep merge manual, but do not add a merge checklist layer.
- Treat `.codex/config.toml` as optional repo-owned configuration, not as the default truth source for every target repository.

## Pitfalls

- It is easy to keep a workflow "lightweight" in words while still making the default install feel protocol-heavy.
- A fixed verification file can become a hidden assumption when the target repository stack is not known yet.
- If `planner` skips the intent-clarification step, the rest of the workflow becomes a plan generator instead of a real entry point.
- If `change-check` only runs existing tests, it can miss the more useful step of adding direct coverage for new behavior.

## Tradeoffs

- Between source-repo defaults and target-repo defaults, I now prefer keeping source defaults and making target config optional.
- Between a separate clarification skill and a richer planner, I prefer the richer planner because it keeps the skill surface small.
- Between static verify hooks and dynamic verification, I prefer dynamic verification because it matches unknown target stacks better.
- Between manual merge and a merge checklist, I keep manual merge and drop the checklist because the checklist adds ceremony without changing responsibility.

## Reusable Lessons

- Lightweight workflow is mostly a subtraction exercise.
- The default path should be easy to explain in one breath.
- Repository contracts matter, but only if they stay smaller than the repo they are meant to help.
- Verification should follow the change, not the other way around.
- An AI workflow is healthier when it can state where human judgment starts.

## Next Steps

- Keep the docs and skills aligned with the lightweight default.
- Continue testing whether `planner` can absorb more intent-clarification without becoming heavy.
- Keep pressure-testing dynamic verification on real target repositories.
- Add more cookbook entries only when they capture a stable pattern, not a one-off decision.
