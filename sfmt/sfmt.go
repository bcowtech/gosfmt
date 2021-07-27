// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sfmt wraps the SFMT SIMD-oriented Fast Mersenne Twister
package sfmt

/*
#cgo CFLAGS: -O3 -fomit-frame-pointer -DNDEBUG -fno-strict-aliasing -std=c99 -msse2 -DHAVE_SSE2 -DSFMT_MEXP=19937
#ifdef WIN32
#define LONG long long
#else
#define LONG long
#endif
#include "SFMT.h"
*/
import "C"

type SFMT19937 struct {
	csfmt *C.sfmt_t
}

func New() *SFMT19937 {
	sfmt19937 := &SFMT19937{
		csfmt: &C.sfmt_t{},
	}

	sfmt19937.Seed(5489)
	return sfmt19937
}

// Seed uses the given 64bit value to initialise the generator state.
// This method is part of the rand.Source interface.
func (sfmt *SFMT19937) Seed(seed int64) {
	C.sfmt_init_gen_rand(sfmt.csfmt, C.uint32_t(seed))
}

// Uint64 generates a (pseudo-)random 64bit value.  The output can be
// used as a replacement for a sequence of independent, uniformly
// distributed samples in the range 0, 1, ..., 2^64-1.  This method is
// part of the rand.Source64 interface.
func (sfmt *SFMT19937) Uint64() uint64 {
	return uint64(C.sfmt_genrand_uint64(sfmt.csfmt))
}

// Int63 generates a (pseudo-)random 63bit value.  The output can be
// used as a replacement for a sequence of independent, uniformly
// distributed samples in the range 0, 1, ..., 2^63-1.  This method is
// part of the rand.Source interface.
func (sfmt *SFMT19937) Int63() int64 {
	return int64(C.sfmt_genrand_uint64(sfmt.csfmt) & 0x7fffffffffffffff)
}
