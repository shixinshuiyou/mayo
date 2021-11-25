# !/bin/bash
jobName='log_user_action_active_user'
depJobName='log_user_action'

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

sql=$(
cat <<EOF
alter table log_user_action_active_user drop partition(date_id='$d')
EOF
)

echo $sql

/usr/bin/hadoop/software/hive/bin/hive -e "$sql"

exitCode=$?
if [ $exitCode -ne 0 ];then
         echo "[ERROR] hive execute failed!"
         exit $exitCode
fi

sql=$(
cat <<EOF
set mapred.job.priority=VERY_HIGH;insert into log_user_action_active_user partition(date_id='$d') select id, uuid, create_time, entity_id, '$create_date' as create_date, if(scene is null, 0, scene) as scene from (SELECT min(t1.id) AS id, t1.uuid, min(t1.create_time) AS create_time, t1.entity_id, min(t1.scene) as scene FROM log_user_action as t1  where t1.date_id = '$d' and uuid != "" group by t1.entity_id, t1.uuid) as tmp
EOF
)

echo $sql

/usr/bin/hadoop/software/hive/bin/hive -e "$sql"

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