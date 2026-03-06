module github.com/centrifugal/centrifuge-go

go 1.26.0

require (
	github.com/centrifugal/protocol v0.9.1
	github.com/gorilla/websocket v1.4.2
	github.com/jpillora/backoff v1.0.0
)

require (
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/segmentio/asm v1.1.4 // indirect
	github.com/segmentio/encoding v0.3.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/centrifugal/protocol => github.com/rocketsciencegg/centrifuge-protocol v0.9.2
