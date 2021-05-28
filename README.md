# gosfmt

gosfmt is a pseudo random number generator(PRNG) based on SFMT. It was originally implemented in golang by gosl. Due to lightweight requirements, we only copy the rnd function from gosl.

[SFMT](http://godoc.org/github.com/qiangxue/fasthttp-routing)
[gosl](https://github.com/cpmech/gosl/)

## Install

```
go get -u -v github.com/bcowtech/gosfmt
```

## Usage

Let's start with a trivial example:

```go
package main

import (
	"fmt"

	"github.com/bcowtech/gosfmt"
)

func main() {
	// MTinit initialises random numbers generators (Mersenne Twister code)
	//  Input:
	//   seed -- seed value; use seed <= 0 to use current time
	gosfmt.MTinit(4321)

	// MTint generates pseudo random integer between low and high using the Mersenne Twister method.
	//  Input:
	//   low  -- lower limit
	//   high -- upper limit
	//  Output:
	//   random integer
	sfmtINT := gosfmt.MTint(1, 10)
	fmt.Println("sfmtINT  : ", sfmtINT)

	// MTints generates pseudo random integers between low and high using the Mersenne Twister method.
	//  Input:
	//   low    -- lower limit
	//   high   -- upper limit
	//  Output:
	//   values -- slice to be filled with len(values) numbers
	sfmtINTs := make([]int, 10)
	gosfmt.MTints(sfmtINTs, 1, 10)
	fmt.Println("sfmtINTs : ", sfmtINTs)

	// MTfloat64 generates pseudo random real numbers between low and high; i.e. in [low, right)
	// using the Mersenne Twister method.
	//  Input:
	//   low  -- lower limit (closed)
	//   high -- upper limit (open)
	//  Output:
	//   random float64
	sfmtF := gosfmt.MTfloat64(1, 10)
	fmt.Println("sfmtF    : ", sfmtF)

	// MTfloat64s generates pseudo random real numbers between low and high; i.e. in [low, right)
	// using the Mersenne Twister method.
	//  Input:
	//   low  -- lower limit (closed)
	//   high -- upper limit (open)
	//  Output:
	//   values -- slice to be filled with len(values) numbers
	sfmtFs := make([]float64, 10)
	gosfmt.MTfloat64s(sfmtFs, 1, 10)
	fmt.Println("sfmtFs   : ", sfmtFs)

	// MTintShuffle shuffles a slice of integers using Mersenne Twister algorithm.
	gosfmt.MTintShuffle(sfmtINTs)
	fmt.Println("sfmtINTs : ", sfmtINTs)
}

```

Output:

```
sfmtINT  :  1
sfmtINTs :  [5 1 3 9 7 4 6 2 1 4]
sfmtF    :  9.722344331392243
sfmtFs   :  [1.9189275176601612 8.615983236661565 3.5797699181835085 4.093636478513859 3.8370310131668512 5.235095756892996 4.0630003969818675 5.4363333680139565 2.0028628090833234 7.933939616108609]
sfmtINTs :  [1 6 4 4 3 5 7 1 2 9]
```
