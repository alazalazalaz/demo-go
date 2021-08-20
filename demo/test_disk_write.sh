#!/bin/bash

datetime1=`date "+%s#%N"`
beginSeconds=`echo $datetime1 | cut -d"#" -f1`   #取出秒
beginNanoseconds=`echo $datetime1 | cut -d"#" -f2`   #取出纳秒

n=0
mb=10
while(( $n<$mb*1024 ))
do
    let "n++"
    #每次输入1KB
    echo "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" >> test_disk_file.log
done

datetime2=`date "+%s#%N"`
endSeconds=`echo $datetime2 | cut -d"#" -f1`   #取出秒
endNanoseconds=`echo $datetime2 | cut -d"#" -f2`   #取出纳秒

diffNs=$((($endSeconds-$beginSeconds)*1000000000+(endNanoseconds-beginNanoseconds)))
echo "${diffNs}ns"
diffMs=$(($diffNs/1000000))
echo "${diffMs}ms"

#mb=10
#diffNs=6370887
diffSeconds=`awk 'BEGIN{printf "%.6f\n",'$diffNs'/1000000000}'`
iops=`awk 'BEGIN{printf "%.2f\n",'$mb'/'$diffSeconds'}'`
echo "${iops}MB/s"
