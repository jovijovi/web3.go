package utils

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) ToHex(n *big.Int) string {
	return fmt.Sprintf("0x%x", n) // or %X or upper case
}

func (u *Utils) HexToUint64(str string) (uint64, error) {
	return ParseUint64orHex(str)
}

func (u *Utils) SameAddress(a, b common.Address) bool {
	return bytes.Compare(a[:], b[:]) == 0
}

func (u *Utils) DifferentAddress(a, b common.Address) bool {
	return bytes.Compare(a[:], b[:]) != 0
}
