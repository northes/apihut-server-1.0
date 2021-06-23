echo "== 开始构建 =="
if test -e apihut-server; then
  echo '文件已存在!'
  read -p "Exit?" input
  exit 0
fi
echo "-> 设置环境变量..."
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
echo $GOOS"-"$GOARCH
echo "-> 编译..."
go build ./
echo "-> 压缩..."
upx --best -k apihut-server
echo "-> 复制配置..."
if [ ! -d "./build/prod/config/" ]; then
  mkdir -p "./build/prod/config/"
fi
cp -u apihut-server ./build/prod/apihut-server
cp -u -p ./config/apihut.yml ./build/prod/config/apihut.yml
echo "-> 复制静态文件..."
cp -u -r -p ./data ./build/prod/data
cp -u -r -p ./templates ./build/prod/templates
cp -u -r -p ./static ./build/prod/static
echo "-> 删除产出文件..."
rm apihut-server
rm apihut-server.~
printf "\n == Done! == \n"
read -r -p "Exit?" input
cd ./build/prod || exit
start .
