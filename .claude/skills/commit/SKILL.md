---
name: commit
description: Write and create a git commit following the AirMouse commit rules. Use when asked to commit, write a commit message, or /commit in this repo.
---

# AirMouse commit

Format:

```
[<branch>] <Verb> <what>
```

The prefix is the current branch name (`git branch --show-current`), which is either:
- the ticket tag, e.g. `AM-12` → `[AM-12] Add flour ratio` (most common),
- or, for branches without a ticket (e.g. epics), a lowercase kebab-case name, e.g. `bakery-menu` → `[bakery-menu] Add menu layout`.

If the branch name is neither (e.g. `main`, uppercase or spaces), warn the user and ask what prefix to use.

## Description rules

- Right after the prefix an imperative verb: Add, Fix, Remove, Update, Rename, Move, Refactor…
- One short line in English, no trailing period. Add a body only if the "why" is not obvious.
- One logical change per commit.

## Fixups

Fix to earlier commit on this branch → `git commit --fixup <hash>`, not "wip"/"fix typo". User squashes before PR.

## Never

- Never push.
- Never rewrite history: no `git rebase`, `git commit --amend`, `git reset` or force push. Squashing fixups (`git rebase -i --autosquash origin/main`) is done by the user — you may only remind them about it before a PR.

## Steps

1. Get the prefix from the branch name.
2. Look at `git status` and `git diff --staged`.
   - **Something is staged** → commit only the staged changes and write the message for them. Don't stage anything else; leave unstaged and untracked files as they are (you may mention them).
   - **Nothing is staged** → stage the relevant files yourself; never stage secrets or generated dirs (`mobile/ios`, `mobile/android`).
3. If the changes are unrelated, propose splitting them into separate commits (if the user staged them, ask before changing what's staged).
4. Decide: fixup of an earlier branch commit, or a new commit.
5. Commit and show the result (`git log -1 --oneline`).

## Examples

```
[AM-2] Update .gitignore
[AM-7] Add shake gesture detection
[cursor-calibration] Fix cursor drift after calibration
fixup! [AM-7] Add shake gesture detection
```
