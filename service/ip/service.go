package ip

import (
	"proxy-collect/config"
	"proxy-collect/dto"
)

var localIpServiceInstance *localIpService

func LocalIpService() *localIpService {
	if localIpServiceInstance == nil {
		LoadLocalIpData()
	}
	return localIpServiceInstance
}

func LoadLocalIpData() {
	// IPData IP库的数据
	var IPData = &fileData{
		FilePath: config.Get().LocalIpDataPath,
	}
	IPData.InitIPData()
	localIpServiceInstance = NewLocalIpService(IPData)
}

func GetIpInfo(host string, port string) *dto.IpInfoDto {
	result, err := LocalIpService().Find(host)
	if err != nil {
		return GetIpInfoByIp138(host, port)
	}
	return result
}

func UpdateLocalIpData() {
	LocalIpService().Data.UpdateLocalData()
}
