package geth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"regexp"
)

var client *ethclient.Client

func init() {
	//client, _ = ethclient.Dial("https://cloudflare-eth.com")
	client, _ = ethclient.Dial("http://127.0.0.1:8545")
}
func GethDemo() {
	log.Println("geth GethDemo...")
	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}

func GethDemo1() {
	log.Println("geth GethDemo1...")
	address := common.HexToAddress("0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199")
	fmt.Println(address.Hex())
}

func GethDemo2() {
	log.Println("geth GethDemo2...")
	account := common.HexToAddress("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	blockNumber := big.NewInt(0)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}

func GethDemo3() {
	log.Println("geth GethDemo3...")
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0x725b29165ed0837f67d03cf9dea4e913a9931527ce6b00008866388966f7d5bd

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x7c934185d1cc7e898aca8ca0350d6771245186d676ddb748c1d33f5a2b71b94f01f940d34dd8c5f677a45854c7e2ae4df7be9027f4f1ab91ca27513fd2b735b8

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x4C05c5844fb13bE6b698A6A6509F02E16267d6Af

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x4c05c5844fb13be6b698a6a6509f02e16267d6af
}

/*
公钥和私钥转换
*/
func GethDemo4() {
	log.Println("geth GethDemo4...")
	// 生成随机私钥
	privatekey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(&privatekey)

	// 转换成字节
	privateKeyBytes := crypto.FromECDSA(privatekey)
	log.Println(privateKeyBytes)

	// 转换成十六进制，
	encode := hexutil.Encode(privateKeyBytes)
	log.Println(encode)
	// 且去掉0x
	log.Println(encode[2:])

	// 公钥是私钥派生，因此私钥可以生成公钥
	public := privatekey.Public()
	log.Println("public: ", public)

	// crypto.PublicKey 类型转换为 *ecdsa.PublicKey
	publicKeyECDSA, ok := public.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: public key is not of type *ecdsa.PublicKey")
	}

	log.Println(publicKeyECDSA)

	// 将*ecdsa.PublicKey转换成字节
	pub := crypto.FromECDSAPub(publicKeyECDSA)
	log.Println(hexutil.Encode(pub))

	// 通过公钥生成公共地址
	// 公共地址其实就是公钥的 Keccak-256 哈希
	/**
	func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
		pubBytes := FromECDSAPub(&p)
		return common.BytesToAddress(Keccak256(pubBytes[1:])[12:])
	}
	*/
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println("公共地址：", address.Hex())

	//公共地址其实就是公钥的 Keccak-256 哈希，然后我们取最后 40 个字符（20 个字节）并用“0x”作为前缀
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pub[1:])
	log.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}

