# DATABASE with me

## Build MySQL container with Docker

Get MySQL from docker hub [MySQL Image](https://hub.docker.com/_/mysql)

```docker
docker pull mysql

docker images -a
```

Start a MySQL instance from image just pulled

```shell
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest

docker exec -it mysql mysql -u root -p

```

After that, login with password="secret"

```mysql
create database simple_school;

drop database simple_school;
```

## Install database migrate CLI

```shell
brew install golang-migrate

migrate --version

migrate --help
```
