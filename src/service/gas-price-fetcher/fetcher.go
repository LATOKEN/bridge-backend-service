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
			gasPrice, err := f.getBscGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for BSC %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "ETH":
			gasPrice, err := f.getEthGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for ETH %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "POS":
			gasPrice, err := f.getPosGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for POS %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "AVAX":
			gasPrice, err := f.getAvaxGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for AVAX %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "FTM":
			gasPrice, err := f.getFtmGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for FTM %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "CRO":
			gasPrice, err := f.getCroGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for CRO %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "ARB":
			gasPrice, err := f.getArbGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for ARB %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "HT":
			gasPrice, err := f.getHtGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for HT %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "ONE":
			gasPrice, err := f.getOneGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for ONE %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		case "OP":
			gasPrice, err := f.getOpGasPrice(cfg)
			if err != nil {
				logrus.Warnf("error fetching gas price for OP %s", err.Error())
				continue
			}
			gasPrices = append(gasPrices, gasPrice)
		default:
			logrus.Warnf("Gas price getter not implemented for ", cfg.ChainName)
		}
	}

	f.storage.SaveGasPriceInfo(gasPrices)
	f.logger.Infoln("New gas prices fetched")
}

func (f *FetcherSrv) getBscGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	// httpClient := &http.Client{
	// 	Timeout: time.Second * 10,
	// }

	// resp, err := f.makeReq(f.chainFetCfgs[storage.BscChain].URL, httpClient)
	// if err != nil {
	// 	logrus.Warnf("fetch BSC gas price error = %s", err)
	// 	return &storage.GasPrice{}
	// }
	// var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "BSC", Price: fmt.Sprintf("%f", 20.0), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getEthGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch ETH gas price error = %s", err)
		return &storage.GasPrice{}, err
	}
	var gasPrice = (*resp)["average"].(float64) / 10
	return &storage.GasPrice{ChainName: "ETH", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getPosGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch POS gas price error = %s", err)
		return &storage.GasPrice{}, err
	}
	var gasPrice = (*resp)["fast"].(float64)
	return &storage.GasPrice{ChainName: "POS", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getAvaxGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch AVAX gas price error = %s", err)
		return &storage.GasPrice{}, err
	}
	gasPrice := (*resp)["data"].(map[string]interface{})["normal"].(map[string]interface{})["price"].(float64) / 1000000000
	return &storage.GasPrice{ChainName: "AVAX", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getFtmGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch FTM gas price error = %s", err)
		return &storage.GasPrice{}, err
	}
	gasPrice := (*resp)["data"].(map[string]interface{})["normal"].(map[string]interface{})["price"].(float64) / 1000000000
	return &storage.GasPrice{ChainName: "FTM", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getCroGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch CRO gas price error = %s", err)
		return &storage.GasPrice{}, err
	}
	gasPrice := (*resp)["data"].(map[string]interface{})["normal"].(map[string]interface{})["price"].(float64) / 1000000000
	return &storage.GasPrice{ChainName: "CRO", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getArbGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)

	if err != nil {
		logrus.Warnf("fetch ARB gas price error = %s", err)
		return &storage.GasPrice{}, err
	}

	gasPrice := (*resp)["data"].(map[string]interface{})["normal"].(map[string]interface{})["price"].(float64) / 1000000000
	return &storage.GasPrice{ChainName: "ARB", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getHtGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch HT gas price error = %s", err)
		return &storage.GasPrice{}, err
	}
	gasPrice := (*resp)["data"].(map[string]interface{})["normal"].(map[string]interface{})["price"].(float64) / 1000000000
	return &storage.GasPrice{ChainName: "HT", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getOneGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := f.makeReq(cfg.URL, httpClient)
	if err != nil {
		logrus.Warnf("fetch ONE gas price error = %s", err)
		return &storage.GasPrice{}, err
	}
	var gasPrice = (*resp)["standard"].(float64)
	return &storage.GasPrice{ChainName: "ONE", Price: fmt.Sprintf("%f", gasPrice), UpdateTime: time.Now().Unix()}, nil
}

func (f *FetcherSrv) getOpGasPrice(cfg *models.FetcherConfig) (*storage.GasPrice, error) {
	return &storage.GasPrice{ChainName: "OP", Price: fmt.Sprintf("%f", 0.001), UpdateTime: time.Now().Unix()}, nil
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
