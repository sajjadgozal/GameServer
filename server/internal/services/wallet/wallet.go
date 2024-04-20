package wallet

import (
	"fmt"

	"github.com/blockcypher/gobcy/v2"
)

// WalletService represents the wallet service
type WalletService struct {
}

// NewWalletService creates a new wallet service
func NewWalletService() *WalletService {
	return &WalletService{}
}

// WalletResponse represents the JSON response structure from BlockCypher API
//{Address:3BF1M1PnTge94QewuWh3B8mRVw8U4SVnb4 Private: Public: Wif: PubKeys:[02c716d071a76cbf0d29c29cacfec76e0ef8116b37389fb7a3e76d6d32cf59f4d3 033ef4d5165637d99b673bcdbb7ead359cee6afd7aaf78d3da9d2392ee4102c8ea 022b8934cc41e76cb4286b9f3ed57e2d27798395b04dd23711981a77dc216df8ca] ScriptType:multisig-2-of-3 OriginalAddress: OAPAddress:}

type WalletResponse struct {
	Address         string   `json:"address"`
	Private         string   `json:"private"`
	Public          string   `json:"public"`
	Wif             string   `json:"wif"`
	PubKeys         []string `json:"pubkeys"`
	Script          string   `json:"script"`
	ScriptType      string   `json:"script_type"`
	OriginalAddress string   `json:"original_address"`
	OAPAddress      string   `json:"oap_address"`
}

func CreateWallet() {
	// apiToken := "45d649d773ce4c1f9f233fd96f062166"

}

func GetBallance() {
	token := "45d649d773ce4c1f9f233fd96f062166"

	address := "DJdJztKHN5TLapbhPHStp7FKadznMfAhcs"

	doge := gobcy.API{token, "dogetest", "test"}
	addr, err := doge.GetAddr(address, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", addr)
}
