// schema/nerdctl.cue — the self-contained, load-gate-only schema for the
// plugin-nerdctl candy. The `engine` provider class carries NO authored input of
// its own (its request/reply envelopes live in the base spec schema —
// spec/schema/engine.cue), so this schema declares the package marker only; it
// exists so the plugin ships the per-plugin `schema/*.cue` contract every plugin
// candy must carry.
package nerdctl

// #EnginePlugin marks the plugin's self-contained schema surface. The engine op
// envelopes (#EngineDescribeRequest/Reply, #EngineBinaryRequest/Reply, …) are
// owned by spec/schema/engine.cue and are deliberately NOT re-declared here.
#EnginePlugin: {
	word: "nerdctl"
}
