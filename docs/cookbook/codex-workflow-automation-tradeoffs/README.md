# How I Think About Codex Workflow Automation

[中文](./README.zh-CN.md)

Related project: [mo2g/codex-composer](https://github.com/mo2g/codex-composer)

## Problem

My starting goal was straightforward: take a normal engineering workflow and make Codex fit it cleanly. The ideal path looked like this:

1. capture the requirement
2. analyze the requirement
3. decide whether it can be split into A and B in parallel
4. let A and B develop independently
5. verify both sides independently
6. commit A and B separately
7. integrate AB
8. run final verification
9. merge back to main

At first glance, this sounds like an automation problem. After more testing, I stopped framing it that way. The real question became: what level of automation stays reliable, reviewable, and economical for the kind of user or repository in front of me?

## Evolution Timeline

### Stage 1. Start from the conventional software workflow

My first instinct was to port a standard engineering model into Codex. If parallel A/B work is good for humans, it should also be good for AI-assisted execution. That assumption made sense structurally, but it hid a key issue: human parallelism already depends on ownership boundaries, review discipline, and shared context. Codex does not remove that requirement.

### Stage 2. Subagents looked like the obvious implementation path

After reading the official documentation, subagents appeared to be the most direct way to realize parallel development. If the goal is "A and B move at the same time," subagents naturally look like the native tool for it.

This was the most tempting phase, because the workflow looked elegant on paper: decompose, dispatch, converge.

### Stage 3. Deeper testing exposed defects and hidden cost

After spending more time testing and learning, I became less convinced that subagents should be the default answer. The problems were not only functional. They were operational:

- execution state became harder to inspect
- failures were harder to localize cleanly
- context duplication pushed token spend up
- integration still needed a human-grade review pass

By contrast, Codex App threads felt more reliable and more inspectable. A thread-first approach is less magical, but the state is easier to understand, and that matters in real work.

### Stage 4. The real debate became full automation vs semi-automation

At that point, the important comparison was no longer "threads or subagents" in isolation. The real design problem was how much automation should exist at all.

The useful spectrum became:

- manual
- semi-automated
- supervised parallel execution
- automation-heavy or near-full-automation

This reframed the problem correctly. Subagents are one implementation option inside a larger automation debate, not the definition of progress by themselves.

### Stage 5. My current position is segment-based, not universal

I do not think there is one correct Codex workflow for everyone. The right answer depends on budget, tolerance for ambiguity, need for auditability, repository maturity, and how much control the user wants to keep.

That is why I now prefer segmenting recommendations instead of prescribing one "best" workflow.

## Design Decisions

My current preferred default is conservative:

- current thread first
- split only when work is truly separable
- prefer Codex App thread plus optional worktree before subagents
- keep explicit `verify` and `commit` gates
- keep merge manual

I prefer thread-first for a few practical reasons:

- execution state is more visible
- human review is easier to stage
- there are fewer hidden coordination assumptions
- accidental cost amplification is lower
- it fits most personal users better than automation-heavy orchestration

This does not mean subagents are never useful. It means I do not want the default workflow to depend on a coordination layer that many repositories and users are not ready to operate well.

## Pitfalls

I would not state "subagents are bad" as a universal conclusion. My current view is narrower: they introduce risks that are easy to underestimate.

- coordination opacity makes it harder to see what really happened where
- debugging gets harder when failures cross execution boundaries
- duplicated context can drive token spend up faster than expected
- apparent parallelism can create false confidence about actual throughput
- review burden often moves downstream to the integration stage instead of disappearing

The biggest mistake is to count only parallel execution time and ignore integration, diagnosis, and review cost.

## Tradeoffs

I think about Codex workflow as an automation spectrum rather than a yes-or-no toggle.

### Manual

The human drives decomposition, execution order, verification, and merge. This is slower, but visibility is highest.

### Semi-automated

Codex helps with planning, drafting, implementation, and review preparation, but the human still controls split points, verification, commit gates, and merge. This is my default recommendation for most personal users.

### Supervised Parallel Execution

Parallel threads or worktrees are introduced when the task is cleanly separable and the integration story is already clear. This can be productive, but only if ownership boundaries are explicit.

### Automation-Heavy / Near-Full-Automation

This mode is attractive when speed matters and the operator is willing to accept more ambiguity, more cost variance, and heavier downstream review. I do not think it is broadly safe as a default.

Higher automation only becomes consistently useful when repository contracts, verification, ownership boundaries, rollback discipline, and review habits are already mature. Without that base, automation mostly shifts complexity around.

## Decision Matrix

| User Segment | Budget Layer | Preferred Workflow Shape | Automation Level | Rationale |
| --- | --- | --- | --- | --- |
| Individual user | Around `$20/month` / low budget | Current thread first, explicit split only when necessary | Manual or semi-automated | Usually cost-sensitive, wants visibility, and benefits more from predictable control than aggressive automation |
| Power user | `$100+/month` / high budget | More willing to test automation-heavy flows or richer parallel execution | Supervised parallel to automation-heavy | Can justify higher spend when speed matters more than auditability or cost variance |
| Team or advanced operator | Variable but maturity-dependent | Thread plus worktree, with supervised parallelism when repo contracts are strong | Semi-automated to supervised parallel | Can absorb coordination overhead if review discipline, verification, and rollback habits already exist |

Budget matters, but not only as a pricing number. It is also a proxy for tolerance: tolerance for token spend, tolerance for ambiguity, and tolerance for control loss.

## Reusable Lessons

- Parallelism is a coordination problem before it is an execution problem.
- Visible state beats hidden orchestration for most real work.
- Reliability and reviewability matter more than looking "native" or "fully automatic."
- Automation should be chosen by user segment and repository maturity, not by novelty.
- A workflow that looks elegant in a demo can still be expensive to debug in production use.

## Next Steps

- Keep validating when subagents are worth their cost instead of rejecting them categorically.
- Keep comparing thread-first and subagent-first flows with real tasks rather than abstract diagrams.
- Refine budget-segmented recommendations as more real usage patterns appear.
- Strengthen repository contracts before recommending more aggressive automation defaults.
