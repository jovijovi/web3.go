package utils

import (
	"math/big"
	"strings"
)

// Reference:
// https://ethdocs.org/en/latest/ether.html
const (
	UnitWeiValue        = int64(1)  // wei
	UnitKWeiValue       = int64(3)  // Kwei (babbage)
	UnitMWeiValue       = int64(6)  // Mwei (lovelace)
	UnitGWeiValue       = int64(9)  // Gwei (shannon)
	UnitMicroEtherValue = int64(12) // microether (szabo)
	UnitMilliEtherValue = int64(15) // milliether (finney)
	UnitEtherValue      = int64(18) // ether
)

const (
	UnitWei         = "wei"
	UnitKWei        = "kwei"
	UnitMWei        = "mwei"
	UnitGWei        = "gwei"
	UnitMicroEther  = "szabo"
	UnitMilliEther  = "finney"
	UnitEther       = "ether"
)

var (
	EthUnit = map[string]int64 {
		UnitWei:        UnitWeiValue,
		UnitKWei:       UnitKWeiValue,
		UnitMWei:       UnitMWeiValue,
		UnitGWei:       UnitGWeiValue,
		UnitMicroEther: UnitMicroEtherValue,
		UnitMilliEther: UnitMilliEtherValue,
		UnitEther:      UnitEtherValue,
	}
)

func convert(val uint64, decimals int64) *big.Int {
	v := big.NewInt(int64(val))
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil)
	return v.Mul(v, exp)
}

// ToWei converts a value to the wei
// Exp:
//   ToWei(200, "gwei")
//   ToWei(200, UnitGWei)
func (u *Utils) ToWei(value uint64, ethUnit string) *big.Int {
	unit, ok := EthUnit[strings.ToLower(ethUnit)]
	if !ok {
		return big.NewInt(-1)
	}

	return convert(value, unit)
}

func (u *Utils) ToDecimals(val uint64, decimals int64) *big.Int {
	return convert(val, decimals)
}

// Equal returns if x equal y
func (u *Utils) Equal(x, y *big.Int) bool {
	// -1 if x <  y
	// 0 if x == y
	// +1 if x >  y
	return x.Cmp(y) == 0
}

// Greater returns if x greater than y
func (u *Utils) Greater(x, y *big.Int) bool {
	return x.Cmp(y) == 1
}

// Less returns if x less than y
func (u *Utils) Less(x, y *big.Int) bool {
	return x.Cmp(y) == -1
}
