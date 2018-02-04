package comm

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestSingleMessageSender_SendMessage(t *testing.T) {
	for _, p := range testPackages {
		expectedBytes, err := toWire(p)
		if assert.Nil(t, err, "toWire returned an error for package %v: %v", p, err) {
			conn := newBufferConn()
			sms := singleMessageSender{connection: conn}
			sms.SendMessage(p)
			actualBytes := conn.Bytes()
			assert.True(t, bytes.Equal(expectedBytes, actualBytes), "singleMessageSender sendMessage did not write toWire to the connection for package %v", p)
		}
	}
}

func TestMultiMessageSender_SendMessage(t *testing.T) {
	for i, p := range testPackages {
		expectedBytes, err := toWire(p)
		if assert.Nil(t, err, "toWire returned an error for package %v: %v", p, err) {
			channels := testPackageChannels[i]

			connNoChan := newBufferConn()
			connNotInDict := newBufferConn()
			connAudio := newBufferConn()
			connMeta := newBufferConn()
			connBoth := newBufferConn()

			mms := &multiMessageSender{
				connections: []net.Conn{connNoChan, connNotInDict, connAudio, connMeta, connBoth},
				channels: map[net.Conn][]Channel{
					connNoChan: {},
					connAudio:  {Channel_AUDIO},
					connMeta:   {Channel_META},
					connBoth:   {Channel_AUDIO, Channel_META},
				}}
			mms.SendMessage(p)

			hasAudio, hasMeta := hasChannels(channels)

			assertSendData(t, connNoChan, expectedBytes, len(channels) == 0, "without channels", p)
			assertSendData(t, connNotInDict, expectedBytes, len(channels) == 0, "not in the channels dictionary", p)

			assertSendData(t, connAudio, expectedBytes, hasAudio, "with AUDIO channel", p)
			assertSendData(t, connMeta, expectedBytes, hasMeta, "with META channel", p)

			assertSendData(t, connBoth, expectedBytes, true, "with both channels", p)
		}
	}
}

func assertSendData(t *testing.T, conn bufferConn, data []byte, shouldSend bool, name string, p proto.Message) {
	if shouldSend {
		assert.True(t, bytes.Equal(data, conn.Bytes()), "multiMessageSender sendMessage did not write toWire to the connection %s for package %v", name, p)
	} else {
		assert.Zero(t, len(conn.Bytes()), "multiMessageSender sendMessage did write to the connection %s for package %v", name, p)
	}
}

func hasChannels(channels []Channel) (hasAudio bool, hasMeta bool) {
	if len(channels) == 0 {
		return true, true
	}
	hasAudio = false
	hasMeta = false

	for _, c := range channels {
		if c == Channel_AUDIO {
			hasAudio = true
		}
		if c == Channel_META {
			hasMeta = true
		}
	}
	return
}

func TestMultiMessageSender_AddConn(t *testing.T) {
	mms := &multiMessageSender{connections: make([]net.Conn, 0), channels: make(map[net.Conn][]Channel, 0)}
	conns := []net.Conn{newBufferConn(), newBufferConn(), newBufferConn()}

	assert.Zero(t, len(mms.connections), "multiMessageSender contains elements before adding any connections")

	for i, c := range conns {
		mms.AddConn(c)
		assert.Equal(t, i+1, len(mms.connections), "multiMessageSender has the wrong number of connections after adding %d connections", i+i)
		assert.True(t, inConnectionsSlice(c, mms.connections), "multiMessageSender connections slice does not contain the added connection after adding %d connections", i+1)
	}
}

func TestMultiMessageSender_DelConn(t *testing.T) {
	conns := []net.Conn{newBufferConn(), newBufferConn(), newBufferConn()}
	mms := &multiMessageSender{connections: make([]net.Conn, len(conns)), channels: make(map[net.Conn][]Channel, 0)}
	copy(mms.connections, conns)

	for i, c := range conns {
		mms.DelConn(c)
		assert.Equal(t, len(conns)-i-1, len(mms.connections), "multiMessageSender has the wrong number of connections after deleting %d connections", i+i)
		assert.False(t, inConnectionsSlice(c, mms.connections), "multiMessageSender connections slice contains the connection after deleting %d connections", i+1)
	}

	assert.Zero(t, len(mms.connections), "multiMessageSender contains elements after deleting all connections")
}

func inConnectionsSlice(c net.Conn, cs []net.Conn) bool {
	for _, e := range cs {
		if c == e {
			return true
		}
	}
	return false
}

func TestMultiMessageSender_Subscribe(t *testing.T) {
	conn := newBufferConn()
	mms := &multiMessageSender{connections: []net.Conn{conn}, channels: make(map[net.Conn][]Channel, 0)}

	assert.False(t, mms.isSubscribed(conn, []Channel{Channel_AUDIO}), "multiMessageSender claims the connection is subscribed to the AUDIO channel")
	assert.False(t, mms.isSubscribed(conn, []Channel{Channel_META}), "multiMessageSender claims the connection is subscribed to the META channel")

	mms.Subscribe(conn, Channel_AUDIO)
	assert.True(t, mms.isSubscribed(conn, []Channel{Channel_AUDIO}), "multiMessageSender claims the connection is not subscribed to the AUDIO channel")
	assert.False(t, mms.isSubscribed(conn, []Channel{Channel_META}), "multiMessageSender claims the connection is subscribed to the META channel")

	mms.Subscribe(conn, Channel_META)
	assert.True(t, mms.isSubscribed(conn, []Channel{Channel_AUDIO}), "multiMessageSender claims the connection is not subscribed to the AUDIO channel")
	assert.True(t, mms.isSubscribed(conn, []Channel{Channel_META}), "multiMessageSender claims the connection is not subscribed to the META channel")
}
