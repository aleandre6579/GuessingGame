echo "** Creating default DB and users"

mariadb -uroot -p$MYSQL_ROOT_PASSWORD --execute \
"
CREATE DATABASE IF NOT EXISTS $MYSQL_DB;
GRANT ALL PRIVILEGES ON $MYSQL_DB.* TO '$MYSQL_USER'@'%';
USE $MYSQL_DB;
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
);
"

echo "** Finished creating default DB and users"
