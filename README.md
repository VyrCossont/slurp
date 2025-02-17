# slurp

A tool for exporting data from and importing data to [Fediverse](https://en.wikipedia.org/wiki/Fediverse) instances. Requires that they support the [Mastodon API](https://docs.joinmastodon.org/) as implemented By [GotoSocial](https://gotosocial.org/). Intended for use with GotoSocial, but should work with other Mastodon-like instances, including Mastodon.

## what

`slurp` has commands for importing and exporting these API objects, in CSV formats compatible with the Mastodon import/export GUI where possible:

- `archive` (import only, only tested with Mastodon archives)
- `blocks`
- `bookmarks`
- `emojis` (the format is `slurp`-specific, and importing them requires admin privileges)
- `filters` (note that Mastodon can't import or export filters yet, so the current format is `slurp`-specific, and handles keyword filters only, not status filters)
- `follows`
- `followers` (export only, and Mastodon can't export followers, so the format is `slurp`-specific)
- `lists`
- `mutes`

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

### importing a Mastodon archive

`slurp` can import [an archive of your statuses and media exported from a Mastodon instance](https://docs.joinmastodon.org/user/moving/#export), backdating statuses to the same date they were originally posted, and loading them in a way that doesn't push them to your followers or generate notifications. In all other respects, this is the same as reposting those statuses manually: the original statuses still exist in their original locations if the original server is still up, their likes and boosts do not transfer, etc.

Backdating requires GoToSocial 0.18.0-rc2 or newer. Importing statuses to non-GTS instances and importing statuses from non-Mastodon instances (for example, Akkoma) have *not* been tested. Try it at your own risk, ideally on a throwaway development instance.

Before proceeding, you might want to copy custom emoji from your old instance. Otherwise, imported statuses that use missing custom emoji will be skipped. The `--inline` option saves the emoji picture data as well as their metadata in the `emoji.json` file, and is optional, but useful for keeping your favorite emoji if the old instance later goes away.

```bash
./slurp --user olduser@old-instance.tld emojis export --inline --file emoji.json

./slurp --user user@instance.tld emojis import --file emoji.json
```

Importing an archive requires that your archive be already uncompressed. (It should contain `actor.json` and `outbox.json` files, and a `media_attachments` folder.) Importing also requires two map files so that an interrupted import can be resumed. It is safe to interrupt an archive import: as long as you have your map files, this shouldn't result in duplicated statuses or media.

Depending on your rate limit settings for both `slurp` and your GTS instance, and how much media you have, an archive import may take minutes or hours.

Assuming you have downloaded your archive and uncompressed it as `archive-XXXXXX-YYYYYYYYYY`, here's how you import it:

```bash
./slurp --user user@instance.tld archive import \
  --file archive-XXXXXX-YYYYYYYYYY \
  --status-map-file archive-XXXXXX-YYYYYYYYYY/status-map.json \
  --attachment-map-file archive-XXXXXX-YYYYYYYYYY/attachment-map.json
```

### prefs

`slurp` respects these environment variables:

- `HTTPS_PROXY`: if you're using an HTTPS proxy for debugging, pass the URL to it here; for example, `http://localhost:9090` for [Proxyman](https://proxyman.io/)'s default configuration. (`HTTPS_PROXY` is common across most Go apps and not specific to `slurp`.)

`slurp` stores its preferences in `~/Library/Application Support/codes.catgirl.slurp/prefs.json` or the equivalent for your OS (usually `~/.config/codes.catgirl.slurp/prefs.json` on Linux), respecting XDG environment variables and their equivalents. You can print that path with this command:

```bash
./slurp prefs path
```

#### rate limits

`slurp`'s default rate limit for any given instance is half of the GTS default, so as not to conflict with a normal client running on the same IP. If you want to increase it (for example, if you've increased the default in your instance's config as well), you can do so with this command, which applies to the default user's instance, or that of whichever user you've set with the `--user` flag:

```bash
./slurp prefs set ratelimit 1.0
```

You can also set burst capacity for an instance (the maximum number of requests that will be allowed to happen at once), although this is generally less useful:

```bash
./slurp prefs set burstcap 300
```

## update Swagger client

Do this when the GotoSocial API changes. This will use the Swagger spec on GotoSocial's `main` branch.

```bash
rm -rf client models
go generate ./...

# apply workaround for https://github.com/go-swagger/go-swagger/issues/2997
patch -u -p1 -i filter-context.diff
```

You can also use `go-swagger` directly instead of through `go generate`, which you'll want to do if using a different branch or tag, or a local copy of the GotoSocial codebase. In the latter case, don't forget to [update your copy's `swagger.yaml`](https://github.com/superseriousbusiness/gotosocial/blob/main/CONTRIBUTING.md#updating-swagger-docs) first.

```bash
go run github.com/go-swagger/go-swagger/cmd/swagger generate client --spec /path/to/gotosocial/docs/api/swagger.yaml
```
