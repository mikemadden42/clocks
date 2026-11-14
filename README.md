# clocks

Display time for various timezones.

## Install

```sh
go install github.com/mikemadden42/clocks@latest
```

Requires Go 1.26 or later.

## Usage

```console
$ clocks
🏠 Local        Fri, 13 Nov 2026 18:46:19 CST
☀️ Los Angeles  Fri, 13 Nov 2026 16:46:19 PST
🗽 New York     Fri, 13 Nov 2026 19:46:19 EST
⏰ UTC          Sat, 14 Nov 2026 00:46:19 UTC
🇬🇧 London       Sat, 14 Nov 2026 00:46:19 GMT
🇦🇺 Sydney       Sat, 14 Nov 2026 11:46:19 AEDT
```

To change the timezones shown, edit the `clocks` list in `main.go`. Zone names come from the
[IANA time zone database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

## Development

```sh
make check    # run golangci-lint
make build    # build snapshot binaries with goreleaser
```
