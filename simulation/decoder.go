package simulation

import (
	tmkv "github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding htlc type
func DecodeStore(kvA, kvB tmkv.Pair) string {
	return ""
}
