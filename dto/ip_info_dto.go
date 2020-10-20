package dto

import "fmt"

type IpInfoDto struct {
	Country string
	City    string
	Region  string
	Isp     string
}

func (d *IpInfoDto) String() string {
	return fmt.Sprintf("country: %s, city: %s, region: %s, isp: %s", d.Country, d.City, d.Region, d.Isp)
}
