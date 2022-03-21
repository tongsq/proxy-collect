package ip

import (
	"proxy-collect/config"
	"proxy-collect/dto"
)

var LocalIpService *localIpService

func init() {
	// IPData IP库的数据
	var IPData = &fileData{
		FilePath: config.Get().LocalIpDataPath,
	}
	IPData.InitIPData()
	LocalIpService = NewLocalIpService(IPData)
}

func GetIpInfo(host string, port string) *dto.IpInfoDto {
	result, err := LocalIpService.Find(host)
	if err != nil && result != nil && result.City != "" {
		return GetIpInfoByIp138(host, port)
	}
	return result
}

func UpdateLocalIpData() {
	LocalIpService.Data.UpdateLocalData()
}
