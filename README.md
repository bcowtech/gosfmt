# gosfmt

gosfmt is a pseudo random number generator(PRNG) based on SFMT. It was originally implemented in golang by gosl. Due to lightweight requirements, we only copy the rnd function from gosl.

[SFMT](http://godoc.org/github.com/qiangxue/fasthttp-routing)

[gosl](https://github.com/cpmech/gosl/)


## Install

```console
go get -u -v github.com/bcowtech/gosfmt
```

## Usage

Let's start with a trivial example:

```go

package main

import (
	"time"

	"github.com/bcowtech/gosfmt"
)

func main() {
	rng := gosfmt.NewSFMT()
	rng.Seed(time.Now().UnixNano())
}


```

Output:

```


```

## Benckmark

```console

go test -benchmem -run=^$ -coverprofile=/var/folders/xy/g72c6j8n34z9lzxx46rk5d1h0000gn/T/vscode-goMk3tDo/go-code-cover -bench . github.com/bcowtech/gosfmt

goos: darwin
goarch: amd64
pkg: github.com/bcowtech/gosfmt
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_gornd_int-12                       	   73124	     15003 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_int-12                        	   17148	     68199 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_int-12                          	   66864	     17668 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63Threadsafe-12            	81340507	        14.74 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63Threadsafe-12            	18778996	        63.21 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63Threadsafe-12              	75479720	        15.95 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63ThreadsafeParallel-12    	21975541	        54.63 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63ThreadsafeParallel-12    	10353994	       116.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63ThreadsafeParallel-12      	21781872	        54.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63Unthreadsafe-12          	79371351	        14.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63Unthreadsafe-12          	18020554	        66.98 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63Unthreadsafe-12            	72136144	        16.85 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Intn1000-12                   	60030170	        20.20 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Intn1000-12                   	16740956	        71.64 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Intn1000-12                     	55122488	        22.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63n1000-12                 	48273963	        24.38 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63n1000-12                 	15962689	        75.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63n1000-12                   	45048855	        27.08 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31n1000-12                 	63473118	        18.87 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31n1000-12                 	16773116	        76.07 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int31n1000-12                   	57507925	        20.88 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32-12                    	70721752	        17.01 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float32-12                    	16940911	        69.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32-12                      	64236172	        18.81 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64-12                    	78786271	        15.47 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64-12                    	17580511	        67.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64-12                      	70578664	        16.82 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Perm3-12                      	19107808	        61.37 ns/op	      24 B/op	       1 allocs/op
Benchmark_sfmt_Perm3-12                      	 4849052	       216.7 ns/op	      24 B/op	       1 allocs/op
Benchmark_mt_Perm3-12                        	17344826	        67.12 ns/op	      24 B/op	       1 allocs/op
Benchmark_rand_Perm30-12                     	 2704623	       416.2 ns/op	     240 B/op	       1 allocs/op
Benchmark_sfmt_Perm30-12                     	  634135	      1959 ns/op	     240 B/op	       1 allocs/op
Benchmark_mt_Perm30-12                       	 2658220	       454.0 ns/op	     240 B/op	       1 allocs/op
Benchmark_rand_Perm30ViaShuffle-12           	 4936770	       240.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Perm30ViaShuffle-12           	  698182	      1713 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Perm30ViaShuffle-12             	 4131399	       292.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_ShuffleOverhead-12            	 3236700	       368.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_ShuffleOverhead-12            	  410712	      2950 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_ShuffleOverhead-12              	 2683315	       448.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Read3-12                      	62283302	        18.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Read3-12                      	28855190	        42.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Read3-12                        	58494046	        20.42 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Read64-12                     	22058172	        53.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Read64-12                     	 2215095	       544.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Read64-12                       	13136886	        91.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Read1000-12                   	 1965436	       611.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Read1000-12                   	  145444	      8698 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Read1000-12                     	 1000000	      1160 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63r1000-12                 	79486406	        14.97 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63r1000-12                 	18196862	        68.66 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63r1000-12                   	71501119	        16.69 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63s30-12                   	 1864542	       649.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63s30-12                   	  550465	      2162 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63s30-12                     	 1721319	       701.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63Shuffle30-12             	 1964617	       612.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63Shuffle30-12             	  565492	      2105 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63huffle30-12                	 1805472	       666.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint32r1000-12                	79895049	        15.01 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint32r1000-12                	18087141	        73.22 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint63r1000-12                  	71377842	        16.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint32s30-12                  	 2208178	       545.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint32s30-12                  	  578996	      2066 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint32s30-12                    	 2016622	       596.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint32Shuffle30-12            	 1963398	       611.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint32Shuffle30-12            	  511730	      2088 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint32Shuffle30-12              	 1812344	       664.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint64r1000-12                	74847852	        16.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint64r1000-12                	17687851	        70.63 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint64r1000-12                  	67786096	        17.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint64s30-12                  	 1987566	       609.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint64s30-12                  	  567224	      2112 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint64s30-12                    	 1836693	       649.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint64Shuffle30-12            	 2025249	       595.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint64Shuffle30-12            	  587175	      2056 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint64Shuffle30-12              	 2007655	       599.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31r1000-12                 	81204918	        14.62 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31r1000-12                 	18714445	        63.37 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int31r1000-12                   	74183540	        16.52 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31s30-12                   	 2421219	       497.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31s30-12                   	  600738	      2082 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int3130-12                      	 2226039	       541.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31Shuffle30-12             	 1988466	       600.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31Shuffle30-12             	  589308	      2035 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int31Shuffle30-12               	 1847608	       661.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Intr1000-12                   	77459973	        15.19 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Intr1000-12                   	17899402	        70.94 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Intr1000-12                     	73752874	        16.69 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Ints30-12                     	 1865827	       626.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Ints30-12                     	  567535	      2334 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Ints30-12                       	 1761175	       684.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_IntShuffle30-12               	 1988996	       589.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_IntShuffle30-12               	  593402	      2024 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_IntShuffle30-12                 	 1854183	       656.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64r1000-12               	76024038	        15.37 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64r1000-12               	18065575	        68.41 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64r1000-12                 	68939498	        17.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64s30-12                 	 2491684	       481.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64s30-12                 	  597176	      2030 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64s0-12                    	 2234696	       544.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64Shuffle30-12           	 1926850	       629.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64Shuffle30-12           	  525816	      2149 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64Shuffle30-12             	 1759837	       685.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32r1000-12               	69166971	        17.60 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float32r1000-12               	15508828	        70.33 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32r1000-12                 	62842567	        19.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32s30-12                 	 2243454	       543.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float3230-12                  	  572166	      2129 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32s0-12                    	 1941075	       597.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32Shuffle30-12           	 1736551	       717.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float32Shuffle30-12           	  565240	      2112 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32Shuffle30-12             	 1633357	       771.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_FlipCoin-12                   	36768331	        33.49 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_FlipCoin-12                   	 8740948	       151.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_FlipCoin-12                     	30733249	        43.30 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 68.4% of statements
ok  	github.com/bcowtech/gosfmt	162.510s

```
