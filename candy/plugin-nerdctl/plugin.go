// Package nerdctl is the out-of-tree charly ENGINE plugin serving the `engine`
// provider class for the nerdctl word — `engine:nerdctl` — plus the `verb:nerdctl`
// / `command:nerdctl` convenience words.
//
// The engine provider is engine-word-keyed DATA + ONE dispatch: this plugin
// answers every engine op (describe / binary / gpu_args / start_plan / unit_emit /
// network_ensure) by delegating to the pure body `container.InvokeEngineOp("nerdctl", …)`
// in the spec module, so the op behavior is written ONCE and served identically by
// the compiled-in `engine:podman` / `engine:docker` providers and this
// out-of-process `engine:nerdctl` plugin.
//
// Dual-placement by construction: the SAME NewProvider()/NewMeta() compile INTO
// charly in-process when listed in compiled_plugins, or cmd/serve serves them
// OUT-OF-PROCESS over go-plugin gRPC when they are not — placement is invisible
// above the registry.
package nerdctl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/opencharly/sdk"
	"github.com/opencharly/spec/container"
	pb "github.com/opencharly/spec/proto"
)

const calver = "2026.266.1500"

// engineWord is both the registry key ("engine:nerdctl") and the engine the pure
// body dispatches on.
const engineWord = "nerdctl"

// NewProvider returns the nerdctl provider.
func NewProvider() pb.ProviderServer { return &provider{} }

// NewMeta advertises the engine:nerdctl capability (empty InputDef — the engine
// class carries no authored plugin_input; its envelopes live in the base
// schema/spec) plus the verb/command convenience words, with a self-contained,
// load-gate-only CUE schema.
func NewMeta() pb.PluginMetaServer {
	return sdk.NewMeta(calver,
		[]sdk.ProvidedCapability{
			{Class: "engine", Word: engineWord, InputDef: ""},
			{Class: "verb", Word: engineWord, InputDef: ""},
			{Class: "command", Word: engineWord, InputDef: ""},
		},
		nil)
}

type provider struct{ pb.UnimplementedProviderServer }

// Invoke answers every engine op by delegating to the pure fabric servant
// container.InvokeEngineOp, which validates the engine word and the op and errors
// on an unknown one. The verb/command words serve the same envelope.
func (provider) Invoke(_ context.Context, req *pb.InvokeRequest) (*pb.InvokeReply, error) {
	out, err := container.InvokeEngineOp(engineWord, req.GetOp(), json.RawMessage(req.GetParamsJson()))
	if err != nil {
		return nil, fmt.Errorf("plugin-nerdctl: engine:%s op %q: %w", engineWord, req.GetOp(), err)
	}
	return &pb.InvokeReply{ResultJson: out}, nil
}
