# ginchat
study golang
 

 # docker run mysql

cat /data/mysql/cnf/my.cnf
```shell
# For advice on how to change settings please see
# http://dev.mysql.com/doc/refman/5.7/en/server-configuration-defaults.html

[mysqld]
#
# Remove leading # and set to the amount of RAM for the most important data
# cache in MySQL. Start at 70% of total RAM for dedicated server, else 10%.
# innodb_buffer_pool_size = 128M
#
# Remove leading # to turn on a very important data integrity option: logging
# changes to the binary log between backups.
# log_bin
#
# Remove leading # to set options mainly useful for reporting servers.
# The server defaults are faster for transactions and fast SELECTs.
# Adjust sizes as needed, experiment to find the optimal values.
# join_buffer_size = 128M
# sort_buffer_size = 2M
# read_rnd_buffer_size = 2M
skip-host-cache
skip-name-resolve
datadir=/var/lib/mysql
socket=/var/run/mysqld/mysqld.sock
secure-file-priv=/var/lib/mysql-files
user=mysql
sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION'

# Disabling symbolic-links is recommended to prevent assorted security risks
symbolic-links=0

#log-error=/var/log/mysqld.log
pid-file=/var/run/mysqld/mysqld.pid
[client]
socket=/var/run/mysqld/mysqld.sock

!includedir /etc/mysql/conf.d/
!includedir /etc/mysql/mysql.conf.d/
```
运行 mysql
 ```shell
 docker run -d -p 3306:3306 --privileged=true  -v /data/mysql/data:/var/lib/mysql  -v /data/mysql/cnf/my.cnf:/etc/my.cnf -e MYSQL_ROOT_PASSWORD=123456 --name mysql uhub.service.ucloud.cn/bigfcat/mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci
 ```
 

# docker 运行 redis
```shell
docker run --name redis -p 6379:6379 -d -v /data/redis/:/data  uhub.service.ucloud.cn/bigfcat/redis:6.2 --requirepass "123456" --appendonly yes
```