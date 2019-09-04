

TO_FILE=rpc.proto ## 生成的文件

SRC_FILE=./proto/proto.go ## go源文件

PACKAGE_NAME=micro ## rpc包名

echo '开始'

gotoproto -toFileSrc $TO_FILE -src $SRC_FILE -packageName $PACKAGE_NAME

echo '完成'