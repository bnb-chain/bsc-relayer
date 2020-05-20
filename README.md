# Relayer service from BBC to BSC

This is an canonical implementation of [bsc-relayer](https://github.com/binance-chain/whitepaper/blob/master/WHITEPAPER.md#bsc-relayers) service to relay cross chain packages from Binance Chain to Binance Smart Chain. It can also monitor double sign behavior on BSC and submit evidence to Binance Chain. Community members are encouraged to implement more implementations according to their own requirements.

## Disclaimer

**The software and related documentation are under active development, 
all subjected to potential future change without notification and not ready for production use. 
The code and security audit have not been fully completed and not ready for any bug bounty.
We advise you to be careful and experiment on the network at your own risk. Stay safe out there.**

## Quick Start

**Note**: Requires [Go 1.13+](https://golang.org/dl/)

### Setup config

1. Edit `config/config.json` and input your private key to `bsc_config.private_key`
2. Transfer enough BNB to this account:
    1. 100:BNB for relayer register
    2. more than 10:BNB for transaction fee.
3. If you want to monitor double sign behavior on BSC, you need to fill more than 1 endpoints in `bsc_config.monitor_data_seed_list`, example `monitor_data_seed_list` for testnet:
    ```json
    {
      "monitor_data_seed_list": [
          "wss://data-seed-prebsc-1-s1.binance.org:8545",
          "wss://data-seed-prebsc-2-s1.binance.org:8545",
          "wss://data-seed-prebsc-1-s2.binance.org:8545",
          "wss://data-seed-prebsc-2-s2.binance.org:8545",
          "wss://data-seed-prebsc-1-s3.binance.org:8545",
          "wss://data-seed-prebsc-2-s3.binance.org:8545"
        ]
    }
    ```
   Besides, you also need to provide proper Binance Chain `mnemonic`.

### Build

#### Build Binary:
```shell script
make build
```

#### Build Docker Image

- Please complete configuration setup in `config/config.json` first, mainly including rpc endpoints, BSC private key and BC mnemonic(if you want to submit double sign evidence)
- Docker build command:
```shell script
docker build -t bsc-relayer:latest .
```

### Run

Run locally:
```shell script
./build/bsc-relayer --config-type local --config-path config/config.json
```

Run docker:
```shell script
docker run -e BBC_NETWORK=1 -e CONFIG_TYPE="local" -d -it bsc-relayer
```

### TroubleShooting

1. Just start bsc-relayer again.

    ```
    panic: failed to sync cross chain protocol from github: Get https://raw.githubusercontent.com/binance-chain/bsc-relayer-config/master/bsc/testnet/protocol.json: EOF
    
    goroutine 1 [running]:
    ```

2. The bsc private key might be empty

    ```
    panic: privateKey of Binance Smart Chain should not be empty
    
    goroutine 1 [running]:
    ```

3. The Binance Chain mnemonic is empty

    ```
    ERROR main missing local mnemonic
    ```

## License

The library is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the [LICENSE](LICENSE) file.