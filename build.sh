export GOOS="windows"
export GOARCH="amd64"
go build -o swtr_windows_amd64.exe -ldflags="-s -w" cmd\surfshark-wireguard-tunnel-generator\main.go

export GOOS="linux"
export GOARCH="amd64"
go build -o swtr_linux_amd64 -ldflags="-s -w" cmd\surfshark-wireguard-tunnel-generator\main.go

export GOOS="darwin"
export GOARCH="amd64"
go build -o swtr_darwin_amd64 -ldflags="-s -w" cmd\surfshark-wireguard-tunnel-generator\main.go

export GOOS="darwin"
export GOARCH="arm64"
go build -o swtr_darwin_arm64 -ldflags="-s -w" cmd\surfshark-wireguard-tunnel-generator\main.go