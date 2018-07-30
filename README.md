### gosntp

#### 使用场景
- 内网中部分机器无法连接网络，需要以一台作为 ntp server，其余作为 ntp client，保持集群中机器时间的一致。

#### 使用方式
- go get github.com/csxuejin/sntp
- cd $GOPATH/src/github.com/csxuejin/sntp
- 运行 `make build` 将程序编译为 linux 环境可执行文件
- 设置 config.json ，其中各个字段的含义解释如下：
    - "server_port"  代表 ntp server 运行端口 
    - "server_ip"  代表 ntp server 所在的 ip，client 需要知道该 ip 才能连接
    - "sync_frequency" 代表 ntp client 多久和 server 通信一次(单位为秒)，每次通信后更新系统时间

- 拷贝 config.json 以及可执行文件 gontp 到服务器，如果要作为 ntp server 则运行 `./gontp server`; 如果要作为 ntp client 则运行 `./gontp client`

以上步骤完成后，client 会和 server 保持时间一致。