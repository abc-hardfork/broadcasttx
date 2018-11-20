# Summary

This Project is used for bitcoin abc hardfork.We support the following features
:

```
queryBsvUtxos
queryBchUtxos
diffUtxos
broadBchRawTransaction
broadBsvRawTransaction
```
# HOWTO

##### ElectrumX 
For ABC:See [Electron-Cash/electrumx](https://github.com/Electron-Cash/electrumx).

For BSV:See [kyuupichan/electrumx](https://github.com/kyuupichan/electrumx).
##### golang
1.10+

# APIS

### 1) /service/broadcast/abc

broadcast bch transaction.

##### Parameters
`rawtx`： raw transaction to broadcast

##### Result

`string`：the hex-encoded

##### Examples

```
00000044
```

---
    
### 2) /service/broadcast/sv

broadcast bsv transaction.

##### Parameters

`rawtx`： raw transaction to broadcast

##### Result

`string`：the hex-encoded

##### Examples

```
00000044
```
---
    
### 3) /utxo/query

broadcast bsv transaction.
##### Parameters
`address`： string

##### Result

`txid`: string pre transaction tx_hash

`vout`: string pre transaction tx_out index

`scriptPubKey`: string script pubkey

`value`: string pre transaction tx_out index

`vout`: string pre transaction amount,unit:satoshi

`blockHeight`: string pre transaction block height

##### Examples

```
{
    "code": 0,
    "message": "",
    "result": [
        {
        "txid": "e01b0d126dbdb4e847fec8056……ea49ad77164aca667b6cfe",//Pre Transaction tx_hash
        "vout": 0,//Pre Transaction tx_out index
        "scriptPubKey": "76a914ec80f6fe2d37……840d7fc34ea9ea88ac",//script pubkey
        "value": 3798054,//Pre Transaction amount,unit:satoshi
        "blockHeight": 556609//Pre Transaction block height
        }
      ]
    }
```
---
### 4) /utxo/bch/query

broadcast bch transaction.
##### Parameters
`address`： string

##### Result

`txid`: string pre transaction tx_hash

`vout`: string pre transaction tx_out index

`scriptPubKey`: string script pubkey

`value`: string pre transaction tx_out index

`vout`: string pre transaction amount,unit:satoshi

`blockHeight`: string pre transaction block height

##### Examples

```
{
    "code": 0,
    "message": "",
    "result": [
        {
        "txid": "e01b0d126dbdb4e847fec8056……ea49ad77164aca667b6cfe",//Pre Transaction tx_hash
        "vout": 0,//Pre Transaction tx_out index
        "scriptPubKey": "76a914ec80f6fe2d37……840d7fc34ea9ea88ac",//script pubkey
        "value": 3798054,//Pre Transaction amount,unit:satoshi
        "blockHeight": 556609//Pre Transaction block height
        }
      ]
    }
```
---
### 5) /utxo/diff

support query bch and bsv utxos,find same utxo...as DupUtxo filed

##### Parameters
`address`： string

##### Result

`txid`: string pre transaction tx_hash

`vout`: string pre transaction tx_out index

`scriptPubKey`: string script pubkey

`value`: string pre transaction tx_out index

`vout`: string pre transaction amount,unit:satoshi

`blockHeight`: string pre transaction block height

##### Examples

```
{
    "code": 0,
    "message": "",
    "result": {
      "BsvUtxo":[
        {
        "txid": "e01b0d126dbdb4e847fec8056……ea49ad77164aca667b6cfe",//Pre Transaction tx_hash
        "vout": 0,//Pre Transaction tx_out index
        "scriptPubKey": "76a914ec80f6fe2d37……840d7fc34ea9ea88ac",//script pubkey
        "value": 3798054,//Pre Transaction amount,unit:satoshi
        "blockHeight": 556609//Pre Transaction block height
        }
      ]},
      "BchUtxo":[
        {
        "txid": "e01b0d126dbdb4e847fec8056……ea49ad77164aca667b6cfe",//Pre Transaction tx_hash
        "vout": 0,//Pre Transaction tx_out index
        "scriptPubKey": "76a914ec80f6fe2d37……840d7fc34ea9ea88ac",//script pubkey
        "value": 3798054,//Pre Transaction amount,unit:satoshi
        "blockHeight": 556609//Pre Transaction block height
        }
      ]},
      "DupUtxo":[
        {
        "txid": "e01b0d126dbdb4e847fec8056……ea49ad77164aca667b6cfe",//Pre Transaction tx_hash
        "vout": 0,//Pre Transaction tx_out index
        "scriptPubKey": "76a914ec80f6fe2d37……840d7fc34ea9ea88ac",//script pubkey
        "value": 3798054,//Pre Transaction amount,unit:satoshi
        "blockHeight": 556609//Pre Transaction block height
        }
      ]}
    }
```
---