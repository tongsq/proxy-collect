# free proxy ip collector

# 免费代理ip收集器、爬虫代理ip池

## Quick Start（使用方法）：

1、get code

    git clone https://github.com/tongsq/proxy-collect.git
    cd proxy-collect
2、start daemon job collect free ip (启动脚本收集免费代理ip)

    go run job.go
3、start api server (启动api服务)

    go run api.go
4、get proxy ip

    curl 127.0.0.1:8090/all?city=上海&duration=100

## Config your yaml (yaml配置)
## Config storage media (设置存储媒介)
### 一、use redis as storage (数据存储到redis)
    
    dao: redis
    redis:
      MaxIdle: 10
      MaxActive: 20
      Network: tcp
      Address: 127.0.0.1:6379
      Password: your password
      
### 二、use mysql as storage (数据存储到mysql)

1、config

    dao: database
    database:
      Dialect: mysql
      Args: user:password@(127.0.0.1:3306)/dbname?charset=utf8&loc=Local
    
2、create table

    CREATE TABLE `proxy` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `host` varchar(255) NOT NULL DEFAULT '',
      `port` int(11) NOT NULL DEFAULT '0',
      `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '0:无效，1：有效',
      `create_time` int(11) NOT NULL DEFAULT '0',
      `update_time` int(11) NOT NULL DEFAULT '0',
      `active_time` int(11) NOT NULL DEFAULT '0',
      `country` varchar(100) NOT NULL DEFAULT '',
      `region` varchar(100) NOT NULL DEFAULT '',
      `city` varchar(100) NOT NULL DEFAULT '',
      `isp` varchar(255) NOT NULL DEFAULT '',
      PRIMARY KEY (`id`) USING BTREE,
      UNIQUE KEY `IDX_HOST_PORT` (`host`,`port`) USING BTREE,
      KEY `IDX_STATUS` (`status`) USING BTREE,
      KEY `IDX_ACTIVE_TIME` (`active_time`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT
