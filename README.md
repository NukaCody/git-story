# Git-Story

EDIT: I'm dumb. `git log --first-parent --date-order --reverse --abbrev-commit` does everything you want and more. Leaving this public so no one makes the same mistake

---

A very silly tool that takes the current HEAD pointer and prints the commit history in reverse order (first to last). I use this with VSCode to see the story of a repository

## Example Usage

Grab the commit history of a repo and log it to a file

```bash
go run main.go --repo https://github.com/hashicorp/consul > consul-hist.log 
```

Print the abbreviated commits from n to m. i.e print the first 10 commits (Or you could just open the file in a text editor).

```bash
head -n 10 consul-hist.log | tail -n +1
```

Go to the beginning (first commit may differ)

```bash
git checkout 0a7996bc
```

Walk through commits one by one

```bash
git checkout HEAD@{1}
```

Go back a commit

```bash
git checkout HEAD^1
```

## FAQ

Q: Can't you do this with bash scripting?
A: Probably, but I aint a masochist.