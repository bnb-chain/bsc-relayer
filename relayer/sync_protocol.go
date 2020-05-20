package relayer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/binance-chain/bsc-relayer/common"
	relayerconfig "github.com/binance-chain/bsc-relayer/config"
	ctypes "github.com/binance-chain/go-sdk/common/types"
)

var crossChainProtocol = newCrossChainProtocol()

func newCrossChainProtocol() CrossChainProtocol {
	return CrossChainProtocol{
		Channels: make(map[common.CrossChainChannelID]relayerconfig.ChannelConfig),
		ABIs:     make(map[string]string),
	}
}

type CrossChainProtocol struct {
	Channels map[common.CrossChainChannelID]relayerconfig.ChannelConfig
	ABIs     map[string]string
}

func SyncProtocol(cfg *relayerconfig.Config, network ctypes.ChainNetwork) error {
	if cfg.CrossChainConfig.ProtocolConfigType == relayerconfig.LocalProtocolConfig {
		for index, abiName := range cfg.CrossChainConfig.ProtocolConfig.ABIFileNameList {
			crossChainProtocol.ABIs[abiName] = cfg.CrossChainConfig.ProtocolConfig.ABIList[index]
		}
		for _, channel := range  cfg.CrossChainConfig.ProtocolConfig.Channels {
			channelID := channel.ChannelID
			crossChainProtocol.Channels[common.CrossChainChannelID(channelID)] = channel
		}
		return nil
	}

	common.Logger.Infof("Sync cross chain protocol from %s", relayerconfig.ProtocolRepository)
	var baseUrl string
	if network == ctypes.ProdNetwork {
		baseUrl = relayerconfig.MainnetChannelConfigUrl
	} else {
		baseUrl = relayerconfig.TestnetChannelConfigUrl
	}

	settingsUrl := baseUrl + "/" + relayerconfig.SettingFileName
	message, err := GetFile(settingsUrl)
	if err != nil {
		return fmt.Errorf("failed to sync cross chain protocol from github: %s",err.Error())
	}
	var protocolConfig relayerconfig.CrossChainProtocolConfig
	err = json.Unmarshal(message, &protocolConfig)
	if err != nil {
		return fmt.Errorf("failed to decode cross chain protocol: %s",err.Error())
	}

	for index, abiName := range protocolConfig.ABIFileNameList {
		crossChainProtocol.ABIs[abiName] = protocolConfig.ABIList[index]
	}

	for _, channel := range protocolConfig.Channels {
		channelID := channel.ChannelID
		crossChainProtocol.Channels[common.CrossChainChannelID(channelID)] = channel
	}

	return nil
}

func GetFile(url string) ([]byte, error) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}

	response, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	message, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}
	return message, nil
}
