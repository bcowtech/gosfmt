// Copyright 2015 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "stdio.h"
#include "SFMT.h"

#ifdef WIN32
#define LONG long long
#else
#define LONG long
#endif

void Seed(sfmt_t * sfmt, LONG seed)
{
    sfmt_init_gen_rand(sfmt,seed);
}

unsigned long Uint64(sfmt_t * sfmt)
{
    return sfmt_genrand_uint64(sfmt);
}
