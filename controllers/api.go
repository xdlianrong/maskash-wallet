package controllers

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	ecc "wallet/ECC"
	"wallet/model"
	"wallet/utils"
)

const (
	ErrorValue   = "value cannot be empty"
	RejectServer = "Server Error"
)

func Register(c echo.Context) error {
	w := new(model.NewWallet)
	// 因为 echo 的 bind 无绑定检查功能
	// echo 强制要求 post 的参数写在 body 里，写在 header 里会绑定不上
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 暂时只能验证是否为空
	if w.Id == "" || w.Name == "" || w.Str == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	}
	// 计算公私钥
	account := ecc.GenerateAccount(w.Str, w.Name, w.Id, w.Str)
	res, err := register(account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RejectServer)
	}
	if res == "Successful!" {
		return c.JSON(http.StatusOK, account.KeyToString())
	} else {
		return c.JSON(http.StatusInternalServerError, RejectServer)
	}

	//_, priv, err := ELGamal.GenerateKeys(w.Str)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, err)
	//}
	// 取哈希  zr:默认注册时，将公钥发给监管者存入公钥池，这里不取hash
	// pub.G1 = new(big.Int)
	// HashInfoBuf := sha256.Sum256([]byte(w.Str))
	// 向监管者提交注册请求，并返回相关信息
	// [32]byte 是一个数组，要把他转换成切片
	//if resp, err := http.PostForm(RegulatorURL+"register", url.Values{"name": {w.Name}, "id": {w.Id}, "Hashky": {account.KeyToString().Publickey}}); err != nil {
	//	c.JSON(http.StatusInternalServerError, err)
	//	return c.JSON(http.StatusInternalServerError, err)
	//} else {
	//	if res, err := ioutil.ReadAll(resp.Body); err != nil {
	//		c.JSON(http.StatusInternalServerError, err)
	//		return c.JSON(http.StatusInternalServerError, err)
	//	} else {
	//		// 判断应该返回的信息
	//		if bytes.Equal(res,[]byte("Successful!")) {
	//			return c.JSON(http.StatusOK, account.KeyToString())
	//		} else {
	//			return c.JSON(http.StatusInternalServerError, RejectServer)
	//		}
	//	}
	//}
}

func register(account ecc.Account) (string, error) {
	data := account.Info
	body, err := ethRPCPost(data, RegulatorURL+"register")
	if err != nil {
		return "", err
	}
	res := string(body)
	if res == "Successful!" {
		fmt.Println("账户" + account.Info.Name + "注册成功")
	} else if res == "Account registered!" {
		fmt.Println("账户" + account.Info.Name + "已注册")
	} else if res == "Fail!" {
		Fatalf("账户" + account.Info.Name + "注册失败")
	}
	return string(body), nil
}

func Buycoin(c echo.Context) error {
	w := new(model.BctoEx)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 向交易所发出购币请求
	body, err := ethRPCPost(w, ExchangeURL+"buy")
	if err != nil {
		return err
	}
	var receipt utils.Receipt
	json.Unmarshal(body, &receipt)

	if receipt.Cmv == "" || receipt.Epkrc1 == "" || receipt.Epkrc2 == "" || receipt.Hash == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	} else {
		// 购买成功,随机数解密
		privKey := utils.CreatePriKey(w.G1, w.G2, w.P, w.H, w.X)
		coin := decryptCoinReceipt(receipt, privKey, w.Amount)
		utils.MineTx(model.EthPort, coin.Hash)
		rpcTx, _ := utils.EthGetTransactionByHash(model.EthPort, receipt.Hash)
		resBody := utils.CoinNTx{
			Coin: coin,
			TX:   rpcTx,
		}
		return c.JSON(http.StatusOK, resBody)
	}
}

func ExchangeCoin(c echo.Context) error {
	w := new(model.ExchangeCoin)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	senderPriv := utils.CreatePriKey(w.SG1, w.SG2, w.SP, w.SH, w.SX)
	reciverPub := utils.CreatePubKey(w.RG1, w.RG2, w.RP, w.RH)
	coin := utils.Coin{
		Cmv:    w.Cmv,
		Vor:    w.Vor,
		Amount: w.Amount,
	}
	amount := coin.Amount
	spend := w.Spend
	if spend > amount {
		return c.JSON(http.StatusOK, errors.New("转出金额不可大于承诺额"))
	}
	senderEthAccount, err := utils.EthAccounts(model.EthPort)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	senderGethAccount := senderEthAccount[0]
	receiverEthAccount, err := utils.EthAccounts(model.EthPort)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	receiverGethAccount := receiverEthAccount[0]
	txHash, err := utils.EthSendTransaction(model.EthPort, senderGethAccount, receiverGethAccount, senderPriv, reciverPub, coin, amount, spend)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	utils.MineTx(model.EthPort, txHash)
	rpcTx, err := utils.EthGetTransactionByHash(model.EthPort, txHash)
	tx := rpcTx.Result
	resCoin := utils.Coin{
		Cmv:    tx.CmR,
		Vor:    decrypt(tx.CmRRC1, tx.CmRRC2, senderPriv),
		Hash:   txHash,
		Amount: amount - spend,
	}
	resBody := utils.CoinNTx{
		Coin: resCoin,
		TX:   rpcTx,
	}
	return c.JSON(http.StatusOK, resBody)
}
func Receive(c echo.Context) error {
	w := new(model.ReceiveData)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(w.Hash) != 66 {
		return c.JSON(http.StatusOK, errors.New("未找到此交易"))
	}
	privKey := utils.CreatePriKey(w.G1, w.G2, w.P, w.H, w.X)
	rpcTx, err := utils.EthGetTransactionByHash(model.EthPort, w.Hash)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	tx := rpcTx.Result
	amountHex := decryptValue(tx.EvsBsC1, tx.EvsBsC2, privKey)
	s, _ := strconv.ParseInt(amountHex, 0, 32)
	returnCoin := utils.Coin{
		Cmv:    tx.CmO,
		Vor:    decrypt(tx.CmSRC1, tx.CmSRC2, privKey),
		Hash:   w.Hash,
		Amount: int(s),
	}
	return c.JSON(http.StatusOK, returnCoin)
}
func decryptCoinReceipt(recript utils.Receipt, priv ecc.PrivateKey, amount int) utils.Coin {
	return utils.Coin{
		Cmv:    recript.Cmv,
		Vor:    decrypt(recript.Epkrc1, recript.Epkrc2, priv),
		Hash:   recript.Hash,
		Amount: amount,
	}
}

//	解密随机数密文
func decrypt(hex0xStringC1 string, hex0xStringC2 string, priv ecc.PrivateKey) string {
	hexData1, _ := hex.DecodeString(hex0xStringC1[2:])
	hexData2, _ := hex.DecodeString(hex0xStringC2[2:])
	C := ecc.CypherText{
		C1: hexData1,
		C2: hexData2,
	}
	M := fmt.Sprintf("0x%x", ecc.ECCDecrypt(priv, C))
	return M
}

//	解密随机数密文
func decryptValue(hex0xStringC1 string, hex0xStringC2 string, priv ecc.PrivateKey) string {
	hexData1, _ := hex.DecodeString(hex0xStringC1[2:])
	hexData2, _ := hex.DecodeString(hex0xStringC2[2:])
	C := ecc.CypherText{
		C1: hexData1,
		C2: hexData2,
	}
	M := fmt.Sprintf("0x%x", ecc.DecryptValue(priv, C))
	return M
}
