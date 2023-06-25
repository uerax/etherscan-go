package etherscan

import (
	"fmt"
	"strings"
)

// Get Ether Balance for a Single Address
// Returns the Ether balance of a given address.
// The result is returned in wei.
// Convert Ethereum units using Unit Converter https://etherscan.io/unitconverter.
//
// Parameters:
//
//	tag: the string pre-defined block parameter, either 'earliest', 'pending' or 'latest', default latest
//	apikey: YourApiKeyToken
//	address: the string representing the address to check for balance
func GetBalance(address, tag, apikey string) (string, error) {
	if apikey == "" || tag == "" || address == "" {
		return "", fmt.Errorf("参数不可为空")
	}
	
	var resp *GetBalanceResp
	err := get(ParseUrl(Url, "module", "account", "action", "balance", "address", address, "tag", tag, "apikey", apikey), &resp)
	if err != nil {
		return "", err
	}
	return resp.Result, err
}

// Get Ether Balance for Multiple Addresses in a Single Call
// Returns the balance of the accounts from a list of addresses.
// The result is returned in wei.
// Convert Ethereum units using Unit Converter https://etherscan.io/unitconverter.
//
// Parameters:
//
//	tag: the string pre-defined block parameter, either 'earliest', 'pending' or 'latest', default latest
//	apikey: YourApiKeyToken
//	address: the string representing the address to check for balance
func GetBalanceMulti(tag, apikey string, address... string) ([]AddrBalance, error) {
	if apikey == "" || tag == "" || len(address) == 0 {
		return nil, fmt.Errorf("参数不可为空")
	}

	var resp *GetBalanceMultiResp
	err := get(ParseUrl(Url, "module", "account", "action", "balancemulti", "address", strings.Join(address, ","), "tag", tag, "apikey", apikey), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}


// Get a list of 'Normal' Transactions By Address
// Returns the list of transactions performed by an address, with optional pagination.
// This API endpoint returns a maximum of 10000 records only.
//
// Parameters:
//	address: the string representing the addresses to check for balance
//	apikey: YourApiKeyToken
//	startblock: the integer block number to start searching for transactions. default '0' optional
//	endblock: the integer block number to stop  searching for transactions. default '9999999999' optional
//	page: the integer page number, if pagination is enabled. default '1' optional
//	offset: the number of transactions displayed per page. default '30' optional
//	sort:the sorting preference, use 'asc' to sort by ascending and 'desc' to sort by descendin Tip: Specify a smaller startblock and endblock range for faster search results. default 'asc' optional
func GetNormalTransactions(apikey, address, startblock, endblock, page, offset, sort string) ([]NormalTransaction, error) {
	if apikey == "" || len(address) == 0 {
		return nil, fmt.Errorf("参数不可为空")
	}
	if startblock == "" {
		startblock = "0"
	}
	if endblock == "" {
		endblock = "9999999999"
	}
	if page == "" {
		page = "1"
	}
	if offset == "" {
		offset = "30"
	}
	if sort == "" {
		sort = "asc"
	}
	var resp *GetNormalTransactionsResp
	err := get(ParseUrl(Url, "module", "account", "action", "txlist", "address", address, "apikey", apikey, "startblock", startblock, "endblock", endblock, "page", page, "offset", offset, "sort", sort), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}


// Get a list of 'Internal' Transactions by Address
// Returns the list of internal transactions performed by an address, with optional pagination.
// This API endpoint returns a maximum of 10000 records only.
//
// Parameters:
//	address: the string representing the addresses to check for balance
//	apikey: YourApiKeyToken
//	startblock: the integer block number to start searching for transactions. default '0' optional
//	endblock: the integer block number to stop  searching for transactions. default '9999999999' optional
//	page: the integer page number, if pagination is enabled. default '1' optional
//	offset: the number of transactions displayed per page. default '30' optional
//	sort:the sorting preference, use 'asc' to sort by ascending and 'desc' to sort by descendin Tip: Specify a smaller startblock and endblock range for faster search results. default 'asc' optional
func GetInternalTransactions(apikey, address, startblock, endblock, page, offset, sort string) ([]InternalTransaction, error) {
	if apikey == "" || len(address) == 0 {
		return nil, fmt.Errorf("参数不可为空")
	}
	if startblock == "" {
		startblock = "0"
	}
	if endblock == "" {
		endblock = "9999999999"
	}
	if page == "" {
		page = "1"
	}
	if offset == "" {
		offset = "30"
	}
	if sort == "" {
		sort = "asc"
	}
	var resp *GetInternalTransactionsResp
	err := get(ParseUrl(Url, "module", "account", "action", "txlistinternal", "address", address, "apikey", apikey, "startblock", startblock, "endblock", endblock, "page", page, "offset", offset, "sort", sort), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

// Get 'Internal Transactions' by Transaction Hash
// Returns the list of internal transactions performed within a transaction.
// This API endpoint returns a maximum of 10000 records only.
//
// Parameters:
//	txhash: the string representing the transaction hash to check for internal transactions
//	apikey: YourApiKeyToken
func GetInternalTransactionsByHash(apikey, txhash string) ([]InternalTransaction, error) {
	if apikey == "" || len(txhash) == 0 {
		return nil, fmt.Errorf("参数不可为空")
	}

	var resp *GetInternalTransactionsResp
	err := get(ParseUrl(Url, "module", "account", "action", "txlistinternal", "txhash", txhash), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}


// Get "Internal Transactions" by Block Range
// Returns the list of internal transactions performed within a block range, with optional pagination.
// This API endpoint returns a maximum of 10000 records only.
//
// Parameters:
//	apikey: YourApiKeyToken
//	startblock: the integer block number to start searching for transactions. default '0' optional
//	endblock: the integer block number to stop  searching for transactions. default '9999999999' optional
//	page: the integer page number, if pagination is enabled. default '1' optional
//	offset: the number of transactions displayed per page. default '30' optional
//	sort:the sorting preference, use 'asc' to sort by ascending and 'desc' to sort by descendin Tip: Specify a smaller startblock and endblock range for faster search results. default 'asc' optional
func GetInternalTransactionsbyByBlockRange(apikey, startblock, endblock, page, offset, sort string) ([]InternalTransaction, error) {
	if apikey == "" || (startblock == "" && endblock == ""){
		return nil, fmt.Errorf("参数不可为空")
	}

	if startblock == "" {
		startblock = "0"
	}
	if endblock == "" {
		endblock = "9999999999"
	}
	if page == "" {
		page = "1"
	}
	if offset == "" {
		offset = "30"
	}
	if sort == "" {
		sort = "asc"
	}
	var resp *GetInternalTransactionsResp
	err := get(ParseUrl(Url, "module", "account", "action", "txlistinternal", "apikey", apikey, "startblock", startblock, "endblock", endblock, "page", page, "offset", offset, "sort", sort), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

// Get a list of 'ERC20 - Token Transfer Events' by Address
// Returns the list of ERC-20 tokens transferred by an address, with optional filtering by token contract.
//
// ERC-20 transfers from an address, specify the address parameter
// ERC-20 transfers from a contract address, specify the contract address parameter
// ERC-20 transfers from an address filtered by a token contract, specify both address and contract address parameters.
// 
// Parameters:
//	apikey: YourApiKeyToken
//	token: the string representing the token contract address to check for balance
//	address: the string representing the address to check for balance
//	startblock: the integer block number to start searching for transactions. default '0' optional
//	endblock: the integer block number to stop  searching for transactions. default '9999999999' optional
//	page: the integer page number, if pagination is enabled. default '1' optional
//	offset: the number of transactions displayed per page. default '30' optional
//	sort:the sorting preference, use 'asc' to sort by ascending and 'desc' to sort by descendin Tip: Specify a smaller startblock and endblock range for faster search results. default 'asc' optional
func GetERC20TokenTransferEvents(apikey, token, address, startblock, endblock, page, offset, sort string) ([]ERC20TokenTransferEvents, error) {
	if apikey == "" || (token == "" && address == ""){
		return nil, fmt.Errorf("参数不可为空")
	}
	if startblock == "" {
		startblock = "0"
	}
	if endblock == "" {
		endblock = "9999999999"
	}
	if page == "" {
		page = "1"
	}
	if offset == "" {
		offset = "30"
	}
	if sort == "" {
		sort = "asc"
	}
	param := ""
	if address == "" {
		param = "contractaddress=" + token + "&"
	} else if token == "" {
		param = "address=" + address + "&"
	} else {
		param = "contractaddress=" + token + "&" + "address=" + address + "&"
	}
	var resp *GetERC20TokenTransferEventsResp
	err := get(ParseUrl(Url + param, "module", "account", "action", "tokentx", "apikey", apikey, "startblock", startblock, "endblock", endblock, "page", page, "offset", offset, "sort", sort), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}


// Get a list of 'ERC721 - Token Transfer Events' by Address
// Returns the list of ERC-721 ( NFT ) tokens transferred by an address, with optional filtering by token contract.
//
// ERC-721 transfers from an address, specify the address parameter
// ERC-721 transfers from a contract address, specify the contract address parameter
// ERC-721 transfers from an address filtered by a token contract, specify both address and contract address parameters.
// 
// Parameters:
//	apikey: YourApiKeyToken
//	token: the string representing the token contract address to check for balance
//	address: the string representing the address to check for balance
//	startblock: the integer block number to start searching for transactions. default '0' optional
//	endblock: the integer block number to stop  searching for transactions. default '9999999999' optional
//	page: the integer page number, if pagination is enabled. default '1' optional
//	offset: the number of transactions displayed per page. default '30' optional
//	sort:the sorting preference, use 'asc' to sort by ascending and 'desc' to sort by descendin Tip: Specify a smaller startblock and endblock range for faster search results. default 'asc' optional
func GetERC721TokenTransferEvents(apikey, token, address, startblock, endblock, page, offset, sort string) ([]ERC721TokenTransferEvents, error) {
	if apikey == "" || (token == "" && address == ""){
		return nil, fmt.Errorf("参数不可为空")
	}
	if startblock == "" {
		startblock = "0"
	}
	if endblock == "" {
		endblock = "9999999999"
	}
	if page == "" {
		page = "1"
	}
	if offset == "" {
		offset = "30"
	}
	if sort == "" {
		sort = "asc"
	}
	param := ""
	if address == "" {
		param = "contractaddress=" + token + "&"
	} else if token == "" {
		param = "address=" + address + "&"
	} else {
		param = "contractaddress=" + token + "&" + "address=" + address + "&"
	}
	var resp *GetERC721TokenTransferEventsResp
	err := get(ParseUrl(Url + param, "module", "account", "action", "tokennfttx", "apikey", apikey, "startblock", startblock, "endblock", endblock, "page", page, "offset", offset, "sort", sort), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}


// Get a list of 'ERC1155 - Token Transfer Events' by Address
// Returns the list of ERC-1155 ( Multi Token Standard ) tokens transferred by an address, with optional filtering by token contract.
//
// ERC-1155 transfers from an address, specify the address parameter
// ERC-1155 transfers from a contract address, specify the contract address parameter
// ERC-1155 transfers from an address filtered by a token contract, specify both address and contract address parameters.
// 
// Parameters:
//	apikey: YourApiKeyToken
//	token: the string representing the token contract address to check for balance
//	address: the string representing the address to check for balance
//	startblock: the integer block number to start searching for transactions. default '0' optional
//	endblock: the integer block number to stop  searching for transactions. default '9999999999' optional
//	page: the integer page number, if pagination is enabled. default '1' optional
//	offset: the number of transactions displayed per page. default '30' optional
//	sort:the sorting preference, use 'asc' to sort by ascending and 'desc' to sort by descendin Tip: Specify a smaller startblock and endblock range for faster search results. default 'asc' optional
func GetERC1155TokenTransferEvents(apikey, token, address, startblock, endblock, page, offset, sort string) ([]ERC1155TokenTransferEvents, error) {
	if apikey == "" || (token == "" && address == ""){
		return nil, fmt.Errorf("参数不可为空")
	}
	if startblock == "" {
		startblock = "0"
	}
	if endblock == "" {
		endblock = "9999999999"
	}
	if page == "" {
		page = "1"
	}
	if offset == "" {
		offset = "30"
	}
	if sort == "" {
		sort = "asc"
	}
	param := ""
	if address == "" {
		param = "contractaddress=" + token + "&"
	} else if token == "" {
		param = "address=" + address + "&"
	} else {
		param = "contractaddress=" + token + "&" + "address=" + address + "&"
	}
	var resp *GetERC1155TokenTransferEventsResp
	err := get(ParseUrl(Url + param, "module", "account", "action", "token1155tx", "apikey", apikey, "startblock", startblock, "endblock", endblock, "page", page, "offset", offset, "sort", sort), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}


// Get list of Blocks Mined by Address
// Returns the list of blocks mined by an address.
// 
// Parameters:
//	apikey: YourApiKeyToken
//	address: the string representing the address to check for balance
//	blocktype: the string pre-defined block type, either blocks for canonical 'blocks' or 'uncles' for uncle blocks only
//	page: the integer page number, if pagination is enabled. default '1' optional
//	offset: the number of transactions displayed per page. default '30' optional
func GetBlocksMined(apikey, address, blocktype, page, offset string) ([]BlocksMined, error) {
	if apikey == "" || address == "" {
		return nil, fmt.Errorf("参数不可为空")
	}
	if page == "" {
		page = "1"
	}
	if offset == "" {
		offset = "30"
	}
	if blocktype == "" {
		blocktype = "blocks"
	}
	var resp *GetBlocksMinedResp
	err := get(ParseUrl(Url, "module", "account", "action", "getminedblocks", "apikey", apikey, "blocktype", blocktype, "page", page, "offset", offset), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}