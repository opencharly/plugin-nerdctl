// plugin-nerdctl's OWN self-contained CUE schema — the SINGLE SOURCE for this plugin's
// declaration surface, served over the Describe channel (there is no schema-less
// plugin). SELF-CONTAINED and PACKAGE-LESS: it references no base def and carries no
// package clause, so the SDK compiles it STANDALONE (the property `cue exp gengotypes`
// needs) AND the host can splice `base ++ plugin` at the load gate.
//
// The plugin serves `engine:nerdctl` plus the `verb:nerdctl` / `command:nerdctl`
// convenience words. The `engine` provider class carries NO authored input of its own —
// its request/reply envelopes live in the base spec schema (spec/schema/engine.cue) —
// so this schema DOCUMENTS the engine word.
#NerdctlPlugin: {
	// The engine word the plugin serves (the registry key `engine:nerdctl`).
	engine: "nerdctl"

	// What the plugin does, in one line (the public-docs surface).
	contract: string & !=""
}
