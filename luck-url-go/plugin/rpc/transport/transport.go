package transport

import (
	"encoding/binary"
	"io"
	"net"
)

const headerLen = 4

// Transport will use TLV protocol
type Transport struct {
	Conn        net.Conn // Conn is a generic stream-oriented network Connection.
	IsConnected bool
}

// NewTransport creates a Transport
func NewTransport(conn net.Conn) *Transport {
	return &Transport{
		Conn:        conn,
		IsConnected: true,
	}
}

// Send TLV encoded data over the network
func (t *Transport) Send(data []byte) error {
	// we will need 4 more byte then the len of data
	// as TLV header is 4bytes and in this header
	// we will encode how much byte of data
	// we are sending for this request.
	buf := make([]byte, headerLen+len(data))
	binary.BigEndian.PutUint32(buf[:headerLen], uint32(len(data)))
	copy(buf[headerLen:], data)
	_, err := t.Conn.Write(buf)
	if err != nil {
		t.IsConnected = false
		return err
	}
	return nil
}

// Read TLV Data sent over the wire
func (t *Transport) Read() ([]byte, error) {
	header := make([]byte, headerLen)
	_, err := io.ReadFull(t.Conn, header)
	if err != nil {
		return nil, err
	}
	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(t.Conn, data)
	if err != nil {
		t.IsConnected = false
		return nil, err
	}
	return data, nil
}

// Close
func (t *Transport) Close() error {
	t.IsConnected = false
	err := t.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
