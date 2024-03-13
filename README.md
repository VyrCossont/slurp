# slurp

A tool for exporting data from and importing data to [Fediverse](https://en.wikipedia.org/wiki/Fediverse) instances. Requires that they support the [Mastodon API](https://docs.joinmastodon.org/) as implemented By [GotoSocial](https://gotosocial.org/). Intended for use with GotoSocial, but should work with other Mastodon-like instances, including Mastodon.

## build

```bash
go mod download
go build .
```

## run

Show help for all commands.

```bash
./slurp help
```

Before running other commands, log in.

You'll be asked to log into your instance in your web browser, and paste the provided authorization code into the prompt. This will save your access token in the system keychain, and that user as the default user in slurp's preferences.

```bash
./slurp --user user@instance.tld auth login
```

Load follows from a previous instance. This will use your stored access token and default user.

```bash
./slurp follows import --file following_accounts.csv
```

Save follows from this instance.

```bash
./slurp follows export --file follows_backup.csv
```

`slurp` respects these environment variables:

- `HTTPS_PROXY`: if you're using an HTTPS proxy for debugging, pass the URL to it here; for example, `http://localhost:9090` for [Proxyman](https://proxyman.io/)'s default configuration. (`HTTPS_PROXY` is common across most Go apps and not specific to `slurp`.)

`slurp` stores its preferences in `~/Library/Application Support/codes.catgirl.slurp/prefs.json` or the equivalent for your OS (usually `~/.config/codes.catgirl.slurp/prefs.json` on Linux), respecting XDG environment variables and their equivalents.

## update Swagger client

Do this when the GotoSocial API changes. This will use the Swagger spec on GotoSocial's `main` branch.

```bash
go get github.com/go-swagger/go-swagger/cmd/swagger
rm -rf client models
go generate ./...

# apply workaround for https://github.com/go-swagger/go-swagger/issues/2997
patch -u -p1 -i fcontext.diff
```

You can also use `go-swagger` directly instead of through `go generate`, which you'll want to do if using a different branch or tag, or a local copy of the GotoSocial codebase. In the latter case, don't forget to [update your copy's `swagger.yaml`](https://github.com/superseriousbusiness/gotosocial/blob/main/CONTRIBUTING.md#updating-swagger-docs) first.

```bash
go run github.com/go-swagger/go-swagger/cmd/swagger generate client --spec /path/to/gotosocial/docs/api/swagger.yaml
```
