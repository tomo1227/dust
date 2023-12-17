# ローカルの ~/.gitconfigをコピーする
git config --global --add safe.directory ${containerWorkspaceFolder}

git config pull.rebase false
git config --global commit.template .commit_template