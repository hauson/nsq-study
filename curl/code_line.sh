# bin/bash

total=0
gofiles=`find . -name "*.go"`
for file1 in $gofiles
do
  num=`wc -l $file1 | awk -F' ' '{print $1}'`
  total=$(($total+$num))
done

echo $total
