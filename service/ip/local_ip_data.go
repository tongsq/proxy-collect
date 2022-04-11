package ip

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tongsq/go-lib/logger"
)

type fileData struct {
	Data     []byte
	FilePath string
	Path     *os.File
	IPNum    int64
}

func (f *fileData) UpdateLocalData() {
	tmpData, err := f.getOnline()
	if err != nil {
		logger.Error("更新本地纯真ip库失败", map[string]interface{}{"err": err})
		return
	} else {
		if err := ioutil.WriteFile(f.FilePath, tmpData, 0644); err == nil {
			logger.FDebug("已将最新的纯真 IP 库保存到本地 %s ", f.FilePath)
		}
		f.loadData(tmpData)
	}
}

// InitIPData 初始化ip库数据到内存中
func (f *fileData) InitIPData() {
	var tmpData []byte
	// 判断文件是否存在
	_, err := os.Stat(f.FilePath)
	if err != nil && os.IsNotExist(err) {
		logger.Info("文件不存在，尝试从网络获取最新纯真 IP 库", nil)
		tmpData, err = f.getOnline()
		if err != nil {
			return
		} else {
			if err := ioutil.WriteFile(f.FilePath, tmpData, 0644); err == nil {
				logger.FDebug("已将最新的纯真 IP 库保存到本地 %s ", f.FilePath)
			}
		}
	} else {
		// 打开文件句柄
		logger.FDebug("从本地数据库文件 %s 打开\n", f.FilePath)
		f.Path, err = os.OpenFile(f.FilePath, os.O_RDONLY, 0400)
		if err != nil {
			return
		}
		defer f.Path.Close()

		tmpData, err = ioutil.ReadAll(f.Path)
		if err != nil {
			logger.Error("load local ip data fail", map[string]interface{}{"err": err})
			return
		}
	}

	f.loadData(tmpData)
}

func (f *fileData) loadData(tmpData []byte) {
	f.Data = tmpData

	buf := f.Data[0:8]
	start := binary.LittleEndian.Uint32(buf[:4])
	end := binary.LittleEndian.Uint32(buf[4:])

	f.IPNum = int64((end-start)/IndexLen + 1)
}

func (f *fileData) getOnline() ([]byte, error) {
	resp, err := http.Get("http://update.cz88.net/ip/qqwry.rar")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		if key, err := f.getKey(); err != nil {
			return nil, err
		} else {
			for i := 0; i < 0x200; i++ {
				key = key * 0x805
				key++
				key = key & 0xff

				body[i] = byte(uint32(body[i]) ^ key)
			}

			reader, err := zlib.NewReader(bytes.NewReader(body))
			if err != nil {
				return nil, err
			}

			return ioutil.ReadAll(reader)
		}
	}
}

func (f *fileData) getKey() (uint32, error) {
	resp, err := http.Get("http://update.cz88.net/ip/copywrite.rar")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return 0, err
	} else {
		// @see https://stackoverflow.com/questions/34078427/how-to-read-packed-binary-data-in-go
		return binary.LittleEndian.Uint32(body[5*4:]), nil
	}
}
