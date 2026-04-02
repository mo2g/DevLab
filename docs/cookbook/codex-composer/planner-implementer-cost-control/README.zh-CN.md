# Planner / Implementer Cookbook（成本控制版）

[English](README.md)

关联项目：[mo2g/codex-composer](https://github.com/mo2g/codex-composer)

## 问题背景

周额度还剩20%的时候，我手动把模型从 gpt-5.4 或者 gpt-5.3-codex 切换到 gpt-5.4-mini 来修复明确的bug任务，体验下来就更加强化了这个想法：可以用更少的额度就能完成这类任务。

当预算需要被认真控制时，最容易浪费成本的做法，就是让最强的模型去做所有事情。对 `codex-composer` 来说，更合理的做法是把高成本推理集中在“规划”和“范围控制”上，把大量边界明确的执行交给更便宜、更快的模型。

这份 cookbook 记录的是一条比较务实的分工线：

* 让 `gpt-5.3-codex` 或 `gpt-5.4` 担任 planner
* 让 `gpt-5.4-mini` 担任 implementer
* 默认采用 `1 planner -> 1 implementer`
* 只有在 handoff 已经稳定、任务可分离、花费仍然可控时，才扩展到 `1 planner -> n implementers`

这不是为了“多 agent”，而是为了让每一轮花费更接近真正产生价值的地方。

仓库当前已经有了比较好的骨架：

* `planner -> implementer -> change-check`
* 人工 review 和人工 merge
* 默认留在当前线程
* 只有在独立审查更清楚时，才拆线程或使用 worktree

这份 cookbook 要补强的，是**模型角色分工、fan-out 准入条件和预算门槛**。

## 设计决策

### 角色定义

| 角色 | 默认模型 | 何时升级 | 主要职责 | 必须避免 |
| --- | --- | --- | --- | --- |
| Planner | `gpt-5.3-codex` | 需求歧义高、跨子系统、风险高时升级到 `gpt-5.4` | 澄清意图、控制范围、拆分 ticket、定义验证路径 | 直接写大补丁，或在范围不清时盲目 fan-out |
| Implementer | `gpt-5.4-mini` | 默认不升级；如果 ticket 已经被切得足够好，通常不需要更强模型 | 执行一个边界明确的 ticket，修改代码、测试，以及 ticket 明确包含的文档 | 擅自重写规划、扩大范围、重新定义需求 |
| Change-check | 默认 `gpt-5.3-codex` | diff 很大或涉及高风险边界时可升级到 `gpt-5.4` | 审查 diff、确认验证、指出剩余风险 | 悄悄接受越界改动，或只看表面测试结果 |

### 基本原则

* 成本控制的基本单位是 **ticket**，不是整段会话。
* 如果任务边界还不清楚，先升级 planner，而不是直接增加 implementer 数量。
* 如果 `gpt-5.4-mini` 反复需要返工，优先怀疑 ticket 切得不够好，而不是立刻把 implementer 升级成更贵的模型。
* 5 小时和 1 周剩余额度只应该当作**操作性门槛**，不应该被写成“官方配额”。
* 讨论记录里的估算只能当作经验参考，不能当作稳定事实。

### Planner -> Implementer 交接格式

建议 planner 和 implementer 之间使用一个短而明确的 handoff block：

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

这个 handoff 的目标是：

* 足够短，避免交接本身消耗太多成本
* 足够具体，避免 `gpt-5.4-mini` 自行发明策略
* 足够边界化，让 implementer 知道什么时候该停下来回交 planner

### 预算门槛

下面这些阈值是**运营启发式**，不是官方硬指标：

| 状态 | 5 小时剩余 | 1 周剩余 | 推荐拓扑 | Planner 策略 |
| --- | --- | --- | --- | --- |
| 绿色 | `>= 70%` | `>= 70%` | `1 -> 1`，在条件很干净时可选择性 `1 -> 2` 或 `1 -> n` | 默认 `gpt-5.3-codex`，只有在歧义或风险明显上升时才切到 `gpt-5.4` |
| 黄色 | `40% - 69%` | `40% - 69%` | 以 `1 -> 1` 为主 | 优先 `gpt-5.3-codex`，fan-out 只留给真正独立的小切片 |
| 红色 | `< 40%` | `< 40%` | 只保留 `1 -> 1`，避免 fan-out | 除非失败代价明显大于花费，否则不要升级到更重的规划 |

### 风险门槛

以下场景默认不建议直接 fan-out：

* 受影响文件重叠明显
* 验证路径不清晰
* 一个 ticket 需要另一个 ticket 先完成才能验证
* planner 还没有把 merge 顺序或独立性说清楚
* 任务看起来可以并行，但实际上只是把协调成本挪到了最后

## 任务分级

### T0 - trivial fix

例如：

* 拼写修正
* 注释调整
* 一行 import 修复
* 显而易见的 snapshot 更新

策略：

* planner 可选
* implementer 可以直接做
* 如果行为发生变化，仍建议走 `change-check`

### T1 - 小型有边界任务

例如：

* 一个模块内的 bug fix
* 一个局部校验分支
* 一组聚焦测试

策略：

* planner：`gpt-5.3-codex`
* implementer：`gpt-5.4-mini`
* 拓扑：`1 -> 1`

### T2 - 中型有边界任务

例如：

* 少量相关文件上的一个功能切片
* 一个局部重构并补周边测试
* 一个 endpoint + service wiring

策略：

* planner：默认 `gpt-5.3-codex`
* 如果歧义较高，升级到 `gpt-5.4`
* implementer：`gpt-5.4-mini`
* 拓扑：先用 `1 -> 1`，只有在稳定后才考虑 `1 -> 2`

### T3 - 跨子系统任务

例如：

* backend + client + docs
* schema + service + compatibility layer
* middleware / handler / permission 联动

策略：

* planner：`gpt-5.4`
* implementer：`gpt-5.4-mini`
* 拓扑：先从 `1 -> 1` 开始，只在真正独立时再扩展

### T4 - 架构或安全敏感任务

例如：

* 深度重构
* 迁移策略设计
* 持久化边界重构
* 并发、隔离、权限、部署链路风险较高的改动

策略：

* planner：`gpt-5.4`
* 必须显式输出风险分析
* 实施前建议有一次人工 checkpoint
* 不要过早 fan-out
* `1 -> n` 只应作为例外情况

## 踩坑

* 很容易把 `gpt-5.4` 当成默认 planner，但这样通常会把预算浪费在常规任务上。
* 如果 implementer 开始承担重新规划的责任，说明 ticket 切分已经失效。
* 只看“能不能并行”，不看“是否独立验证”，最后通常会把协调成本转移到 review 阶段。
* 把讨论记录里的估算直接写成硬配额，会让文档看起来更确定，但实际上更误导。
* 如果 change-check 只跑已有测试，不补直接覆盖，新行为很容易漏掉。

## 权衡

* `gpt-5.3-codex` 和 `gpt-5.4` 之间，我更偏向前者作为默认 planner，因为更强的模型应该只在歧义或风险真的升高时再用。
* `gpt-5.4-mini` 和更强 implementer 之间，我更偏向前者作为默认执行器，因为 implementer 的价值在于便宜、快、边界清楚。
* `1 -> 1` 和 `1 -> n` 之间，我更偏向先把 `1 -> 1` 跑稳，再扩展并行。
* current thread 和 worktree 之间，我更偏向 current thread first，因为可见性更高，协调成本更低。
* 固定“配额表”和“操作性门槛”之间，我更偏向后者，因为这类限制会变化，而模型分工原则更稳定。
* 如果 `gpt-5.4-mini` 开始产生大量返工，正确动作通常不是盲目升级 implementer，而是缩小 ticket 或提升 planner 的拆分质量。

## 可复用经验

* 成本控制的关键不是“少用模型”，而是“把贵的推理只花在最值得花的地方”。
* planner 的工作是减少歧义、控制风险、切 ticket，不是替 implementer 直接完成实现。
* implementer 的工作是把已定义好的 ticket 做出来，不是重新解释需求。
* fan-out 只有在文件边界、验证边界和责任边界都清楚时才真的有价值。
* 如果一个工作流不能说清楚人在什么地方接管判断，它就还不够适合真实使用。

## 下一步

* 在真实任务里验证 `1 -> 1` 的稳定性，再决定是否引入 `1 -> 2`。
* 统计 planner、implementer、change-check 的返工率，而不是只看并行数量。
* 把更好的 handoff 模板沉淀成仓库级约定。
* 继续观察不同模型在真实任务里的边界变化，再调整默认策略。
