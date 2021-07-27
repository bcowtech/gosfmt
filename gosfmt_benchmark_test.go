package gosfmt_test

import (
	"math/rand"
	"testing"

	"github.com/bcowtech/gosfmt"
	"github.com/seehuhn/mt19937"
)

var benchNsamples int64
var benchResult int64
var gorand = gosfmt.New(rand.NewSource(5489))
var mt19937rand = gosfmt.New(mt19937.New())

func init() {
	gorand.Seed(5489)
	gosfmt.Seed(5489)
	mt19937rand.Seed(5489)
	benchNsamples = 1000
}

func Benchmark_gornd_int(b *testing.B) {
	var res int64
	lo, hi := 0, 50
	for i := 0; i < b.N; i++ {
		for j := int64(0); j < benchNsamples; j++ {
			res = int64(gorand.Int()%(hi-lo+1) + lo)
		}
	}
	benchResult = res
}

func Benchmark_sfmt_int(b *testing.B) {
	var res int64
	lo, hi := 0, 50
	for i := 0; i < b.N; i++ {
		for j := int64(0); j < benchNsamples; j++ {
			res = int64(gosfmt.Int()%(hi-lo+1) + lo)
		}
	}
	benchResult = res
}

func Benchmark_mt_int(b *testing.B) {
	var res int64
	lo, hi := 0, 50
	for i := 0; i < b.N; i++ {
		for j := int64(0); j < benchNsamples; j++ {
			res = int64(mt19937rand.Int()%(hi-lo+1) + lo)
		}
	}
	benchResult = res
}

// Benchmarks -

func Benchmark_rand_Int63Threadsafe(b *testing.B) {
	for n := b.N; n > 0; n-- {
		gorand.Int63()
	}
}

func Benchmark_sfmt_Int63Threadsafe(b *testing.B) {
	for n := b.N; n > 0; n-- {
		gosfmt.Int63()
	}
}

func Benchmark_mt_Int63Threadsafe(b *testing.B) {
	for n := b.N; n > 0; n-- {
		mt19937rand.Int63()
	}
}

func Benchmark_rand_Int63ThreadsafeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gorand.Int63()
		}
	})
}

func Benchmark_sfmt_Int63ThreadsafeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gosfmt.Int63()
		}
	})
}

func Benchmark_mt_Int63ThreadsafeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mt19937rand.Int63()
		}
	})
}

// Benchmarks - New

func Benchmark_rand_Int63Unthreadsafe(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Int63()
	}
}

func Benchmark_sfmt_Int63Unthreadsafe(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Int63()
	}
}

func Benchmark_mt_Int63Unthreadsafe(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Int63()
	}
}

func Benchmark_rand_Intn1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Intn(1000)
	}
}

func Benchmark_sfmt_Intn1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Intn(1000)
	}
}

func Benchmark_mt_Intn1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Intn(1000)
	}
}

func Benchmark_rand_Int63n1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Int63n(1000)
	}
}

func Benchmark_sfmt_Int63n1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Int63n(1000)
	}
}

func Benchmark_mt_Int63n1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Int63n(1000)
	}
}

func Benchmark_rand_Int31n1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Int31n(1000)
	}
}

func Benchmark_sfmt_Int31n1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Int31n(1000)
	}
}

func Benchmark_mt_Int31n1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Int31n(1000)
	}
}

func Benchmark_rand_Float32(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Float32()
	}
}

func Benchmark_sfmt_Float32(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Float32()
	}
}

func Benchmark_mt_Float32(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Float32()
	}
}

func Benchmark_rand_Float64(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Float64()
	}
}

func Benchmark_sfmt_Float64(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Float64()
	}
}

func Benchmark_mt_Float64(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Float64()
	}
}

func Benchmark_rand_Perm3(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Perm(3)
	}
}

func Benchmark_sfmt_Perm3(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Perm(3)
	}
}

func Benchmark_mt_Perm3(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Perm(3)
	}
}

func Benchmark_rand_Perm30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Perm(30)
	}
}

func Benchmark_sfmt_Perm30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Perm(30)
	}
}

func Benchmark_mt_Perm30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Perm(30)
	}
}

func Benchmark_rand_Perm30ViaShuffle(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		for i := range p {
			p[i] = i
		}
		r.Shuffle(30, func(i, j int) { p[i], p[j] = p[j], p[i] })
	}
}

func Benchmark_sfmt_Perm30ViaShuffle(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		for i := range p {
			p[i] = i
		}
		r.Shuffle(30, func(i, j int) { p[i], p[j] = p[j], p[i] })
	}
}

func Benchmark_mt_Perm30ViaShuffle(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		for i := range p {
			p[i] = i
		}
		r.Shuffle(30, func(i, j int) { p[i], p[j] = p[j], p[i] })
	}
}

// Benchmark_sfmt_ShuffleOverhead uses a minimal swap function
// to measure just the shuffling overhead.

func Benchmark_rand_ShuffleOverhead(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Shuffle(52, func(i, j int) {
			if i < 0 || i >= 52 || j < 0 || j >= 52 {
				b.Fatalf("bad swap(%d, %d)", i, j)
			}
		})
	}
}

