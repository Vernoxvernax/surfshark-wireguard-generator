export GOOS="windows"
export GOARCH="amd64"
CGO_ENABLED=1 go build -o swtr_${GOOS}_${GOARCH}.exe -ldflags="-s -w" cmd/surfshark-wireguard-tunnel-generator/main.go

export GOOS="linux"
export GOARCH="amd64"
CGO_ENABLED=1 go build -o swtr_${GOOS}_${GOARCH} -ldflags="-s -w" cmd/surfshark-wireguard-tunnel-generator/main.go

export GOOS="linux"
export GOARCH="arm64"
go build -o swtr_${GOOS}_${GOARCH} -ldflags="-s -w" cmd/surfshark-wireguard-tunnel-generator/main.go

export GOOS="darwin"
export GOARCH="amd64"
CGO_ENABLED=1 go build -o swtr_${GOOS}_${GOARCH} -ldflags="-s -w" cmd/surfshark-wireguard-tunnel-generator/main.go

export GOOS="darwin"
export GOARCH="arm64"
CGO_ENABLED=1 go build -o swtr_${GOOS}_${GOARCH} -ldflags="-s -w" cmd/surfshark-wireguard-tunnel-generator/main.go
