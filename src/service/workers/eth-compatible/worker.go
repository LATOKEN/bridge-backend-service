package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"time"

	ethBr "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/eth"
	laBr "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/la"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	jsonrpc "github.com/ybbus/jsonrpc/v2"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/models"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
)

// Erc20Worker ...
type Erc20Worker struct {
	provider           string
	chainName          string
	chainID            int64
	destinationChainID string
	storage            *storage.DataBase
	logger             *logrus.Entry // logger
	config             *models.WorkerConfig
	client             *ethclient.Client
	contractAddr       common.Address
	proxyContractAddr  common.Address
}

// NewErc20Worker ...
func NewErc20Worker(logger *logrus.Logger, cfg *models.WorkerConfig, db *storage.DataBase) *Erc20Worker {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		panic("new eth client error")
	}

	privKey, err := utils.GetPrivateKey(cfg)
	if err != nil {
		panic(fmt.Sprintf("generate private key error, err=%s", err.Error()))
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Sprintf("get public key error, err=%s", err.Error()))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if !bytes.Equal(cfg.WorkerAddr.Bytes(), fromAddress.Bytes()) {
		panic(fmt.Sprintf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.WorkerAddr, fromAddress,
		))
	}

	// init token addresses
	return &Erc20Worker{
		chainName:          cfg.ChainName,
		chainID:            cfg.ChainID,
		destinationChainID: cfg.DestinationChainID,
		logger:             logger.WithField("worker", cfg.ChainName),
		provider:           cfg.Provider,
		config:             cfg,
		client:             client,
		contractAddr:       cfg.ContractAddr,
		proxyContractAddr:  cfg.ProxyContractAddr,
		storage:            db,
	}
}

// GetChainName returns chain ID
func (w *Erc20Worker) GetChainName() string {
	return w.chainName
}

// GetChainName returns chain ID
func (w *Erc20Worker) GetChainID() string {
	return string(w.chainID)
}

//returns destinationChainID to be checked to execute proposal
func (w *Erc20Worker) GetDestinationID() string {
	return w.destinationChainID
}

// GetStartHeight returns start blockchain height from config
func (w *Erc20Worker) GetStartHeight() (int64, error) {
	if w.config.StartBlockHeight == 0 {
		return w.GetHeight()
	}
	return w.config.StartBlockHeight, nil
}

// GetConfirmNum returns numbers of blocks after them tx will be confirmed
func (w *Erc20Worker) GetConfirmNum() int64 {
	return w.config.ConfirmNum
}

func (w *Erc20Worker) ExecuteProposalEth(depositNonce uint64, originChainID [8]byte, destinationChainID [8]byte, resourceID [32]byte, receiptAddr string, amount string) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := ethBr.NewEthBr(w.proxyContractAddr, w.client)
	if err != nil {
		return "", err
	}
	value, _ := new(big.Int).SetString(amount, 10)
	tx, err := instance.ExecuteProposal(auth, originChainID, destinationChainID, depositNonce, resourceID, common.HexToAddress(receiptAddr), value)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

func (w *Erc20Worker) ExecuteProposalLa(depositNonce uint64, originChainID [8]byte, destinationChainID [8]byte, resourceID [32]byte, receiptAddr string, amount string) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := laBr.NewLabr(w.proxyContractAddr, w.client)
	if err != nil {
		return "", err
	}
	value, _ := new(big.Int).SetString(amount, 10)
	tx, err := instance.ExecuteProposal(auth, originChainID, destinationChainID, depositNonce, resourceID, common.HexToAddress(receiptAddr), value)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// // GetStatus returns status of relayer account(balance eg)
// func (w *Erc20Worker) GetStatus(symbol string) (interface{}, error) {
// 	ethStatus := &EthStatus{}

// 	allowance, err := w.Allowance(symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Allowance = QuoBigInt(allowance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	balance, err := w.Erc20Balance(w.Config.WorkerAddr, symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Erc20Balance = QuoBigInt(balance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	ethBalance, err := w.EthBalance(w.Config.WorkerAddr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.EthBalance = QuoBigInt(ethBalance, GetBigIntForDecimal(18)).String()

