package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/latoken/bridge-backend-service/src/models"
	"github.com/latoken/bridge-backend-service/src/service/storage"
	"github.com/sirupsen/logrus"
)

// FetcherSrv
type FetcherSrv struct {
	logger       *logrus.Entry
	storage      *storage.DataBase
	chainFetCfgs []*models.FetcherConfig
}

// CreateFetcherSrv
func CreateFetcherSrv(logger *logrus.Logger, db *storage.DataBase, chainFetCfgs []*models.FetcherConfig) *FetcherSrv {
	return &FetcherSrv{
		logger:       logger.WithField("layer", "fetcher"),
		storage:      db,
		chainFetCfgs: chainFetCfgs,
	}
}

func (f *FetcherSrv) Run() {
	f.logger.Infoln("Fetcher srv started")
	go f.collector()
}

func (f *FetcherSrv) collector() {
	for {
		f.getAllGasPrice()
		time.Sleep(60 * time.Second)
	}
}

func (f *FetcherSrv) getAllGasPrice() {
	gasPrices := make([]*storage.GasPrice, 0, len(f.chainFetCfgs))

	for _, cfg := range f.chainFetCfgs {

		gasPrice, err := f.getGasPrice(cfg, cfg.ChainName)
		if err != nil {
			f.logger.Warnf("error fetching gas price for %s %s", cfg.ChainName, err.Error())
			continue
		}
		gasPrices = append(gasPrices, gasPrice)

	}

	f.storage.SaveGasPriceInfo(gasPrices)
	f.logger.Infoln("New gas prices fetched")
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (f *FetcherSrv) getGasPrice(cfg *models.FetcherConfig, chainName string) (*storage.GasPrice, error) {

	var gasPrice = 0.0

	f.logger.Infoln("chainName:: ", chainName)

	if chainName == "OP" {
		gasPrice = 0.001
	} else {
		httpClient := &http.Client{
			Timeout: time.Second * 10,
		}

		resp, err := f.makeReq(cfg.URL, httpClient)
		if err != nil {
			f.logger.Warnf("fetch %s gas price error = %s", chainName, err)
			return &storage.GasPrice{}, err
		} else if resp == nil {
			return &storage.GasPrice{}, fmt.Errorf("Fetched empty response for %s", chainName)
		}

		// if stringInSlice(chainName, []string{"ETH"}) {
		// 	gasPrice = (*resp)["average"].(float64) / 10

		if stringInSlice(chainName, []string{"POS"}) {
			gasPrice = (*resp)["fast"].(float64)

		} else if stringInSlice(chainName, []string{"AVAX", "FTM", "HT", "CRO", "ARB", "ETH"}) {
			gasPrice = (*resp)["data"].(map[string]interface{})["normal"].(map[string]interface{})["price"].(float64) / 1000000000

		} else if stringInSlice(chainName, []string{"ONE", "BSC"}) {
			gasPrice = (*resp)["standard"].(float64)
		} else if stringInSlice(chainName, []string{"OP"}) {
			gasPrice = 0.001
		}
	}

	f.logger.Println("gasPrice: ", gasPrice)

	if gasPrice == 0.0 {
		return &storage.GasPrice{}, fmt.Errorf("Gas price getter not implemented for %s chain", chainName)
	}

	return &storage.GasPrice{ChainName: chainName, Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

// MakeReq HTTP request helper
func (f *FetcherSrv) makeReq(url string, c *http.Client) (*map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := f.doReq(req, c)
	if err != nil {
		return nil, err
	}

	t := make(map[string]interface{})
	er := json.Unmarshal(resp, &t)
	if er != nil {
		return nil, er
	}

	return &t, err
}

// helper
// doReq HTTP client
func (f *FetcherSrv) doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
