# sdd-hello-world

A working example of
[spec-driven development with a coding agent](https://meshintelligence.substack.com/p/spec-driven-development-with-github?utm_source=github&utm_campaign=sdd-hello-world),
small enough to read on a projector. Four slash commands drive the whole
cycle in Claude Code or
[opencode](https://meshintelligence.substack.com/p/how-to-glm-52-on-opencode?utm_source=github&utm_campaign=sdd-hello-world):
[GitHub issues carry the memory between sessions](https://meshintelligence.substack.com/p/how-to-use-github-as-long-term-memory?utm_source=github&utm_campaign=sdd-hello-world),
[each task gets its own git worktree](https://meshintelligence.substack.com/p/how-to-use-git-worktrees-with-coding?utm_source=github&utm_campaign=sdd-hello-world),
and a pull request you review closes the loop. It is the demo repository for
the talk
[The Human Is the Loop](talks/2026-08-11-human-is-the-loop/), and the test
fixture cobbler-scaffold scaffolds against in its end-to-end suite.

The specifications come first. There is no `cmd/` tree at rest — the Go
binary is what the agent generates from `docs/specs/` when you run the demo.

## Try the demo yourself

### 1. Fork it

**Fork this repository to your own account first.** The demo files issues,
pushes branches, opens and merges pull requests, and closes issues. All of
that needs write access, and you have none here. Working against this
account, `/gh-issue-push` fails at the first step.

```bash
gh repo fork petar-djukic/sdd-hello-world --clone
```

Or fork on the GitHub website, then clone your fork.

### 2. Clone coding-skills beside it

The four commands the demo runs do not live in this repository. `.claude`
and `.opencode` are **relative symlinks** to `../coding-skills/`, so
[coding-skills](https://github.com/petar-djukic/coding-skills) has to sit in
the same parent directory as your clone:

```bash
cd ..                                                        # the parent of your clone
git clone https://github.com/petar-djukic/coding-skills.git
```

You should end up with the two side by side:

```
~/GITHUB/
  sdd-hello-world/     <- your fork
    .claude    -> ../coding-skills/.claude
    .opencode  -> ../coding-skills/.opencode
  coding-skills/       <- the commands
```

### 3. Check that the commands resolved

```bash
cd sdd-hello-world
ls .claude/commands/
```

You should see `make-work.md`, `gh-issue-push.md`, `gh-issue-pop.md`, and
`do-work.md` among others. An empty listing or "No such file or directory"
means the symlink is dangling — `coding-skills` is not where the link
expects it.

If you cloned `coding-skills` somewhere else and would rather not move it,
repoint the links locally:

```bash
ln -sfn /path/to/coding-skills/.claude   .claude
ln -sfn /path/to/coding-skills/.opencode .opencode
```

Leave that change out of your commits — the tracked links are relative on
purpose.

### 4. Run the loop

Open Claude Code (or opencode) in the repository and
[work the cycle](https://meshintelligence.substack.com/p/how-to-loop-engineering?utm_source=github&utm_campaign=sdd-hello-world):

| Command | What it does |
|---|---|
| `/make-work` | reads the specs, verifies the tracker, proposes a batch — you approve |
| `/gh-issue-push` | writes the approved plan as self-contained GitHub issues |
| `/gh-issue-pop <n>` | decomposes an issue into sub-issues on its own git worktree |
| `/do-work` | implements one sub-issue, validates it, commits, opens the PR |

You merge the PR. The agent never merges its own.

Scope `/make-work` when you run it, or it will propose from the whole
roadmap. For the same batch the talk demos:

```
/make-work
scope it to release 02.0: one epic, and two issues under it,
one per source file. Nothing else.
```

The beat-by-beat runbook, including the
[recurring issue](https://meshintelligence.substack.com/p/how-to-use-github-to-give-a-coding?utm_source=github&utm_campaign=sdd-hello-world)
that returns the repository to a clean slate afterwards, is in
[talks/2026-08-11-human-is-the-loop/DEMO-PLAN.md](talks/2026-08-11-human-is-the-loop/DEMO-PLAN.md).

## Why spec-driven development needs a fixture like this

Cobbler-scaffold's end-to-end tests need a stable, minimal Go project to
scaffold against. Pointing them at a production SDD project couples test
stability to ongoing development. This repository is the simplest possible
Go binary wrapped in the full SDD structure, which isolates that validation
from production changes — and makes the whole loop visible in one screen.

## Repository structure

```
sdd-hello-world/
  docs/VISION.yaml            -- goals and boundaries
  docs/ARCHITECTURE.yaml      -- components
  docs/road-map.yaml          -- releases and their use cases
  docs/specs/                 -- SRDs, use cases, test suites
  docs/constitutions/         -- design, testing, interface rules the agent obeys
  docs/prompts/               -- scaffolded prompt templates
  magefiles/orchestrator.go   -- cobbler-scaffold mage targets
  configuration.yaml          -- orchestrator configuration
  talks/                      -- The Human Is the Loop: deck and demo runbook
  .claude, .opencode          -- symlinks to ../coding-skills (see above)
```

`cmd/sdd-hello-world/` appears once release 02.0 is implemented, and the
demo reset removes it again.

## Build and test

Only after the binary has been generated:

```bash
go build -o bin/sdd-hello-world ./cmd/sdd-hello-world
./bin/sdd-hello-world
# Output: Hello, World!
```

With the orchestrator scaffolded:

```bash
mage build
mage init
mage -l
```

## The talk

The talk page lives at
[petar-djukic.github.io/sdd-hello-world](https://petar-djukic.github.io/sdd-hello-world/).
`talks/2026-08-11-human-is-the-loop/` holds the slides, the demo runbook,
and the prompt cards used on stage. Its
[README](talks/2026-08-11-human-is-the-loop/README.md) covers serving the
deck. Every skill the talk demos has a write-up on
[meshintelligence.substack.com](https://meshintelligence.substack.com?utm_source=github&utm_campaign=sdd-hello-world):

| Write-up | What it covers |
|---|---|
| [How to Loop Engineering](https://meshintelligence.substack.com/p/how-to-loop-engineering?utm_source=github&utm_campaign=sdd-hello-world) | getting the model to write its own prompts |
| [How to Use GitHub as Long-Term Memory for Coding Agents](https://meshintelligence.substack.com/p/how-to-use-github-as-long-term-memory?utm_source=github&utm_campaign=sdd-hello-world) | issues as the memory between sessions |
| [How to Use Git Worktrees with Coding Agents](https://meshintelligence.substack.com/p/how-to-use-git-worktrees-with-coding?utm_source=github&utm_campaign=sdd-hello-world) | one workspace per task |
| [How to Use GitHub to Give a Coding Agent Recurring Work](https://meshintelligence.substack.com/p/how-to-use-github-to-give-a-coding?utm_source=github&utm_campaign=sdd-hello-world) | recurring issues and temporal agency |
| [How to Code with GLM 5.2 on OpenCode](https://meshintelligence.substack.com/p/how-to-glm-52-on-opencode?utm_source=github&utm_campaign=sdd-hello-world) | the second tool the demo hands the same issue to |
| [Spec-Driven Development with GitHub, Claude, and GLM on OpenCode](https://meshintelligence.substack.com/p/spec-driven-development-with-github?utm_source=github&utm_campaign=sdd-hello-world) | the whole cycle, demonstrated on this repository |
