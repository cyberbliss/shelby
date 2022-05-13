# Development notes
`make build` to build the latest in this directory. Can then test with `./shelby info`

## Releasing
```
export GITHUB_TOKEN=1234xxx
export HB_TOKEN=123xxx
```
1. Commit and push your changes to a remote branch
2. PR to merge branch into main
3. Tag commit:
```
git tag -a vn.n.n
git push origin vn.n.n
```

4. `goreleaser release --rm-dist`