package medius

import (
	"encoding/binary"
	"errors"
)

func ReadNextRTMessage(buf []byte) *RTMessage {
	id := buf[0]
	length := binary.LittleEndian.Uint16(buf[1:3])
	b := buf[3:length]
	return &RTMessage{
		ID:        id,
		Length:    length,
		Encrypted: false,
		Bytes:     b,
	}
}

func ProcessRTMessage(rtMessage *RTMessage) ([]byte, error) {
	rt, ok := rtMessages[rtMessage.ID]
	if !ok {
		return nil, errors.New("unknown rt message type")
	}
	return rt.Process()
}

func ProcessMediusMessage(mediusMessage *MediusMessage) ([]byte, error) {
	mediusType, ok := mediusMessages[mediusMessage.Type]
	if !ok {
		return nil, errors.New("unknown medius message type")
	}
	medius, ok := mediusType[mediusMessage.ID]
	if !ok {
		return nil, errors.New("unknown medius message id")
	}
	return medius.Process()
}
