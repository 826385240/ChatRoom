#!/bin/bash

check_dirs='
./
'

function get_sub_dirs()
{
    if [ ! $1 = "" ];then
        find $1 -type d 
    else
        echo "错误,不能输入空目录!"
    fi
}

function get_sub_dirs_by_depth()
{
    if [ ! $# -lt 2 ];then
        find $1 -maxdepth $2 -type d | sed 's/\.\///g' | sed "/\(^\W\|^$\)/d" 
    else
        echo "错误,必须输入遍历子目录的层级!"
    fi
}


function svn_st()
{
    for dir in ${check_dirs};do
        svn st ${dir} | grep -E "(\.proto$|\.go$|\.sh$)" | sed  '/^?/d'
    done
}


echo '+-------------------------------------------------------------------+'
echo '|修改的文件                                                         |'
echo '+-------------------------------------------------------------------+'

echo $(svn_st | sed 's/^[MACD!]//g')

echo '+-------------------------------------------------------------------+'
echo '|详细列表                                                           |'
echo '+-------------------------------------------------------------------+'

svn_st 

echo '+-------------------------------------------------------------------+'
echo '|新增文件                                                           |'
echo '+-------------------------------------------------------------------+'

svn st | grep "\(\.sh\>$\|\.go\>$\|\.proto\>$\|\.cpp\>$\|\.h\>$\)" | grep '^?' | sed 's/^?//g'

echo '+--------------------------------+------------------+---------------+'
echo '|代码目录                        | 添加的行数       | 文件数        |'
echo '+--------------------------------+------------------+---------------+'

totallines='0'
totalfiles='0'
for dir in ${check_dirs}; do
    for subdirs in $(get_sub_dirs_by_depth ${dir} 1); do
        linenum=$(svn diff ${dir}/${subdirs} | sed '/^+/! d' | wc -l)
        filenum=$(svn st ${dir}/${subdirs} | grep -v '^?' | wc -l)

        if [ ${linenum} -gt 0  ];then
            printf "| %-30s | %-16s | %-13s |\n" ${subdirs} ${linenum} ${filenum}
        fi

        totallines=$(expr ${totallines} + ${linenum})
        totalfiles=$(expr ${totalfiles} + ${filenum})
    done
done

printf "| %-32s | %-16s | %-13s |\n" "总计" ${totallines} ${totalfiles}
echo '+--------------------------------+------------------+---------------+'
