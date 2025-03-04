//go:build !windows
// +build !windows

package keystore

import (
	"github.com/comdevx/go-tron-sdk/pkg/address"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func RecoverPubkey(hash []byte, signature []byte) (address.Address, error) {

	if signature[64] >= 27 {
		signature[64] -= 27
	}

	sigPublicKey, err := secp256k1.RecoverPubkey(hash, signature)
	if err != nil {
		return nil, err
	}
	pubKey, err := UnmarshalPublic(sigPublicKey)
	if err != nil {
		return nil, err
	}

	addr := address.PubkeyToAddress(*pubKey)
	return addr, nil
}
