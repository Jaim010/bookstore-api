#!/bin/bash
rm ./.git-hooks/hooks/*.sample > /dev/null 2>&1
cp -TRv ./.git-hooks/hooks ./.git/hooks

git config --local commit.template ./.git-hooks/.commit-msg-template