// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,darwin,dragonfly,freebsd,netbsd,openbsd,!appengine

package siegreader

import (
	"log"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func mmapable(sz int64) bool {
	if int64(int(sz+4095)) != sz+4095 {
		return false
	}
	return true
}

func mmapFile(f *os.File, sz int64) []byte {
	st, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := st.Size()
	// mmapable call goes here
	n := int(size)
	if n == 0 {
		return nil
	}
	data, err := syscall.Mmap(int(f.Fd()), 0, (n+4095)&^4095, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("mmap %s: %v", f.Name(), err)
	}
	return data[:n]
}