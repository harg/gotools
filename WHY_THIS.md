# Forked golang tools repo with vendored deps

1. Fork the repo : https://github.com/golang/tools.git

2. Add remote upstream repo

`git remote add upstream https://github.com/golang/tools.git`

3. Fetch all branches from upstream

`git fetch upstream`

4. Rebase from upstream

`git rebase upstream/master`

5. commit and push

6. Build gopls

``
cd .\gopls\  
go build -ldflags "-s -w" -o gopls.exe
`

7. Build goimports

``
cd .\cmd\goimports\  
go build -ldflags "-s -w" -o goimports.exe
`
