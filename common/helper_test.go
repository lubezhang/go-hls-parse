package common

import (
	"fmt"
	"os"
	"testing"

	"github.com/lubezhang/hls-parse/types"

	"github.com/stretchr/testify/assert"
)

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestProtocolStrToArray(t *testing.T) {
	var strHls = `#EXTM3U
    


	
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1064000
	1000kbps.m3u8
	
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=564000
	/500kbps.m3u8
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=282000
	/path1/250kbps.m3u8
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=2128000
	path2/2000kbps.m3u8
	#EXT-X-ENDLIST`

	arrHls, _ := ProtocolStrToArray(&strHls)
	assetObj := assert.New(t)
	assetObj.Equal(arrHls[0], "#EXTM3U", "协议根元素必须是#EXTM3U")
	assetObj.NotEqual(arrHls[1], "", "协议对象包含无用数据")
	assetObj.Equal(arrHls[1], "#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1064000")
}

func TestExtractTag(t *testing.T) {
	assetObj := assert.New(t)
	assetObj.Equal(ExtractTag("#EXTM3U"), types.ProtocolTagExtm3U, "协议根元素EXTM3U标签解析错误")
	assetObj.Equal(ExtractTag("#EXT-X-PLAYLIST-TYPE:VOD"), types.ProtocolTagPlaylistType, "协议标签 - EXT-X-PLAYLIST-TYPE解析错误")
	assetObj.Equal(ExtractTag("#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1064000"), types.ProtocolTagStreamInf, "协议标签 - EXT-X-STREAM-INF 解析错误")
	assetObj.Equal(ExtractTag("#EXTINF:2.043"), types.ProtocolTagExtinf, "协议标签 - EXTINF 解析错误")
	assetObj.Equal(ExtractTag("#EXT-X-KEY:METHOD=AES-128,URI=sdfsdf"), types.ProtocolTagKey, "协议标签 - EXT-X-KE 解析错误")
	assetObj.Equal(ExtractTag("1000kbps.m3u8"), types.ProtocolTagValue, "协议标签 - 协议标签对应的数据")
	assetObj.Equal(ExtractTag("#EXT-X-KEY1:METH"), types.ProtocolTagNil, "协议标签 - 不支持次协议标签")
}

func TestDestructureParams(t *testing.T) {
	assetObj := assert.New(t)

	params, err := DestructureParams("#EXT-X-PLAYLIST-TYPE:VOD")
	if err == nil {
		assetObj.Equal(len(params.Array), 1)
		assetObj.Equal(params.Array[0], "VOD")
	} else {
		t.Errorf(err.Error())
	}

	params1, err1 := DestructureParams("#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1064000")
	if err1 == nil {
		assetObj.Equal(len(params1.Map), 2)
		assetObj.Equal(params1.Map["PROGRAM-ID"], "1")
		assetObj.Equal(params1.Map["BANDWIDTH"], "1064000")
	} else {
		t.Errorf(err.Error())
	}

	tag1 := "#EXT-X-KEY:METHOD=AES-128,URI=\"https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/key.key\""
	params2, err2 := DestructureParams(tag1)
	if err2 == nil {
		assetObj.Equal(len(params2.Map), 2)
		assetObj.Equal(params2.Map["METHOD"], "AES-128")
		assetObj.Equal(params2.Map["URI"], "https://ts4.chinalincoln.com:9999/20210419/OvroTYry/1000kb/hls/key.key")
	} else {
		t.Errorf(err2.Error())
	}
}

func TestJoinUrl(t *testing.T) {
	assetObj := assert.New(t)
	url1, _ := JoinUrl("1000kbps.m3u8", "http://www.baidu.com")
	assetObj.Equal(url1, "http://www.baidu.com/1000kbps.m3u8")

	url2, _ := JoinUrl("1000kbps.m3u8", "http://www.baidu.com/path1/path2")
	assetObj.Equal(url2, "http://www.baidu.com/path1/path2/1000kbps.m3u8")

	url3, _ := JoinUrl("/1000kbps.m3u8", "http://www.baidu.com/path1/path2")
	assetObj.Equal(url3, "http://www.baidu.com/1000kbps.m3u8")

	url4, _ := JoinUrl("http://www.baidu.com/1000kbps.m3u8", "http://www.2baidu.com/path1/path2")
	assetObj.Equal(url4, "http://www.baidu.com/1000kbps.m3u8")

	url5, _ := JoinUrl("/1000kbps.m3u8", "http://www.baidu.com/path1/path2/a.html")
	assetObj.Equal(url5, "http://www.baidu.com/1000kbps.m3u8")

	// TODO 用例待完善
	// url6, _ := JoinUrl("1000kbps.m3u8", "http://www.baidu.com/path1/path2/a.html")
	// assetObj.Equal(url6, "http://www.baidu.com/path1/path2/1000kbps.m3u8")
}
