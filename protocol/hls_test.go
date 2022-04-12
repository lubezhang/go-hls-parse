package protocol

import (
	"hls-parse/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHlsParse(t *testing.T) {
	assetObj := assert.New(t)

	var strHls = `#EXTM3U

	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=1064000
	1000kbps.m3u8

	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=564000
	/500kbps.m3u8
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=282000
	/path1/250kbps.m3u8
	#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=2128000
	path2/2000kbps.m3u8`

	hls, err := Parse(&strHls)
	assetObj.Nil(err)
	assetObj.Equal(hls.ExtM3u, "#EXTM3U")
	assetObj.Equal(hls.PlayListType, types.PlayListTypeMaster)
}
