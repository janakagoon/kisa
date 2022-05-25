package lib

// #include <stdlib.h>
// #include "relay.h"
import "C"

import (
	"errors"
	"sync"
)

type RelayStatus int

var writeLock sync.Mutex
var readLock sync.Mutex

const (
	RelayOff RelayStatus = 0
	RelayOn  RelayStatus = 1
)

func UnsafeSetRelay(board, relay int64, status RelayStatus) error {

	boardCInt := C.int(board)

	relayCInt := C.int(relay)

	statusCBool := C.int(status)

	writeLock.Lock()
	errorCode := C.set_relay(boardCInt, relayCInt, statusCBool)
	writeLock.Unlock()

	if errorCode != 0 {
		return errors.New("error calling set_relay()")
	}

	return nil
}

func GetRelay(board, relay int64) (RelayStatus, error) {
	boardCInt := C.int(board)

	relayCInt := C.int(relay)

	statusCInt := C.int(0)

	readLock.Lock()
	errorCode := C.get_relay(boardCInt, relayCInt, &statusCInt)
	readLock.Unlock()

	if errorCode != 0 {
		return RelayOff, errors.New("error calling get_relay()")
	}

	result := RelayOff
	if statusCInt == C.int(1) {
		result = RelayOn
	}

	return result, nil
}
