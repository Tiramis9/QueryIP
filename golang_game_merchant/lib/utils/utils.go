package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	ERR_SYSTEM_ERROR    = "system error"
	ERR_AUTH_FAILED     = "auth failed"
	ERR_LACK_PARAMETERS = "Lack of necessary parameters"
	KC_RAND_KIND_NUM    = 0    // 纯数字
	KC_RAND_KIND_LOWER  = 1    // 小写字母
	KC_RAND_KIND_UPPER  = 2    // 大写字母
	KC_RAND_KIND_ALL    = 3    // 数字、大小写字母
	DEFAULT_PAGE        = "1"  //默认第几页
	DEFAULT_PAGECOUNT   = "10" //默认一页多少数量
)

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
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
func CreateOrderNo(userId int) string {
	now := time.Now().Unix()
	nowStr := strconv.FormatInt(now, 10)
	userIdStr := strconv.Itoa(userId)
	randInt := RandInt64(100, 999)
	randStr := strconv.FormatInt(randInt, 10)
	return nowStr + userIdStr + randStr
}

func Log(content interface{}, _type string, file_name string) {
	if file_name == "" {
		file_name = "mylog.log"
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

func IsIntContains(list []int, element int) bool {
	for i := range list {
		if list[i] == element {
			return true
		}
	}

	return false
}

//MD5加密
func Md5S(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
