package basic

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"time"
)

/**
string -> other type
*/
func StrToInt64(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 64)
	return i, err
}

func StrToInt32(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 32)
	return i, err
}

func StrToInt(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	return i, err
}

func StrToFloat64(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 64)
	return f, err
}

func StrToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func StrToByte(s string) []byte {
	return []byte(s)
}

func StrToTime(t string) (time.Time,error){
	// s := "2019-01-08 13:50:30" //外部传入的时间字符串
	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05"                          //常规类型
	// timeTemplate2 := "2006/01/02 15:04:05"                          //其他类型
	// timeTemplate3 := "2006-01-02"                                   //其他类型
	// timeTemplate4 := "15:04:05"                                     //其他类型
	stamp, err := time.ParseInLocation(timeTemplate1, t, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return stamp,err
}

/**
  other type -> string
*/

func ByteToStr(s []byte) string {
	return string(s)
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Uint64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func BoolToStr(i bool) string {
	return strconv.FormatBool(i)
}

func TimeToStr(t time.Time) string{
	return t.Format("2006-01-02 15:04:05")
}

/**
int 转换
*/

func IntToInt32(i int) int32 {
	return int32(i)
}

func IntToInt64(i int) int64 {
	return int64(i)
}

func IntTofloat64(i int) float64 {
	return float64(i)
}

/**
int32 转换
*/
func Int32ToInt(i int32) int {
	return int(i)
}

func Int32ToInt64(i int32) int64 {
	return int64(i)
}

/**
int64 转换
*/
func Int64ToInt(i int64) int {
	return int(i)
}

func Int64ToInt32(i int64) int32 {
	return int32(i)
}

func Map2array(data []map[string]string) ([][]string, error) {
	var err error = nil
	if len(data) == 0 {
		err = errors.New("数组不能为空")
		return nil, err
	}
	map0 := data[0]
	var list0 []string
	for k0, _ := range map0 {
		list0 = append(list0, k0)
	}
	var res [][]string
	res = append(res, list0)
	for _, maping := range data {
		list := []string{}
		for _, lv := range list0 {
			list = append(list, maping[lv])
		}
		res = append(res, list)
	}
	return res, err
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	fieldNum := t.NumField()

	var data = make(map[string]interface{})
	for i := 0; i < fieldNum; i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Map2Json(dict map[string]string) ([]byte, error) {
	jsonStr, err := json.Marshal(dict)
	return jsonStr, err
}

func Json2Map(str string) (map[string]string, error) {
	// str := "{\"address\":\"北京\",\"username\":\"kongyixueyuan\"}"
	mymap := make(map[string]string)
	err := json.Unmarshal([]byte(str), &mymap)
	return mymap, err

}

func Struct2json(v interface{}) ([]byte, error) {
	jsonStr, err := json.Marshal(v)
	return jsonStr, err
}
