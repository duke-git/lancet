# Lancet Contributing Guide

Hi! Thank you for choosing Lancet.

Lancet is a powerful, efficient, and reusable util function library of go. It makes Go dev easier by taking the hassle out of working with concurrency, net, math, slice, string, etc.

We are excited that you are interested in contributing to lancet. Before submitting your contribution though, please make sure to take a moment and read through the following guidelines.

## Issue Guidelines

- Issues are exclusively for bug reports, feature requests and design-related topics. Other questions may be closed directly.

- Before submitting an issue, please check if similar problems have already been issued.

- Please specify which version of Lancet and Go you are using, and provide OS information. [Go Playground](https://go.dev/play/) is recommended to build a live demo so that your issue can be reproduced clearly.

## Pull Request Guidelines

- Fork this repository to your own account. Do not create branches here.

- Commit info should be formatted as `type(scope): info about commit`. eg. `fix(package): [scrollbar] fix xxx bug`.

  1. type: type must be one of [chore, docs, feat, fix, refactor, release, test].

  2. scope: scope must be one of [package, file, internal].

  3. header: header must not be longer than 72 characters.

- Rebase before creating a PR to keep commit history clear.

- Before submitting a PR, please execute the unit test command: `go test -v ./...` to ensure that all unit test tasks should pass.

- Make sure PRs are created to `v2` branch instead of `master` branch.

- If your PR fixes a bug, please provide a description about the related bug.

- Merging a PR takes two maintainers: one approves the changes after reviewing, and then the other reviews and merges.
