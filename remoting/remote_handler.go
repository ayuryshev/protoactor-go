package remoting

import "github.com/AsynkronIT/protoactor-go/process"

func remoteHandler(pid *process.PID) (process.Process, bool) {
	ref := newRemoteProcess(pid)
	return ref, true
}
