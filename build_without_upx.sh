echo "== 开始构建 =="
echo "-> 设置环境变量..."
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
echo $GOOS"-"$GOARCH
echo "-> 编译..."
go build ./
echo "== Done! =="
read -r -p "Exit?" input