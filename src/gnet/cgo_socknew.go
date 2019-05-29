// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +/*build cgo,!netgo*/
// +/*build android linux solaris*/

package gnet

/*
#include <sys/types.h>
#include <sys/socket.h>

#include <netinet/in.h>
*/
//import "C"

import (
	"syscall"
	"unsafe"
)

func cgoSockaddrInet4(ip IP) *struct_sockaddr {
	sa := syscall.RawSockaddrInet4{Family: syscall.AF_INET}
	copy(sa.Addr[:], ip)
	return (*struct_sockaddr)(unsafe.Pointer(&sa))
}

func cgoSockaddrInet6(ip IP, zone int) *struct_sockaddr {
	sa := syscall.RawSockaddrInet6{Family: syscall.AF_INET6, Scope_id: uint32(zone)}
	copy(sa.Addr[:], ip)
	return (*struct_sockaddr)(unsafe.Pointer(&sa))
}