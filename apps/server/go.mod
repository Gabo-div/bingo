module github.com/Gabo-div/bingo/apps/backend-main

go 1.24.5

require (
	connectrpc.com/connect v1.18.1
	github.com/Gabo-div/bingo/packages/protobuf v0.0.0
	github.com/go-chi/chi/v5 v5.2.2
	github.com/lestrrat-go/jwx/v3 v3.0.9
)

require (
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/lestrrat-go/blackmagic v1.0.4 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/httprc/v3 v3.0.0 // indirect
	github.com/lestrrat-go/option v1.0.1 // indirect
	github.com/lestrrat-go/option/v2 v2.0.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/Gabo-div/bingo/packages/protobuf v0.0.0 => ../../packages/protobuf

replace github.com/Gabo-div/bingo/packages/database v0.0.0 => ../../packages/database
