## Dashlights
[![Go Report
Card](https://goreportcard.com/badge/github.com/erichs/dashlights)](https://goreportcard.com/report/github.com/erichs/dashlights)

> turn ENV vars into dashboard diagnostic lights for your shell

Dashlights is a simple binary that you drop somewhere in your PATH. It lets you
turn this:

```shell
~
❯ export DASHLIGHT_HEXAGRAM_4DCB_FGRED="get up and stretch a bit" DASHLIGHT_LINK_1F517="VPN is up"

```

into this:

```shell
~
䷋  🔗
❯
```

## Why?

From time to time, you may want to bring attention to the fact that you're
operating in a certain state or mode, and highlight this in your shell.
Dashlights provides a simple interface between shell functions or
aliases and a dash display, via exported ENV vars.

Any time you export a var of the following form in your shell, `dashlights` can
parse it and render a UTF-8 glyph, with optional foreground and background coloring applied:

```
 prefix            utf8 hexcode
DASHLIGHTS_BULBNAME_HEXSTRING_OPTIONALCOLORS...
             label              one or more color codes
```

## Usage

```
Usage: dashlights [--obd] [--list] [--clear]

Options:
  --obd, -d              On-Board Diagnostics: display diagnostic info if provided.
  --list, -l             List supported color attributes.
  --clear, -c            Shell code to clear set dashlights.
  --help, -h             display this help and exit
```
