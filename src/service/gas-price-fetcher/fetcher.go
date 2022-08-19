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

//FetcherSrv
type FetcherSrv struct {
	logger       *logrus.Entry
	storage      *storage.DataBase
	chainFetCfgs []*models.FetcherConfig
}

//CreateFetcherSrv
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
		switch cfg.ChainName {
		case "BSC":
			gasPrice, err := f.getGasPrice(cfg, "BSC")
			if err != nil {
				f.logger.Warnf("error fetching gas price for BSC %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "ETH":
			gasPrice, err := f.getGasPrice(cfg, "ETH")
			if err != nil {
				f.logger.Warnf("error fetching gas price for ETH %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "POS":
			gasPrice, err := f.getGasPrice(cfg, "POS")
			if err != nil {
				f.logger.Warnf("error fetching gas price for POS %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "AVAX":
			gasPrice, err := f.getGasPrice(cfg, "AVAX")
			if err != nil {
				f.logger.Warnf("error fetching gas price for AVAX %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "FTM":
			gasPrice, err := f.getGasPrice(cfg, "FTM")
			if err != nil {
				f.logger.Warnf("error fetching gas price for FTM %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "CRO":
			gasPrice, err := f.getGasPrice(cfg, "CRO")
			if err != nil {
				f.logger.Warnf("error fetching gas price for CRO %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "ARB":
			gasPrice, err := f.getGasPrice(cfg, "ARB")
			if err != nil {
				f.logger.Warnf("error fetching gas price for ARB %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "HT":
			gasPrice, err := f.getGasPrice(cfg, "HT")
			if err != nil {
				f.logger.Warnf("error fetching gas price for HT %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "ONE":
			gasPrice, err := f.getGasPrice(cfg, "ONE")
			if err != nil {
				f.logger.Warnf("error fetching gas price for ONE %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "OP":
			gasPrice, err := f.getGasPrice(cfg, "OP")
			if err != nil {
				f.logger.Warnf("error fetching gas price for OP %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		default:
			f.logger.Warnf("Gas price getter not implemented for ", cfg.ChainName)
		}
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

	var gasPrice float64

	if chainName == "BSC" {
		gasPrice = 20.0
		return &storage.GasPrice{ChainName: chainName, Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
	} else if chainName == "OP" {
		gasPrice = 0.001
		return &storage.GasPrice{ChainName: chainName, Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)

	if err != nil {
		f.logger.Warnf("fetch %s gas price error = %s", chainName, err)
		return &storage.GasPrice{}, err
	} else if resp == nil {
		return &storage.GasPrice{}, fmt.Errorf("Wrong gas price fetched for POS")
	}

	if stringInSlice(chainName, []string{"ETH"}) {
		gasPrice = (*resp)["average"].(float64) / 10

	} else if stringInSlice(chainName, []string{"POS"}) {
		gasPrice = (*resp)["fast"].(float64)

	} else if stringInSlice(chainName, []string{"AVAX", "FTM", "HT", "CRO", "ARB"}) {
		gasPrice = (*resp)["data"].(map[string]interface{})["normal"].(map[string]interface{})["price"].(float64) / 1000000000

	} else if stringInSlice(chainName, []string{"ONE"}) {
		gasPrice = (*resp)["standard"].(float64)
	} else if stringInSlice(chainName, []string{"OP"}) {
		gasPrice = 0.001
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
