module github.com/centrifugal/centrifuge-go/examples

go 1.26.0

replace github.com/centrifugal/centrifuge-go => ../

require (
	github.com/centrifugal/centrifuge-go v0.3.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
)

require (
	github.com/centrifugal/protocol v0.9.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/segmentio/asm v1.1.4 // indirect
	github.com/segmentio/encoding v0.3.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/centrifugal/protocol => github.com/rocketsciencegg/centrifuge-protocol v0.9.2