// 	return ethStatus, nil
// }

// GetBlockAndTxs ...
func (w *Erc20Worker) GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error) {
	// var head *Header
	// rpcClient := jsonrpc.NewClient(w.provider)

	// resp, err := rpcClient.Call("eth_getBlockByNumber", "latest", false)
	// if err != nil {
	// 	w.logger.Errorln("while call eth_getBlockByNumber, err = ", err)
	// 	return nil, err
	// }

	// if err := resp.GetObject(&head); err != nil {
	// 	w.logger.Errorln("while GetObject, err = ", err)
	// 	return nil, err
	// }

	// if head == nil {
	// 	return nil, fmt.Errorf("not found")
	// }

	// if height >= int64(head.Number) {
	// 	return nil, fmt.Errorf("not found")
	// }

	// logs, err := w.getLogs(height, int64(head.Number))
	// if err != nil {
	// 	w.logger.Errorf("while getEvents(from = %d, to = %d), err = %v", height, int64(head.Number), err)
	// 	return nil, err
	// }

	// return &models.BlockAndTxLogs{
	// 	Height:          int64(head.Number),
	// 	BlockHash:       head.Hash.String(),
	// 	ParentBlockHash: head.ParentHash.Hex(),
	// 	BlockTime:       int64(head.Time),
	// 	TxLogs:          logs,
	// }, nil

	client, err := ethclient.Dial(w.provider)
	if err != nil {
		w.logger.Errorln("Error while dialing the client = ", err)
		return nil, err
	}

	clientResp, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		w.logger.Errorln("Error while fetching the block header = ", err)
		return nil, err

	}

	if height >= clientResp.Number.Int64() {
		return nil, fmt.Errorf("not found")
	}

	logs, err := w.getLogs(height, clientResp.Number.Int64())
	if err != nil {
		w.logger.Errorf("while getEvents(block number from %d to %d), err = %v", height, clientResp.Number, err)
		return nil, err
	}

	client.Close()

	return &models.BlockAndTxLogs{
		Height:          clientResp.Number.Int64(),
		BlockHash:       clientResp.Hash().String(),
		ParentBlockHash: clientResp.ParentHash.Hex(),
		BlockTime:       int64(clientResp.Time),
		TxLogs:          logs,
	}, nil
}

// GetFetchInterval ...
func (w *Erc20Worker) GetFetchInterval() time.Duration {
	return time.Duration(w.config.FetchInterval) * time.Second
}

// getLogs ...
func (w *Erc20Worker) getLogs(curHeight, nextHeight int64) ([]*storage.TxLog, error) {
	//	topics := [][]common.Hash{{DepositEventHash, ProposalEventHash, ProposalVoteHash}}
	if curHeight == 0 || nextHeight-curHeight > 3500 {
		curHeight = nextHeight - 1
	}
	logs, err := w.client.FilterLogs(context.Background(), ethereum.FilterQuery{
		// BlockHash: &blockHash,
		FromBlock: big.NewInt(curHeight + 1),
		ToBlock:   big.NewInt(nextHeight),
		// Topics:    topics,
		Addresses: []common.Address{w.contractAddr},
		Topics:    [][]common.Hash{},
	})

	w.logger.Info(logs)

	if err != nil {
		w.logger.Info("ERR")
		w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("get event log error, err=%s", err)
		return nil, err
	}

	models := make([]*storage.TxLog, 0, len(logs))
	for _, log := range logs {
		w.logger.Infof("WORKER(%s) NEW EVENT: %v\n\n", w.chainName, log)
		event, err := w.parseEvent(&log)
		if err != nil {
			w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("parse event log error, err=%s", err)
			continue
		}
		if event == nil {
			continue
		}

		txLog := event.ToTxLog(w.chainName)
		txLog.Chain = w.chainName
		txLog.Height = int64(log.BlockNumber)
		txLog.EventID = log.TxHash.Hex()
		txLog.BlockHash = log.BlockHash.Hex()
		txLog.TxHash = log.TxHash.Hex()
		txLog.Status = storage.TxStatusInit

		models = append(models, txLog)
	}

	return models, nil
}

