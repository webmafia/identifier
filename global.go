package identifier

import (
	"github.com/webmafia/identifier/coder"
	"github.com/webmafia/identifier/coder/alpha"
	"github.com/webmafia/identifier/node"
)

var (
	gen *node.Node
	cod *coder.Coder
)

func init() {
	var err error

	if gen, err = node.NewNode(0); err != nil {
		panic(err)
	}

	// Use an alphabet without vowels and similar characters
	if cod, err = coder.New(alpha.NewAlphabet("bcdfghjkmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ02345679").Shuffle(1337)); err != nil {
		return
	}
}
