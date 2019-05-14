package util

import (
	"crypto/rand"
	"fmt"
	"forum/util/config"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	regular = "^((1[0-9][0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$"
)

var SysLogs = logrus.New()
var ClientLogs = logrus.New()
var StaffLogs = logrus.New()
var validate *validator.Validate

//发送验证码
func SentMsg(mobile string, code string) int {
	pwd := config.Mes.Pwd
	name := config.Mes.Name
	content := "短信验证码为：" + code + "，请勿将验证码提供给他人。"

	httpurl := config.Mes.Url

	url := httpurl + "?name=" + name + "&pwd=" + pwd + "&content=" + content + "&mobile=" + mobile + "&sign=易编学&type=pt"
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		CheckErr(err)
	}
	mes := SplitMesString(string(body))
	return mes

}

//打印系统日志1
func FmtSyslog(msg string) {
	SysLogs.Println(msg)
}

//打印用户日志
func FmtClientLog(msg string) {
	ClientLogs.Println(msg)
}

//打印员工日志
func FmtStaffLog(msg string) {
	ClientLogs.Println(msg)
}

//手机号正则验证
func CheckPhone(p string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(p)
}

//随机生成6位数字
func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

//生成uuid
func GetUuid() string {
	uuid := uuid.NewV4()
	return uuid.String()
}

//返回一个Validate实例
func GetValidate() *validator.Validate {
	validate = validator.New() //默认导入的是v8 手动改v9
	return validate
}
func CheckErr(err error) {
	if err != nil {
		SysLogs.Println(err)
		_, file, line, _ := runtime.Caller(1)
		SysLogs.Println("Worning: " + file + ":" + strconv.Itoa(line))
		_, file, line, _ = runtime.Caller(2)
		SysLogs.Println("Worning: " + file + ":" + strconv.Itoa(line))
	}
}

/**
 * Struct to map: Convert a struct to a map[string]interface{}
 *
 * @param {struct} input Struct
 * @return {map} example: {"Name":"gopher", "ID":123456, "Enabled":true}
 */
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

/**
 * 判断结构体里是否有对应字段
 *
 * @param {struct} input Struct
 * @return {map} example: {"Name":"gopher", "ID":123456, "Enabled":true}
 */
func CheckStruct(obj interface{}, a string) bool {
	var slice []string
	m := Struct2Map(obj)
	for k := range m {
		slice = append(slice, k)
	}
	for _, v := range slice {
		if v == a {
			return true
		}
	}
	return false
}

/**
 * 时间戳转时间
 *
 * @param  时间戳
 * @return  time格式
 */
func IntToTime(t int64) time.Time {
	tm := time.Unix(t, 0)
	return tm
}

/**
 * 根据传进来的年龄 返回对应的出生年月
 *
 * @param  时间戳
 * @return  time格式
 */
func IntTime(num int) time.Time {
	year := time.Now().Year()
	birth := year - num
	data := strconv.Itoa(birth) + "-01-01 00:00:00"
	t, _ := time.Parse("2006-01-02 15:04:05", data)
	return t
}

/**
 * 获取token
 *
 */
func GetToken(c *gin.Context) string {
	//根据token 从redis中获取
	token := c.Request.Header.Get("token") //获取token
	return token
}

/**
 * 根据当前时间算给给定的出生日期算年龄
 *
 * @param  出生日期
 * @return  年龄
 */
func DateToAge(birth time.Time) int {
	str := birth.Format("2006-01-02 15:04:05")
	gety := string(str[0:4])
	year := time.Now().Year()
	num, _ := strconv.Atoi(gety)
	age := year - num
	if age <= 0 {
		age = 1
	}
	return age
}

/**
 * 根据年龄算出出生日期
 *
 * @param  出生日期
 * @return  年龄
 */
func AgeToBir(bir int) time.Time {
	year := time.Now().Year()
	birth := year - bir
	data := string(birth) + "-01" + "-01"
	t, _ := time.Parse("2006-01-02 15:04:05", data)
	return t
}

/**
 * 字符串切割
 *
 * @param  错误信息
 * @return  根据；切割后的最后一段信息
 */
func SplitErrString(err error) string {
	a := strings.Split(err.Error(), ":")
	return a[len(a)-1]
}

/**
 * 字符串切割
 *
 * @param  短信返回信息
 * @return  根据；切割后的最后一段信息
 */
func SplitMesString(msg string) int {
	a := strings.Split(msg, ",")
	num, err := strconv.Atoi(a[0])
	if err != nil {
		CheckErr(err)
	}
	return num
}

/**
 * 初始化系统日志
 *
 * @param  路经 文件名保存时间 切割时间
 *
 */
func ConfigsystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		SysLogs.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	SysLogs.AddHook(lfHook)
}

/**
 * 初始化客户日志
 *
 * @param  路经 文件名保存时间 切割时间
 *
 */
func ConfigClientLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		ClientLogs.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	ClientLogs.AddHook(lfHook)
}

/**
 * 初始化员工日志
 *
 * @param  路经 文件名保存时间 切割时间
 *
 */
func ConfigStaffLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		StaffLogs.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	StaffLogs.AddHook(lfHook)
}

/**
 * Duration to string
 *
 * @param  duration格式
 * @return 转换之后的string
 */
func ShortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}

//返回参数封装
func Response(status int, msg interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["status"] = status
	m["msg"] = msg
	return m
}

// IsEmpty 判断一个值是否为空(0, "", false, 空数组等)。
// []string{""}空数组里套一个空字符串，不会被判断为空。
func IsEmpty(expr interface{}) bool {
	if expr == nil {
		return true
	}

	switch v := expr.(type) {
	case bool:
		return !v
	case int:
		return 0 == v
	case int8:
		return 0 == v
	case int16:
		return 0 == v
	case int32:
		return 0 == v
	case int64:
		return 0 == v
	case uint:
		return 0 == v
	case uint8:
		return 0 == v
	case uint16:
		return 0 == v
	case uint32:
		return 0 == v
	case uint64:
		return 0 == v
	case string:
		return len(v) == 0
	case float32:
		return 0 == v
	case float64:
		return 0 == v
	case time.Time:
		return v.IsZero()
	case *time.Time:
		return v.IsZero()
	}

	// 符合 IsNil 条件的，都为 Empty
	if IsNil(expr) {
		return true
	}

	// 长度为 0 的数组也是 empty
	v := reflect.ValueOf(expr)
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return 0 == v.Len()
	}

	return false
}

// IsNil 判断一个值是否为 nil。
// 当特定类型的变量，已经声明，但还未赋值时，也将返回 true
func IsNil(expr interface{}) bool {
	if nil == expr {
		return true
	}

	v := reflect.ValueOf(expr)
	k := v.Kind()

	return k >= reflect.Chan && k <= reflect.Slice && v.IsNil()
}
