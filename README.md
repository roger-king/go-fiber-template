# Go-Fiber-Template

GoLang API server that utilizes Fiber and Google's Wire DI.

## How-To

In order to re-use this template, we need to perform some manual steps: (TODO: to automate some steps plus some extra codegen)

1. Do a global find and replace of `github.com/roger-king/gh-template`.
   - This is needed because of how go does imports!
   - WIP: automating this step!
2. Ensure you have docker installed!

## Usage

```bash
   # run docker-compose command
   make start

   # generate wire_gen file
   make wire

   # production build
   make build
```

## Docs:

- [Fiber](https://gofiber.io/)
- [Wire](https://github.com/google/wire)
- [Testify](https://github.com/stretchr/testify)
