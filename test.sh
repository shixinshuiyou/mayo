# !/bin/bash
jobName='log_user_action_active_user_to_mysql'
depJobName='log_user_action_active_user'

mysqlHost='10.132.94.14'
mysqlPort='13310'
mysqlDbname='uas'
mysqlUsername='uas'
mysqlPassword='67f635273ef0ec1d'
mysqlTable="hive_log_user_action_active_user"


if [ $# -eq 0 ];
then
        d=`date '+%Y%m%d' -d '-24 hours' `
else
        d=$1
fi

create_date=$(date -d ${d} +'%Y-%m-%d')

while :
do
	if !  /usr/bin/hadoop/software/hadoop/bin/hadoop fs -test -e "/home/hdp-gfjb-wxcard/job/sync/$d/$depJobName.done"
        then
                echo "$depJobName data of $d doesn't exist"
                sleep 60
        else
                break
        fi
done

/home/hdp-gfjb-wxcard/software/sqoop/bin/sqoop eval --connect jdbc:mysql://$mysqlHost:$mysqlPort/$mysqlDbname --username $mysqlUsername --password $mysqlPassword --query "delete from $mysqlTable Where create_date = '$create_date'"

exitCode=$?
if [ $exitCode -ne 0 ];then
         echo "[ERROR] hive execute failed!"
         exit $exitCode
fi

/home/hdp-gfjb-wxcard/software/sqoop/bin/sqoop export --connect jdbc:mysql://$mysqlHost:$mysqlPort/$mysqlDbname --username $mysqlUsername --password $mysqlPassword  --table $mysqlTable --export-dir /home/hdp-gfjb-wxcard/hive/warehouse/hdp_gfjb_wxcard.db/$depJobName/date_id=$d --input-fields-terminated-by '\0001' --columns id,uuid,create_time,entity_id,create_date,scene

exitCode=$?
if [ $exitCode -ne 0 ];then
         echo "[ERROR] hive execute failed!"
         exit $exitCode
else
	if ! /usr/bin/hadoop/software/hadoop/bin/hadoop fs -test -e "/home/hdp-gfjb-wxcard/job/sync/$d/$jobName.done"
	then
	/usr/bin/hadoop/software/hadoop/bin/hadoop fs -touchz "/home/hdp-gfjb-wxcard/job/sync/$d/$jobName.done"
        fi
fi
