build:
	GOOS=linux GOARCH=amd64 go build -o gontp main.go
	scp config.json root@vm3:/root/ntp
	scp gontp root@vm3:/root/ntp
