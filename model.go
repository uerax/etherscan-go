package etherscan

// GetBalance
type GetBalanceResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// GetBalanceMulti
type GetBalanceMultiResp struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []AddrBalance `json:"result"`
}

type AddrBalance struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

// GetNormalTransactions
type GetNormalTransactionsResp struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Result  []NormalTransaction `json:"result"`
}

type NormalTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxreceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodID          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
}

// GetInternalTransactions
// GetInternalTransactionsByHash
// GetInternalTransactionsByBlockRange
type GetInternalTransactionsResp struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  []InternalTransaction `json:"result"`
}

type InternalTransaction struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Gas             string `json:"gas"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	GasUsed         string `json:"gasUsed"`
	TraceId         string `json:"traceId"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}


// GetERC20TokenTransferEvents
type GetERC20TokenTransferEventsResp struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  []ERC20TokenTransferEvents `json:"result"`
}

type ERC20TokenTransferEvents struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

// GetERC721TokenTransferEvents
type GetERC721TokenTransferEventsResp struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  []ERC721TokenTransferEvents `json:"result"`
}

type ERC721TokenTransferEvents struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

// GetERC1155TokenTransferEvents
type GetERC1155TokenTransferEventsResp struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  []ERC1155TokenTransferEvents `json:"result"`
}

type ERC1155TokenTransferEvents struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	From              string `json:"from"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenValue        string `json:"tokenValue"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	Confirmations     string `json:"confirmations"`
}

// GetBlocksMined
type GetBlocksMinedResp struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  []BlocksMined `json:"result"`
}

type BlocksMined struct {
	BlockNumber string `json:"blockNumber"`
	TimeStamp   string `json:"timeStamp"`
	BlockReward string `json:"blockReward"`
}