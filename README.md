# harfork

This Project is used for bitcoin abc hardfork.We support list apis:
    
### 1、/service/broadcast/abc
broadcast bch transaction

### 2、/service/broadcast/sv
broadcast bsv transaction

### 3、/utxo/query
query bsv utxos
  
    #parameter
    address string
    
    #result
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

### 4、/utxo/bch/query
query bch utxos
  
    #parameter
    address string
    
    #result
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
    
### 5、/utxo/diff
support query bch and bsv utxos,find same utxo...as DupUtxo filed
  
    #parameter
    address string
    
    #result
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
    
