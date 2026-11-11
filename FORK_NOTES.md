# Fork notes — local-first Ollama

This is a **local-first fork** of [ollama/ollama](https://github.com/ollama/ollama).
It tracks upstream and keeps a small, reviewable patch set. The guiding rule:
**nothing leaves your machine unless you explicitly opt in.**

## What's different from upstream

### Cloud is OFF by default

Upstream enables cloud features (remote inference, web search, model
recommendations fetched from `ollama.com`, account sign-in) by default and asks
you to opt *out* with `OLLAMA_NO_CLOUD=1`. This fork inverts that: cloud is
**disabled by default**, and you opt *in*.

| Goal | Upstream | This fork |
| --- | --- | --- |
| Default behavior | cloud on | **cloud off** |
| Turn cloud **on** | (default) | `OLLAMA_ENABLE_CLOUD=1` or `OLLAMA_NO_CLOUD=0` |
| Turn cloud **off** | `OLLAMA_NO_CLOUD=1` | (default) |

The inversion lives in a single place — `envconfig.NoCloud()` — so it
propagates automatically to every cloud-gated path:

- the `ollama.com` **model-recommendations** fetch (`server/model_recommendations.go`)
- the cloud **policy** layer used by the OpenAI/Anthropic middleware (`internal/cloud/policy.go`)
- remote/cloud **inference passthrough** for `:cloud` models

The CLI's `:cloud` "did you mean?" suggestion — which probes `ollama.com` after
a failed pull — is also gated, so a failed local pull stays a purely local
error (`cmd/cloud_suggest.go`).

### Re-enabling cloud

Set either environment variable before starting the server:

```shell
OLLAMA_ENABLE_CLOUD=1 ollama serve   # explicit opt-in
# or, upstream-style:
OLLAMA_NO_CLOUD=0 ollama serve
```

A `~/.ollama/server.json` with `"disable_ollama_cloud": true` still forces cloud
off even when opted in.

## Deliberately left unchanged (for now)

- **`NoCloudEnv()` keeps upstream semantics** ("is `OLLAMA_NO_CLOUD` truthy?").
  The desktop app's cloud-config state machine and `NoCloudSource()` read it
  directly; the default inversion is confined to `NoCloud()` so app-side source
  detection is unaffected.
- **The desktop app's own cloud toggle default** is unchanged — this first pass
  targets the server + CLI (the headless `ollama` binary). Making the GUI
  local-first by default is a follow-up.

## Keeping in sync with upstream

```shell
git remote add upstream https://github.com/ollama/ollama.git
git fetch upstream
git rebase upstream/main      # patch set is small by design
```

## Building / testing

This fork pins the same Go toolchain as upstream (see `go.mod`). Build and test
with the standard commands:

```shell
go build .
go test ./envconfig/ ./cmd/ ./middleware/ ./server/
```
