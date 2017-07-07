#!/bin/bash
while true;
do
        #需要定时执行的命令，注意结尾的&，进入后台执行，不阻塞
        /xxpath/bin/php /yypath/test.php >> /dev/null 2>&1 &
        #每秒执行
        sleep 1;
done
