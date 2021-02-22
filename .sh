#!/usr/bin/env bash

name="kf001"
dir_for()
{
  if [ $1 == go ]; then
    count=0
    declare -a arr
    for file in in ./cmd/${name}/*
      do
      if [ -d "$file" ]
      then
        if [ ${file#./cmd/${name}/} != "plugin" ]; then
          echo "   $count)  ${file#./cmd/${name}/}"
          arr[$count]=${file#./cmd/${name}/}
          count=$[$count + 1]
        fi
      fi
      done
    read -e -p "选择要运行的程序编号... " num
    if [ ${arr[$num]} != "" ]; then
      echo "开始启动 ..."
      cd ./cmd/${name}/${arr[$num]}

      echo "go watch---"

      gowatch -o ../../../tmp/${arr[$num]}"Main" -p ./main.go -args="--config=config-dev.conf"
   fi
 elif [ $1 == proto ]; then
    declare -a arr
    for file in in ./model/proto/*
    do
      if [ -f "$file" ]
      then
        arr[$count]=${file#./model/proto/}
        count=$[$count + 1]
      fi
    done
    cd ./model/proto/
    # shellcheck disable=SC2068
    for var in ${arr[@]}
    do
      array=(${var//_/ })
      echo "开始转换 ..."
      if [ ${array[0]} != "" ]; then
        sudo mkdir -p -m 755 ../imp/${array[0]}
        echo "protoc --go_out=plugins=tarsrpc:../imp/${array[0]} $var"
        protoc --go_out=plugins=tarsrpc:../imp/${array[0]} $var
      fi
      echo $var
    done
 elif [ $1 == gormt ]; then
    read -e -p "输入要转换的数据库表名... " num
    cd ./model/
    gormt.exe -d $num
 elif [ $1 == swag ]; then
    count=0
    declare -a arr
    for file in in ./cmd/${name}/*
      do
      if [ -d "$file" ]
      then
        echo "   $count)  ${file#./cmd/${name}/}"
        arr[$count]=${file#./cmd/${name}/}
        count=$[$count + 1]
      fi
      done
    read -e -p "选择要生成文档的程序编号... " num
    if [ ${arr[$num]} != "" ]; then
      cd ./cmd/${name}/${arr[$num]}
      swag init --parseDependency
    fi
 fi
}
# gowatch -o ./cmd/${name}/clientUp/admin.exe -p ./cmd/${name}/clientUp/main.go
# swag init  --dir=./cmd/${name}/clientUp/main.go --parseDependency=true
echo 'QINWONG GO 便捷调试'
echo "   1) 启动GO服务"
echo "   2) proto到go"
# echo "   3) 生成数据库模型"
echo "   4) 生成swag文档"
echo "   0) 退出"
read -p "填写数字然后按回车: " -e option_1
case $option_1 in
			1)
			  dir_for "go"
				exit
				;;
			2)
        dir_for "proto"
				exit
				;;
			3)
        dir_for "gormt"
				exit
				;;
      4)
        dir_for "swag"
				exit
				;;
			0) exit;;
		esac
		exit;;
esac