package types

import (
	_ "github.com/33cn/chain33/system/address" //init address
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

//Header block header
type Header struct {
	ParentHash  string      `json:"parentHash"       gencodec:"required"`
	UncleHash   string      `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    string      `json:"miner"            gencodec:"required"`
	Root        string      `json:"stateRoot"        gencodec:"required"`
	TxHash      string      `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash string      `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom `json:"logsBloom"        gencodec:"required"`
	Difficulty  string      `json:"difficulty"       gencodec:"required"`
	Number      string      `json:"number"           gencodec:"required"`
	GasLimit    string      `json:"gasLimit"         gencodec:"required"`
	GasUsed     string      `json:"gasUsed"          gencodec:"required"`
	Time        string      `json:"timestamp"        gencodec:"required"`
	Extra       []byte      `json:"extraData"        gencodec:"required"`
	MixDigest   string      `json:"mixHash"`
	Nonce       string      `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee string `json:"baseFeePerGas" rlp:"optional"`
}

// Transaction LegacyTx is the transaction data of regular Ethereum transactions.
type Transaction struct {
	BlockHash        string `json:"blockHash,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	From             string `json:"from,omitempty"`
	To               string `json:"to,omitempty"`
	Hash             string `json:"hash,omitempty"`
	Data             string `json:"input,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`
	Value            string `json:"value,omitempty"`
	Type             string `json:"type,omitempty"`
	V                string `json:"v,omitempty"`
	R                string `json:"r,omitempty"`
	S                string `json:"s,omitempty"`
}

//Transactions txs types
type Transactions []*Transaction

//Block ETH 交易结构体
type Block struct {
	*Header
	Uncles []*Header `json:"uncles"`
	//TODO 保留ETH的交易结构类型还是替换为chain33的Transaction
	Transactions interface{} `json:"transactions"`
	Hash         string      `json:"hash"`
}

// Receipt represents the results of a transaction.
type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type      string `json:"type,omitempty"`
	PostState []byte `json:"root"`
	//0x1 表示交易成功，0x2 表示交易失败
	Status            string `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" gencodec:"required"`
	//Bloom             Bloom  `json:"logsBloom"         gencodec:"required"`
	Logs interface{} `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash          string `json:"transactionHash" gencodec:"required"`
	ContractAddress string `json:"contractAddress"`
	GasUsed         uint64 `json:"gasUsed" gencodec:"required"`
	To              string `json:"to,omitempty"`
	From            string `json:"from,omitempty"`
	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash string `json:"blockHash,omitempty"`
	//十六进制
	BlockNumber      string `json:"blockNumber,omitempty"`
	TransactionIndex string `json:"transactionIndex"`
}

//CallMsg eth api param
type CallMsg struct {
	From     string          `json:"from"`
	To       string          `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons. "input" is the
	// newer name and should be preferred by clients.
	Data  string `json:"data"`
	Input string `json:"input"`
}

//EvmLog evm log
type EvmLog struct {
	Topic []hexutil.Bytes `json:"topic,omitempty"`
	Data  hexutil.Bytes   `json:"data,omitempty"`
}

//Peer peer info
type Peer struct {
	ID         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	NetWork    *Network   `json:"netWork,omitempty"`
	Protocols  *Protocols `json:"protocols,omitempty"`
	Self       bool       `json:"self,omitempty"`
	Ports      *Ports     `json:"ports,omitempty"`
	Encode     string     `json:"encode,omitempty"`
	ListenAddr string     `json:"listenAddr,omitempty"`
}

//Network networkinfo
type Network struct {
	LocalAddress  string `json:"localAddress,omitempty"`
	RemoteAddress string `json:"remoteAddress,omitempty"`
}

//Protocols  peer protocols
type Protocols struct {
	EthProto *EthProto `json:"eth,omitempty"`
}

//EthProto eth proto
type EthProto struct {
	Difficulty uint32 `json:"difficulty,omitempty"`
	Head       string `json:"head,omitempty"`
	Version    string `json:"version,omitempty"`
	NetworkID  int    `json:"network,omitempty"`
}

//Ports ...
type Ports struct {
	Discovery int32 `json:"discovery,omitempty"`
	Listener  int32 `json:"listener,omitempty"`
}

//SubLogs ...
//logs Subscription
type SubLogs struct {
	Address string   `json:"address,omitempty"`
	Topics  []string `json:"topics,omitempty"`
}

// EvmLogInfo  ...
type EvmLogInfo struct {
	Address          string      `json:"address,omitempty"`
	BlockHash        string      `json:"blockHash,omitempty"`
	BlockNumber      string      `json:"blockNumber,omitempty"`
	LogIndex         string      `json:"logIndex,omitempty"`
	Topics           interface{} `json:"topics,omitempty"`
	TransactionHash  string      `json:"transactionHash,omitempty"`
	TransactionIndex string      `json:"transactionIndex,omitempty"`
}