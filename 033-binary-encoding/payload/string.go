package payload

import (
	types "binaryencoding/types"
	"encoding/binary"
	"errors"
	"io"
)

type String string

func NewString() *String {
	return new(String)
}

func (m String) Bytes() []byte {
	return []byte(m)
}

func (m String) String() string {
	return string(m)
}

func (m String) WriteTo(w io.Writer) (n int64, err error) {
	// Write Type
	err = binary.Write(w, binary.BigEndian, types.StringType)
	if err != nil {
		return
	}
	n = 1
	// Write Length
	err = binary.Write(w, binary.BigEndian, uint32(len(m)))
	if err != nil {
		return
	}
	n += 4

	// Write Value
	o, err := w.Write([]byte(m))
	if err != nil {
		return
	}
	n += int64(o)
	return
}

func (m *String) ReadFrom(r io.Reader) (int64, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ) // 1-byte type
	if err != nil {
		return 0, err
	}
	var n int64 = 1
	if typ != types.StringType {
		return n, errors.New("invalid String")
	}

	var size uint32
	err = binary.Read(r, binary.BigEndian, &size) // 4-byte size
	if err != nil {
		return n, err
	}
	n += 4
	if size > types.MaxPayloadSize {
		return n, types.ErrMaxPayloadSize
	}

	buf := make([]byte, size)
	o, err := r.Read(buf) // payload
	if err != nil {
		return n, err
	}
	*m = String(buf)

	return n + int64(o), nil
}
