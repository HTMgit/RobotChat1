#!/bin/sh

ulimit -n 1024000
ulimit -c unlimited

p='robot_chat'

KillServer()
{
    pid=`ps x | grep "$p" | sed -e '/mykill/d' | sed -e '/grep/d' | sed -e '/tail/d' | awk '{print $1}' `
    pid=`echo $pid | awk '{print $1}'`
    while [ ! -z "$pid"]
    do
        kill -INT $pid
        sleep 5
        pid=`ps s | grep "$p" | sed -e '/grep/d' | sed -e '/tail/d' |awk '{print $1}'`
        pid=`echo $pid | awk '{print $1}'`
    done
}

case $1 in 
    start)
        KillServer
        sleep 1
        nohup ./$p --config ./config.toml >> ./out.log 2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p | grep -v grep
        ;;
    stop)
        KillServer
        sleep 1
        echo ""
        ps -elf | grep $p | grep -v grep
        ;;
    restart)
        KillServer
        sleep 1
        nohup ./$p --config ./config.toml  >> ./out.log  2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p | grep -v grep
        ;;
    *)
        KillServer
        sleep 1
        nohup ./$p --config ./config.toml  >> ./out.log  2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p | grep -v grep
        ;;
esac