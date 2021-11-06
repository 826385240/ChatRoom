#!/bin/bash

#新消息最好往后加,尽量保持与原消息的兼容
declare -A HighId=()
declare -A CurMsgId=()
declare -A DstToMsg=()           #用于生成   MSG_XXX_SC = 1000
declare -A PkgToMsgs=()          #用于生成   return &pkg.MSG_XXX_SC{}
declare -A PkgToFiles=()         #用于生成   import的包路径

#==============================可以修改生成消息的相关配置================================
#消息id根据发送方向和序列化方式占用高5位值(高位最大32)
#protobuffer占用高位值的1-10范围,其他占用11-32范围
HighId["TOFU"]=1
HighId["TOSC"]=2
HighId["SC"]=3
HighId["CS"]=4
MAX_PROTO_HIGH=10
MAX_MSG_NUM=2048

#消息回调处理者 格式: 包名.处理者类型
MsgExecutor=("tcptask.TcpTask" "tcpclient.TcpClient")

protoOutDir=protoout
cmdDir=cmdid
rootDir=$(pwd)
protoDir=$rootDir/proto
msgFile=$rootDir/src/$cmdDir/msgId.go
protoPkgPath=github.com/golang/protobuf/proto
#==============================可以修改生成消息的相关配置================================

function writeMsgRange()
{
    printf "//以下为各个消息段的范围\n" >> $msgFile
	for key in ${!HighId[@]}  
	do  
        local high=${HighId[$key]}
        let minId=high*MAX_MSG_NUM
        let maxId=minId+MAX_MSG_NUM-1
        printf "const MIN_$(echo ${key} | tr 'a-z' 'A-Z')_ID = ${minId}\n"  >> $msgFile
        printf "const MAX_$(echo ${key} | tr 'a-z' 'A-Z')_ID = ${maxId}\n\n"  >> $msgFile
    done
    let maxProtoId=(MAX_PROTO_HIGH+1)*MAX_MSG_NUM-1
    printf "//ProtoBuffer消息占用的消息ID的上限\n" >> $msgFile
    printf "const MAX_PROTO_MSG_ID = ${maxProtoId}\n\n"  >> $msgFile
}

function writeMsgBegin()
{
    printf "//***********************生成消息ID***********************\n" >> $msgFile
    printf "const (\n" >> $msgFile
}

function writeMsgEnd()
{
    printf ")\n\n" >> $msgFile
}

function genOneServerMsg()
{
    local dst=$1
    local msgName=$2

    if [ "$dst" != "" -a "$msgName" != "" ];then
        if [ "${CurMsgId[$dst]}" == "" ];then
            CurMsgId["$dst"]=0
        fi

        local curId=CurMsgId["$dst"]

        genOneServerMsgStr $msgName ${HighId["$dst"]} curId 
        DstToMsg["$dst"]="${DstToMsg[$dst]}$msgStr"

        let curId=curId+1
        CurMsgId["$dst"]=$curId
    fi
}

function genOneServerMsgStr()
{
    local msgName=$1
    local highId=$2
    local lastmsgId=$3

    #将消息id高若干位和低若干位部分组合
    let msgId=highId*MAX_MSG_NUM+lastmsgId+1

    msgStr="\t$msgName = $msgId\n"
}

function getMsgDst()
{
    if [ "$1" == "" ];then
        curDst=""
        return
    fi

    curDst=$(echo "$1" | grep -E -o "[a-zA-Z]+$")
}

function preCheck()
{
    if [ "$$(pwd)" == "" ];then
        echo "错误,$(pwd)环境变量没有指定!"
        exit 1
    fi
    if [ ! -d $protoDir ];then 
        echo "错误,找不了proto目录!"
        exit 1
    fi

    cd $protoDir 
    rm -fr $msgFile
}

function genOneFileMsg()
{
    #解析proto文件
    local pkgName=$(cat $protoDir/$1 | grep -E "^\s*package\ +[0-9a-zA-Z]+\s*;$" | grep -E -o "[0-9a-zA-Z]+\s*;$" | grep -E -o "[0-9a-zA-Z]+")
    local msgName=$2

    if [ "$pkgName" != "" ];then
        PkgToMsgs["$pkgName"]="${PkgToMsgs[$pkgName]}${msgName} "
    fi
}

