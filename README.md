proxy-collect - GO Proxy Collector
------
### GO语言实现的ip代理池

#### [English README](README.en.md)

功能
------

* 支持redis或mysql作为存储
* 自动抓取互联网上的免费代理ip
* 定时检测ip可用性，失效ip复检、记录ip存活时间
* 获取ip归属地等信息
* 通过api获取可用代理ip列表

## 使用方法：

#### 安装

    git clone https://github.com/tongsq/proxy-collect.git
    cd proxy-collect

#### 启动服务
1、启动脚本收集免费代理ip

    go run ./cmd -S=job -C=conf.yaml
2、启动api服务

    go run ./cmd -S=api -C=conf.yaml

#### 通过接口获取可用ip

    curl 127.0.0.1:8090/all?city=上海&duration=100

#### 启动多个服务

    go run ./cmd -S=job -S=api -C=conf.yaml
#### 一键启动

    go run ./cmd -S=all -C=conf.yaml

## yaml配置
### 设置存储媒介
A、数据存储到redis

* 修改配置文件conf.yaml如下
```
    dao: redis
    redis:
      MaxIdle: 10
      MaxActive: 20
      Network: tcp
      Address: 127.0.0.1:6379
      Password: your password
``` 
B、数据存储到mysql

* 修改配置文件conf.yaml如下
```
    dao: database
    database:
      Dialect: mysql
      Args: user:password@(127.0.0.1:3306)/dbname?charset=utf8&loc=Local
```
* 创建表
```
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
      `check_count` int(11) NOT NULL DEFAULT '10',
      `source` varchar(50) NOT NULL DEFAULT '',
      `proto` varchar(20) NOT NULL DEFAULT 'http',
      `user` varchar(50) NOT NULL DEFAULT '',
      `password` varchar(50) NOT NULL DEFAULT '',
      PRIMARY KEY (`id`) USING BTREE,
      UNIQUE KEY `IDX_HOST_PORT_PROTO` (`host`,`port`, `proto`) USING BTREE,
      KEY `IDX_STATUS` (`status`) USING BTREE,
      KEY `IDX_ACTIVE_TIME` (`active_time`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT
```
# 待开发
- [X] 支持socket5等其它协议代理采集、较验
- [X] 支持配置日志分级
- [ ] 支持开启隧道代理服务：用户名密码验证、限流功能

#感谢
* 如果您觉得程序还不错，不妨点个星以鼓励本人继续努力，后续会更新更多ip源，谢谢！
* 如果您对程序有任何建议和意见或者有更好的免费ip源，也欢迎提交issue。
