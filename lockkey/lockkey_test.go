package lockkey

import (
	"fmt"
	"testing"
)

func TestGetKey(t *testing.T) {
	key1 := GetKey(LockDocumentToProcess, "623746527")
	fmt.Println(key1)
	// output: LOCK:Watcher:623746527

	key2 := GetKey("abc", "xyz", "mnp")
	fmt.Println(key2)
	// output: LOCK:abc:xyz:mnp

	key3 := GetKey("")
	fmt.Println(key3)
	// output: LOCK:

	key4 := GetKey("", "")
	fmt.Println(key4)
	// output: LOCK::
}
