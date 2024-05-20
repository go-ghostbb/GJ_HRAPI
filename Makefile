pack:
	@echo "打包靜態文件..."
	@echo y| gb pack resource,manifest/i18n internal/packed/data.go -n=packed
build:
	@echo "開始編譯..."
	@go mod tidy
	@go build -o bin/hrapi main.go
build-dev:
	@echo "開始編譯(dev)..."
	@go mod tidy
	@go build -o bin/hrapi-dev main.go
docker:
	@echo "開始打包 Docker Image - hrapi"
	@make pack
	@docker buildx build --platform linux/amd64 -t giajiu/hrapi -f ./manifest/docker/Dockerfile .
dlv-dev:
	@dlv --listen=:12366 --headless=true --api-version=2 --continue --accept-multiclient exec bin/hrapi-dev
dev:
	@air
gen:
	@go run internal/cmd/gen/gen.go
docker-dev:
	@echo "開始打包 Docker Image - hrapi"
	@make pack
	@docker buildx build --platform linux/amd64 -t giajiu/hrapi:dev -f ./manifest/docker/Dockerfile .
	@docker push giajiu/hrapi:dev
