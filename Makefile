.PHONY: start

# 开启项目
.PHONY: start
start:
	@sh ./scripts/start.sh

# 编译框架
.PHONY: frame
frame:
	@sh ./scripts/frame.sh

# 生成gRPC相关文件
.PHONY: grpc
grpc:
	@buf format -w api/proto
	@buf lint api/proto
	@buf generate api/proto

# 清理项目依赖
.PHONY: tidy
tidy:
	@go mod tidy -v
