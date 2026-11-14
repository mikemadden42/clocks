# TODO

Findings from a project review (2026-11-13).

## Config problems

- [x] **Migrate `.golangci.yml` to the v2 schema.** `golangci-lint config verify` (v2.15.0) rejects it:
  ```
  "version" ... got number, want string
  'disable-all' not allowed
  'exclude-dirs' not allowed
  'linters-settings' not allowed
  ```
  Use `version: "2"`, `linters.default: none`, `linters.settings.revive`,
  `formatters.settings.goimports.local-prefixes: [github.com/mikemadden42]` and
  `linters.exclusions.paths`. Until then the revive `exported` settings and the
  excluded directories don't take effect.
- [x] **Add a package comment to `main.go`.** `make check` reports
  `main.go:1:1: package-comments: should have a package comment (revive)`.
- [x] **Fix `.goreleaser.yaml` so `goreleaser check` passes.**
  - Add `version: 2` at the top (`only version: 2 configuration files are supported, yours is version: 0`).
  - Rename the deprecated `snapshot.name_template` to `snapshot.version_template`.

## Possible runtime bug

- [x] **Embed timezone data.** `time.LoadLocation` reads the OS timezone database. Windows
  doesn't have one in that format, so Go falls back to `$GOROOT/lib/time/zoneinfo.zip`, which
  only exists where Go is installed. Slim containers often lack timezone data too. Released
  binaries could then print `Error loading timezone` for every zone except UTC. CI passes on
  `windows-latest` only because those runners have Go installed. Fix (about 450 KB larger binary):
  ```go
  import _ "time/tzdata"
  ```

## Version mismatches

- [x] **Pick one minimum Go version.** `go.mod` says `go 1.21`, the commit history says 1.21–1.23,
  and CI tests 1.23–1.25. Go 1.21–1.23 are past end of life. Set a floor such as `go 1.25` in
  `go.mod` and test from there up to the current release.
- [ ] **Rename the module to `github.com/mikemadden42/clocks`.** The module is currently `clocks`,
  so the goimports `local-prefixes` setting has no effect, and
  `go install github.com/mikemadden42/clocks@latest` doesn't work.

## Smaller issues

- [x] Run the linter (`golangci/golangci-lint-action`) and `goreleaser check` in CI. CI runs
  `go test`, but there are no test files.
- [ ] Remove the duplicated timezone definitions in `main.go` (the map plus the `keys` slice) by
  using a single ordered slice of `{zone, label}` structs. Drop the unneeded `timezone := key` copy.
- [ ] Align the output columns (labels differ in length and emoji width varies), or drop the
  weekday/date when it matches local time.
- [ ] Remove the leftover goreleaser hooks: `go generate ./...` does nothing here, and
  `go mod tidy` isn't needed without dependencies.
- [ ] Expand the README with install instructions and sample output.