// GetHeight ..
func (w *Erc20Worker) GetHeight() (int64, error) {
	header, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, nil
	}
	return header.Number.Int64(), nil
}

// GetSentTxStatus ...
func (w *Erc20Worker) GetSentTxStatus(hash string) storage.TxStatus {
	// _, isPending, err := w.client.TransactionByHash(context.Background(), common.HexToHash(hash))
	// if err != nil {
	// 	w.logger.Errorln("GetSentTxStatus, err = ", err)
	// 	return storage.TxSentStatusNotFound
	// }

	// if isPending {
	// 	return storage.TxSentStatusPending
	// }

	txReceipt, err := w.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return storage.TxSentStatusNotFound
	}

	if txReceipt.Status == types.ReceiptStatusFailed {
		return storage.TxSentStatusFailed
	}

	return storage.TxSentStatusSuccess
}

func (w *Erc20Worker) GetTxCountLatest() (uint64, error) {
	var result hexutil.Uint64
	rpcClient := jsonrpc.NewClient(w.provider)

	resp, err := rpcClient.Call("eth_getTransactionCount", w.config.WorkerAddr.Hex(), "latest")
	if err != nil {
		return 0, err
	}
	if err := resp.GetObject(&result); err != nil {
		return 0, err
	}
	return uint64(result), nil
}

// GetTransactor ...
func (w *Erc20Worker) getTransactor() (auth *bind.TransactOpts, err error) {
	privateKey, err := utils.GetPrivateKey(w.config)
	if err != nil {
		return nil, err
	}

	var nonce uint64
	if w.chainName == storage.LaChain || w.chainName == storage.PosChain {
		nonce, err = w.GetTxCountLatest()
		if err != nil {
			return nil, err
		}
	} else {
		nonce, err = w.client.PendingNonceAt(context.Background(), w.config.WorkerAddr)
		if err != nil {
			return nil, err
		}
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(w.chainID))
	if err != nil {
		return nil, err
	}

	var gasPrice float64
	if w.chainName == storage.LaChain {
		gasPrice = 1
	} else {
		gasPriceGWei, _ := strconv.ParseFloat(w.storage.GetGasPrice(w.chainName).Price, 64)
		if gasPriceGWei > 0 {
			gasPrice = gasPriceGWei * 1000000000
		}
	}
	auth.GasPrice = big.NewInt((int64(gasPrice)))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(w.config.GasLimit) // in units

	return auth, nil
}

// EthBalance ...
func (w *Erc20Worker) EthBalance(address common.Address) (*big.Int, error) {
	return w.client.BalanceAt(context.Background(), address, nil)
}

// GetWorkerAddress ...
func (w *Erc20Worker) GetWorkerAddress() string {
	return w.config.WorkerAddr.String()
}

//GetGasPrice
// func (w *Erc20Worker) GetGasPrice() (uint64, error) {
// 	var result hexutil.Uint64
// 	rpcClient := jsonrpc.NewClient(w.provider)
// 	resp, err := rpcClient.Call("eth_gasPrice")
// 	if err != nil {
// 		return 0, err
// 	}
// 	if err := resp.GetObject(&result); err != nil {
// 		println("err in obj")
// 		return 0, err
// 	}
// 	return uint64(result), nil
// 	time.Sleep(time.Second * 30)
// }

// GetColdWalletAddress ...
func (w *Erc20Worker) GetColdWalletAddress() string {
	return w.config.ColdWalletAddr.String()
}

// IsSameAddress ...
func (w *Erc20Worker) IsSameAddress(addrA string, addrB string) bool {
	return bytes.Equal(common.FromHex(addrA), common.FromHex(addrB))
}

// SendAmount ...
func (w *Erc20Worker) SendAmount(address string, amount *big.Int) (string, error) {
	return "", fmt.Errorf("not implemented") // TODO
}
