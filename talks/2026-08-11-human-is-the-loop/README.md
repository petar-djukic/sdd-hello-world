# The Human Is the Loop

Slides and demo runbook for the 2026-08-11 Ottawa ML/AI meetup talk.
The talk demos this repository live.

## Serve the deck

```bash
cd talks/2026-08-11-human-is-the-loop
present -notes -http 127.0.0.1:3999
```

Open http://127.0.0.1:3999/human-is-the-loop.slide and press `N` for the
speaker-notes window — the per-slide cues, gates, and fallbacks live there.
The server must run from this directory (the deck references `figures/` and
the prompt cards by relative path).

No `present` on the machine:

```bash
go install golang.org/x/tools/cmd/present@latest
```

## Terminal in the browser

The whole talk runs from one browser window — deck, terminal, and GitHub
as tabs. The terminal tab is ttyd:

```bash
ttyd -W zsh
```

Open http://127.0.0.1:7681 (`-W` makes it writable). Run the demo commands
there; alt-tab never leaves the browser.

## Pre-talk checklist

1. `gh auth status` — logged in as petar-djukic.
2. `git -C ~/GITHUB/sdd-hello-world status -sb` — clean, on main.
3. `gh issue list` — standing issues only (#22, #23, #77, and the current
   reset instance). Anything else: run the reset first.
4. Confirm the reset instance number on the Skill 4 live slide and in
   `prompt-reset.txt` matches the open reset issue.
5. Pre-load browser tabs: the deck (:3999), the ttyd terminal (:7681),
   this repo's issues page, the recurring upgrade issue (#77), and
   Nokia-Bell-Labs/declarative-agents/issues.

## What's here

| File | Purpose |
|---|---|
| `human-is-the-loop.slide` | the deck (go-present markdown) |
| `DEMO-PLAN.md` | the runbook: beat table, rehearsal findings, reset procedure |
| `prompt-make-work.txt` | Skill 1 live card — plan the backlog |
| `prompt-pop-work.txt` | Skill 2 live card — pop and do the work |
| `prompt-opencode.txt` | Skill 3 live card — the tool switch |
| `prompt-reset.txt` | Skill 4 live card — reset via the recurring issue |
| `chat-vs-loop.html`, `loop-pane.html` | two-pane slide layouts (`.html` includes) |
| `qr-overview.html`, `qr-takehome.html` | QR placement (meshintelligence.substack.com) |
| `figures/` | d2 sources and rendered SVGs |

The four prompt cards mirror the live slides one to one — what is on screen
is what gets typed. Beat-by-beat flow, timings, and failure insurance:
[DEMO-PLAN.md](DEMO-PLAN.md).

## Write-ups

Each of the four skills on stage has a longer written treatment:
[How to Loop Engineering](https://meshintelligence.substack.com/p/how-to-loop-engineering?utm_source=github&utm_campaign=sdd-hello-world),
[How to Use GitHub as Long-Term Memory for Coding Agents](https://meshintelligence.substack.com/p/how-to-use-github-as-long-term-memory?utm_source=github&utm_campaign=sdd-hello-world),
[How to Use Git Worktrees with Coding Agents](https://meshintelligence.substack.com/p/how-to-use-git-worktrees-with-coding?utm_source=github&utm_campaign=sdd-hello-world),
and
[How to Use GitHub to Give a Coding Agent Recurring Work](https://meshintelligence.substack.com/p/how-to-use-github-to-give-a-coding?utm_source=github&utm_campaign=sdd-hello-world).
The talk itself is written up as
[Spec-Driven Development with GitHub, Claude, and GLM on OpenCode](https://meshintelligence.substack.com/p/spec-driven-development-with-github?utm_source=github&utm_campaign=sdd-hello-world),
and the tool switch in Skill 3 as
[How to Code with GLM 5.2 on OpenCode](https://meshintelligence.substack.com/p/how-to-glm-52-on-opencode?utm_source=github&utm_campaign=sdd-hello-world).
