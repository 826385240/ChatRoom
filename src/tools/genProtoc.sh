
protoOutDir=protoout
outDir=./src/$protoOutDir
protoDir=./proto

cd $GOPATH
files=$(ls $protoDir)
declare -A pkgToFiles=()

function preCheck()
{
    if [ "$GOPATH" == "" ];then
        echo "错误,GOPATH环境变量没有指定!"
        exit 1
    fi
    if [ ! -d $protoDir ];then
        echo "错误,找不到存放协议的目录!"
        exit 1
    fi

    if [ ! -d $outDir ];then
        mkdir $outDir 
    fi
}

function analyzePkgs()
{
    for file in $(echo $files)
    do
        local pkgName=$(cat $protoDir/$file | grep -E "^\s*package\ +[0-9a-zA-Z]+\s*;$" | grep -E -o "[0-9a-zA-Z]+\s*;$" | grep -E -o "[0-9a-zA-Z]+")
        
        if [ "$pkgName" != "" ];then
            pkgToFiles["$pkgName"]="${pkgToFiles[$pkgName]}$file "
        fi
    done
}

function checkDirs()
{
	for pkg in ${!pkgToFiles[@]}  
	do  
        if [ ! -d $outDir/$pkg ];then
            mkdir $outDir/$pkg 
        fi

        if [ ! -d $protoDir/$pkg ];then
            mkdir $protoDir/$pkg 
        fi
    done
}

function genProto()
{
	for pkg in ${!pkgToFiles[@]}  
	do  
        local pkgFils=${pkgToFiles[$pkg]}
        for file in $(echo $pkgFils)
        do
            cp -fr $protoDir/$file $protoDir/$pkg/
            protoc --go_out=$outDir/$pkg $protoDir/$pkg/*.proto

            local genOutPath=$outDir/$pkg/proto
            if [ -d $genOutPath ];then
                mv $genOutPath/$pkg/* $outDir/$pkg
                rm -fr $genOutPath
            fi
        done
    done

    #目录是空的,则删除
	for pkg in ${!pkgToFiles[@]}  
    do
        local filesnum=$(ls $outDir/$pkg | wc -l)
        if [ $filesnum -eq 0 ];then
            rm -fr $outDir/$pkg
        fi

        if [ -d $protoDir/$pkg ];then
            rm -fr $protoDir/$pkg
        fi
    done
}

function goInstallPkgs()
{
	for pkg in ${!pkgToFiles[@]}  
    do
        go install -gcflags "-N -l" $protoOutDir/$pkg
    done
}

function main()
{
    preCheck
    analyzePkgs
    checkDirs

    #生效消息协议
    genProto
    #go编译生成包文件
    goInstallPkgs

    echo "生成和编译proto代码完毕!"
}

main
