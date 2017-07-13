#!/bin/bash
while true;
do
        #需要定时执行的命令
        process_name='/xxpath/bin/php /yypath/test.php'
        #进程数检查，避免异常情况阻塞启用更多的进程
        process_num=$(ps -ef|grep "$process_name"|grep -v grep|wc -l)
        echo "Process Count: $process_num"
        if [ $process_num -gt 1 ]
        then
            ps -ef|grep "$process_name"|grep -v grep|awk '{print $2}'|xargs kill -USR1
            echo "Killed: $process_num"
        fi
        #注意结尾的&，进入后台执行，不阻塞
        $($process_name >> /dev/null 2>&1 &)
        #每秒执行
        sleep 1;
done
