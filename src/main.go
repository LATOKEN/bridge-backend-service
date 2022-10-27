package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/latoken/bridge-backend-service/src/app"
	"github.com/latoken/bridge-backend-service/src/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewViperConfig()
	srvURL := cfg.ReadServiceConfig()
	laCfg := cfg.ReadLachainConfig()
	chainCfgs := cfg.ReadWorkersConfig()
	chainFetCfgs := cfg.ReadFetcherConfig()
	dbConfig := cfg.ReadDBConfig()
	dbURL := fmt.Sprintf(dbConfig.URL, dbConfig.DBHOST, dbConfig.DBPORT, dbConfig.DBUser, dbConfig.DBName, dbConfig.DBPassword, dbConfig.DBSSL)
	resourceIDs := cfg.ReadResourceIDs()
	// init logrus logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
	// set logger level
	level, err := logrus.ParseLevel(cfg.GetString("logger-level"))
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	// os.MkdirAll("./logs", os.ModePerm)
	// logFile, err := os.OpenFile("./logs/backend.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
	// }
	// logger.SetOutput(logFile)

	// Set connection to onlife_business database
	db, err := gorm.Open(dbConfig.DBDriver, dbURL)
	if err != nil {
		logger.WithFields(logrus.Fields{"dbURL": dbURL}).Fatalf("Set connection to PostgreSQL: ", err.Error())
	}
	defer db.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sign := <-c
		logger.Infof("System signal: %+v\n", sign)
		cancel()
	}()

	app := app.NewApp(logger, srvURL, db, laCfg, chainCfgs, chainFetCfgs, resourceIDs)

	//run App
	app.Run(ctx)
}

// client, err := ethclient.Dial("https://ropsten.infura.io/v3/8b31a268c67b4a6f839db69b8e9a9cdc")
// if err != nil {
// 	panic(fmt.Sprintf("new eth client error, err=%s", err.Error()))
// }

// header, err := client.HeaderByNumber(context.Background(), big.NewInt(10114174))
// if err != nil {
// 	panic(fmt.Sprintf("new eth client error, err=%s", err.Error()))
// }

// // header2, err := client.HeaderByHash(context.Background(), common.HexToHash("0xa7fbb9720bd34065559550c7270f35918720bc9562247f112dcf21f49eb3f81a"))
// // if err != nil {
// // 	panic(fmt.Sprintf("new eth client error, err=%s", err.Error()))
// // }

// fmt.Println(header.Hash().Hex(), header.Number)
