package utils_test

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/influxdata/influxdb/pkg/testing/assert"

	"github.com/jovijovi/web3.go/utils"
)

func TestToWei(t *testing.T) {
	u := utils.NewUtils()
	wei := u.ToWei(200, utils.UnitGWei)
	t.Log("ToWei=", wei)
	t.Log("WeiHex=", u.ToHex(wei))

	wei2 := u.ToWei(200, "gwei")
	t.Log("ToWei=", wei2)
	t.Log("WeiHex=", u.ToHex(wei2))

	assert.Equal(t, true, u.Equal(wei, wei2))

	wei3 := u.ToWei(200, "not_exist_unit")
	t.Log("ToWei=", wei3)
	assert.Equal(t, big.NewInt(-1), wei3)
}

func TestUtils_Greater(t *testing.T) {
	u := utils.NewUtils()
	result1 := u.Greater(big.NewInt(1), big.NewInt(2))
	assert.Equal(t, false, result1)

	result2 := u.Greater(big.NewInt(2), big.NewInt(1))
	assert.Equal(t, true, result2)
}

func TestUtils_Less(t *testing.T) {
	u := utils.NewUtils()
	result1 := u.Less(big.NewInt(1), big.NewInt(2))
	assert.Equal(t, true, result1)

	result2 := u.Less(big.NewInt(2), big.NewInt(1))
	assert.Equal(t, false, result2)
}

func TestSignMethod(t *testing.T) {

	funcName := "transfer(address,uint256)"
	id := crypto.Keccak256([]byte(funcName))[:4]
	fmt.Printf("id %x\n", id)
}

func TestSameAddr(t *testing.T) {
	addr1 := common.HexToAddress("0x0000000000fC95fD88A4c46d9d7984A56289c52A")
	addr2 := common.HexToAddress("0x0000000000fc95fd88a4c46d9d7984a56289c52a")

	fmt.Printf("addr1 %v, addr2 %v\n", addr1, addr2)

	fmt.Printf("addr1 == addr2 %t\n", addr1 == addr2)

	fmt.Printf("bytes cmp addr1 == addr2 %t\n", bytes.Compare(addr1[:], addr2[:]) == 0)
}

func BenchmarkTestCompare1(b *testing.B) {
	addr1 := common.HexToAddress("0x0000000000fC95fD88A4c46d9d7984A56289c52A")
	addr2 := common.HexToAddress("0x0000000000fc95fd88a4c46d9d7984a56289c52a")

	for i := 0; i < b.N; i++ {
		if addr1 == addr2 {
			continue
		}
	}
}

func BenchmarkTestCompare2(b *testing.B) {
	addr1 := common.HexToAddress("0x0000000000fC95fD88A4c46d9d7984A56289c52A")
	addr2 := common.HexToAddress("0x0000000000fc95fd88a4c46d9d7984a56289c52a")

	for i := 0; i < b.N; i++ {
		if bytes.Compare(addr1[:], addr2[:]) == 0 {
			continue
		}
	}
}

func BenchmarkTestCompare3(b *testing.B) {
	addr1 := common.HexToAddress("0x0000000000fC95fD88A4c46d9d7984A56289c52A")
	addr2 := common.HexToAddress("0x0000000000fc95fd88a4c46d9d7984a56289c52a")

	for i := 0; i < b.N; i++ {
		if addr1.Hex() == addr2.Hex() {
			continue
		}
	}
}

func BenchmarkTestCompare4(b *testing.B) {
	addr1 := common.HexToAddress("0x0000000000fC95fD88A4c46d9d7984A56289c52A")
	addr2 := common.HexToAddress("0x0000000000fc95fd88a4c46d9d7984a56289c52a")

	for i := 0; i < b.N; i++ {
		if string(addr1[:]) == string(addr2[:]) {
			continue
		}
	}
}
