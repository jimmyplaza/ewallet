# Ethereum Hot Wallet

## Endpoint
-  nodeInfo
-  blockInfo
-  transationInfo
-  startMiner
-  stopMiner
-  sendTrans

## Environment
   Mac os
   Golang 1.6+

## Usage

``` 
1：Run Geth at RPC mode 

	   geth --rpc --rpcapi "eth,admin,personal,miner"
  
2：Run Go Binary 

	   cd ewallet/api
	   go build (will generate a binary called "api", or just use the exist binary "api")
	   ./api    ( Starting HTTP service on port 3000 ...)

3： API document can be found at:

    	  http://localhost:3000/apidoc
	
```