function walkAllFiles()
{
    for protofile in $(ls)
    do
        for msgName in $(cat $protofile | grep -E "^\s*message\s+MSG_\w+_" | grep -E -o "MSG_(_|\w)+")
        do
            #protoc会默认将与数字相连的小写字母转换为大写,比如 11zoo 转换为 11Zoo
            local needDeal=$(echo $msgName | grep -E "[0-9]+[a-z]")
            if [ "$needDeal" != "" ];then
                msgName=$(echo $msgName | sed -n 's/[0-9]\+\w/\U&/ p')
            fi

            #开始生产一个消息
            getMsgDst $msgName
            if [ "$curDst" != "" -a "${HighId[$curDst]}" != "" ];then
                genOneServerMsg $curDst $msgName
                genOneFileMsg $protofile $msgName
            fi
        done
    done
}

function writeToMsgFile()
{
	for key in ${!DstToMsg[@]}  
	do  
        printf "\t//下面定义 $key 后缀的消息\n" >> $msgFile 
		printf "${DstToMsg[$key]}" >> $msgFile 
	done 
}


function analyzePkgs()
{
    for file in $(ls $protoDir)
    do
        local pkgName=$(cat $protoDir/$file | grep -E "^\s*package\ +[0-9a-zA-Z]+\s*;$" | grep -E -o "[0-9a-zA-Z]+\s*;$" | grep -E -o "[0-9a-zA-Z]+")
        
        if [ "$pkgName" != "" ];then
            PkgToFiles["$pkgName"]="${PkgToFiles[$pkgName]}$file "
        fi
    done
}

function writeImportPkgBegin()
{
    printf "//***********************导入依赖的包***********************\n" >> $msgFile
    printf "package cmd\n\n" >> $msgFile
    printf "import (\n" >> $msgFile 
    printf "\t\"$protoPkgPath\"\n" >> $msgFile 
}

function writeImportPkgEnd()
{
    printf ")\n\n" >> $msgFile 
}

function writeImportMsgPkg()
{
	for key in ${!PkgToFiles[@]}  
	do  
		printf "\t\"$protoOutDir/$key\"\n" >> $msgFile 
    done
}

function writeGenAndConvertMsg()
{
    #根据消息id生成proto消息对象函数
    printf "//GenMsgById函数通过消息id生成对应的proto对象\n" >> $msgFile 
    printf "func GenMsgById(msgId uint16) (proto.Message, unsafe.Pointer) {\n" >> $msgFile 
    printf "\tswitch msgId {\n" >> $msgFile 
	
	for pkg in ${!PkgToMsgs[@]}  
	do  
        for msgName in ${PkgToMsgs[$pkg]}
        do
            printf "\tcase ${msgName}:\n" >> $msgFile
            printf "\t\t{\n" >> $msgFile
            printf "\t\t\tmsgPtr := &${pkg}.${msgName}{}\n" >> $msgFile 
            printf "\t\t\treturn msgPtr, unsafe.Pointer(msgPtr)\n" >> $msgFile
            printf "\t\t}\n" >> $msgFile
        done
    done
    printf "\t}\n" >> $msgFile
    printf "\treturn nil, nil\n}\n\n" >> $msgFile 

    #根据消息id转换指针为proto对象函数
    printf "//ConvertMsgById函数通过消息id生成对应的proto对象\n" >> $msgFile 
    printf "func ConvertMsgById(msgId uint16, msgPtr unsafe.Pointer) proto.Message {\n" >> $msgFile 
    printf "\tswitch msgId {\n" >> $msgFile 
	
	for pkg in ${!PkgToMsgs[@]}  
	do  
        for msgName in ${PkgToMsgs[$pkg]}
        do
            printf "\tcase ${msgName}:\n" >> $msgFile
            printf "\t\treturn (*${pkg}.${msgName})(msgPtr)\n" >> $msgFile
        done
    done
    printf "\t}\n" >> $msgFile
    printf "\treturn nil\n}\n\n" >> $msgFile 
}

function writeImportHandlerPkg(){
	printf "\t\"lib/common\"\n" >> $msgFile 
    printf "\t\"lib/handler\"\n" >> $msgFile 
}

