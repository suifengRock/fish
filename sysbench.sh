sysbench --test=oltp --mysql-host="$DB_PORT_3306_TCP_ADDR"  \
--mysql-user="$DB_ENV_MYSQL_USER" --mysql-password="$DB_ENV_MYSQL_PASSWORD" \
--mysql-db="$DB_ENV_MYSQL_DATABASE" --oltp-table-size=100000000 prepare

sysbench --max-requests=1000 --test=oltp \
--mysql-host="$DB_PORT_3306_TCP_ADDR"  \
--mysql-user="$DB_ENV_MYSQL_USER" --mysql-password="$DB_ENV_MYSQL_PASSWORD" \
--mysql-db="$DB_ENV_MYSQL_DATABASE" \
--num-threads=100 run

#sysbench --test=oltp --mysql-host="$DB_PORT_3306_TCP_ADDR"  \
#--mysql-user="$DB_ENV_MYSQL_USER" --mysql-password="$DB_ENV_MYSQL_PASSWORD" \
#--mysql-db="$DB_ENV_MYSQL_DATABASE" cleanup