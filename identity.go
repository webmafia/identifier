package identity

import (
	"fmt"
	"sync"
	"time"
)

const bits = 22

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	mu      sync.Mutex
	epoch   time.Time
	epochTs int64
	time    int64
	node    int64
	seq     int64

	nodeMax   int64
	nodeMask  int64
	seqMask   int64
	timeShift uint8
	nodeShift uint8

	coder Coder
}

type Options struct {
	// Epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64

	// NodeBits holds the number of bits to use for Node. Default 7, max 22
	NodeBits uint8

	Coder Coder
}

func (opt *Options) setDefaults() {
	if opt.Epoch <= 0 {
		opt.Epoch = 1288834974657
	}

	if opt.NodeBits == 0 {
		opt.NodeBits = 7
	}

	if opt.Coder == nil {
		opt.Coder = stringCoder{}
	}
}

func NewNode(nodeId int64, options ...Options) (n *Node, err error) {
	var opt Options

	if len(options) > 0 {
		opt = options[0]
	}

	opt.setDefaults()

	if opt.NodeBits > bits {
		err = fmt.Errorf("node bits cannot be more than %d", bits)
	}

	n = &Node{}
	n.epochTs = opt.Epoch
	n.nodeShift = bits - opt.NodeBits
	n.timeShift = bits
	n.node = nodeId
	n.nodeMax = -1 ^ (-1 << opt.NodeBits)
	n.nodeMask = n.nodeMax << n.nodeShift
	n.seqMask = -1 ^ (-1 << n.nodeShift)

	if n.node < 0 || n.node > n.nodeMax {
		err = fmt.Errorf("node ID must be between 0 and %d", n.nodeMax)
		return
	}

	n.coder = opt.Coder
	n.setup()

	return
}

func (n *Node) setup() {
	curTime := time.Now()

	// add time.Duration to curTime to make sure we use the monotonic clock if available
	n.epoch = curTime.Add(time.Unix(n.epochTs/1000, (n.epochTs%1000)*1000000).Sub(curTime))
}

func (n *Node) Generate() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Since(n.epoch).Milliseconds()

	if now == n.time {
		n.seq = (n.seq + 1) & n.seqMask

		if n.seq == 0 {
			for now <= n.time {
				now = time.Since(n.epoch).Milliseconds()
			}
		}
	} else {
		n.seq = 0
	}

	n.time = now

	r := ID((now)<<n.timeShift |
		(n.node << n.nodeShift) |
		(n.seq),
	)

	return r
}

func (n *Node) Time(id ID) time.Time {
	return time.UnixMilli((int64(id) >> n.timeShift) + n.epochTs)
}

func (n *Node) ToString(id ID) string {
	s, _ := n.coder.Encode([]uint64{uint64(id)})

	return s
}
