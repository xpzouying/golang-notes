# MySQL Binlog解析


## 前言

由于之前的项目接触各种数据库较少，包括MySQL。常常在

最近在项目中接触到


## 启动MySQL

```yaml
version: '3.3'
services:
  db:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_DATABASE: 'db'
      # So you don't have to use root, but you can if you like
      MYSQL_USER: 'user'
      # You can use whatever password you like
      MYSQL_PASSWORD: 'password'
      # Password for root access
      MYSQL_ROOT_PASSWORD: 'password'
    ports:
      # <Port exposed> : < MySQL Port running inside container>
      - '3306:3306'
    expose:
      # Opens port 3306 on the container
      - '3306'
      # Where our data will be persisted
    volumes:
      - ./mysql:/var/lib/mysql
```


查看binlog路径：

```bash
# 查看binlog信息
show variables like 'log_%';

# 查看binlog数据偏移
SHOW master status;


# 查看binlog日志
SHOW binary logs;
```
