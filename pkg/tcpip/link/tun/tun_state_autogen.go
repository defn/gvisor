// automatically generated by stateify.

package tun

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (d *Device) StateTypeName() string {
	return "pkg/tcpip/link/tun.Device"
}

func (d *Device) StateFields() []string {
	return []string{
		"Queue",
		"endpoint",
		"notifyHandle",
		"flags",
	}
}

// +checklocksignore
func (d *Device) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.Queue)
	stateSinkObject.Save(1, &d.endpoint)
	stateSinkObject.Save(2, &d.notifyHandle)
	stateSinkObject.Save(3, &d.flags)
}

func (d *Device) afterLoad() {}

// +checklocksignore
func (d *Device) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.Queue)
	stateSourceObject.Load(1, &d.endpoint)
	stateSourceObject.Load(2, &d.notifyHandle)
	stateSourceObject.Load(3, &d.flags)
}

func (f *Flags) StateTypeName() string {
	return "pkg/tcpip/link/tun.Flags"
}

func (f *Flags) StateFields() []string {
	return []string{
		"TUN",
		"TAP",
		"NoPacketInfo",
	}
}

func (f *Flags) beforeSave() {}

// +checklocksignore
func (f *Flags) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.TUN)
	stateSinkObject.Save(1, &f.TAP)
	stateSinkObject.Save(2, &f.NoPacketInfo)
}

func (f *Flags) afterLoad() {}

// +checklocksignore
func (f *Flags) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.TUN)
	stateSourceObject.Load(1, &f.TAP)
	stateSourceObject.Load(2, &f.NoPacketInfo)
}

func (r *tunEndpointRefs) StateTypeName() string {
	return "pkg/tcpip/link/tun.tunEndpointRefs"
}

func (r *tunEndpointRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *tunEndpointRefs) beforeSave() {}

// +checklocksignore
func (r *tunEndpointRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *tunEndpointRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*Device)(nil))
	state.Register((*Flags)(nil))
	state.Register((*tunEndpointRefs)(nil))
}
