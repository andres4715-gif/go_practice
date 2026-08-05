---
name: git-helper
description: Pushes changes to the repository when the user says "git_push". Runs git add ., a git commit with a message matching the changes, and git push origin to the current branch.
---

# Git Helper

When the user says **`git_push`**, push the changes to the repository following these steps in order:

## Steps

0. Important: the commit message should be in English 

1. **Check what changed** so you can write a good commit message:
   ```bash
   git status
   git diff --stat
   ```

2. **Stage all changes**:
   ```bash
   git add -A
   ```

3. **Find the current branch**:
   ```bash
   git rev-parse --abbrev-ref HEAD
   ```

4. **Commit** with a message that matches the actual changes (not generic).
   The message should describe *what* changed, short and clear:
   ```bash
   git commit -s -m "<message matching the changes>"
   ```

5. **Push** to `origin` on the current branch:
   ```bash
   git push origin <branch-name>
   ```
6. Once the process is finished you should show up the commit message as: "The commit message is: "

## Rules for the commit message

- Make it reflect what actually changed (check `git diff --stat` / `git status`).
- Keep it short and descriptive. Examples:
  - `add pointers example`
  - `fix indentation in slice`
  - `add git-helper skill`
- Avoid vague messages like "changes" or "update".

## Notes

- If there is nothing to commit (`git status` is clean), tell the user and do not create an empty commit.
- Use the branch name returned by `git rev-parse --abbrev-ref HEAD`; do not assume `master` or `main`.
- Do not add co-authors or signatures to the message unless the user asks for it.
