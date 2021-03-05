package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"wallet/ELGamal"
	"wallet/model"
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
	fmt.Println("绑定完成")
	if w.Id == "" || w.Name == "" || w.Str == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	}
	// 计算公私钥
	account := ELGamal.GenerateAccount(w.Str, w.Name, w.Id, w.Str)
	if res := register(account); res == "Successful!" {
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

func register(account ELGamal.Account) string {
	data := account.Info
	body := ethRPCPost(data, RegulatorURL+"register")
	res := string(body)
	if res == "Successful!" {
		fmt.Println("账户" + account.Info.Name + "注册成功")
	} else if res == "Account registered!" {
		fmt.Println("账户" + account.Info.Name + "已注册")
	} else if res == "Fail!" {
		Fatalf("账户" + account.Info.Name + "注册失败")
	}
	return string(body)
}

// 39.105.58.136
func Buycoin(c echo.Context) error {
	w := new(model.BctoEx)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 向交易所发出购币请求
	body := ethRPCPost(w, ExchangeURL+"buy")
	var receipt Receipt
	json.Unmarshal(body, &receipt)
	if receipt.Cmv == "" || receipt.Epkrc1 == "" || receipt.Epkrc2 == "" || receipt.Hash == "" {
		return c.JSON(http.StatusBadRequest, ErrorValue)
	} else {
		// 购买成功,随机数解密在前端进行
		return c.JSON(http.StatusOK, receipt)
	}
}

func ExchangeCoin(c echo.Context) error {
	w := new(model.ExchangeCoin)
	if err := c.Bind(w); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	coin := decryptCoinReceipt(w.Receipt, w.Priv)
	return c.JSON(http.StatusOK, coin)
}

func decryptCoinReceipt(recript Receipt, priv ELGamal.PrivateKey) Coin {
	return Coin{
		Cmv:  recript.Cmv,
		Vor:  decrypt(recript.Epkrc1, recript.Epkrc2, priv),
		Hash: recript.Hash,
	}
}

//	解密随机数密文
func decrypt(hex0xStringC1 string, hex0xStringC2 string, priv ELGamal.PrivateKey) string {
	hexData1, _ := hex.DecodeString(hex0xStringC1[2:])
	hexData2, _ := hex.DecodeString(hex0xStringC2[2:])
	C := ELGamal.CypherText{
		C1: hexData1,
		C2: hexData2,
	}
	M := fmt.Sprintf("0x%x", ELGamal.Decrypt(priv, C))
	return M
}
