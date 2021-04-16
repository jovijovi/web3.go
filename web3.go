package web3

import (
	"github.com/jovijovi/web3.go/eth"
	"github.com/jovijovi/web3.go/rpc"
	"github.com/jovijovi/web3.go/utils"
)

type Web3 struct {
	Eth   *eth.Eth
	Utils *utils.Utils
	c     *rpc.Client
}

func NewWeb3(provider string) (*Web3, error) {
	c, err := rpc.NewClient(provider)
	if err != nil {
		return nil, err
	}
	e := eth.NewEth(c)
	e.SetChainId(1)
	u := utils.NewUtils()
	w := &Web3{
		Eth:   e,
		Utils: u,
		c:     c,
	}
	return w, nil
}

func (w *Web3) Version() (string, error) {
	var out string
	err := w.c.Call("web3_clientVersion", &out)
	return out, err
}
