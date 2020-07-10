package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/binance-chain/bsc-relayer/common"
	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	Config *config.Config
	db     *gorm.DB
}

func NewAdmin(config *config.Config) *Admin {
	return &Admin{
		Config: config,
	}
}

func (admin *Admin) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{"statistic"},
	}

	jsonBytes, err := json.MarshalIndent(endpoints, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		common.Logger.Error(err)
	}
}

func (admin *Admin) StatisticHandler(w http.ResponseWriter, r *http.Request) {

	jsonBytes, err := json.MarshalIndent(endpoints, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (admin *Admin) Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/", admin.Endpoints)
	router.HandleFunc("/statistic", admin.StatisticHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         admin.Config.AdminConfig.ListenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	common.Logger.Infof("start admin server at %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("start admin server error, err=%s", err.Error()))
	}
}
