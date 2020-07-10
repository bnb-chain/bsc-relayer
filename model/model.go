package model

import (
	"github.com/jinzhu/gorm"
)

const (
	SyncBlockHeader = "SyncHeader"
	DeliverPackage  = "DeliverPackage"

	Created = "Created"
	Success = "Success"
	Failure = "Failed"
)

type RelayTransaction struct {
	Id int64

	TxHash string
	Type   string

	ChannelId uint8
	Sequence  uint64
	BCHeight  uint64

	TxStatus   string
	TxGasPrice uint64
	TxGasLimit uint64
	TxUsedGas  uint64
	TxFee      uint64
	TxHeight   uint64

	CreateTime int64
	UpdateTime int64
}

func (RelayTransaction) TableName() string {
	return "relay_transaction"
}

func InitTables(db *gorm.DB) {
	if !db.HasTable(&RelayTransaction{}) {
		db.CreateTable(&RelayTransaction{})
		db.Model(&RelayTransaction{}).AddIndex("idx_tx_height", "tx_height")
		db.Model(&RelayTransaction{}).AddIndex("idx_tx_status", "tx_status")
		db.Model(&RelayTransaction{}).AddIndex("idx_tx_create_time", "create_time")
	}
}
