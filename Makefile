g:
	git add .
	git commit -m"自动提交 git 代码"
	git push

dev:
	make build && make run

build:
	protoc -I . --gogofaster_out=plugins=micro:. proto/goods/goods.proto
