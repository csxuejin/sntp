build:
	GOOS=linux GOARCH=amd64 go build -o gontp main.go
	scp config.json root@vm2:/root/ntp
	tar czvf gontp.tar gontp
	scp gontp.tar root@vm2:/root/ntp
	ssh root@vm2 "cd /root/ntp && tar zxvf gontp.tar && rm -rf gontp.tar"
	rm -rf gontp.tar
