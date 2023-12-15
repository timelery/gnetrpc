package protocol

import "errors"

var (
	magicNumberInvalid = errors.New("magic number invalid")
	versionInvalid     = errors.New("version number invalid")
)

const (
	MagicNumber byte = 0xFF
	Version     byte = 0x01
)

type Header struct {
	MagicNumber   byte
	Version       byte
	HeartBeat     byte
	SerializeType byte
}

func (h *Header) Check() error {
	if h.MagicNumber != MagicNumber {
		return magicNumberInvalid
	}
	if h.Version != Version {
		return versionInvalid
	}
	return nil
}
