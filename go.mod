module github.com/centrifugal/centrifuge-go

go 1.26.0

require (
	github.com/centrifugal/protocol v0.9.1
	github.com/gorilla/websocket v1.5.3
	github.com/jpillora/backoff v1.0.0
	github.com/stretchr/testify v1.12.1
)

require (
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.2 // indirect
	github.com/segmentio/asm v1.2.1 // indirect
	github.com/segmentio/encoding v0.5.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/sys v0.49.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

replace github.com/centrifugal/protocol => github.com/rocketsciencegg/centrifuge-protocol v0.9.2
