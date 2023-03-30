package common

import "github.com/cosmos/cosmos-sdk/types/bech32"

func ValidateBech32Address(address string) error {
	_, _, err := bech32.DecodeAndConvert(address)
	return err
}
