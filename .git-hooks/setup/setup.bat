del .\.git\hooks\*.sample
copy .\.git-hooks\hooks .\.git\hooks

git config --local commit.template .\.git-hooks\.commit-msg-template