package utils

import "sync"

// SyncMapLength returns the length of syncMap.
func SyncMapLength(syncMap sync.Map) int {
	var length int

	syncMap.Range(func(k, v interface{}) bool {
		length++
		return true
	})

	return length
}