func Benchmark_sfmt_ShuffleOverhead(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Shuffle(52, func(i, j int) {
			if i < 0 || i >= 52 || j < 0 || j >= 52 {
				b.Fatalf("bad swap(%d, %d)", i, j)
			}
		})
	}
}

func Benchmark_mt_ShuffleOverhead(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Shuffle(52, func(i, j int) {
			if i < 0 || i >= 52 || j < 0 || j >= 52 {
				b.Fatalf("bad swap(%d, %d)", i, j)
			}
		})
	}
}

func Benchmark_rand_Read3(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	buf := make([]byte, 3)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_sfmt_Read3(b *testing.B) {
	r := gosfmt.NewSFMT()
	buf := make([]byte, 3)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_mt_Read3(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	buf := make([]byte, 3)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_rand_Read64(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	buf := make([]byte, 64)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_sfmt_Read64(b *testing.B) {
	r := gosfmt.NewSFMT()
	buf := make([]byte, 64)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_mt_Read64(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	buf := make([]byte, 64)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_rand_Read1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	buf := make([]byte, 1000)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_sfmt_Read1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	buf := make([]byte, 1000)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func Benchmark_mt_Read1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	buf := make([]byte, 1000)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

// gosl method Benchmarks

func Benchmark_rand_Int63r1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Int63r(0, 1000)
	}
}

func Benchmark_sfmt_Int63r1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Int63r(0, 1000)
	}
}

func Benchmark_mt_Int63r1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Int63r(0, 1000)
	}
}

func Benchmark_rand_Int63s30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]int64, 30)
		r.Int63s(p, 0, 1000)
	}
}

func Benchmark_sfmt_Int63s30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]int64, 30)
		r.Int63s(p, 0, 1000)
	}
}

func Benchmark_mt_Int63s30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]int64, 30)
		r.Int63s(p, 0, 1000)
	}
}

func Benchmark_rand_Int63Shuffle30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]int64, 30)
		for i := range p {
			p[i] = int64(i)
		}
		r.Int63Shuffle(p)
	}
}

func Benchmark_sfmt_Int63Shuffle30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]int64, 30)
		for i := range p {
			p[i] = int64(i)
		}
		r.Int63Shuffle(p)
	}
}

func Benchmark_mt_Int63huffle30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]int64, 30)
		for i := range p {
			p[i] = int64(i)
		}
		r.Int63Shuffle(p)
	}
}

func Benchmark_rand_Uint32r1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Uint32r(0, 1000)
	}
}

func Benchmark_sfmt_Uint32r1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Uint32r(0, 1000)
	}
}

func Benchmark_mt_Uint32r1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Uint32r(0, 1000)
	}
}

func Benchmark_rand_Uint32s30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]uint32, 30)
		r.Uint32s(p, 0, 1000)
	}
}

func Benchmark_sfmt_Uint32s30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]uint32, 30)
		r.Uint32s(p, 0, 1000)
	}
}

func Benchmark_mt_Uint32s30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]uint32, 30)
		r.Uint32s(p, 0, 1000)
	}
}

func Benchmark_rand_Uint32Shuffle30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]uint32, 30)
		for i := range p {
			p[i] = uint32(i)
		}
		r.Uint32Shuffle(p)
	}
}

func Benchmark_sfmt_Uint32Shuffle30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]uint32, 30)
		for i := range p {
			p[i] = uint32(i)
		}
		r.Uint32Shuffle(p)
	}
}

func Benchmark_mt_Uint32Shuffle30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]uint32, 30)
		for i := range p {
			p[i] = uint32(i)
		}
		r.Uint32Shuffle(p)
	}
}

func Benchmark_rand_Uint64r1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Uint64r(0, 1000)
	}
}

func Benchmark_sfmt_Uint64r1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Uint64r(0, 1000)
	}
}

func Benchmark_mt_Uint64r1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Uint64r(0, 1000)
	}
}

func Benchmark_rand_Uint64s30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]uint64, 30)
		r.Uint64s(p, 0, 1000)
	}
}

func Benchmark_sfmt_Uint64s30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]uint64, 30)
		r.Uint64s(p, 0, 1000)
	}
}

func Benchmark_mt_Uint64s30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]uint64, 30)
		r.Uint64s(p, 0, 1000)
	}
}

func Benchmark_rand_Uint64Shuffle30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]uint64, 30)
		for i := range p {
			p[i] = uint64(i)
		}
		r.Uint64Shuffle(p)
	}
}

func Benchmark_sfmt_Uint64Shuffle30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]uint64, 30)
		for i := range p {
			p[i] = uint64(i)
		}
		r.Uint64Shuffle(p)
	}
}

func Benchmark_mt_Uint64Shuffle30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]uint64, 30)
		for i := range p {
			p[i] = uint64(i)
		}
		r.Uint64Shuffle(p)
	}
}

func Benchmark_rand_Int31r1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Int31r(0, 1000)
	}
}

func Benchmark_sfmt_Int31r1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Int31r(0, 1000)
	}
}

