.PHONY: all clean

all: out/checkin-linux-386 out/checkin-linux-amd64 out/checkin-darwin-amd64 out/checkin-windows-386 out/checkin-windows-amd64

clean:
	rm -rf out

out/checkin-linux-386: main.go go.mod go.sum **/*.go
	GOOS=linux GOARCH=386 go build -i -o $@

out/checkin-linux-amd64: main.go go.mod go.sum **/*.go
	GOOS=linux GOARCH=amd64 go build -i -o $@

out/checkin-darwin-amd64: main.go go.mod go.sum **/*.go
	GOOS=darwin GOARCH=amd64 go build -i -o $@

out/checkin-windows-386: main.go go.mod go.sum **/*.go
	GOOS=windows GOARCH=386 go build -i -o $@

out/checkin-windows-amd64: main.go go.mod go.sum **/*.go
	GOOS=windows GOARCH=amd64 go build -i -o $@
