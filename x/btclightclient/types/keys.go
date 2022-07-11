package types

import (
	bbl "github.com/babylonchain/babylon/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "btclightclient"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_btclightclient"
)

var (
	HeadersPrefix       = []byte{0x0}                // reserve this namespace for headers
	HeadersObjectPrefix = append(HeadersPrefix, 0x0) // where we save the concrete header bytes
	HashToHeightPrefix  = append(HeadersPrefix, 0x1) // where we map hash to height
	HashToWorkPrefix    = append(HeadersPrefix, 0x2) // where we map hash to height
	TipPrefix           = append(HeadersPrefix, 0x3) // where we store the tip
)

func HeadersObjectKey(height uint64, hash *bbl.BTCHeaderHashBytes) []byte {
	he := sdk.Uint64ToBigEndian(height)
	hashBytes := hash.MustMarshal()

	heightPrefix := append(HeadersObjectPrefix, he...)
	return append(heightPrefix, hashBytes...)
}

func HeadersObjectHeightKey(hash *bbl.BTCHeaderHashBytes) []byte {
	return append(HashToHeightPrefix, hash.MustMarshal()...)
}

func HeadersObjectWorkKey(hash *bbl.BTCHeaderHashBytes) []byte {
	return append(HashToWorkPrefix, hash.MustMarshal()...)
}

func TipKey() []byte {
	return TipPrefix
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
