package relayer

import (
	"bytes"
	"context"
	"encoding/hex"
	"sync"

	"github.com/binance-chain/bsc-double-sign-sdk/client"
	"github.com/binance-chain/bsc-double-sign-sdk/types/bsc"
	"github.com/ethereum/go-ethereum"
	ethereumtype "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/clist"

	"github.com/binance-chain/bsc-relayer/common"
	"github.com/binance-chain/bsc-relayer/executor"
)

const (
	maxCacheHeader = 100
)

type DoubleSignMonitor struct {
	mux                   sync.Mutex
	ethereumHeaderNumbers []*clist.CList
	ethereumHeaders       []map[int64]*ethereumtype.Header
	bbcExecutor           *executor.BBCExecutor
}

func isDoubleSignHeaders(headers [2]*bsc.Header) bool {
	if headers[0] == nil || headers[1] == nil {
		return false
	}
	if headers[0].Number != (headers[1].Number) {
		return false
	}
	if bytes.Equal(headers[0].ParentHash[:], headers[1].ParentHash[:]) {
		return false
	}
	signature1, err := headers[0].GetSignature()
	if err != nil {
		return false
	}
	signature2, err := headers[1].GetSignature()
	if err != nil {
		return false
	}
	if bytes.Equal(signature1, signature2) {
		return false
	}

	signer1, err := headers[0].ExtractSignerFromHeader()
	if err != nil {
		return false
	}
	signer2, err := headers[1].ExtractSignerFromHeader()
	if err != nil {
		return false
	}
	if !bytes.Equal(signer1.Bytes(), signer2.Bytes()) {
		return false
	}

	return true
}

func (monitor *DoubleSignMonitor) doubleSignChecker(channelNumber int, header *ethereumtype.Header) {
	monitor.mux.Lock()
	defer monitor.mux.Unlock()

	if header == nil {
		return
	}

	for _, headerMap := range monitor.ethereumHeaders {
		previousHeader, ok := headerMap[header.Number.Int64()]
		if !ok {
			continue
		}

		var twoHeader [2]*bsc.Header
		twoHeader[0] = client.EthHeaderToBscHeader(header)
		twoHeader[1] = client.EthHeaderToBscHeader(previousHeader)

		ok = isDoubleSignHeaders(twoHeader)
		if !ok {
			continue
		}

		signature0, _ := twoHeader[0].GetSignature()
		signature1, _ := twoHeader[1].GetSignature()
		common.Logger.Infof("found double sign evidence: height %d, first signature %s, second signature %s",
			header.Number.Int64(), hex.EncodeToString(signature0), hex.EncodeToString(signature1))

		tx, err := monitor.bbcExecutor.SubmitEvidence(twoHeader[:])
		if err != nil {
			common.Logger.Errorf("submit evidence err: %s", err.Error())
			continue
		}
		common.Logger.Infof("submit evidence success, tx result: %v", tx)
	}

	// cleanup old headers
	if monitor.ethereumHeaderNumbers[channelNumber].Len() >= maxCacheHeader {
		elem := monitor.ethereumHeaderNumbers[channelNumber].Front()
		monitor.ethereumHeaderNumbers[channelNumber].Remove(elem)
		elem.DetachPrev()
		removedHeaderNumber, ok := elem.Value.(int64)
		if !ok {
			common.Logger.Error("found unexpect element %v", elem.Value)
		}
		delete(monitor.ethereumHeaders[channelNumber], removedHeaderNumber)
	}
	monitor.ethereumHeaderNumbers[channelNumber].PushBack(header.Number.Int64())
	monitor.ethereumHeaders[channelNumber][header.Number.Int64()] = header
}

func (r *Relayer)doubleSignMonitorDaemon() {
	var bscProviderList = r.cfg.BSCConfig.MonitorDataSeedList
	doubleSignMonitor := DoubleSignMonitor{
		ethereumHeaders:       make([]map[int64]*ethereumtype.Header, len(bscProviderList)),
		ethereumHeaderNumbers: make([]*clist.CList, len(bscProviderList)),
		bbcExecutor:           r.bbcExecutor,
	}

	var headerChannelList []chan *ethereumtype.Header
	var subscriptionList []ethereum.Subscription
	for index, provider := range bscProviderList {
		headerChannel := make(chan *ethereumtype.Header, maxCacheHeader)
		bscClient, err := ethclient.Dial(provider)
		if err != nil {
			panic(err)
		}
		subscription, err := bscClient.SubscribeNewHead(context.Background(), headerChannel)
		if err != nil {
			panic(err)
		}
		subscriptionList = append(subscriptionList, subscription)
		headerChannelList = append(headerChannelList, headerChannel)

		doubleSignMonitor.ethereumHeaderNumbers[index] = clist.New()
		doubleSignMonitor.ethereumHeaders[index] = make(map[int64]*ethereumtype.Header, maxCacheHeader)
	}
	defer func() {
		for _, subscription := range subscriptionList {
			subscription.Unsubscribe()
		}
	}()

	go func() {
		for index, headerChannel := range headerChannelList {
			go func(channelNumber int, channel chan *ethereumtype.Header) {
				for header := range channel {
					doubleSignMonitor.doubleSignChecker(channelNumber, header)
				}
			}(index, headerChannel)
		}
	}()
	select {}
}
