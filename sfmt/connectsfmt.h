// Copyright 2015 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
#include "SFMT.h"

#ifndef GOSL_RANDOM_H
#define GOSL_RANDOM_H

#ifdef WIN32
#define LONG long long
#else
#define LONG long
#endif

void Seed(sfmt_t * sfmt, LONG seed);
unsigned long Uint64(sfmt_t * sfmt);

#endif // GOSL_RANDOM_H
