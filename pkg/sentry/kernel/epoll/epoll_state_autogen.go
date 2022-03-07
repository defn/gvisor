// automatically generated by stateify.

package epoll

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (f *FileIdentifier) StateTypeName() string {
	return "pkg/sentry/kernel/epoll.FileIdentifier"
}

func (f *FileIdentifier) StateFields() []string {
	return []string{
		"File",
		"Fd",
	}
}

func (f *FileIdentifier) beforeSave() {}

// +checklocksignore
func (f *FileIdentifier) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.File)
	stateSinkObject.Save(1, &f.Fd)
}

func (f *FileIdentifier) afterLoad() {}

// +checklocksignore
func (f *FileIdentifier) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadWait(0, &f.File)
	stateSourceObject.Load(1, &f.Fd)
}

func (p *pollEntry) StateTypeName() string {
	return "pkg/sentry/kernel/epoll.pollEntry"
}

func (p *pollEntry) StateFields() []string {
	return []string{
		"pollEntryEntry",
		"id",
		"userData",
		"waiter",
		"mask",
		"flags",
		"epoll",
		"readySeq",
	}
}

func (p *pollEntry) beforeSave() {}

// +checklocksignore
func (p *pollEntry) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.pollEntryEntry)
	stateSinkObject.Save(1, &p.id)
	stateSinkObject.Save(2, &p.userData)
	stateSinkObject.Save(3, &p.waiter)
	stateSinkObject.Save(4, &p.mask)
	stateSinkObject.Save(5, &p.flags)
	stateSinkObject.Save(6, &p.epoll)
	stateSinkObject.Save(7, &p.readySeq)
}

// +checklocksignore
func (p *pollEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.pollEntryEntry)
	stateSourceObject.LoadWait(1, &p.id)
	stateSourceObject.Load(2, &p.userData)
	stateSourceObject.Load(3, &p.waiter)
	stateSourceObject.Load(4, &p.mask)
	stateSourceObject.Load(5, &p.flags)
	stateSourceObject.Load(6, &p.epoll)
	stateSourceObject.Load(7, &p.readySeq)
	stateSourceObject.AfterLoad(p.afterLoad)
}

func (e *EventPoll) StateTypeName() string {
	return "pkg/sentry/kernel/epoll.EventPoll"
}

func (e *EventPoll) StateFields() []string {
	return []string{
		"Queue",
		"files",
		"readyList",
		"waitingList",
		"disabledList",
		"readySeq",
	}
}

func (e *EventPoll) beforeSave() {}

// +checklocksignore
func (e *EventPoll) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	if !state.IsZeroValue(&e.FilePipeSeek) {
		state.Failf("FilePipeSeek is %#v, expected zero", &e.FilePipeSeek)
	}
	if !state.IsZeroValue(&e.FileNotDirReaddir) {
		state.Failf("FileNotDirReaddir is %#v, expected zero", &e.FileNotDirReaddir)
	}
	if !state.IsZeroValue(&e.FileNoFsync) {
		state.Failf("FileNoFsync is %#v, expected zero", &e.FileNoFsync)
	}
	if !state.IsZeroValue(&e.FileNoopFlush) {
		state.Failf("FileNoopFlush is %#v, expected zero", &e.FileNoopFlush)
	}
	if !state.IsZeroValue(&e.FileNoIoctl) {
		state.Failf("FileNoIoctl is %#v, expected zero", &e.FileNoIoctl)
	}
	if !state.IsZeroValue(&e.FileNoMMap) {
		state.Failf("FileNoMMap is %#v, expected zero", &e.FileNoMMap)
	}
	stateSinkObject.Save(0, &e.Queue)
	stateSinkObject.Save(1, &e.files)
	stateSinkObject.Save(2, &e.readyList)
	stateSinkObject.Save(3, &e.waitingList)
	stateSinkObject.Save(4, &e.disabledList)
	stateSinkObject.Save(5, &e.readySeq)
}

// +checklocksignore
func (e *EventPoll) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Queue)
	stateSourceObject.Load(1, &e.files)
	stateSourceObject.Load(2, &e.readyList)
	stateSourceObject.Load(3, &e.waitingList)
	stateSourceObject.Load(4, &e.disabledList)
	stateSourceObject.Load(5, &e.readySeq)
	stateSourceObject.AfterLoad(e.afterLoad)
}

func (l *pollEntryList) StateTypeName() string {
	return "pkg/sentry/kernel/epoll.pollEntryList"
}

func (l *pollEntryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *pollEntryList) beforeSave() {}

// +checklocksignore
func (l *pollEntryList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *pollEntryList) afterLoad() {}

// +checklocksignore
func (l *pollEntryList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *pollEntryEntry) StateTypeName() string {
	return "pkg/sentry/kernel/epoll.pollEntryEntry"
}

func (e *pollEntryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *pollEntryEntry) beforeSave() {}

// +checklocksignore
func (e *pollEntryEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *pollEntryEntry) afterLoad() {}

// +checklocksignore
func (e *pollEntryEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*FileIdentifier)(nil))
	state.Register((*pollEntry)(nil))
	state.Register((*EventPoll)(nil))
	state.Register((*pollEntryList)(nil))
	state.Register((*pollEntryEntry)(nil))
}
