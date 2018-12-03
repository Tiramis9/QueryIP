package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"

	"regexp"

	"log"
	"os"
	"strconv"
	"time"
)

const (
	ERR_SYSTEM_ERROR       = "system error"
	ERR_NET_ERROR          = "network error"
	ERR_AUTH_FAILED        = "auth failed"
	ERR_SIGN_ERROR         = "sign error"
	ERR_APPID_ERROR        = "appid error"
	ERR_SECRET_ERROR       = "secret error"
	ERR_LACK_PARAMETERS    = "Lack of necessary parameters"
	ERR_BALANCE_NOT_ENOUGH = "balance is not enough"
	ERR_ILLEGAL_IP         = "illegal ip"
	KC_RAND_KIND_NUM       = 0    // 纯数字
	KC_RAND_KIND_LOWER     = 1    // 小写字母
	KC_RAND_KIND_UPPER     = 2    // 大写字母
	KC_RAND_KIND_ALL       = 3    // 数字、大小写字母
	DEFAULT_PAGE           = "1"  //默认第几页
	DEFAULT_PAGECOUNT      = "10" //默认一页多少数量
	LOGIN_EXPIRED_TIME     = 3600 //token过期时间
	//redis key值
	INFO_USER  = "info_user"
	INFO_MERCH = "info_merch"
	GAME_AG    = "AG"
	GAME_BBIN  = "BBIN"
	GAME_SB    = "SB"
	GAME_AB    = "AB"
)

var AccountType = [...]string{GAME_AG, GAME_BBIN, GAME_SB, GAME_AB}

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func CheckEmptyStr(s string, i string) string {
	if s == "" {
		return i
	} else {
		return s
	}
}

//正则匹配,匹配规则可以一直加
func RegexpMatch(pattern_type string, source string) bool {
	pattern_list := map[string]string{}
	pattern_list["ip"] = "(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3}"
	pattern_list["email"] = "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$"
	pattern_list["qq"] = "^[1-9]\\d{4,10}$"
	pattern := pattern_list[pattern_type]
	reg := regexp.MustCompile(pattern)
	if res := reg.FindAllString(source, -1); res == nil {
		return false
	} else {
		return true
	}
}

//7天前时间戳的字符串
func SevenDay() string {
	res := time.Now().Unix() - 7*86400
	ress := strconv.FormatInt(res, 10)
	return ress
}

//当时时间戳的字符串
func Now() string {
	res := time.Now().Unix()
	ress := strconv.FormatInt(res, 10)
	return ress
}

//随机数(100,999)
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

//根据用户id生成订单号
func CreateOrderNo(user_id int) string {
	now := time.Now().Unix()
	now_str := strconv.FormatInt(now, 10)
	user_id_str := strconv.Itoa(user_id)
	rand_int := RandInt64(100, 999)
	rand_str := strconv.FormatInt(rand_int, 10)
	return now_str + user_id_str + rand_str
}

func Log(content interface{}, _type string, file_name string) {
	if file_name == "" {
		now := time.Now()
		time_str := now.Format("2006-01-02")
		file_name = time_str + ".log"
	}
	var logFile *os.File
	var err error
	var error2 error
	logFile, err = os.OpenFile(file_name, os.O_APPEND, 0666)
	defer logFile.Close()
	if err != nil && os.IsNotExist(err) {
		logFile, error2 = os.Create(file_name)
		if error2 != nil {
			log.Fatalln("open file error")
		}
	}
	defer logFile.Close()
	debuglog := log.New(logFile, "[info]", log.Llongfile)
	debuglog.SetPrefix("[" + _type + "]")
	debuglog.Println(content)
}

func Debug(content interface{}) {
	Log(content, "dubug", "")
}

func Error(content interface{}) {
	Log(content, "error", "")
}

