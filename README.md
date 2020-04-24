# ConCom - Conventional Committer :memo:

![Build](https://github.com/mstrangfeld/concom/workflows/Build/badge.svg)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/mstrangfeld/concom?label=version&sort=semver)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mstrangfeld/concom)
![License](https://img.shields.io/github/license/mstrangfeld/concom)

A CLI for committing [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).

1. Choose the type of your commit
2. Define the scope (can be empty)
3. Choose whether your commit does include breaking changes
4. Write the commit message in your predefined `$EDITOR`

```
> concom
Use the arrow keys to navigate: ↓ ↑ → ←
? What type is your commit?:
  ▸ feat - A new feature
    fix
    docs
    build
    ci
    perf
    refactor
    style
    test
    chore
```

```
> concom
✔ feat - A new feature
✔ Scope:
```

```
> concom
✔ feat - A new feature
Scope: █
? Breaking? [y/N]
```

## Installation

Via Go get:

```
> go get -u github.com/mstrangfeld/concom
```

## Tips and Tricks

### Vim
+ To discard a commit you have to close vim wih a non-zero exit value
    + Type `:cq`

## License

    Copyright (C) 2020  Marvin Strangfeld

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
