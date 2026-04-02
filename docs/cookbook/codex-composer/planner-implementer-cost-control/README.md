# Planner / Implementer Cookbook (Cost Control)

[中文](README.zh-CN.md)

Related project: [mo2g/codex-composer](https://github.com/mo2g/codex-composer)

## Problem

When my weekly quota dropped to 20% remaining, I manually switched the model from gpt-5.4 or gpt-5.3-codex to gpt-5.4-mini to fix clearly defined bug tasks. From that experience, it further reinforced my view that this kind of work can be completed with much lower quota consumption.

When cost matters, the easiest way to burn budget is to let the strongest model do everything. For `codex-composer`, a more practical split is to spend expensive reasoning on planning and scope control, then let a cheaper and faster model handle the bulk of bounded implementation work.

This cookbook records a pragmatic division of labor:

* use `gpt-5.3-codex` or `gpt-5.4` as the planner
* use `gpt-5.4-mini` as the implementer
* default to `1 planner -> 1 implementer`
* only expand to `1 planner -> n implementers` after the handoff has proven stable, the work is clearly separable, and spend is still under control

The goal is not "more agents."
The goal is to make each unit of spend land closer to the part of the work that actually creates value.

The repository already has a good skeleton:

* `planner -> implementer -> change-check`
* explicit human review and manual merge
* current-thread-first by default
* split threads or use a worktree only when reviewability or isolation clearly improves risk management

What this cookbook adds is a clearer policy for **model roles, fan-out admission, and budget-aware gating**.

## Design Decisions

### Role model

| Role | Default model | When to upgrade | Primary responsibility | Must avoid |
| --- | --- | --- | --- | --- |
| Planner | `gpt-5.3-codex` | Upgrade to `gpt-5.4` when ambiguity, cross-system coupling, or risk is high | clarify intent, bound scope, split tickets, define verification | writing the final patch directly or fanning out blindly |
| Implementer | `gpt-5.4-mini` | No upgrade by default; if the ticket is well-bounded, the mini model is usually enough | execute one bounded ticket, edit code, tests, and explicitly included docs | re-planning architecture or widening scope |
| Change-check | `gpt-5.3-codex` by default | Upgrade to `gpt-5.4` for large or safety-sensitive diffs | inspect diff, verify commands, call out residual risk | accepting scope creep or only checking superficial tests |

### Guiding principles

* The cost-control unit is the **ticket**, not the whole conversation.
* If scope is still unclear, upgrade the planner before adding more implementers.
* If `gpt-5.4-mini` keeps failing because the ticket is too broad, the first fix is usually to shrink the ticket, not to throw a larger model at the same shape of work.
* Use the 5-hour and weekly remaining views as **operational gates**, not as official quotas.
* Any credit estimate from discussion notes should be treated as reference material, not as a stable fact.

### Planner -> Implementer handoff

Use a short, explicit handoff block between planner and implementer:

```text
ticket_id: T1
parent_plan_id: P2026-04-01-001
model_role: implementer
scope:
  include:
    - src/...
    - test/...
  exclude:
    - installer
    - unrelated docs
success_criteria:
  - ...
verification:
  - npm test -- ...
stop_conditions:
  - requires migration
  - touches auth model
  - needs cross-module redesign
deliverables:
  - code change
  - tests
  - short completion report
```

The handoff should be:

* short enough to stay cheap
* specific enough that `gpt-5.4-mini` does not invent policy
* bounded enough that the implementer knows when to stop and return to the planner

### Budget gates

These thresholds are **operational heuristics**, not official platform quotas:

| State | 5h remaining | Weekly remaining | Recommended topology | Planner policy |
| --- | --- | --- | --- | --- |
| Green | `>= 70%` | `>= 70%` | `1 -> 1`, with selective `1 -> 2` or `1 -> n` only when the split is very clean | default to `gpt-5.3-codex`; move to `gpt-5.4` only when ambiguity or risk rises materially |
| Yellow | `40% - 69%` | `40% - 69%` | keep `1 -> 1` as the default | prefer `gpt-5.3-codex`; reserve fan-out for truly independent slices |
| Red | `< 40%` | `< 40%` | stay at `1 -> 1`, avoid fan-out | avoid heavier planning unless the cost of failure is clearly higher than the cost of spend |

### Risk gates

Do not fan out by default when:

* file ownership overlaps in a meaningful way
* the verification path is not clear
* one ticket depends on another ticket being complete before it can be tested
* the planner has not stated merge order or independence clearly
* the task only looks parallel, but the real cost shifts into integration and review

## Task Sizing

### T0 - trivial fix

Examples:

* typo fix
* comment tweak
* one-line import correction
* obvious snapshot update

Policy:

* planner optional
* implementer may proceed directly
* still run `change-check` if behavior changes

### T1 - bounded small task

Examples:

* one bug fix in one module
* one validation branch
* a focused test set

Policy:

* planner: `gpt-5.3-codex`
* implementer: `gpt-5.4-mini`
* topology: `1 -> 1`

### T2 - bounded medium task

Examples:

* one feature slice across a few related files
* one local refactor plus nearby tests
* one endpoint plus service wiring

Policy:

* planner: default to `gpt-5.3-codex`
* upgrade to `gpt-5.4` if ambiguity is high
* implementer: `gpt-5.4-mini`
* topology: start with `1 -> 1`; only consider `1 -> 2` after the pattern has proven stable

### T3 - cross-subsystem task

Examples:

* backend + client + docs
* schema + service + compatibility layer
* middleware / handler / permission work that spans layers

Policy:

* planner: `gpt-5.4`
* implementer: `gpt-5.4-mini`
* topology: start at `1 -> 1`; only expand when the tickets are truly independent

### T4 - architecture or safety-sensitive task

Examples:

* deep refactor
* migration strategy
* persistence boundary redesign
* concurrency, isolation, auth, or deployment-risk work

Policy:

* planner: `gpt-5.4`
* require an explicit risk section
* add a human checkpoint before implementation
* do not fan out early
* treat `1 -> n` as exceptional

## Pitfalls

* It is easy to make `gpt-5.4` the default planner, but that usually burns budget on ordinary work.
* If the implementer starts re-planning the work, the ticket split has already failed.
* Optimizing only for parallelism, without requiring independent verification, usually just moves coordination cost into the review stage.
* Turning discussion estimates into hard quotas makes the document look more certain than it really is.
* If `change-check` only re-runs existing tests and never adds direct coverage, new behavior is easy to miss.

## Tradeoffs

* Between `gpt-5.3-codex` and `gpt-5.4` as planner, I prefer the former by default because the stronger model should be reserved for genuine ambiguity or risk.
* Between `gpt-5.4-mini` and a larger implementer, I prefer the mini model by default because the implementer wins by being cheap, fast, and bounded.
* Between `1 -> 1` and `1 -> n`, I prefer stabilizing `1 -> 1` first and expanding only when the split is genuinely independent.
* Between current-thread-first and worktree-first, I prefer current-thread-first because visibility is better and coordination overhead is lower.
* Between fixed quota thinking and operational gating, I prefer gating because the UI signals change, while the role split remains durable.
* If `gpt-5.4-mini` starts producing a lot of rework, the right response is usually to shrink the ticket or improve planner decomposition, not to blindly upgrade the implementer.

## Reusable Lessons

* Cost control is not about using fewer models; it is about spending expensive reasoning only where it pays back.
* The planner should reduce ambiguity, manage risk, and split tickets, not directly finish the implementation.
* The implementer should build the defined ticket, not reinterpret the requirement.
* Fan-out is only valuable when file boundaries, verification boundaries, and ownership boundaries are all clear.
* A workflow that cannot explain when human judgment takes over is not yet ready for real use.

## Next Steps

* Validate whether `1 -> 1` stays stable on real work before allowing `1 -> 2`.
* Measure planner, implementer, and change-check rework rate instead of only counting parallelism.
* Turn a better handoff block into a repository-level convention.
* Keep observing how model behavior changes on real tasks before adjusting the default policy.
