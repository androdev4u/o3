package o3

import "time"

type pktType uint32

const (
	// SENDINGMSG is the packet type of a message packet from client to server
	SENDINGMSG pktType = 0x1
	// DELIVERINGMSG is the packet type of a message packet from server to client
	DELIVERINGMSG pktType = 0x2
	// SERVERACK is the packet type of a server ack for a messsage sent by the client
	SERVERACK pktType = 0x81
	// CLIENTACK is the packet type of a client ack for a messsage delivered by the server
	CLIENTACK pktType = 0x82
	// CONNESTABLISHED is the packet type of a pkt send by ther server when all MSGs have been delivered
	CONNESTABLISHED pktType = 0xd0
	// MSGHEADERLENGHT is the length of a ThreemaMessageHeader
	MSGHEADERLENGHT uint8 = 64
)

// ThreemaMessageHeader contains fields that every type of message needs
type messagePacket struct {
	PktType    pktType
	Sender     IdString
	Recipient  IdString
	ID         uint64
	Time       time.Time
	Flags      msgFlags
	PubNick    PubNick
	Nonce      nonce
	Ciphertext []byte
	Plaintext  []byte
}

type ackPacket struct {
	PktType  pktType
	SenderID IdString
	MsgID    uint64
}

type connEstPacket struct {
	PktType pktType
}

type clientHelloPacket struct {
	ClientSPK   [32]byte
	NoncePrefix [16]byte
}

type serverHelloPacket struct {
	NoncePrefix [16]byte
	Ciphertext  [64]byte
}

type authPacket struct {
	Ciphertext [144]byte
}

//plain content of an auth packet
type authPacketPayload struct {
	Username          IdString
	SysData           [32]byte
	ServerNoncePrefix [16]byte
	RandomNonce       nonce
	Ciphertext        [48]byte
}
