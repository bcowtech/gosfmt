package sfmt_test

import (
	"testing"

	"github.com/bcowtech/gosfmt/sfmt"
	"github.com/cpmech/gosl/chk"
	"github.com/stretchr/testify/assert"
)

var (
	seed     = int64(5489)
	seed5489 = uint64(226931099713899959)
)

func Test_SFMT(t *testing.T) {
	//verbose()
	chk.PrintTitle("sfmt. Uint64")
	assert := assert.New(t)

	sfmt := sfmt.New()
	sfmt.Seed(seed)
	assert.Equal(seed5489, sfmt.Uint64())

	sfmt.Seed(seed)
	assert.Equal(int64(seed5489&0x7fffffffffffffff), sfmt.Int63())
}
