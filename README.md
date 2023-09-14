# Guide


1. 建立module檔案
```go
// go mod init <module>
go mod init module
```

be like:
```go
// ./go.mod底下會新增這個檔案
module module

go 1.21.0

```

2. Import source githubc 上面的go-ethereum/ethclient
```go
go get github.com/ethereum/go-ethereum/ethclient
```

3. visit go.mod

檢查是否有多出這些套件
```go
require (
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/bits-and-blooms/bitset v1.5.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.10.0 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.3.0 // indirect
	github.com/deckarep/golang-set/v2 v2.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/c-kzg-4844 v0.3.1 // indirect
	github.com/ethereum/go-ethereum v1.13.0 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/holiman/uint256 v1.2.3 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/supranational/blst v0.3.11 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/exp v0.0.0-20230810033253-352e893a4cad // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
```

4. Write main.go

```golang
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//ethereum env adress
	client, err := ethclient.Dial("http://127.0.0.1:7457")
	fmt.Println(client)

	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
		panic(err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}
}
```

5. Run main.go
```go
go run main.go

// or
if you use vscode, you can press F5(Debug:)

if you have error, you should make a launch.json file

it will talk to VScode what should do, if user press F5

if success:

  ~/smart-contract-with-golang # go run main.go                                        
client : &{0x140001d6090}
Success! you are connected to the Ethereum Network

```

* launch.json
```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${file}"
        }
    ]
}
```





