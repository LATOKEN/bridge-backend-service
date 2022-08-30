package config

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/latoken/bridge-backend-service/src/models"
	"github.com/latoken/bridge-backend-service/src/service/storage"

	"github.com/ethereum/go-ethereum/common"
)

// Reads Service params from config.json
func (v *viperConfig) ReadServiceConfig() string {
	return fmt.Sprintf("%s:%s", v.GetString("service.host"), v.GetString("service.port"))
}

// reads ethereum chain apis to fetch gas price
func (v *viperConfig) ReadFetcherConfig() []*models.FetcherConfig {
	chains := v.GetStringSlice("chains")
	chainFetCgfs := make([]*models.FetcherConfig, 0)

	for _, chain := range chains {
		chainFetCgfs = append(chainFetCgfs, v.readFetcherConfig(chain))
	}

	return chainFetCgfs
}

func (v *viperConfig) readFetcherConfig(name string) *models.FetcherConfig {
	return &models.FetcherConfig{
		ChainName: name,
		URL:       v.GetString(fmt.Sprintf("gas_price_api.%s", name)),
	}
}

// ReadLachainConfig reads lachain chain params from config.json
func (v *viperConfig) ReadLachainConfig() *models.WorkerConfig {
	return v.readWorkerConfig("LA")
}

func (v *viperConfig) ReadWorkersConfig() []*models.WorkerConfig {
	chains := v.GetStringSlice("chains")
	chainCfgs := make([]*models.WorkerConfig, 0)

	for _, chain := range chains {
		chainCfgs = append(chainCfgs, v.readWorkerConfig(chain))
	}

	return chainCfgs
}

// readETHWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) readWorkerConfig(name string) *models.WorkerConfig {
	return &models.WorkerConfig{
		NetworkType:           v.GetString(fmt.Sprintf("workers.%s.type", name)),
		ChainName:             strings.ToUpper(name),
		WorkerAddr:            common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.worker_addr", name))),
		PrivateKey:            v.GetString(fmt.Sprintf("workers.%s.private_key", name)),
		Provider:              v.GetString(fmt.Sprintf("workers.%s.provider", name)),
		ContractAddr:          common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.contract_addr", name))),
		AMUSDTContractAddr:    common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.amUSDT_token_addr", name))),
		AmTokenHandlerAddress: common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.amToken_handler_addr", name))),
		NativeResourceID:      v.GetString(fmt.Sprintf("workers.%s.native_resource_id", name)),
		GasLimit:              v.GetInt64(fmt.Sprintf("workers.%s.gas_limit", name)),
		GasPrice:              big.NewInt(v.GetInt64(fmt.Sprintf("workers.%s.gas_price", name))),
		FetchInterval:         v.GetInt64(fmt.Sprintf("workers.%s.fetch_interval", name)),
		ConfirmNum:            v.GetInt64(fmt.Sprintf("workers.%s.confirm_num", name)),
		StartBlockHeight:      v.GetInt64(fmt.Sprintf("workers.%s.start_block_height", name)),
		DestinationChainID:    v.GetString(fmt.Sprintf("workers.%s.dest_id", name)),
	}
}

// Reads storage params from config.json
func (v *viperConfig) ReadDBConfig() *models.StorageConfig {
	return &models.StorageConfig{
		URL:        v.GetString("storage.url"),
		DBDriver:   v.GetString("storage.driver"),
		DBHOST:     v.GetString("storage.host"),
		DBPORT:     v.GetInt64("storage.port"),
		DBSSL:      v.GetString("storage.ssl_mode"),
		DBName:     v.GetString("storage.db_name"),
		DBUser:     v.GetString("storage.user"),
		DBPassword: v.GetString("storage.password"),
	}
}

func (v *viperConfig) ReadResourceIDs() []*storage.ResourceId {
	tokens := v.GetStringSlice("all_tokens")
	resouceIDs := make([]*storage.ResourceId, len(tokens))
	for index, name := range tokens {
		resouceIDs[index] = &storage.ResourceId{
			Name: name,
			ID:   v.GetString(fmt.Sprintf("resourceIDs.%s", name)),
		}
	}
	return resouceIDs
}

func (v *viperConfig) ReadChains() []string {
	return v.GetStringSlice("chains")
}
