# slurp

A tool for copying posts between [Fediverse](https://en.wikipedia.org/wiki/Fediverse) instances for testing purposes. Requires that they support the [Mastodon API](https://docs.joinmastodon.org/). Intended for use with [GotoSocial](https://gotosocial.org/) but will probably work with other Mastodon-like instances.

## build

```bash
go mod download
go build ./cmd/slurp
```

## run

`slurp` currently takes all of its configuration through environment variables:

- `SRC_INSTANCE`: the source instance from which to read the public timeline (required)
- `SRC_TOKEN`: OAuth bearer token for the source instance (optional: some instances still allow unauthenticated public timeline access)
- `DST_INSTANCE`: the destination instance to which to copy posts (required)
- `DST_TOKEN`: OAuth bearer token for the destination instance (required)
- `COUNT`: copy this many posts (optional: defaults to 100)
- `PAGE_SIZE`: fetch this many posts at a time (optional: defaults to 40, the Mastodon maximum)
- `DEBUG`: set to `true` to enable verbose debug logging (optional: defaults to false)
- `HTTPS_PROXY`: if you're using an HTTPS proxy for debugging, pass the URL to it here; for example, `http://localhost:9090` for [Proxyman](https://proxyman.io/)'s default configuration. (`HTTPS_PROXY` is common across most Go apps and not specific to `slurp`.)

```bash
SRC_INSTANCE=mastodon.social DST_INSTANCE=instance.example DST_TOKEN=XXXXXX COUNT=10 go run ./cmd/slurp
```

By default, posts with media, bot posts, non-public posts, and non-English posts are skipped. You can change this by editing the `skip()` function and recompiling `slurp`.

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
