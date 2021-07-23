module github.com/thetatoken/theta/query

go 1.14

require (
	github.com/onsi/gomega v1.14.0 // indirect
	github.com/thetatoken/theta v0.0.0
	github.com/thetatoken/theta/common v0.0.0
	github.com/thetatoken/theta/utils v0.0.0
	// github.com/ybbus/jsonrpc v2.1.2+incompatible
	github.com/ybbus/jsonrpc v1.1.1
)

replace (
	github.com/thetatoken/theta v0.0.0 => ../theta-protocol-ledger
	github.com/thetatoken/theta/common v0.0.0 => ../theta-protocol-ledger/common
	github.com/thetatoken/theta/rpc/lib/rpc-codec/jsonrpc2 v0.0.0 => ../theta-protocol-ledger/rpc/lib/rpc-codec/jsonrpc2/
	github.com/thetatoken/theta/utils v0.0.0 => ../thetaoffchaingo_utils
)
