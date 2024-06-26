// automatically generated by stateify.

//go:build !false
// +build !false

package kernel

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (p *pidNamespaceData) StateTypeName() string {
	return "pkg/sentry/kernel.pidNamespaceData"
}

func (p *pidNamespaceData) StateFields() []string {
	return []string{}
}

func (p *pidNamespaceData) beforeSave() {}

// +checklocksignore
func (p *pidNamespaceData) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
}

func (p *pidNamespaceData) afterLoad() {}

// +checklocksignore
func (p *pidNamespaceData) StateLoad(stateSourceObject state.Source) {
}

func init() {
	state.Register((*pidNamespaceData)(nil))
}
