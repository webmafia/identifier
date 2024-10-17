package identity

import (
	"github.com/webmafia/identity/coder"
	"github.com/webmafia/identity/node"
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

	if cod, err = coder.NewCoder(coder.ShuffleAlpha(1337, "bcdfghjkmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ02345679")); err != nil {
		return
	}
}
