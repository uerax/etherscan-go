package etherscan

import (
	"fmt"
	"strings"
)

// Get Contract Creator and Creation Tx Hash
// Returns a contract's deployer address and transaction hash it was created, up to 5 at a time.
//
// Parameters:
//	apikey: YourApiKeyToken
//	address: the contract address , up to 5 at a time
func GetContractCreator(apikey string, tokens... string) ([]ContractCreator, error) {
	if apikey == "" || len(tokens) == 0 {
		return nil, fmt.Errorf("参数不可为空")
	}
	token := strings.Join(tokens, ",")
	var resp *GetContractCreatorResp
	err := get(ParseUrl(Url, "module", "contract", "action", "getcontractcreation", "contractaddresses", token, "apikey", apikey), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, err
}