/*
keystore 是一个包含经过加密了的钱包私钥
keystore，每个文件只能包含一个钱包密钥对
*/
func GethDemo5() {
	log.Println("geth GethDemo5...")
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	log.Println("ks: ", ks)
	password := "hahahaha"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("account: ", account.Address.Hex())

	// 原本的密钥存储文件
	file := "./wallets/UTC--2025-02-28T14-47-50.330729400Z--eae2d0894975026ad37ea13396e78b2f5b0b1982"
	// 创建新的密钥存储库对象ks
	ks = keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	// 读取文件成字节数组
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	newpassword := "newpassword"
	// 根据原有的密钥生成新的，并修改密码
	acount, err := ks.Import(jsonBytes, password, newpassword)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("account: ", acount.Address.Hex())
	// 删除原来的密钥存储文件
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

/*
判断用户是否为智能合约地址
*/
func GethDemo6() {
	log.Println("geth GethDemo6...")
	// 通过正则来校验地址
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	log.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	log.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	//如果地址存储了字节码，则是一个智能合约地址，否则认为是一个以太坊地址
	//将合约部署在本地链上 npx hardhat ignition deploy ./ignition/modules/Lock.js --network localhost
	// 本地合约地址 0x5FbDB2315678afecb367f032d93F642f64180aa3
	address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	log.Println("address: ", address.Hex())
	byteCode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(byteCode) > 0
	log.Println("isContract: ", isContract)
}

/*
交易区块
*/
func GethDemo7() {
	log.Println("geth GethDemo7...")
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("header: ", header.Number.String())
	blockNumber := big.NewInt(3)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(block.Number().Uint64()) // 块号
	log.Println(block.Time())            // 块时间戳
	log.Println(block.Hash().Hex())
	log.Println(block.Difficulty())        // 区块难度
	log.Println(len(block.Transactions())) // 区块的交易数目

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(count)
}

/*
查询交易
*/
func GethDemo8() {
	log.Println("geth GethDemo8...")
	block, err := client.BlockByNumber(context.Background(), big.NewInt(4))
	if err != nil {
		log.Fatal(err)
	}
	transactions := block.Transactions()
	log.Println("count: ", len(transactions))
	for _, tx := range transactions {
		log.Println(tx.Hash().Hex())
		log.Println(tx.Value().String())
		log.Println(tx.Gas())
		log.Println(tx.GasPrice().Uint64())
		log.Println(tx.Nonce())
		log.Println(tx.Data())
		log.Println("to address:", tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var signer types.Signer
		switch tx.Type() {
		case types.LegacyTxType:
			// 普通 EIP - 155 交易
			signer = types.NewEIP155Signer(chainID)
		case types.AccessListTxType:
			// EIP - 2930 带访问列表的交易
			signer = types.NewEIP2930Signer(chainID)
		case types.DynamicFeeTxType:
			// EIP - 1559 动态手续费交易
			signer = types.NewLondonSigner(chainID)
		default:
			log.Fatalf("unsupported transaction type: %d", tx.Type())
		}

		sender, err := types.Sender(signer, tx)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("sender: ", sender.Hex())

		// 交易收据
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("receipt: ", receipt.Status)
	}

	//有报错 json: cannot unmarshal non-string into Go value of type hexutil.Uint
	blockHash := common.HexToHash("0x2527c9d7d95a741c640f7e72525e5fbf83c567fa7e1080bd2e80b378d9d08d2d")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("count: ", count)
}

// 以太币转账
func GethDemo9() {
	log.Println("geth GethDemo9...")
	privateKey, err := crypto.HexToECDSA("47c99abed3324a2707c28affff1267e45918ec8c3f20b8aa892e8b065d2942dd")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: public key is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 每笔交易都需要一个 nonce。 根据定义，nonce 是仅使用一次的数字。
	// 如果是发送交易的新帐户，则该随机数将为“0”。
	// 来自帐户的每个新事务都必须具有前一个 nonce 增加 1 的 nonce

	log.Println("address: ", fromAddress.Hex())

	// 发送用户余额
	fromAddressBalance, err := client.PendingBalanceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("fromAddressBalance: ", fromAddressBalance)

	// 获取指定地址的待处理 nonce 值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("nonce: ", nonce)

	// 燃气价格必须以 wei 为单位设定
	// 18位 1 ETH 为 10^18 wei
	value := big.NewInt(1000000000000000000)
	// 燃气应设上限为“21000”单位
	gasLimit := uint64(21000)
	log.Println(value)
	log.Println("gasLimit: ", gasLimit)

	// 根据'x'个先前块来获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("gasPrice: ", gasPrice)

	// 转入用户
	toAddress := common.HexToAddress("0xcd3B766CCDd6AE721141F452C550Ca635964ce71")

	// 创建交易
	transaction := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//使用发件人的私钥对事务进行签名
	signTx, err := types.SignTx(transaction, types.NewEIP155Signer(chainID), privateKey)

	// SendTransaction 广播交易
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}

	// 交易hash值
	log.Println("transaction sent:", signTx.Hash())

	// 接受用户余额
	toAddressBalance, err := client.PendingBalanceAt(context.Background(), toAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("toAddressBalance: ", toAddressBalance)
}

func getFromAddress(hexKey string) (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return privateKey, crypto.PubkeyToAddress(*publicKeyECDSA)
}

func getPrice() *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

/*
代币的转账
*/

func GethDemo10() {
	log.Println("geth GethDemo10...")
	// 以 ERC-20 为例

	// 代币发送用户
	privateKey, fromAddress := getFromAddress("f214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897")

	// 接收用户
	toAddress := common.HexToAddress("0xdF3e18d64BC6A983f673Ab319CCaE4f1a57C7097")
	// 智能合约地址
	tokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	// 计算转账函数签名
	// 生成函数签名的 Keccak256 哈希。 然后我们只使用前 4 个字节来获取方法 ID
	transferFnSignature := []byte("transfer(address,uint256)")
	keccak256 := sha3.NewLegacyKeccak256()
	keccak256.Write(transferFnSignature)
	methodID := keccak256.Sum(nil)[:4]
	log.Println("methodID: ", hexutil.Encode(methodID))

	// 将接收代币的地址左填充到 32 字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	log.Println("paddedAddress: ", hexutil.Encode(paddedAddress))

	// 发送10个eth
	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10)

	// 将代币量左填充到32个字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	log.Println("paddedAmount: ", hexutil.Encode(paddedAmount))

	// 对以上字节进行拼接
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 通过EstimateGas 可以估算所需燃气量
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	gasLimit = gasLimit * 120 / 100
	log.Println("gasLimit: ", gasLimit)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	// 构建交易事务类型
	tx := types.NewTransaction(nonce, tokenAddress, amount, gasLimit, getPrice(), data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("transaction sent:", signTx.Hash().Hex())
}

func GethDemo11() {
	log.Println("geth GethDemo11...")
	// 连接到 Hardhat 本地链，默认端口是 8545
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	log.Println("Successfully connected to the Ethereum client.")

	// 创建一个用于接收新区块头的通道
	headers := make(chan *types.Header)

	// 订阅新的区块头
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Failed to subscribe to new headers: %v", err)
	}
	log.Println("Successfully subscribed to new block headers.")

	// 产生新交易可以观察数据有否有新区块产生
	// 循环监听事件
	for {
		select {
		case err := <-sub.Err():
			// 处理订阅错误
			log.Fatalf("Subscription error: %v", err)
		case header := <-headers:
			// 处理新的区块头
			log.Printf("New block header received. Block hash: %s, Block number: %d", header.Hash().Hex(), header.Number.Uint64())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatalf("Failed to get block: %v", err)
			}
			log.Println(block.Hash().Hex())
			log.Println(block.Number().Uint64())
			log.Println(block.Time())
			log.Println(block.Nonce())
			log.Println(len(block.Transactions()))
		}
	}
}

