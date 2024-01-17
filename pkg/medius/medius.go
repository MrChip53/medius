package medius

type IRTMessage interface {
	GetRequestID() uint8
	Process() ([]byte, error)
}

type IMediusMessage interface {
	GetType() uint
	GetID() uint
	Process() ([]byte, error)
}

type RTMessage struct {
	ID        uint8
	Length    uint16
	Encrypted bool
	Bytes     []byte
}

type MediusMessage struct {
	Type  uint8
	ID    uint8
	Bytes []byte
}

var rtMessages = map[uint8]IRTMessage{}
var mediusMessages = map[uint8]map[uint8]IMediusMessage{}
