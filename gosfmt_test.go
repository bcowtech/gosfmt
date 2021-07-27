package gosfmt_test

import (
	"testing"

	"github.com/bcowtech/gosfmt"
	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/stretchr/testify/assert"
)

var (
	seed5489 = int64(5489)
	gen5489  = uint64(226931099713899959)
)

func init() {
	io.Verbose = false
	chk.Verbose = false
	// verbose()
}

func verbose() {
	io.Verbose = true
	chk.Verbose = true
}

func Test_sfmt01_Int(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt01. Int")
	assert := assert.New(t)

	gosfmt.Seed(seed5489)
	assert.Equal(uint64(gen5489), gosfmt.Uint64())

	gosfmt.Seed(seed5489)
	assert.Equal(int64(gen5489&0x7fffffffffffffff), gosfmt.Int63())

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int()
		io.Pforan("gosfmt.Int  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Intn(1000)
		io.Pforan("gosfmt.Intn = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Intr(-100, 100)
		io.Pforan("gosfmt.Intr = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int, 3)
		gosfmt.Ints(gen, -100, 100)
		io.Pforan("gosfmt.Ints = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Perm(3)
		io.Pforan("gosfmt.Perm = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int, 3)
		for i := range gen {
			gen[i] = int(i)
		}
		gosfmt.IntShuffle(gen)
		io.Pforan("gosfmt.IntShuffle = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int, 3)
		for i := range gen {
			gen[i] = i
		}
		gosfmt.Shuffle(3, func(i, j int) { gen[i], gen[j] = gen[j], gen[i] })
		io.Pforan("gosfmt.Shuffle = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	buf := make([]byte, 3)
	gen, _ := gosfmt.Read(buf)
	io.Pforan("gosfmt.Read = %v\n", gen)
}

func Test_sfmt02_Int63(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt02. Int63")

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int63()
		io.Pforan("gosfmt.Int63  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int63n(1000)
		io.Pforan("gosfmt.Int63n = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int63r(0, 1000)
		io.Pforan("gosfmt.Int63r = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int, 3)
		gosfmt.Ints(gen, 0, 1000)
		io.Pforan("gosfmt.Int63s = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int64, 3)
		for i := range gen {
			gen[i] = int64(i)
		}
		gosfmt.Int63Shuffle(gen)
		io.Pforan("gosfmt.Int63Shuffle = %v\n", gen)
	}
}

func Test_sfmt03_Uint32(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt03. Uint32")

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Uint32()
		io.Pforan("gosfmt.Uint32  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Uint32r(0, 1000)
		io.Pforan("gosfmt.Uint32r = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]uint32, 3)
		gosfmt.Uint32s(gen, 0, 1000)
		io.Pforan("gosfmt.Uint32s = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]uint32, 3)
		for i := range gen {
			gen[i] = uint32(i)
		}
		gosfmt.Uint32Shuffle(gen)
		io.Pforan("gosfmt.Int63Shuffle = %v\n", gen)
	}
}

func Test_sfmt04_Uint64(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt04. Uint64")

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Uint64()
		io.Pforan("gosfmt.Uint64  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Uint64r(0, 1000)
		io.Pforan("gosfmt.Uint64r = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]uint64, 3)
		gosfmt.Uint64s(gen, 0, 1000)
		io.Pforan("gosfmt.Uint64s = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]uint64, 3)
		for i := range gen {
			gen[i] = uint64(i)
		}
		gosfmt.Uint64Shuffle(gen)
		io.Pforan("gosfmt.Int63Shuffle = %v\n", gen)
	}
}

func Test_sfmt05_Int31(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt05. Int32")

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int31()
		io.Pforan("gosfmt.Int31  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int31n(1000)
		io.Pforan("gosfmt.Int31n = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int31r(-100, 100)
		io.Pforan("gosfmt.Int31r = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int32, 3)
		gosfmt.Int31s(gen, -100, 100)
		io.Pforan("gosfmt.Int31s = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int32, 3)
		for i := range gen {
			gen[i] = int32(i)
		}
		gosfmt.Int31Shuffle(gen)
		io.Pforan("gosfmt.Int31Shuffle = %v\n", gen)
	}
}

func Test_sfmt06_Int63(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt06. Int63")

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int63()
		io.Pforan("gosfmt.Int63  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int63n(1000)
		io.Pforan("gosfmt.Int63n = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Int63r(-100, 100)
		io.Pforan("gosfmt.Int63r = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int64, 3)
		gosfmt.Int63s(gen, -100, 100)
		io.Pforan("gosfmt.Int63s = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]int64, 3)
		for i := range gen {
			gen[i] = int64(i)
		}
		gosfmt.Int63Shuffle(gen)
		io.Pforan("gosfmt.Int63Shuffle = %v\n", gen)
	}
}

func Test_sfmt07_Float64(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt07. Float64")

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Float64()
		io.Pforan("gosfmt.Float64  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Float64r(-100.0, 100.0)
		io.Pforan("gosfmt.Float64r = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]float64, 3)
		gosfmt.Float64s(gen, -100.0, 100.0)
		io.Pforan("gosfmt.Float64s = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]float64, 10)
		for i := range gen {
			gen[i] = float64(i) / 10.0
		}
		gosfmt.Float64Shuffle(gen)
		io.Pforan("gosfmt.Float64Shuffle = %v\n", gen)
		coin := make([]bool, len(gen))
		for i := range coin {
			coin[i] = gosfmt.FlipCoin(gen[i])
		}
		io.Pforan("gosfmt.FlipCoin = %v\n", coin)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.NormFloat64()
		io.Pforan("gosfmt.NormFloat64  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.ExpFloat64()
		io.Pforan("gosfmt.ExpFloat64  = %v\n", gen)
	}
}

func Test_sfmt08_Float32(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt08. Float32")

	nsamples := 10
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Float32()
		io.Pforan("gosfmt.Float32  = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := gosfmt.Float32r(-100.0, 100.0)
		io.Pforan("gosfmt.Float32r = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]float32, 3)
		gosfmt.Float32s(gen, -100.0, 100.0)
		io.Pforan("gosfmt.Float32s = %v\n", gen)
	}
	gosfmt.Seed(seed5489)
	for i := 0; i < nsamples; i++ {
		gen := make([]float32, 10)
		for i := range gen {
			gen[i] = float32(i) / 10.0
		}
		gosfmt.Float32Shuffle(gen)
		io.Pforan("gosfmt.Float32Shuffle = %v\n", gen)
	}
}
