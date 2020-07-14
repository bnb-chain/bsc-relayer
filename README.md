# Relayer service from BBC to BSC

This is an canonical implementation of [bsc-relayer](https://github.com/binance-chain/whitepaper/blob/master/WHITEPAPER.md#bsc-relayers) service to relay cross chain packages from Binance Chain to Binance Smart Chain. It can also monitor double sign behavior on BSC and submit evidence to Binance Chain. Community members are encouraged to implement more implementations according to their own requirements.

## Disclaimer

**The software and related documentation are under active development, 
all subject to potential future change without notification and not ready for production use. 
The code and security audit have not been fully completed and not ready for any bug bounty.
We advise you to be careful and experiment on the network at your own risk. Stay safe out there.**

## Quick Start

**Note**: Requires [Go 1.13+](https://golang.org/dl/)

### Setup config

1. Edit `config/config.json` and input your private key to `bsc_config.private_key`
2. Transfer enough BNB to this account:
    1. 100:BNB for relayer register
    2. More than 10:BNB for transaction fee.
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
   
4. The supported db platforms include `mysql` and `sqlite3`. This is an example db config:
    ```json
    {
       "dialect": "mysql",
       "db_path": "relayer:12345678@(localhost:3306)/bsc_relayer?charset=utf8&parseTime=True&loc=Local"
    }
    ```
   If you don't want to specify a db for your bsc-relayer, just leave `db_path` to empty.
   
5. Send alert telegram message when the balance of relayer account is too low. This is an example alert config:
    ```json
    {
        "enable_alert": true,
        "telegram_bot_id": "1377262449:AAH70B43ES75uiKyyiyUh3fRCy6JrvK4O6c",
        "telegram_chat_id": "@bsc_relayer",
        "balance_threshold": "1000000000000000000"
    }
    ```
   Please refer to [telegram_bot](https://www.home-assistant.io/integrations/telegram_bot) to setup your telegram bot. If you don't want this feature, just set `enable_alert` to false.

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

1. Please fill proper bsc private key to `bsc_config.private_key`, example private key: `5576779EB3F28F1067BE07AC643A81A8C74E6C55EDE38BEEF68BE1E9D4C1CA3C`

    ```
    panic: privateKey of Binance Smart Chain should not be empty
    
    goroutine 1 [running]:
    ```

2. Please fill your mnemonic to `bbc_config.mnemonic`, example mnemonic: `witness pitch peasant bird year sponsor conduct push enhance melt betray spare police region strategy hammer potato lecture cloud business habit student vehicle allow`

    ```
    ERROR main missing local mnemonic
    ```

### Monitor Relayer Status

To enable this function, you must specify proper db config for your relayer. Suppose `8080` is the admin port: 
```shell script
curl localhost:8080/status
```

Example response:
```json
{
    "total_tx": 82,
    "success_tx": 36,
    "failed_tx": 46,
    "sync_header_tx": 41,
    "deliver_package_tx": 41,
    "accumulated_total_tx_fee": "0.14336928:BNB",
    "accumulated_success_tx_fee": "0.1030536:BNB",
    "accumulated_failed_tx_fee": "0.04031568:BNB",
    "update_time": "2020-07-14 06:59:39 PM"
}
```

## License

The library is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the [LICENSE](LICENSE) file.