.PHONY: fmt proxy

fmt:
	@find . -name '*.go' -not -path "./vendor/*" | xargs gofumpt -w -extra
	@find . -name '*.go' -not -path "./vendor/*" | xargs -n 1 -t goimports-reviser -rm-unused -set-alias -company-prefixes "github.com/sliveryou" -project-name "github.com/sliveryou/submail-go-sdk"
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shfmt -w -s -i 2 -ci -bn -sr

proxy:
	@go env -w GO111MODULE="on"
	@go env -w GOPROXY="https://goproxy.cn,direct"
