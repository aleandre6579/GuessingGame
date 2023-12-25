echo "** Creating default DB and users"

mariadb -uroot -p$MYSQL_ROOT_PASSWORD --execute \
"
CREATE DATABASE IF NOT EXISTS $MYSQL_DB;
GRANT ALL PRIVILEGES ON $MYSQL_DB.* TO '$MYSQL_USER'@'%';
"

echo "** Finished creating default DB and users"
