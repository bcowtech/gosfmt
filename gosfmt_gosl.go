package gosfmt

// Int63r generates pseudo random int64 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random int64
func Int63r(low, high int64) int64 {
	return globalRand.Int63r(low, high)
}

// Int63s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func Int63s(values []int64, low, high int64) {
	globalRand.Int63s(values, low, high)
}

// Int63huffle shuffles a slice of integers
func Int63Shuffle(values []int64) {
	globalRand.Int63Shuffle(values)
}

// Uint32 is int range generates pseudo random uint32 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random uint32
func Uint32r(low, high uint32) uint32 {
	return globalRand.Uint32r(low, high)
}

// Uint32s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func Uint32s(values []uint32, low, high uint32) {
	globalRand.Uint32s(values, low, high)
}

// Uint32huffle shuffles a slice of integers
func Uint32Shuffle(values []uint32) {
	globalRand.Uint32Shuffle(values)
}

// Uint64r generates pseudo random uint64 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random uint64
func Uint64r(low, high uint64) uint64 {
	return globalRand.Uint64r(low, high)
}

// Uint64s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func Uint64s(values []uint64, low, high uint64) {
	globalRand.Uint64s(values, low, high)
}

// Uint64huffle shuffles a slice of integers
func Uint64Shuffle(values []uint64) {
	globalRand.Uint64Shuffle(values)
}

// Int31r is int range generates pseudo random int32 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random int32
func Int31r(low, high int32) int32 {
	return globalRand.Int31r(low, high)
}

// Int31s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func Int31s(values []int32, low, high int32) {
	globalRand.Int31s(values, low, high)
}

// Int31huffle shuffles a slice of integers
func Int31Shuffle(values []int32) {
	globalRand.Int31Shuffle(values)
}

// Intr is int range generates pseudo random integer between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random integer
func Intr(low, high int) int {
	return globalRand.Intr(low, high)
}

// Ints generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func Ints(values []int, low, high int) {
	globalRand.Ints(values, low, high)
}

// IntShuffle shuffles a slice of integers
func IntShuffle(values []int) {
	globalRand.IntShuffle(values)
}

// Float64r generates a pseudo random real number between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   random float64
func Float64r(low, high float64) float64 {
	return globalRand.Float64r(low, high)
}

// Float64s generates pseudo random real numbers between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   values -- slice to be filled with len(values) numbers
func Float64s(values []float64, low, high float64) {
	globalRand.Float64s(values, low, high)
}

// Float64Shuffle shuffles a slice of float point numbers
func Float64Shuffle(values []float64) {
	globalRand.Float64Shuffle(values)
}

// Float32r generates a pseudo random real number between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   random float32
func Float32r(low, high float32) float32 {
	return globalRand.Float32r(low, high)
}

// Float32s generates pseudo random real numbers between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   values -- slice to be filled with len(values) numbers
func Float32s(values []float32, low, high float32) {
	globalRand.Float32s(values, low, high)
}

// Float32Shuffle shuffles a slice of float point numbers
func Float32Shuffle(values []float32) {
	globalRand.Float32Shuffle(values)
}

// FlipCoin generates a Bernoulli variable; throw a coin with probability p
func FlipCoin(p float64) bool {
	return globalRand.FlipCoin(p)
}
