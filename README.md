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

go test -benchmem -run=^$ -coverprofile=/var/folders/xy/g72c6j8n34z9lzxx46rk5d1h0000gn/T/vscode-goaou7XQ/go-code-cover -bench . github.com/bcowtech/gosfmt

goos: darwin
goarch: amd64
pkg: github.com/bcowtech/gosfmt
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_gornd_int-12                       	   73202	     15791 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_int-12                        	   15212	     71712 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_int-12                          	   67243	     18001 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63Threadsafe-12            	78682668	        14.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63Threadsafe-12            	18221886	        71.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63Threadsafe-12              	74292028	        16.15 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63ThreadsafeParallel-12    	21323655	        53.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63ThreadsafeParallel-12    	10321776	       117.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63ThreadsafeParallel-12      	22092015	        54.01 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63Unthreadsafe-12          	79980430	        14.49 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63Unthreadsafe-12          	18556160	        66.20 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63Unthreadsafe-12            	73620218	        16.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Intn1000-12                   	61189861	        19.52 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Intn1000-12                   	16257384	        76.74 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Intn1000-12                     	55045237	        21.70 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63n1000-12                 	49623834	        23.55 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63n1000-12                 	16625440	        73.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63n1000-12                   	44415165	        26.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31n1000-12                 	65823196	        18.20 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31n1000-12                 	17409052	        71.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int31n1000-12                   	59718427	        20.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32-12                    	72005583	        16.51 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float32-12                    	16364518	        72.03 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32-12                      	63803806	        18.58 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64-12                    	77025218	        15.10 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64-12                    	17934555	        68.44 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64-12                      	73565535	        16.65 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Perm3-12                      	19398682	        62.11 ns/op	      24 B/op	       1 allocs/op
Benchmark_sfmt_Perm3-12                      	 5744672	       216.5 ns/op	      24 B/op	       1 allocs/op
Benchmark_mt_Perm3-12                        	18561678	        64.86 ns/op	      24 B/op	       1 allocs/op
Benchmark_rand_Perm30-12                     	 3099498	       393.3 ns/op	     240 B/op	       1 allocs/op
Benchmark_sfmt_Perm30-12                     	  642010	      1923 ns/op	     240 B/op	       1 allocs/op
Benchmark_mt_Perm30-12                       	 2689364	       446.5 ns/op	     240 B/op	       1 allocs/op
Benchmark_rand_Perm30ViaShuffle-12           	 4970145	       236.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Perm30ViaShuffle-12           	  712490	      1712 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Perm30ViaShuffle-12             	 4036831	       295.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_ShuffleOverhead-12            	 3325021	       360.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_ShuffleOverhead-12            	  416008	      2963 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_ShuffleOverhead-12              	 2684896	       445.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Read3-12                      	65097250	        18.59 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Read3-12                      	26259727	        41.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Read3-12                        	59372941	        20.29 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Read64-12                     	22312156	        53.11 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Read64-12                     	 2225148	       543.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Read64-12                       	13283383	        89.91 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Read1000-12                   	 1978146	       600.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Read1000-12                   	  148468	      9049 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Read1000-12                     	 1000000	      1174 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63r1000-12                 	78956234	        14.90 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63r1000-12                 	18170481	        66.14 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63r1000-12                   	71224512	        16.50 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63s30-12                   	 1863955	       632.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63s30-12                   	  559028	      2181 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63s30-12                     	 1742422	       687.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int63Shuffle30-12             	 1985632	       601.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int63Shuffle30-12             	  536121	      2095 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int63huffle30-12                	 1789054	       661.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint32r1000-12                	79910392	        14.75 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint32r1000-12                	17898226	        66.03 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint32r1000-12                  	73332720	        16.58 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint32s30-12                  	 2245064	       539.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint32s30-12                  	  586099	      2095 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint32s30-12                    	 2033931	       584.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint32Shuffle30-12            	 1990003	       605.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint32Shuffle30-12            	  562398	      2076 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint32Shuffle30-12              	 1849230	       640.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint64r1000-12                	77057907	        15.70 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint64r1000-12                	17709919	        69.93 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint64r1000-12                  	67924168	        17.60 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint64s30-12                  	 1948162	       620.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint64s30-12                  	  553335	      2157 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint64s30-12                    	 1802133	       657.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Uint64Shuffle30-12            	 1957382	       604.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Uint64Shuffle30-12            	  578160	      2062 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Uint64Shuffle30-12              	 1837713	       655.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31r1000-12                 	81066088	        14.67 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31r1000-12                 	18439201	        65.48 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int31r1000-12                   	73665444	        16.48 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31s30-12                   	 2395334	       506.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31s30-12                   	  581770	      2041 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int3130-12                      	 2178031	       545.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Int31Shuffle30-12             	 1982134	       600.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Int31Shuffle30-12             	  572739	      2068 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Int31Shuffle30-12               	 1835955	       664.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Intr1000-12                   	79433973	        15.11 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Intr1000-12                   	18393291	        65.54 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Intr1000-12                     	65093133	        16.75 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Ints30-12                     	 1907890	       647.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Ints30-12                     	  546342	      2200 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Ints30-12                       	 1739860	       695.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_IntShuffle30-12               	 1992492	       601.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_IntShuffle30-12               	  576374	      2071 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_IntShuffle30-12                 	 1798989	       656.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64r1000-12               	78130340	        14.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64r1000-12               	16853383	        66.92 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64r1000-12                 	70114152	        16.63 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64s30-12                 	 2539143	       471.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64s30-12                 	  571390	      2028 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64s0-12                    	 2294743	       522.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float64Shuffle30-12           	 1952976	       604.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float64Shuffle30-12           	  571812	      2143 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float64Shuffle30-12             	 1795771	       677.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32r1000-12               	70598168	        17.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float32r1000-12               	16206597	        72.43 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32r1000-12                 	59814129	        18.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32s30-12                 	 2259800	       523.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float3230-12                  	  513606	      2107 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32s0-12                    	 1958778	       584.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_Float32Shuffle30-12           	 1889010	       603.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_Float32Shuffle30-12           	  586758	      2074 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_Float32Shuffle30-12             	 1803433	       664.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_rand_FlipCoin-12                   	38587446	        31.41 ns/op	       0 B/op	       0 allocs/op
Benchmark_sfmt_FlipCoin-12                   	 8452344	       144.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_mt_FlipCoin-12                     	33111992	        36.04 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 68.4% of statements
ok  	github.com/bcowtech/gosfmt	161.350s

```
