# Demo plan: sdd-hello-world

Repo: https://github.com/petar-djukic/sdd-hello-world
App: minimal Go hello-world binary wrapped in the full SDD structure
(VISION, ARCHITECTURE, road-map, SRDs, use cases, constitutions).
Chosen over a purpose-built demo repo because its history is real:
77 issues, defect storms, and a recurring upgrade issue with eight
logged runs since March.

## Why this repo carries the talk

- **Real work queued.** road-map.yaml has rel02.0 pending (generate the
  hello-world binary from srd001), rel03.0 pending (CLI flags, two use
  cases), rel04.0 pending (config support, spare). A live /make-work
  proposes real work, not staged work.
- **Real recurring history.** Issue #77 "Recurring: Upgrade to latest
  cobbler-scaffold" carries eight Previous Runs entries (Mar 6 → Apr 6).
  Closing it live spawns the next instance — Skill 5 with receipts.
- **Real memory history.** The tracker holds actual defect cascades
  (#37–#51, schema errors filed by the pipeline) — Skill 4's claims
  point at scrollable evidence.
- Commands are already in .claude/commands/ (symlinked from
  coding-skills): make-work, gh-issue-push, gh-issue-pop, do-work,
  gh-release-push. .opencode is wired for the tool-switch beat.

## The demo queue (from road-map.yaml, worked live)

Any prefix of the queue is a coherent demo:

1. **rel02.0-uc001 — build hello world from spec** — the opener.
   Generates cmd/sdd-hello-world/{main.go,version.go} from
   srd001-hello-world-binary. Smallest round, whole loop visible.
2. **rel03.0 — CLI enhancements** — two use cases (uc002
   greeting-and-version, uc003 help-and-errors), so the pop beat shows
   multiple sub-issues on one worktree branch. The "second full round"
   if time allows.
3. **Release beat** — after rel02.0 (or rel03.0) merges,
   /gh-release-push tags the release; run the binary live.
4. **Recurring beat** — close #77, watch the next instance open with a
   dated Previous Runs line. Do NOT work the spawned upgrade issue live;
   the close/spawn is the demo.
5. **rel04.0 — config support** — spare; mention, don't run.

## Flow of the talk demo

| Beat | Terminal | GitHub screen |
|---|---|---|
| 1 | `/make-work` **+ scope line** → proposes epic + 2 issues | — |
| 2 | approve plan | — |
| 3 | `/gh-issue-push` | read the issue body: files, criteria, LOC |
| 4 | `/gh-issue-show`, then `/gh-issue-pop <epic>` | parent issue: sub-issues appear |
| 5 | `git worktree list` | branches page: the new remote branch |
| 6 | `/do-work` ×2 (one per sub-issue) | PR: diff, checks |
| 7 | merge PR | sub-issues close |
| 8 | switch tools: `opencode` in the worktree, `/do-work` | — |
| 9 | (time) second round: rel03.0, two use cases | — |
| 10 | (time) `/gh-release-push` | releases page; run the binary |
| 11 | close #77 (recurring upgrade) | next instance opens; Previous Runs log |
| 12 | `/gh-issue-pop 89`, `/do-work` — the reset | next reset instance opens |

The opencode switch (beat 8) is the riskiest beat: rehearse it first. If
opencode cannot resume cleanly, fallback is narrating the markdown command
file + GitHub state ("everything it needs is here, nothing was in the tool").
Beat 8 needs an open sub-issue to resume onto — do it before the last
`/do-work` pass, or narrate instead.

## What the 2026-08-10 rehearsal established

Ran the full loop against sdd-hello-world. Results:

- **`/make-work` takes no `$ARGUMENTS`.** Its sizing rule (300–700 production
  lines, up to 10 tasks) is written for real features; hello-world is nine
  lines of Go, so unscoped it proposes either one lumpy task or work from
  rel03.0/rel04.0. The scope line in `prompt-make-work.txt` is what produces
  the epic-plus-two shape. Type it with the command.
- **Worktrees land at `../gh-<number>-<slug>`** — a sibling of the repo
  directory, not `../sdd-hello-world-gh-N`.
- The rel02.0 split is spec-honest: srd001 AC3 and AC4 are "main.go exists"
  and "version.go exists", so one issue per file traces cleanly.
- Artifacts from the first run: epic **#80**, sub-issues **#81**/**#82**,
  PR **#83**. All cleared by reset run **#79** (2026-08-10, verified end to
  end: issues closed not-planned, PR closed unmerged, branch and worktree
  pruned, recurrence spawned **#84**). The re-run filed epic **#85** with
  sub-issues **#86**/**#87** — the queue that is popped-ready for the talk.
- All four acceptance criteria passed live: `go build` exit 0, binary prints
  `Hello, World!` and exits 0, both files present, `go.mod` dependency-free.

## Reset between rehearsals

The reset is itself a recurring issue —
[#89 "Recurring: Reset demo to clean slate"](https://github.com/petar-djukic/sdd-hello-world/issues/89)
(third instance; runs #79 and #84 verified, including the revert path; also the Skill 4 live beat). Run it after each
rehearsal: `/gh-issue-pop 89`, then `/do-work`. The contract: revert demo
commits on main, restore road-map.yaml statuses, close demo issues, prune
worktrees and `gh-*` branches. Closing it spawns the next instance, and
every rehearsal lands a Previous Runs line.

- Verify after reset: no `cmd/` tree, and `git worktree list` shows only
  the main checkout.
- If #77 (upgrade) was closed in rehearsal, the spawned instance becomes
  the live one to close.

Pre-rehearsal baseline SHA for sdd-hello-world: `d7bfab5`

## Failure insurance

- Rehearse today (Aug 10); restore per the reset list above.
- If live generation stalls: the rehearsal PR is in the history —
  narrate it from GitHub instead ("this is yesterday's run").
- Ollama/network not needed: Claude Code + gh only.
- `gh auth status` verified (petar-djukic, keyring) and clone at
  ~/GITHUB/sdd-hello-world is clean on main.
- Have the #77 issue page pre-loaded in a browser tab — it works as a
  standalone exhibit even if nothing live runs.

## Slides

- Deck: `human-is-the-loop.slide` (go present, legacy syntax).
- Install: `go install golang.org/x/tools/cmd/present@latest`,
  run `present` in this directory, open the printed URL.
- Figures: `figures/` — article SVGs (commands, worktrees, recurring
  loop, memory-vs-CI) plus two talk-only charts (specdd, issue-pr).

## Link sweep

Done 2026-08-10: all six take-home article links verified against minted
Substack slugs; demo-repo link updated to sdd-hello-world.
