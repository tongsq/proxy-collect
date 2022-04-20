proxy-collect - GO Proxy Collector
------
### GO语言实现的ip代理池，隧道代理

#### [English README](README.en.md)

功能
------

* 通过api获取可用代理ip列表
* 隧道代理服务：支持开启多个端口监听，支持多协议，多级转发，用户名密码较验，用户限流
* 支持redis或mysql作为存储
* 自动抓取互联网上的免费代理ip
* 定时检测ip可用性，失效ip复检、记录ip存活时间
* 获取ip归属地等信息

## 使用方法：

#### 安装

    git clone https://github.com/tongsq/proxy-collect.git
    cd proxy-collect

#### 启动服务
1、启动脚本收集免费代理ip

    go run ./cmd -S=job -C=conf.yaml
2、启动api服务

    go run ./cmd -S=api -C=conf.yaml
3、启动隧道代理服务

    go run ./cmd -S=tunnel -C=conf.yaml
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
### 隧道代理配置
```
tunnel: 
  tunnelLevel: 1 //代理转发层级
  refresh: 10    //代理池自动刷新时间，单位秒
  debug: false   //是否debug
  strategy: round //指定节点选择策略，round - 轮询，random - 随机, fifo - 自上而下。默认值为round。
  maxFails: 100 //指定节点连接的最大失败次数，当与一个节点建立连接失败次数超过此设定值时，此节点会被标记为死亡节点
  failTimeout: 10 //定死亡节点的超时时间，当一个节点被标记为死亡节点后，在此设定的时间间隔内不会被选择使用
tunnels: //隧道代理监听的端口，可配多个
  -
    proto: http //协议：http、https、socks5等等
    host: 0.0.0.0 //host
    port: 8888 //监听端口
    users:  //授权用户，可配多个, 不开启用户较验不要配这个
      -
        username: root //用户名
        password: 123 //密码
        limiter: 1,10,2 //限流规则,1:1秒内，10：可访问10次，2: 限制2个并发
```
# 待开发
- [X] 支持socket5等其它协议代理采集、较验
- [X] 支持配置日志分级
- [X] 支持开启隧道代理服务：用户名密码验证、限流功能
- [ ] 记录代理响应时间

# 感谢
* 如果您觉得程序还不错，不妨点个星以鼓励一下，后续会更新更多ip源，谢谢！
* 如果您对程序有任何建议和意见或者有更好的免费ip源，欢迎提交issue。
* 感谢大神[@ginuerzh](https://github.com/ginuerzh) 的开源项目提供隧道代理支持：[gost](https://github.com/ginuerzh/gost)