func transaction() *types.Transaction {
	// 代币发送用户
	privateKey, fromAddress := getFromAddress("701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82")

	// 接收用户
	toAddress := common.HexToAddress("0x71bE63f3384f5fb98995898A86B02Fb2426c5788")
	// 智能合约地址
	tokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	// 计算转账函数签名
	// 生成函数签名的 Keccak256 哈希。 然后我们只使用前 4 个字节来获取方法 ID
	transferFnSignature := []byte("transfer(address,uint256)")
	keccak256 := sha3.NewLegacyKeccak256()
	keccak256.Write(transferFnSignature)
	methodID := keccak256.Sum(nil)[:4]
	log.Println("methodID: ", hexutil.Encode(methodID))

	// 将接收代币的地址左填充到 32 字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	log.Println("paddedAddress: ", hexutil.Encode(paddedAddress))

	// 发送10个eth
	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10)

	// 将代币量左填充到32个字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	log.Println("paddedAmount: ", hexutil.Encode(paddedAmount))

	// 对以上字节进行拼接
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 通过EstimateGas 可以估算所需燃气量
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	// gas不够 加点上限
	gasLimit = gasLimit * 12 / 10
	log.Println("gasLimit: ", gasLimit)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	// 构建交易事务类型
	tx := types.NewTransaction(nonce, tokenAddress, amount, gasLimit, getPrice(), data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return signTx
}

func GethDemo12() {
	log.Println("geth GethDemo12...")

	signTx := transaction()
	/*
		对以太坊交易进行签名，并将签名后的交易数据编码为十六进制字符串，以便后续可以将该字符串发送到以太坊网络进行交易广播
	*/
	err := client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}
	ts := types.Transactions{signTx}
	var buf bytes.Buffer
	ts.EncodeIndex(0, &buf)
	rawTxHex := hex.EncodeToString(buf.Bytes())
	log.Println("rawTxHex: ", rawTxHex)
}

func GethDemo13() {
	log.Println("geth GethDemo13...")
	rawTx := "f8b20684464c03bf826567945fbdb2315678afecb367f032d93f642f64180aa3880de0b6b3a7640000b844a9059cbb000000000000000000000000df3e18d64bc6a983f673ab319ccae4f1a57c70970000000000000000000000000000000000000000000000000de0b6b3a764000082f4f6a09f70ae0fdea61ad1493392bcdde53dcfa330ff57c1a100b8a2956b6fb6043671a04617ce8add6f82d9d6e505b11fe3609f4a1ea0c99190aeb720c937f172bb80e6"
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		log.Fatal(err)
	}
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tx sent:", tx.Hash().Hex())
}
