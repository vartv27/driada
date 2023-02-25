#!/usr/bin/osascript

#/etc crontab -e
# * * * * * /Users/baseuser/Desktop/code/notification.sh
#/Users/baseuser/Desktop/code/notification_enter.sh   automator

#DD_0=`date '+ %A, %B %d, %Y. %T '`
#echo $DD_0 >> /Users/baseuser/Desktop/code/cron_log.txt

#
#while read y
#do
#fread=$y
#done < /Users/baseuser/Desktop/code/cron_driada.txt
#echo $fread
#if [ $fread == "already_open" ]
#then
#nothing=1
#else
#echo "already_open" > /Users/baseuser/Desktop/code/cron_driada.txt
#fi
bt="result"
cd /Users/romanstr/Downloads/go/driada/
STT=`/usr/local/go/bin/go run hello.go`
#say $STT
#DD_1=`date '+ %A, %B %d, %Y.'`
say " сегодня" $DD_1;

VAR1='"'$STT'"'
echo $VAR1
bt=`osascript -e 'tell app "System Events" to display dialog '"$VAR1"' buttons {"Повторить", "Новая задача", "ok"} default button 1  '`
echo $bt

while [ "$bt" = "button returned:Повторить" ];
do
    say $STT
    bt=`osascript -e 'tell app "System Events" to display dialog '"$VAR1"' buttons {"Повторить", "Новая задача", "ok"} default button 1  '`
done

google="https://docs.google.com/spreadsheets/d/1PKORfLXcTv9CK1gBbPPO4RDY0_XJA6n83Mgz2nO0oy4/edit#gid=1160736350"
if [ "$bt" = "button returned:Новая задача" ]; then
    open $google
fi