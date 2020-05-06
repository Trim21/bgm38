package vote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseOptions(t *testing.T) {
	voteOptions := "vote: true\nmulti: false\nfilter:\n  user_group: [1, 2, 3, 10]\noptions:\n  - hello\n  - world\n"
	o, err := parseOption(voteOptions)
	assert.Nil(t, err)
	assert.Equal(t, []string{"hello", "world"}, o.Options)
}
