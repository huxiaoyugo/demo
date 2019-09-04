// +build windows

package osutil

import (
	"os"
)

type SignalHandleFunc func(sig os.Signal) (ret bool)

func ListenSignal(handler SignalHandleFunc, signals ...os.Signal) {
	select {}
}

func ListenQuitAndDump() {
	select {}
}

func QuitAndDump(sig os.Signal) bool {
	return false
}
