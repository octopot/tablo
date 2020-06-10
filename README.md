> # 🧐 Tablo
>
> The one point of view to all your task boards.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]
[![Coverage][coverage.icon]][coverage.page]

## 💡 Idea

...

A full description of the idea is available [here][design.page].

## 🏆 Motivation

...

## 🤼‍♂️ How to

### Development

```bash
$ make init
$ source bin/activate
# PATH         -> bin/:PATH
# go get -mod= -> get
...
$ deactivate
```

## 🧩 Installation

### Homebrew

```bash
$ brew install octolab/tap/tablo
```

### Binary

```bash
$ curl -sSfL https://raw.githubusercontent.com/octopot/tablo/master/bin/install | sh
# or
$ wget -qO-  https://raw.githubusercontent.com/octopot/tablo/master/bin/install | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get github.com/octopot/tablo/cmd/client@latest
$ go get github.com/octopot/tablo/cmd/server@latest
# or use egg tool
$ egg tools add github.com/octopot/tablo/cmd/client@latest
$ egg tools add github.com/octopot/tablo/cmd/server@latest
```

> [egg][] is an `extended go get`.

### Bash and Zsh completions

```bash
$ tablo completion bash > /path/to/bash_completion.d/tablo.sh
$ tablo completion zsh  > /path/to/zsh-completions/_tablo.zsh
# or autodetect
$ source <(tablo completion)
```

> See `kubectl` [documentation](https://kubernetes.io/docs/tasks/tools/install-kubectl/#enabling-shell-autocompletion).

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/octopot/tablo
[build.icon]:       https://travis-ci.com/octopot/tablo.svg?branch=master
[coverage.page]:    https://codeclimate.com/github/octopot/tablo/test_coverage
[coverage.icon]:    https://api.codeclimate.com/v1/badges/0d0bd50ec0cb47760ea5/test_coverage
[design.page]:      https://www.notion.so/octolab/Tablo-a56b2f5410be487f953df0db79b20f2f?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://github.com/octopot/tablo
[template.page]:    https://github.com/octomation/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue

[egg]:              https://github.com/kamilsk/egg