function writeImportMsgCBPkg()
{
    cd $rootDir/src

    for executor in ${MsgExecutor[*]} 
    do
        local pkgName=$(echo $executor | grep -E -o "^[0-9a-zA-Z]+")
        local exeName=$(echo $executor | grep -E -o "[0-9a-zA-Z]+$")
        #解析go文件
        for goFile in $(find ./lib ./mainserver | grep -E  "\.go$") 
        do
            local curPkg=$(cat $goFile | grep -E "^\s*package\ +[0-9a-zA-Z]+\s*$" | grep -E -o "[0-9a-zA-Z]+\s*$" | grep -E -o "[0-9a-zA-Z]+")
            if [ "$curPkg" != "" -a "$curPkg" == "$pkgName" ];then
                local exist=$(cat $goFile | grep -E -o "^\s*type\s+$exeName\s+struct")
                if [ "$exist" != "" ];then
                    #获得消息回调者的包路径
                    local importPath=$(echo $goFile | sed "s/^\.\///g" | sed "s/\w\+\.go$//g" | sed "s/\/$//g")
                    printf "\t\"$importPath\"\n" >> $msgFile 

                    break
                fi
            fi
        done
    done
    printf "\t\"unsafe\"\n" >> $msgFile 

    cd $rootDir
}

function writeMsgHandler(){
    printf "//***********************消息回调处理***********************\n" >> $msgFile

    printf "var msgCallBackHandler *handler.Handler\n\n" >> $msgFile
    printf "func InitCbHandler() *handler.Handler {\n" >> $msgFile
	printf "\tmsgCallBackHandler = handler.NewHandler()\n" >> $msgFile
    printf "\treturn msgCallBackHandler\n" >> $msgFile
    printf "}\n\n" >> $msgFile


    printf "func ExecCallBack(u unsafe.Pointer, m com.MsgToLogicPtr) bool {\n" >> $msgFile
	printf "\tif m != nil {\n" >> $msgFile
    printf "\t\tif m.MsgId <= MAX_PROTO_MSG_ID {\n" >> $msgFile
    printf "\t\t\treturn msgCallBackHandler.ExecHandler(m.MsgId, u, m.ProtoPtr)\n" >> $msgFile
    printf "\t\t} else {\n" >> $msgFile
    printf "\t\t\treturn msgCallBackHandler.ExecHandler(m.MsgId, u, unsafe.Pointer(&m.MsgPtr.Data[0]))\n" >> $msgFile
    printf "\t\t}\n" >> $msgFile
	printf "\t}\n" >> $msgFile
	printf "\treturn false\n" >> $msgFile
    printf "}\n\n" >> $msgFile
}

function writeMsgExecFlag(){
    printf "const (\n" >> $msgFile

    local curId=1
    for executor in ${MsgExecutor[*]} 
    do
        local exeName=$(echo $executor | grep -E -o "[0-9a-zA-Z]+$" | tr 'a-z' 'A-Z')
        if [ "$exeName" != "" ];then
            printf "\tEXEC_FLAG_${exeName} = $curId\n" >> $msgFile
            let curId=curId*2
        fi
    done

    printf ")\n\n" >> $msgFile
}

function writeConvertConn(){
    printf "func ConvertConnById(o com.ConnToLogicPtr) com.IBaseTcpConn {\n" >> $msgFile
    for executor in ${MsgExecutor[*]} 
    do
        local exeName=$(echo $executor | grep -E -o "[0-9a-zA-Z]+$")
        local upper=$(echo $exeName | tr 'a-z' 'A-Z')
        printf "\tif o.ConnType == EXEC_FLAG_${upper} {\n" >> $msgFile
        printf "\t\treturn (*${executor})(o.Conn)\n" >> $msgFile
        printf "\t}\n" >> $msgFile
    done
	printf "\treturn nil\n" >> $msgFile
    printf "}\n\n" >> $msgFile
}

