package ethdemo

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"gopractice/projects/ethdemo/nbc"
	"log"
	"math/big"
	"os"
	"path/filepath"
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
	contractAddr = common.HexToAddress("0x0e22701968Dcafc0a7bb8892E554D1b2eCE11Be7")
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
	t.Skip()
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
	t.Skip()
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
	t.Skip()
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

func TestClient_TransactionReceipt(t *testing.T) {
	t.Skip()
	txHash := common.HexToHash("0xb774bba726b0404f2c19e5b13b26a9f09fd15f33e0b976ae808886f1e7872fe9")

	receipt, err := client.GetClient().TransactionReceipt(context.Background(), txHash)

	if err != nil {
		t.Fatalf("TransactionReceipt fail. | err: %s", err)
		return
	}

	t.Logf("receipt: %+v", receipt)
}

func TestClient_TransactionInBlock(t *testing.T) {
	t.Skip()
	blockHash := common.HexToHash("0x3ea16d8aee483e24966b9236317409ca704104b881f5825f77cff1c8252ad5e3")
	txIndex := uint(33)

	tx, err := client.GetClient().TransactionInBlock(context.Background(), blockHash, txIndex)

	if err != nil {
		t.Fatalf("TransactionInBlock fail. | err: %s", err)
		return
	}

	tx.GasPrice()

	t.Logf("transaction. | value: %s | nonce: %d | gas: %d | gasFee: %s | gasPrice: %s", tx.Value(), tx.Nonce(), tx.Gas(), tx.GasFeeCap(), tx.GasPrice())
}

func TestClient_NonceAt(t *testing.T) {
	t.Skip()
	nonce, err := client.GetClient().NonceAt(context.Background(), account, nil)

	if err != nil {
		t.Fatalf("TransactionInBlock fail. | err: %s", err)
		return
	}

	t.Logf("nonce: %d", nonce)
}

func TestClient_PendingTransactionCount(t *testing.T) {
	t.Skip()
	pendingCount, err := client.GetClient().PendingTransactionCount(context.Background())

	if err != nil {
		t.Fatalf("PendingTransactionCount fail. | err: %s", err)
		return
	}

	t.Logf("nonce: %d", pendingCount)
}

// 转账
func TestClient_SendTransaction(t *testing.T) {
	// 交易发送方
	// 获取私钥方式一，通过keystore文件
	fromKeystore, err := os.ReadFile("./private_key.keystore")
	require.NoError(t, err)
	fromKey, err := keystore.DecryptKey(fromKeystore, "123456")
	privateKey := fromKey.PrivateKey
	fromPubkey := privateKey.PublicKey
	fromAddr := crypto.PubkeyToAddress(fromPubkey)

	// 获取私钥方式二，通过私钥字符串
	//privateKey, err := crypto.HexToECDSA("")

	// 交易接收方
	toAddr := common.HexToAddress("0xEca1596D49a2325e1b80126d1F93756705A3b9ce")

	// 数量
	amount := big.NewInt(100000000000000000)

	// gasLimit
	var gasLimit uint64 = 300000

	gasPrice, err := client.GetClient().SuggestGasPrice(context.Background())

	if err != nil {
		t.Fatalf("SuggestGasPrice fail. | err: %s", err)
		return
	}

	// gasPrice
	//gasPrice = big.NewInt(250000000000)

	// nonce获取
	//nonce := uint64(10)
	nonce, err := client.GetClient().PendingNonceAt(context.Background(), fromAddr)
	//nonce, err := client.GetClient().TransactionCount(context.Background(), account)

	// 认证信息组装
	//auth := bind.NewKeyedTransactor(fromPrivkey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))

	if err != nil {
		t.Fatalf("NewKeyedTransactorWithChainID fail. | err: %s", err)
		return
	}
	//auth,err := bind.NewTransactor(strings.NewReader(mykey),"111")
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = amount // in wei
	//auth.Value = big.NewInt(100000)     // in wei
	auth.GasLimit = gasLimit // in units
	//auth.GasLimit = uint64(0) // in units
	auth.GasPrice = gasPrice
	auth.From = fromAddr

	// 交易创建
	//tx := types.NewTransaction(nonce,toAddr,amount,gasLimit,gasPrice,[]byte{})
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddr,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     []byte{},
	})

	// 交易签名
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	require.NoError(t, err)

	// 交易发送
	err = client.GetClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatalf("SendTransaction fail. err: %s", err)
	}

	// 等待挖矿完成
	receipt, err := bind.WaitMined(context.Background(), client.GetClient(), signedTx)

	if err != nil {
		t.Fatalf("WaitMined fail. | err: %s", err)
		return
	}

	t.Logf("receipt: %+v", receipt)
}

