#!/bin/bash

libDir=lib
srcDir=$GOPATH/src
libPath=$srcDir/$libDir

function installLibs()
{
    local libs=$(ls $libPath)
    for lib in $(echo $libs)
    do
        go install -gcflags "-N -l" $libDir/$lib
    done
}

function installAll()
{
    installLibs
    go install -gcflags "-N -l" cmdid 
    go install -gcflags "-N -l" mainserver 
    go install -gcflags "-N -l" mainclient 
}

function preCheck()
{
    if [ "$GOPATH" == "" ];then
        echo "错误,GOPATH环境变量没有指定!"
        exit 1
    fi
}

function main()
{
    preCheck
    installAll
}

main