function writeMsgCallBack()
{
	for pkg in ${!PkgToMsgs[@]}  
	do  
        local msgs=${PkgToMsgs[$pkg]}
        for msgName in ${msgs} 
        do
            printf "\n//以下是 ${msgName} 相关的回调处理代码\n" >> $msgFile 
            printf "type ${msgName}_CB struct {\n" >> $msgFile 
            printf "\texecFlag          int\n" >> $msgFile 
            #声明Functor的所有可能的函数
            for executor in ${MsgExecutor[*]} 
            do
                local exeName=$(echo $executor | grep -E -o "[0-9a-zA-Z]+$")
                printf "\t${exeName}RealExec func(u *${executor}, m *${pkg}.${msgName})\n" >> $msgFile 
            done
            printf "}\n\n" >> $msgFile 

            #生成执行此回调的所有可能函数
            for executor in ${MsgExecutor[*]} 
            do
                local exeName=$(echo $executor | grep -E -o "[0-9a-zA-Z]+$")
                local upper=$(echo $exeName | tr 'a-z' 'A-Z')
                printf "func (this *${msgName}_CB) ${exeName}Exec(u unsafe.Pointer, m unsafe.Pointer) {\n" >> $msgFile 
                printf "\texePtr := (*${executor})(u)\n" >> $msgFile 
                printf "\tmsgPtr := (*${pkg}.${msgName})(m)\n" >> $msgFile 
                printf "\tthis.${exeName}RealExec(exePtr, msgPtr)\n" >> $msgFile 
                printf "}\n" >> $msgFile 

                printf "\nfunc Reg_${exeName}_${msgName}(f func(u *${executor}, m *${pkg}.${msgName})) {\n"  >> $msgFile 
	            printf "\tbExist := msgCallBackHandler.IsExist($msgName)\n"  >> $msgFile 
	            printf "\tvar cb *${msgName}_CB = nil\n"  >> $msgFile 
	            printf "\tif bExist {\n"  >> $msgFile 
		        printf "\t\ti := msgCallBackHandler.GetHandler(${msgName})\n"  >> $msgFile 
		        printf "\t\tcb = i.(*${msgName}_CB)\n"  >> $msgFile 
		        printf "\t\tcb.${exeName}RealExec = f\n"  >> $msgFile 
	            printf "\t} else {\n"  >> $msgFile 
	            printf "\t\tcb = &${msgName}_CB{${exeName}RealExec: f}\n"  >> $msgFile 
	            printf "\t}\n"  >> $msgFile 
                printf "\tcb.execFlag = cb.execFlag | EXEC_FLAG_${upper}\n" >> $msgFile
	            printf "\tmsgCallBackHandler.RegHandler(${msgName}, cb)\n"  >> $msgFile 
                printf "}\n\n"  >> $msgFile 
            done

            #生成回调函数的调度函数Exec
            printf "func (this *${msgName}_CB) Exec(u unsafe.Pointer, m unsafe.Pointer) {\n" >> $msgFile 
            for executor in ${MsgExecutor[*]} 
            do
                local exeName=$(echo $executor | grep -E -o "[0-9a-zA-Z]+$")
                local upper=$(echo $exeName | tr 'a-z' 'A-Z')
                printf "\tif this.execFlag&EXEC_FLAG_${upper} > 0 {\n" >> $msgFile 
                printf "\t\tthis.${exeName}Exec(u, m)\n" >> $msgFile 
                printf "\t\treturn\n" >> $msgFile 
                printf "\t}\n"  >> $msgFile 
            done
            printf "}\n" >> $msgFile 
        done
    done
}

function goInstallPkgs()
{
    go fmt $msgFile 
    go install $cmdDir
}

function main()
{
    local beginTime=$(date +%s)
    #临时保存数据变量
    msgStr=""
    curDst=""

    preCheck

    #主要的分析函数
    analyzePkgs
    walkAllFiles                

    #写入导入的包
    writeImportPkgBegin
    writeImportHandlerPkg
    writeImportMsgCBPkg
    writeImportMsgPkg
    writeImportPkgEnd

    #写入消息
    writeMsgBegin
    writeToMsgFile
    writeMsgEnd
    writeMsgRange

    #写入根据id获取消息代码
    writeGenAndConvertMsg

    #写入回调管理器
    writeMsgHandler
    writeMsgExecFlag
    writeConvertConn
    writeMsgCallBack

    #编译包文件
    goInstallPkgs

    local endTime=$(date +%s)
    echo "生成消息id完毕,耗时$(expr ${endTime} - ${beginTime})秒!"
}

main
