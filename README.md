# Relayer service from BBC to BSC

This is an canonical implementation of [bsc-relayer](https://github.com/binance-chain/whitepaper/blob/master/WHITEPAPER.md#bsc-relayers) service to relay cross chain packages from Binance Chain to Binance Smart Chain. It can also monitor double sign behavior on BSC and submit evidence to Binance Chain. Community members are encouraged to implement more implementations according to their own requirements.

## Quick Start

**Note**: Requires [Go 1.19+](https://golang.org/dl/)

### Setup config

1. Edit `config/config.json` and input your private key to `bsc_config.private_key`
2. Transfer enough BNB to this account - more than 10:BNB for transaction fee.
3. Setup relayer mode:
    ```json
    {
      
      "competition_mode": true
      
    }
    ```
   If `competition_mode` is true, `bsc-relayer` will monitor cross chain packages in every block and try to deliver the packages immediately. Otherwise, `bsc-relayer` will check if there are undelivered packages in every `clean_up_block_interval` blocks and batch deliver these packages. Competition mode can accelerate package delivery but will cost more transaction fee. In contrary, in non-competition mode, package delivery will be slower but less transaction fee will be cost. 
   
4. If you want to monitor double sign behavior on BSC, you need to fill more than 1 endpoints in `bsc_config.monitor_data_seed_list`.

    3.1 Testnet `monitor_data_seed_list`:
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
   
   3.2 Mainnet: `monitor_data_seed_list`:
   ```json
   {
     "monitor_data_seed_list": [
           "wss://bsc-dataseed1.binance.org:443",
           "wss://bsc-dataseed2.binance.org:443",
           "wss://bsc-dataseed3.binance.org:443",
           "wss://bsc-dataseed4.binance.org:443",
           "wss://bsc-dataseed1.defibit.io:443",
           "wss://bsc-dataseed2.defibit.io:443",
           "wss://bsc-dataseed3.defibit.io:443",
           "wss://bsc-dataseed4.defibit.io:443",
           "wss://bsc-dataseed1.ninicoin.io:443",
           "wss://bsc-dataseed2.ninicoin.io:443",
           "wss://bsc-dataseed3.ninicoin.io:443",
           "wss://bsc-dataseed4.ninicoin.io:443"
       ]
   }
   ```
   Besides, you also need to provide proper Binance Chain `mnemonic`.
   
5. The supported db platforms include `mysql` and `sqlite3`. This is an example db config:
    ```json
    {
       "dialect": "mysql",
       "db_path": "relayer:12345678@(localhost:3306)/bsc_relayer?charset=utf8&parseTime=True&loc=Local"
    }
    ```
   If you don't want to specify a db for your bsc-relayer, just leave `db_path` to empty.
   
6. Send alert telegram message when the balance of relayer account is too low. This is an example alert config:
    ```json
    {
        "enable_alert": true,
        "enable_heart_beat": false,
        "interval": 300,
        "telegram_bot_id": "your_bot_id",
        "telegram_chat_id": "your_chat_id",
        "balance_threshold": "1000000000000000000",
        "sequence_gap_threshold": 10
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

Run locally for testnet:
```shell script
./build/bsc-relayer --bbc-network-type 0 --config-type local --config-path config/config.json
```

Run locally for mainnet:
```shell script
./build/bsc-relayer --bbc-network-type 1 --config-type local --config-path config/config.json
```

Run docker for testnet:
```shell script
docker run -e BBC_NETWORK=0 -e CONFIG_TYPE="local" -d -it bsc-relayer
```

Run docker for mainnet:
```shell script
docker run -e BBC_NETWORK=1 -e CONFIG_TYPE="local" -d -it bsc-relayer
```

### TroubleShooting

1. Please fill proper bsc private key to `bsc_config.private_key`, example private key: `your_private_key`

    ```
    panic: privateKey of Binance Smart Chain should not be empty
    
    goroutine 1 [running]:
    ```

2. Please fill your mnemonic to `bbc_config.mnemonic`, example mnemonic: `your_mnemonic`

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