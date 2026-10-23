package centrifuge

import (
	"testing"

	"github.com/centrifugal/protocol"
	"github.com/stretchr/testify/require"
)

func TestNewCommandEncoder(t *testing.T) {
	encoder := newCommandEncoder(protocol.TypeJSON)

	t.Run("encodes-unity-v1-connect", func(t *testing.T) {
		cmd := &protocol.Command{
			Id:     1,
			Method: protocol.Command_CONNECT,
			Params: protocol.Raw(`{"token":"unity-token"}`),
		}
		data, err := encoder.Encode(cmd)
		require.NoError(t, err)
		require.JSONEq(t, `{"id":1,"params":{"token":"unity-token"}}`, string(data))
	})

	t.Run("encodes-unity-v1-subscribe", func(t *testing.T) {
		cmd := &protocol.Command{
			Id:     2,
			Method: protocol.Command_SUBSCRIBE,
			Params: protocol.Raw(`{"channel":"server:123","recover":true,"offset":42,"epoch":"test"}`),
		}
		data, err := encoder.Encode(cmd)
		require.NoError(t, err)
		require.JSONEq(t, `{"id":2,"method":1,"params":{"channel":"server:123","recover":true,"offset":42,"epoch":"test"}}`, string(data))
	})
}

func TestNewReplyDecoder(t *testing.T) {
	t.Run("decodes-unity-v1-connect", func(t *testing.T) {
		decoder := newReplyDecoder(protocol.TypeJSON, []byte(`{"id":1,"result":{"client":"unity-client"}}`))
		reply, err := decoder.Decode()
		require.NoError(t, err)
		require.Equal(t, &protocol.Reply{
			Id:     1,
			Result: protocol.Raw(`{"client":"unity-client"}`),
		}, reply)
	})

	t.Run("decodes-unity-v1-subscribe", func(t *testing.T) {
		decoder := newReplyDecoder(protocol.TypeJSON, []byte(`{"id":2,"result":{"recoverable":true,"offset":42,"epoch":"test"}}`))
		reply, err := decoder.Decode()
		require.NoError(t, err)
		require.Equal(t, &protocol.Reply{
			Id:     2,
			Result: protocol.Raw(`{"recoverable":true,"offset":42,"epoch":"test"}`),
		}, reply)
	})

	t.Run("decodes-unity-v1-publication", func(t *testing.T) {
		decoder := newReplyDecoder(protocol.TypeJSON, []byte(`{"result":{"type":0,"channel":"server:123","data":{"data":{"message":"hello"}}}}`))
		reply, err := decoder.Decode()
		require.NoError(t, err)
		require.Equal(t, &protocol.Reply{
			Result: protocol.Raw(`{"type":0,"channel":"server:123","data":{"data":{"message":"hello"}}}`),
		}, reply)
	})

	t.Run("rejects-malformed-json", func(t *testing.T) {
		decoder := newReplyDecoder(protocol.TypeJSON, []byte(`{"id":1,"result":`))
		_, err := decoder.Decode()
		require.Error(t, err)
	})
}

func TestNewResultDecoder(t *testing.T) {
	decoder := newResultDecoder(protocol.TypeJSON)

	t.Run("decodes-unity-v1-connect-result", func(t *testing.T) {
		result := &protocol.ConnectResult{}
		err := decoder.Decode([]byte(`{"client":"unity-client"}`), result)
		require.NoError(t, err)
		require.Equal(t, &protocol.ConnectResult{Client: "unity-client"}, result)
	})

	t.Run("decodes-unity-v1-subscribe-result", func(t *testing.T) {
		result := &protocol.SubscribeResult{}
		err := decoder.Decode([]byte(`{"recoverable":true,"offset":42,"epoch":"test"}`), result)
		require.NoError(t, err)
		require.Equal(t, &protocol.SubscribeResult{
			Recoverable: true,
			Offset:      42,
			Epoch:       "test",
		}, result)
	})

	t.Run("rejects-malformed-json", func(t *testing.T) {
		result := &protocol.ConnectResult{}
		err := decoder.Decode([]byte(`{"client":`), result)
		require.Error(t, err)
	})
}