func HttpPostForm(reqUrl string, params map[string]interface{}) ([]byte, error) {
	fmt.Println(reqUrl)
	query := url.Values{}
	for key, value := range params {
		query.Add(key, fmt.Sprintf("%v", value))
	}
	fmt.Println(query.Encode())
	resp, err := http.PostForm(reqUrl, query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//Post 请求方法
func HttpPost(reqUrl string, params map[string]string) ([]byte, error) {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}
	resp, err := http.PostForm(reqUrl, query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//Post 请求方法
func HttpPostProxy(reqUrl string, params map[string]interface{}) ([]byte, error) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}
	transport := &http.Transport{Proxy: proxy}
	c := &http.Client{Transport: transport}
	query := url.Values{}
	for key, value := range params {
		query.Add(key, fmt.Sprintf("%v", value))
	}
	resp, err := c.PostForm(reqUrl, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

func HttpGet(getUrl string) (string, error) {
	fmt.Println(getUrl)
	resp, err := http.Get(getUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//结构体转json方法
func ToJson(input interface{}) string {
	jsons, errs := json.Marshal(input) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	return string(jsons)
}

//md5方法
func Md5V(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func Md5B(str string) []byte {
	w := md5.New()
	io.WriteString(w, str)
	return w.Sum(nil)
}

//接口字段排序方法
func AbSign(post map[string]string, md5key string) string {
	strs := []string{}
	var sign_str string
	//将键值存入数组中
	for k, _ := range post {
		strs = append(strs, k)
	}
	//键值按自然顺序排序
	sort.Strings(strs)
	for _, val := range strs {
		sign_str += val + "=" + post[val] + "&"
	}
	sign_s := sign_str[0 : len(sign_str)-1]
	//将要签名的字符串
	sign_s = sign_s + md5key
	fmt.Println(sign_s)
	md5_str := Md5V(string(sign_s))
	fmt.Println(md5_str)
	return md5_str
}

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

//AllBet 3des加密
func Encrypt3DES(src []byte, key []byte) []byte {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	src = padding(src, block.BlockSize())
	iv, _ := base64.StdEncoding.DecodeString("AAAAAAAAAAA=")
	blockmode := cipher.NewCBCEncrypter(block, iv)
	blockmode.CryptBlocks(src, src)
	return src
}

//AllBet 3des解密
func Decrypt3DES(src []byte, key []byte) []byte {
	block, _ := des.NewTripleDESCipher(key)
	iv, _ := base64.StdEncoding.DecodeString("AAAAAAAAAAA=")
	fmt.Println("iv:", string(iv))
	blockmode := cipher.NewCBCDecrypter(block, iv)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

//类似php 的http_build_query
func Http_build_query(data map[string]string) string {
	var str string
	for k, v := range data {
		str = str + k + "=" + v + "&"
	}
	str2 := str[0 : len(str)-1]
	return string(str2)
	/*query := url.Values{}
	for k, v := range data {
		query.Add(k, fmt.Sprintf("%v", v))
	}
	return query.Encode()*/
}

//AllBet的sign
func AbDesSign(data map[string]string, deskey string, md5key string) (string, string) {
	key, _ := base64.StdEncoding.DecodeString(deskey)
	fmt.Println(key)
	fmt.Println(deskey[0:24])
	fmt.Println(Http_build_query(data))
	//fmt.Println([]byte(Http_build_query(data)))
	crypted := Encrypt3DES([]byte(Http_build_query(data)), key)
	crypted_str := base64.StdEncoding.EncodeToString(crypted)
	fmt.Println("crypted:", crypted_str)
	crypted_md5 := crypted_str + md5key
	fmt.Println(crypted_md5)
	sign := base64.StdEncoding.EncodeToString(Md5B(crypted_md5))
	fmt.Println("sign:", sign)
	return crypted_str, sign
}

//Allbet订单号
func AbOrderSn(property_id string) string {
	//property_id(7位数字)
	//tail(13位数字)
	timestamp := time.Now().UnixNano() / 1e6
	tail := strconv.FormatInt(timestamp, 10)
	fmt.Println(tail)
	return property_id + tail
}

func HttpPostJson(url string, data interface{}) ([]byte, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
