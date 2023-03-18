package ethdemo

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

var client *Client
var account common.Address
var contractAddr common.Address

func init() {
	var err error

	// goerli: https://eth-goerli.g.alchemy.com/v2/dNfDfdwK13xMv9xVPcQ7GEaroWNpTch6
	// main: https://eth-mainnet.g.alchemy.com/v2/L1GdDpXQwe_eqs6QI4ewxwEdROCklTza
	conf := &Config{Addr: "https://eth-goerli.g.alchemy.com/v2/dNfDfdwK13xMv9xVPcQ7GEaroWNpTch6"}

	client, err = NewClient(conf)

	if err != nil {
		log.Fatal(err)
	}

	account = common.HexToAddress("0x5B7f33E9f0B24465cBD575d3cb354D286a9DF576")
	contractAddr = common.HexToAddress("0x0e22701968dcafc0a7bb8892e554d1b2ece11be7")
}

func TestClient_BalanceAt(t *testing.T) {
	t.Skip()
	balance, err := client.GetClient().BalanceAt(context.Background(), account, nil)

	if err != nil {
		t.Fatalf("TestClient_BalanceAt BalanceAt fail. | err: %s", err)
		return
	}

	t.Logf("balance: %s", balance.String())
}

func TestClient_StorageAt(t *testing.T) {
	/* 计算变量名称哈希值方法，没有测试过
		str := "myVariable"
	    hash := sha3.NewLegacyKeccak256()
	    hash.Write([]byte(str))
	    fmt.Printf("Keccak-256 hash of '%s': %x\n", str, hash.Sum(nil))
	*/

	t.Skip()

	// 连接以太坊网络
	clientNew, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/L1GdDpXQwe_eqs6QI4ewxwEdROCklTza")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum network. | err: %s", err)
		return
	}

	// 要获取状态变量的合约地址和变量名
	contractAddress := common.HexToAddress("0x7b07017e5c47df09b5f6ea1c12a6799b742ba9bc")
	variableName := common.HexToHash("0x5f16f6ed")

	// 获取最新的块号
	latestBlock, err := clientNew.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to get latest block. | err: %s", err)
		return
	}
	blockNumber := latestBlock.Number()

	// 使用 StorageAt 方法获取状态变量值
	value, err := clientNew.StorageAt(context.Background(), contractAddress, variableName, blockNumber)
	if err != nil {
		t.Fatalf("Failed to get storage value. | err: %s", err)
		return
	}

	// 将 value 转换为大整数类型
	valueInt := new(big.Int)
	valueInt.SetBytes(value)

	t.Logf("Storage value. | valuse: %s", valueInt.String())
}

func TestClient_CodeAt(t *testing.T) {
	t.Skip()
	code, err := client.GetClient().CodeAt(context.Background(), contractAddr, nil)

	if err != nil {
		t.Fatalf("TestClient_BalanceAt BalanceAt fail. | err: %s", err)
		return
	}

	t.Logf("code: %s", string(code))
}

func TestClient_BlockByNumber(t *testing.T) {
	// 选择您要查找的块号
	blockNumber := big.NewInt(8663586)

	// 从节点中检索块
	block, err := client.GetClient().BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		t.Logf("Failed to retrieve block. | err: %s", err)
		return
	}

	// goerli测试网络哈希值计算会有问题
	blockHash := block.Hash().Hex()
	// 打印块哈希
	t.Logf("blockHash: %s", blockHash)
}

func TestClient_BlockByHash(t *testing.T) {
	// 选择您要查找的块号
	blockHash := common.HexToHash("0xcd4dd4435c881ea0ceed43770078cf3adf644d0d9c05c0e7e17521a0b469b9df")

	// 从节点中检索块
	block, err := client.GetClient().BlockByHash(context.Background(), blockHash)
	if err != nil {
		t.Logf("Failed to retrieve block. | err: %s", err)
		return
	}

	blockNumber := block.Number()
	// 打印块哈希
	t.Logf("blockNumber: %s", blockNumber.String())
}

func TestClient_HeaderByNumber(t *testing.T) {
	// 获取最新的块
	block, err := client.GetClient().BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Errorf("Failed to get the latest block: %v", err)
		return
	}

	// 计算块的哈希值
	blockHash := block.Hash().Hex()
	num := block.Number()

	t.Logf("Block hash: %s | number: %s", blockHash, num.String())
}