func Benchmark_mt_Int31r1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Int31r(0, 1000)
	}
}

func Benchmark_rand_Int31s30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]int32, 30)
		r.Int31s(p, 0, 1000)
	}
}

func Benchmark_sfmt_Int31s30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]int32, 30)
		r.Int31s(p, 0, 1000)
	}
}

func Benchmark_mt_Int3130(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]int32, 30)
		r.Int31s(p, 0, 1000)
	}
}

func Benchmark_rand_Int31Shuffle30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]int32, 30)
		for i := range p {
			p[i] = int32(i)
		}
		r.Int31Shuffle(p)
	}
}

func Benchmark_sfmt_Int31Shuffle30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]int32, 30)
		for i := range p {
			p[i] = int32(i)
		}
		r.Int31Shuffle(p)
	}
}

func Benchmark_mt_Int31Shuffle30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]int32, 30)
		for i := range p {
			p[i] = int32(i)
		}
		r.Int31Shuffle(p)
	}
}

func Benchmark_rand_Intr1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Intr(0, 1000)
	}
}

func Benchmark_sfmt_Intr1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Intr(0, 1000)
	}
}

func Benchmark_mt_Intr1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Intr(0, 1000)
	}
}

func Benchmark_rand_Ints30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		r.Ints(p, 0, 1000)
	}
}

func Benchmark_sfmt_Ints30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		r.Ints(p, 0, 1000)
	}
}

func Benchmark_mt_Ints30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		r.Ints(p, 0, 1000)
	}
}

func Benchmark_rand_IntShuffle30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		for i := range p {
			p[i] = int(i)
		}
		r.IntShuffle(p)
	}
}

func Benchmark_sfmt_IntShuffle30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		for i := range p {
			p[i] = int(i)
		}
		r.IntShuffle(p)
	}
}

func Benchmark_mt_IntShuffle30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		for i := range p {
			p[i] = int(i)
		}
		r.IntShuffle(p)
	}
}

func Benchmark_rand_Float64r1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Float64r(0.001, 1000)
	}
}

func Benchmark_sfmt_Float64r1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Float64r(0.001, 1000)
	}
}

func Benchmark_mt_Float64r1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Float64r(0.001, 1000)
	}
}

func Benchmark_rand_Float64s30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]float64, 30)
		r.Float64s(p, 0.001, 1000)
	}
}

func Benchmark_sfmt_Float64s30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]float64, 30)
		r.Float64s(p, 0.001, 1000)
	}
}

func Benchmark_mt_Float64s0(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]float64, 30)
		r.Float64s(p, 0.001, 1000)
	}
}

func Benchmark_rand_Float64Shuffle30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]float64, 30)
		for i := range p {
			p[i] = float64(i)
		}
		r.Float64Shuffle(p)
	}
}

func Benchmark_sfmt_Float64Shuffle30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]float64, 30)
		for i := range p {
			p[i] = float64(i)
		}
		r.Float64Shuffle(p)
	}
}

func Benchmark_mt_Float64Shuffle30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]float64, 30)
		for i := range p {
			p[i] = float64(i)
		}
		r.Float64Shuffle(p)
	}
}

func Benchmark_rand_Float32r1000(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.Float32r(0.001, 1000)
	}
}

func Benchmark_sfmt_Float32r1000(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.Float32r(0.001, 1000)
	}
}

func Benchmark_mt_Float32r1000(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.Float32r(0.001, 1000)
	}
}

func Benchmark_rand_Float32s30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]float32, 30)
		r.Float32s(p, 0.001, 1000)
	}
}

func Benchmark_sfmt_Float3230(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]float32, 30)
		r.Float32s(p, 0.001, 1000)
	}
}

func Benchmark_mt_Float32s0(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]float32, 30)
		r.Float32s(p, 0.001, 1000)
	}
}

func Benchmark_rand_Float32Shuffle30(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		p := make([]float32, 30)
		for i := range p {
			p[i] = float32(i)
		}
		r.Float32Shuffle(p)
	}
}

func Benchmark_sfmt_Float32Shuffle30(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		p := make([]float32, 30)
		for i := range p {
			p[i] = float32(i)
		}
		r.Float32Shuffle(p)
	}
}

func Benchmark_mt_Float32Shuffle30(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		p := make([]float32, 30)
		for i := range p {
			p[i] = float32(i)
		}
		r.Float32Shuffle(p)
	}
}

func Benchmark_rand_FlipCoin(b *testing.B) {
	r := gosfmt.New(rand.NewSource(5489))
	for n := b.N; n > 0; n-- {
		r.FlipCoin(r.Float64())
	}
}

func Benchmark_sfmt_FlipCoin(b *testing.B) {
	r := gosfmt.NewSFMT()
	for n := b.N; n > 0; n-- {
		r.FlipCoin(r.Float64())
	}
}

func Benchmark_mt_FlipCoin(b *testing.B) {
	r := gosfmt.New(mt19937.New())
	for n := b.N; n > 0; n-- {
		r.FlipCoin(r.Float64())
	}
}
