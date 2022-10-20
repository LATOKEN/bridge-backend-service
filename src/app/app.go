package app

import (
	"context"
	"net/http"
	"time"

	"github.com/latoken/bridge-backend-service/src/models"
	rlr "github.com/latoken/bridge-backend-service/src/service"
	"github.com/latoken/bridge-backend-service/src/service/storage"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// App ...
type App struct {
	logger  *logrus.Logger
	router  *mux.Router
	server  *http.Server
	relayer *rlr.BridgeSRV
}

// NewApp is initializes the app
func NewApp(logger *logrus.Logger, addr string, db *gorm.DB,
	laCfg *models.WorkerConfig, chainCfgs []*models.WorkerConfig, chainFetCfgs []*models.FetcherConfig,
	resourceIDs []*storage.ResourceId) *App {
	// create new app
	inst := &App{
		logger:  logger,
		router:  mux.NewRouter(),
		server:  &http.Server{Addr: addr},
		relayer: rlr.CreateNewBridgeSRV(logger, db, laCfg, chainCfgs, chainFetCfgs, resourceIDs),
	}
	// set router
	inst.router = mux.NewRouter()
	inst.setRouters()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{""})

	inst.server.Handler = handlers.CORS(headers, methods, origins)(inst.router)

	inst.relayer.Run()

	return inst
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}

func (a *App) setRouters() {
	a.Get("/", a.Endpoints)
	a.Get("/status", a.StatusHandler)
	a.Get("/gas-price/{chain}", a.GasPriceHandler)
	a.Get("/tx-sent/{tx_hash}", a.TxSentHandler)
	// a.Get("/resend_tx/{id}", a.ResendTxHandler)
	// a.Get("/set_mode/{mode}", a.SetModeHandler)
}

// Run the app on it's router
func (a *App) Run(ctx context.Context) {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal(err)
		}
	}()

	a.logger.Infof("Bridge backend service has started on %s\nPress ctrl + C to exit.", a.server.Addr)

	<-ctx.Done()

	a.logger.Infoln("Bridge backend service has stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctxShutDown); err != nil {
		a.logger.Fatalf("Shutdown: %v\n", err)
	}
}
