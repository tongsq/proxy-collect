package dao

import "proxy-collect/dao/mysql"

var ProxyDao proxyDaoInterface = mysql.NewMysqlProxyDao()
