package accesstoken

import (
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
	"bytes"
	"okni.com/siapool/util"
	"io/ioutil"
	"time"
	"strconv"
	"crypto/md5"
	"encoding/base64"
	"log"
	"fmt"
)

var (
	ErrorToken      = errors.New("not found record")
	ErrorTokenId    = errors.New("id is not allow")
	ErrorTokenParam = errors.New("list team param error")
)

// 后台房间信令获取access token
type AccessTokenParam struct {
	Version              string   	`form:"version"`
	Seq                  string   	`form:"seq"`
	AppId                string    	`form:"app_id"`
	BizType              string    	`form:"biz_type"`
	Token                string   	`form:"token"`
}

// 后台房间信令获取access token
type AccessToken struct {
	Version              uint32   	`json:"version"`
	Seq                  uint64   	`json:"seq"`
	AppId                int32    	`json:"app_id"`
	BizType              int32    	`json:"biz_type"`
	Token                string   	`json:"token"`
}

// TrapGetAccessToken will get accesstoken
// @Summary get accesstoken
// @Accept json
// @Tags accesstoken
// @Security Bearer
// @Produce  json
// @Param Version query string true "version"
// @Param Seq query string true "seq"
// @Param AppId query string true "app_id"
// @Param BizType query string true "biz_type"
// @Param Token query string false "token"
// @Resource accesstoken
// @Router /accesstoken [post]
// @Success 200 {string} string "token message"
func TrapGetAccessToken(c *gin.Context)  {
	//获取参数
	var param AccessTokenParam
	if err := c.ShouldBindQuery(&param); err != nil{
		c.JSON(400, c.AbortWithError(400, err))
		return
	}
	access := AccessToken{}
	v,_ := strconv.Atoi(param.Version)
	access.Version = uint32(v)

	s,_ := strconv.Atoi(param.Seq)
	access.Seq = uint64(s)

	a,_ := strconv.Atoi(param.AppId)
	access.AppId = int32(a)

	b,_ := strconv.Atoi(param.BizType)
	access.BizType = int32(b)

	// 获取token
	if t, err := buildToken(); err!= nil{
		log.Println("get token error")
		c.JSON(400, c.AbortWithError(400, err))
		return
	} else {
		access.Token = t
		log.Println("param.Token : ", access.Token)
		time.Sleep(100)
	}

	//正式环境：
	//curl -X POST https://liveroomsvr{APPID}-api.zego.im/cgi/token -d 'json_str'
	url := "https://liveroomsvr-test.zego.im/cgi/token"

	if k, err := json.Marshal(access); err != nil{
		c.JSON(400, c.AbortWithError(400, err))
		return
	}else{
		log.Println("app_id:",access.AppId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(k))
		if err != nil{
			c.JSON(400, c.AbortWithError(400, err))
			return
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
		var client *http.Client
		timeoutIntv := util.MustParseDuration("10m")
		client = &http.Client{
			Timeout: timeoutIntv,
		}

		resp,err := client.Do(req)
		if err!=nil{
			c.JSON(400, c.AbortWithError(400, err))
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		c.JSON(http.StatusOK, gin.H{
			"token message": string(body),
		})
	}
}

func buildToken() (string, error) {
	current_time := time.Now().Unix()
	expired_time := time.Now().Add(2*time.Hour).Unix()

	log.Println("当前时间: " + strconv.Itoa(int(current_time)))
	log.Println("过期时间: " + strconv.Itoa(int(expired_time)))

	appid := "512206300" //分配给客户的appid
	serverSecret := "81265ff28cf0ec842366aa157deb86ce" // 联系即构技术支持获取
	nonce := "12345678123456781" // 建议使用随机串

	aStr := appid + serverSecret + nonce +  strconv.Itoa(int(expired_time))
	log.Println("未hash的串: ", aStr)

	//aStr = "51220630081265ff28cf0ec842366aa157deb86ce123456781234567811535527836"
	var data []byte = []byte(aStr)
	ret := md5.Sum(data)
	hashStr := MarshalMd5(ret[:])
	log.Println("hashStr后的字符串:", hashStr)

	//tokenInfo := map[string]interface{}{"ver":1,"hash":hashStr,"nonce":nonce,"expired":1535517067}
	//tokenInfo := map[string]interface{}{"expired":1535517067,"nonce":nonce,"hash":hashStr,"ver":1}

	var tokenInfo Info
	tokenInfo.Ver = 1
	tokenInfo.Hash = hashStr
	tokenInfo.Nonce = nonce
	tokenInfo.Expired = int(expired_time)
	log.Println(tokenInfo)
	if mjson, err := json.Marshal(tokenInfo);err!=nil{
		log.Println(err)
		return "", err
	}else{
		token := string(mjson)
		log.Println("json string:", token)
		//Base64 加密
		//temp := ConvertToString(token,"utf-8","ascii")
		//log.Println("json string temp:", temp)
		encoded := base64.StdEncoding.EncodeToString([]byte(string(mjson)))
		//encoded := base64.StdEncoding.EncodeToString(temp)
		log.Println("access token:", encoded)
		return encoded, nil
	}
}

func MarshalMd5(v []byte) string {
	var tmp string
	l := len(v)
	if l > 0 {
		for i:=0; i<l; i++{
			tmp += fmt.Sprintf("%x", v[i])
		}
	}
	return tmp
}


type Info struct{
	Ver int			`json:"ver"`
	Hash string		`json:"hash"`
	Nonce string	`json:"nonce"`
	Expired int		`json:"expired"`
}


