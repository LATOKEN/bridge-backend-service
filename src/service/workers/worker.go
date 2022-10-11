package workers

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/latoken/bridge-backend-service/src/models"
	"github.com/latoken/bridge-backend-service/src/service/storage"
)

// IWorker ...
type IWorker interface {
	// GetChain returns unique name of the chain(like LA, ETH and etc)
	GetChainID() string
	GetChainName() string
	GetDestinationID() string
	// GetWokrerAddress returns worker address
	GetWorkerAddress() string
	// GetStartHeight returns blockchain start height for watcher
	GetStartHeight() (int64, error)
	// GetConfirmNum returns numbers of blocks after them tx will be confirmed
	GetConfirmNum() int64
	// GetHeight returns current height of chain
	GetHeight() (int64, error)
	// GetBlockAndTxs returns block info and txs included in this block
	GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error)
	// GetFetchInterval returns fetch interval of the chain like average blocking time, it is used in observer
	GetFetchInterval() time.Duration
	GetGasPrice() float64
	GetConfig() *models.WorkerConfig
	// gets tx status from chain
	GetSentTxStatus(hash string, nonce string) storage.TxStatus
	GetStatus() (*models.WorkerStatus, error)
	// IsSameAddress returns is addrA the same with addrB
	IsSameAddress(addrA string, addrB string) bool
	//Executes Swap on ETH based chains
	ExecuteProposalEth(depositNonce uint64, originChainID [8]byte, destinationChainID [8]byte, resourceID [32]byte, receiptAddr string, amount string, bytes []byte) (string, string, error)
	//Executes Swap on Lachain
	ExecuteProposalLa(depositNonce uint64, originChainID [8]byte, destinationChainID [8]byte, resourceID [32]byte, receiptAddr string, amount string, bytes []byte) (string, string, error)
	//to get Liquidity Index for aave tokens
	GetLiquidityIndex(handlerAddress, usdtAddress common.Address) ([]byte, error)
	//updates withdraw swap status on lachain
	UpdateSwapStatusOnChain(depositNonce uint64, originChainID [8]byte, destinationChainID [8]byte, resourceID [32]byte, receiptAddr string, outAmount, inAmount *big.Int, bytes []byte, status uint8) (string, string, error)
	//gets decimals from token address by taking resource id
	GetDecimalsFromResourceID(resourceID string) (uint8, error)
}
