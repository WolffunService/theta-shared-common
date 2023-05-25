package lockkey

import (
	"strings"
)

const (
	lockPrefix = "LOCK"

	// Watcher's locks
	LockDocumentToProcess = "Watcher"
)

// GetKey cách dùng:
//
//	 docID := "626wjhfkwfhkwfhj"
//	 key := lockkey.GetKey(lockKey.LockDocumentToProcess, docID)
//		lock, err := thetanlock.Lock(key)
//		if err != nil {
//			return err
//		}
//		defer lock.Unlock()
//
//	 Output: key="LOCK:Watcher:626wjhfkwfhkwfhj"
func GetKey(subNames ...string) string {
	return lockPrefix + ":" + strings.Join(subNames, ":")
}
