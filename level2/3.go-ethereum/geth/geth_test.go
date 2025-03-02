package geth

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestGeth(t *testing.T) {

	//GethDemo()
	//GethDemo1()
	//GethDemo2()
	//GethDemo3()
	//GethDemo4()
	//GethDemo5()
	//GethDemo6()
	//GethDemo7()
	//GethDemo8()
	//GethDemo9()
	GethDemo10()
	//GethDemo11()
	//GethDemo12()
	//GethDemo13()
}

// # 部署合约
// npx hardhat ignition deploy ./ignition/modules/Lock.js
// # 进行交易
// npx hardhat run ./test/TodoList.js --network localhost
