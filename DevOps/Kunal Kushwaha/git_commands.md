# Git commands

`git init` (Initialize a new Git repository) <br>
`git status` (Check the status of the repository) <br>
`git add <file>` (Stage changes for the next commit) <br>
`git add .` (Stage all changes for the next commit) <br>
`git commit -m "message"` (Commit staged changes with a message) <br>
`git commit -am "message"` (Stage and commit changes in one step) <br>
`git restore --staged <file>` (Unstage a file) <br>
`git log` (View commit history) <br>
`git log --oneline` (View commit history in a condensed format) <br>
`git reset <commit>` (Reset to a specific commit; NOTE: each commit is built on top of each other. We cannot remove a commit from the middle.) <br>
`git stash` (Temporarily save changes that are not ready to be committed) <br>
`git stash pop` (Reapply stashed changes) <br>
`git stash clear` (Clear all stashed changes) <br>
`git remote add origin <repository_url>` (Add a remote repository) <br>
`git remote -v` (View remote repositories) <br>
`git push origin main` (Push changes to the remote repository on the main branch) <br>