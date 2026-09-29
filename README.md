# plugin-nerdctl

The out-of-tree charly ENGINE plugin serving `engine:nerdctl` — the containerd
engine backend for OpenCharly pods, plus the `verb:nerdctl` / `command:nerdctl`
convenience words.

The plugin is a standalone Go module: charly's loader host-builds `./cmd/serve`
and connects it **out-of-process** over go-plugin gRPC via the plugin SDK. It
answers every engine op by delegating to the pure body
`container.InvokeEngineOp("nerdctl", …)` in the spec module — the SAME body the
compiled-in `engine:podman` / `engine:docker` providers serve — so the engine
class has one implementation and two placements.

## What it provides

| Capability | Surface |
|---|---|
| `engine:nerdctl` | the engine provider-class op envelope: `describe` / `binary` / `gpu_args` / `start_plan` / `unit_emit` / `network_ensure` |
| `verb:nerdctl` | the `nerdctl:` convenience verb |
| `command:nerdctl` | the `charly nerdctl` convenience command |

Composes with the `layer-nerdctl` candy (the packages + rootless
containerd/buildkit/CNI stack) to deploy a pod with `engine: nerdctl`.

## How to use it

Compose the plugin candy alongside the layer candy in a box's `candy:` list:

```yaml
- '@github.com/opencharly/plugin-nerdctl/candy/plugin-nerdctl:<tag>'
- '@github.com/opencharly/layer-nerdctl/candy/layer-nerdctl:<tag>'
```

Then select the engine on the deploy:

```yaml
engine: nerdctl
```

## Layout

- `candy/plugin-nerdctl/` — the plugin module: `plugin.go` (provider + meta),
  `schema/plugin.cue`, and `cmd/serve/main.go`.
- `charly.yml` — the root project manifest (`discover: box` + `discover: candy`).
- The candy's `plugin-nerdctl-skill:` entity is the projected source for the
  `/charly-internals:plugin-nerdctl` skill.

## Related

- Owning skill: `/charly-internals:plugin-nerdctl` — the engine:nerdctl provider
  reference (projected from this candy's own `skill:` entity).
- `/charly-internals:plugin` — the plugin/provider model and placement.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.
