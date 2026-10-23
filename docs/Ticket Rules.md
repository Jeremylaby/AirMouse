# Rulebook

This is a rulebook for us as a team on how to work with the board. If you really want, you can also use it as context for an LLM, so we don't end up with a bunch of bloated text here.

I designed these rules based on how I do it at work, to give things some structure but keep it simple, so it doesn't slow down development.

## Language

I want tickets and the rest of the documentation in English to keep this a bit more professional. When we try to translate our inner thoughts, we actually have to think more deeply about the topic structure and find better words for some terms. Also, translation usually involves some kind of AI, so there's a chance it'll rephrase our thoughts in a more understandable way. And we limit the risk that we'll just write down everything that comes to mind.

One more thing: before using AI, try to explain your task in your own words, just in English, because AI tends to use overly complicated words or generalize the task too much.

## Keep it simple

We're not a corporation, so we don't want long technical stories about simple stuff. Keep it simple.

For example: a ticket about adding a new gesture to our air mouse. Don't write a whole story about how we came up with the idea, why we think we should add it and how much it'll improve our app. Just write that we want to add the [gesture name] gesture, describe it, and say what it's for, e.g. "so we can use the app for [use case]".

We don't need a complicated ticket structure like this:

- Motivation and Business Impact
- User Story Statement
- User Interaction / Design / User Flow
- Functional Requirements and Developer Notes
- QA Notes
- Not in Scope, Questions and Answers
- Links and References
- Affected Ventures
- Affected Platforms

We only keep two parts: User Story and Acceptance Criteria.

## Ticket structure

### Title

The title should be a short description of what's going to be done, just a few words. When the issue is created, it'll automatically get an [AM-issue number] prefix, e.g. "[AM-12] Baking bread". Epics don't get the prefix.

### User Story

A straightforward description of what we want to do and why, plus any useful links, like technical/UI designs or research results (preferably as links).

### Acceptance Criteria

Can be left blank for most tickets. We don't want to write obvious things here, like "the bread can't be burned" or "code has to pass lint and CI tests" (if we have any). Only add the less obvious stuff, for example a communication design pattern between platforms that you discussed in a meeting with the rest of the team.

### Example

**Title:** [AM-12] Baking bread

**User Story:** We want to make tasty bread for our bakery, because we like good food and our customers are hungry. This bread is an essential product for our bakery project.

**Acceptance Criteria:**

- Has to be made from a mix of flour 1 and flour 2 in a 1:3 ratio.

## Tickets as a sanity check

Tickets are also a way to double-check that our idea makes sense, actually holds together, fits our project, and isn't too complicated for one task.

## Epics vs Tickets

If an epic is small enough, it can be a ticket, and if a ticket is big enough, it should be an epic. For example, if we're baking a very simple bread from a recipe we found online, we don't need to make it an epic. But if we need to design that bread, discuss proportions, ingredients etc., it should be an epic split into tickets. Same goes for code.

GitHub Issues doesn't give us many options for epics, so we can add a custom "Epic" label and connect sub-issues. The label is mainly there so epics don't get the [AM-issue number] prefix. What really matters is connecting the sub-issues.

## Board

Link: https://github.com/users/Jeremylaby/projects/2/views/1

|Backlog|In Progress|In Review|OK|Done|
|---|---|---|---|---|
|Stuff to do in the future.|Stuff we're working on right now. I'd suggest no more than 2–3 tickets at a time, including the ones in review.|Stuff that's ready for the other person on the project to check and test. We're not going to do any strict reviews here, so we don't slow down development. The point is just to catch smaller or bigger mistakes.|Stuff that got approved in review and is ready to merge, or accepted concept files.|Stuff that's finished and merged.|

## Branches

Not everything needs to be an issue, but everything has to be on a separate branch. I'll block pushing directly to `main`.

If you do something outside the board, let the other person know so they can check it before you merge it to `main`.

### Branch naming

I don't want to force anything here, but I'd suggest using either the ticket title (e.g. `baking-bread`) or the issue tag (e.g. `AM-12`). For epics, use the name instead (e.g. `bakery-menu`), so it's easier to tell them apart from ticket branches.

### Working with epics

When working on an epic, create an epic branch, branch off it for each sub-issue, and merge the sub-issues back into it. We don't commit directly to epic branches. They're only there to work on a bigger part of the project.

## Merging

We use rebase and merge, so commits from your branch land on `main` one by one and the history stays linear. If something breaks, we can revert the commits that caused it.

## Commits

Since every commit lands on `main`, keep them clean:

- Start with the issue tag in square brackets, e.g. `[AM-12] Add flour ratio`.
- Right after the tag, always start with a verb that says what the commit does (Add, Fix, Remove, Update…).
- Keep messages short.
- Before opening a PR, squash "wip" and "fix typo" commits into meaningful ones (`git rebase -i`).

### Fixup commits

The easiest way to avoid "fix typo" commits is a fixup. When you fix something in an earlier commit on your branch, stage the fix and run:

```sh
git commit --fixup <hash>
```

This creates a normal commit named `fixup! <title of that commit>`. Nothing is changed or merged yet. It's just a note that this commit belongs to the other one.

Right before opening a PR (and again after fixes from review), squash them:

```sh
git fetch origin
git rebase -i --autosquash origin/main
```

The rebase only rewrites your current branch, `main` stays untouched. `origin/main` is just the reference point: Git takes the commits that are on your branch but not on `main`, merges each `fixup!` commit into the commit it belongs to, and puts the result on top of the latest `main`. In the editor that opens, just save and close.

If the branch was already pushed, push it with `git push --force-with-lease`. Only do this on your own branch, never on `main` or an epic branch. For sub-issues of an epic, use the epic branch as the base instead of `origin/main`.

Ask Stanisław Barycki if there's already a git hook that automatically adds the branch name in `[]` to the commit message.
