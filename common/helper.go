package common

import (
	"fmt"
	"hls-parse/types"
	"regexp"
	"strconv"
	"strings"
)

func StringToInt(str string) (result int, err error) {
	result, err = strconv.Atoi(str)
	return
}

// 将字符串形式的协议转换为数组，并清理无用的数据
//
// @param strHls *string sdf
func ProtocolStrToArray(strHls *string) (result []string, err error) {
	arrHls := strings.Split(*strHls, "\n")
	for _, v := range arrHls {
		val := strings.TrimSpace(v)
		if len(val) != 0 { // 去除空行
			result = append(result, val)
		}
	}
	return
}

// 提取协议标签类型
func ExtractTag(protoLine string) (result types.ProtocolTagType) {
	reg, _ := regexp.Compile("^#E([^:])+")
	if reg.MatchString(protoLine) {
		tag1 := reg.FindString(protoLine)
		switch tag1 {
		case "#EXTM3U":
			result = types.ProtocolTagExtm3U
		case "#EXT-X-STREAM-INF":
			result = types.ProtocolTagStreamInf
		case "#EXTINF":
			result = types.ProtocolTagExtinf
		case "#EXT-X-PLAYLIST-TYPE":
			result = types.ProtocolTagPlaylistType
		case "#EXT-X-KEY":
			result = types.ProtocolTagKey
		case "#EXT-X-ENDLIST":
			result = types.ProtocolTagEndlist
		default:
			result = types.ProtocolTagNil
		}
	} else {
		result = types.ProtocolTagValue
	}
	return
}

// 将字符串形式的协议，解构成方面使用的结构化对象
func DestructureParams(protoLine string) (params types.ProtocolParams, err error) {
	if ExtractTag(protoLine) == types.ProtocolTagNil {
		return types.ProtocolParams{}, fmt.Errorf("不符合协议规范")
	}
	reg, _ := regexp.Compile("^#E([^:])+:")
	list := reg.Split(protoLine, -1) // 拆分协议中的参数字符串
	if len(list) < 2 {
		return types.ProtocolParams{}, fmt.Errorf("没有参数")
	}
	strParams := strings.TrimSpace(list[1])
	// TODO 去掉单引号和双引号
	arrParams := strings.Split(strParams, ",") // 拆分参数字符串
	params.Map = make(map[string]string)

	for _, v := range arrParams {
		param := strings.Split(strings.TrimSpace(v), "=")
		if len(param) == 2 { // key/value 形式
			params.Map[param[0]] = param[1]
		} else { // 数组形式
			params.Array = append(params.Array, param[0])
		}
	}

	fmt.Println(arrParams)
	err = nil
	return
}