func TestClient_PendingNonceAt(t *testing.T) {
	t.Skip()
	pendingNonce, err := client.GetClient().PendingNonceAt(context.Background(), account)

	if err != nil {
		t.Fatalf("PendingTransactionCount fail. | err: %s", err)
		return
	}

	t.Logf("nonce: %d", pendingNonce)
}

func TestClient_StoreKey(t *testing.T) {
	t.Skip()
	keyPass := &KeyStorePassphrase{
		scryptN: keystore.StandardScryptN,
		scryptP: keystore.LightScryptP,
	}

	privateKey, _ := crypto.HexToECDSA("your private key")

	err := keyPass.StoreKey("./xiaolong.keystore", &keystore.Key{PrivateKey: privateKey}, "123456")

	if err != nil {
		t.Fatalf("StoreKey fail. | err: %s", err)
		return
	}
}

type KeyStorePassphrase struct {
	scryptN int
	scryptP int
}

func (ks KeyStorePassphrase) StoreKey(filename string, key *keystore.Key, auth string) error {
	keyjson, err := keystore.EncryptKey(key, auth, ks.scryptN, ks.scryptP)
	if err != nil {
		return err
	}
	// Write into temporary file
	tmpName, err := WriteTemporaryKeyFile(filename, keyjson)
	if err != nil {
		return err
	}

	return os.Rename(tmpName, filename)
}

func WriteTemporaryKeyFile(file string, content []byte) (string, error) {
	// Create the keystore directory with appropriate permissions
	// in case it is not present yet.
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return "", err
	}
	// Atomic write: create a temporary hidden file first
	// then move it into place. TempFile assigns mode 0600.
	f, err := os.CreateTemp(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return "", err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", err
	}
	f.Close()
	return f.Name(), nil
}

func TestClient_ContractCall(t *testing.T) {
	nbcContract, err := nbc.NewNbcToken(contractAddr, client.GetClient())

	if err != nil {
		t.Fatalf("NewNbcToken fail. | err: %s", err)
		return
	}

	nftCount, err := nbcContract.BalanceOf(nil, account)

	if err != nil {
		t.Fatalf("BalanceOf fail. | err: %s", err)
		return
	}

	t.Logf("owner: %s have %s nft", account.String(), nftCount.String())
}

func TestClient_SafeMint(t *testing.T) {
	nbcContract, err := nbc.NewNbcToken(contractAddr, client.GetClient())

	if err != nil {
		t.Fatalf("NewNbcToken fail. | err: %s", err)
		return
	}

	// 交易发送方
	// 获取私钥方式一，通过keystore文件
	fromKeystore, err := os.ReadFile("./private.keystore")
	require.NoError(t, err)
	fromKey, err := keystore.DecryptKey(fromKeystore, "123456")
	privateKey := fromKey.PrivateKey
	fromPubkey := privateKey.PublicKey
	fromAddr := crypto.PubkeyToAddress(fromPubkey)

	t.Logf("account: %s", fromAddr.String())

	// gasLimit
	var gasLimit uint64 = 300000

	gasPrice, err := client.GetClient().SuggestGasPrice(context.Background())

	if err != nil {
		t.Fatalf("SuggestGasPrice fail. | err: %s", err)
		return
	}

	// nonce获取
	//nonce := uint64(10)
	nonce, err := client.GetClient().PendingNonceAt(context.Background(), fromAddr)

	// 认证信息组装
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))

	if err != nil {
		t.Fatalf("NewKeyedTransactorWithChainID fail. | err: %s", err)
		return
	}
	//auth,err := bind.NewTransactor(strings.NewReader(mykey),"111")
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.Value = big.NewInt(100000)     // in wei
	auth.GasLimit = gasLimit // in units
	//auth.GasLimit = uint64(0) // in units
	auth.GasPrice = gasPrice
	auth.From = fromAddr

	opts := &bind.TransactOpts{
		From:      auth.From,
		Nonce:     auth.Nonce,
		Signer:    auth.Signer,
		Value:     nil,
		GasPrice:  auth.GasPrice,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  auth.GasLimit,
		Context:   nil,
		NoSend:    false,
	}

	tx, err := nbcContract.SafeMint(opts, account, "QmVJ2bVQYukBFLhKixZ64LPPF5fbw1jYj4anupPKmfci9e")
	if err != nil {
		t.Fatalf("SafeMint fail. err: %s", err)
	}

	// 等待挖矿完成
	receipt, err := bind.WaitMined(context.Background(), client.GetClient(), tx)

	if err != nil {
		t.Fatalf("WaitMined fail. | err: %s", err)
		return
	}

	t.Logf("receipt: %+v", receipt)
}
