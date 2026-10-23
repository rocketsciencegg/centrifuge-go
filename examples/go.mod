module github.com/centrifugal/centrifuge-go/examples

go 1.26.0

replace github.com/centrifugal/centrifuge-go => ../

require (
	github.com/centrifugal/centrifuge-go v0.8.5
	github.com/golang-jwt/jwt/v5 v5.3.1
)

require (
	github.com/centrifugal/protocol v0.9.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.2 // indirect
	github.com/segmentio/asm v1.2.1 // indirect
	github.com/segmentio/encoding v0.5.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

replace github.com/centrifugal/protocol => github.com/rocketsciencegg/centrifuge-protocol v0.9.2
