package types

// ChordQuerySuccessorMessage describes a message sent to request the successor of a key.
//
// - implements types.Message
type ChordQuerySuccessorMessage struct {
	// RequestID must be a unique identifier. Use xid.New().String() to generate
	// it.
	RequestID string

	// Source is the address of the peer that initiate the query
	Source string

	// Key is the key to query
	Key uint
}

// ChordReplySuccessorMessage describes a reply message to the ChordQuerySuccessorMessage, it includes
// the indicator ReplyPacketID (which query it replies to), and the answer to the query, Successor.
//
// - implements types.Message
type ChordReplySuccessorMessage struct {
	// ReplyPacketID is the PacketID this reply is for
	ReplyPacketID string

	// Successor is the answer to the query, i.e., which successor the query key belongs to
	Successor string
}

// ChordQueryPredecessorMessage describes a message sent to request the predecessor of the node
//
// - implements types.Message
type ChordQueryPredecessorMessage struct{}

// ChordReplyPredecessorMessage describes a reply message to the ChordQueryPredecessorMessage
//
// - implements types.Message
type ChordReplyPredecessorMessage struct {
	// Predecessor is the answer to the query, i.e., which predecessor the node is storing
	Predecessor string
}

// ChordNotifyMessage describes a notifyMessage in the case that we believe that we should
// be the predecessor of another node
//
// - implements types.Message
type ChordNotifyMessage struct{}

// ChordRingLenMessage describes a query message to find out the total number of nodes inside
// a Chord ring. It is initiated by Source and Length is the cumulative length.
//
// - implements types.Message
type ChordRingLenMessage struct {
	// RequestID must be a unique identifier. Use xid.New().String() to generate
	// it.
	RequestID string

	// Source is the source who initiate the request
	Source string

	// Length is the current length of the ring
	Length uint
}

// ChordClearPredecessorMessage instructs our successor to remove us from its predecessor field,
// this message is used when some nodes leave the system
//
// - implements types.Message
type ChordClearPredecessorMessage struct{}

// ChordSkipSuccessorMessage instructs our predecessor to remove us from its successor field, this
// message is used when some nodes leave the system
//
// - implements types.Message
type ChordSkipSuccessorMessage struct {
	// The new successor that our predecessor should use
	Successor string
}

// ChordPingMessage pings a chord peer, and check for its liveliness
//
// - implements types.Message
type ChordPingMessage struct {
	// RequestID must be a unique identifier. Use xid.New().String() to generate
	// it.
	RequestID string
}

// ChordPingReplyMessage replies a ping message
type ChordPingReplyMessage struct {
	// ReplyPacketID is the PacketID this reply is for
	ReplyPacketID string
}
