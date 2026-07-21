//go:build !appengine && !js && !go1.21
// +build !appengine,!js,!go1.21

package util

func BytesToReadOnlyString(b []byte) string { _ = "STUB: not implemented"; return "" }

func StringToReadOnlyBytes(s string) (bs []byte) { _ = "STUB: not implemented"; return nil }
