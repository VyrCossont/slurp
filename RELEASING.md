# Releasing

## Steps

1. Ensure you have the most current main branch
   - This command assumes [slurp's](https://github.com/VyrCossont/slurp) remote
     is origin.
   - `git checkout main && git pull origin main`
2. Create a valid [semver](https://semver.org/) tag
   - `git tag -a v${VERSION} -m "$RELEASE_MESSAGE_HERE"`
   - `git push origin v${VERSION}`
3. Authenticate to GitHub and set GITHUB_TOKEN
   - `export GITHUB_TOKEN=$YOUR_TOKEN_HERE`
4. Run goreleaser
   - `goreleaser release`

You will now have created a github release!

## References

- [goreleaser quick start](https://goreleaser.com/quick-start/)
- [goreleaser documentation](https://goreleaser.com/customization/)
