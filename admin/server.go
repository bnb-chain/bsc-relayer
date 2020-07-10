package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/binance-chain/bsc-relayer/common"
	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type Admin struct {
	Config *config.Config
	db     *gorm.DB
}

func NewAdmin(db *gorm.DB, config *config.Config) *Admin {
	return &Admin{
		db:     db,
		Config: config,
	}
}

func (admin *Admin) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{"status"},
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

func (admin *Admin) StatusHandler(w http.ResponseWriter, r *http.Request) {
	statistic := model.Statistic{}
	if admin.db != nil {
		admin.db.First(&statistic)
	}

	statisticResponse := struct {
		TotalTx          uint64 `json:"total_tx"`
		SuccessTx        uint64 `json:"success_tx"`
		FailedTx         uint64 `json:"failed_tx"`
		SyncHeaderTx     uint64 `json:"sync_header_tx"`
		DeliverPackageTx uint64 `json:"deliver_package_tx"`

		AccumulatedTotalTxFee   string `json:"accumulated_total_tx_fee"`
		AccumulatedSuccessTxFee string `json:"accumulated_success_tx_fee"`
		AccumulatedFailedTxFee  string `json:"accumulated_failed_tx_fee"`
		UpdateTime              string `json:"update_time"`
	}{
		TotalTx:                 statistic.TotalTx,
		SuccessTx:               statistic.SuccessTx,
		FailedTx:                statistic.FailedTx,
		SyncHeaderTx:            statistic.SyncHeaderTx,
		DeliverPackageTx:        statistic.DeliverPackageTx,
		AccumulatedTotalTxFee:   decimal.NewFromInt(int64(statistic.AccumulatedTotalTxFee)).Div(decimal.NewFromInt(1e18)).String() + ":BNB",
		AccumulatedSuccessTxFee: decimal.NewFromInt(int64(statistic.AccumulatedSuccessTxFee)).Div(decimal.NewFromInt(1e18)).String() + ":BNB",
		AccumulatedFailedTxFee:  decimal.NewFromInt(int64(statistic.AccumulatedFailedTxFee)).Div(decimal.NewFromInt(1e18)).String() + ":BNB",
		UpdateTime:              time.Unix(int64(statistic.UpdateTime), 0).Format("2006-01-02 03:04:05 PM"),
	}

	jsonBytes, err := json.MarshalIndent(statisticResponse, "", "    ")
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
	router.HandleFunc("/status", admin.StatusHandler)

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
