# Exercise 1 and Exercise 2

1 - "Execute os comandos para criar os 4 bancos de dados listados com containers, e use como se tivesse instalado eles localmente na sua máquina (Não esquece de garantir que não vai perder os dados caso o container seja excluido).

- MongoDB
- MariaDB
- PostgreSQL
- Redis" - kubedev community exercises

2 - "Certo, você conseguiu criar 4 bancos na sua máquina utilizando containers. Mas tem uma coisa, não adianta só conectar a aplicação no banco quando se está desenvolvendo, é preciso também acessar o banco, executar comandos e consultar a base. Então vamos fazer isso da forma KubeDev de ser, com containers !!! Cada banco de dados tem uma ferramenta administrativa com interface web que você pode usar.

- MongoDB ⇒ Mongo Express
- MariaDB ⇒ phpmyadmin
- PostgreSQL ⇒ PgAdmin
- Redis ⇒ redis-commander" - kubedev community exercises

## Mongo

Exercise 2 using MongoDB Compass instead Mongo Express

### Create a docker volume. Docker will manage a folder in host filesystem where to store the same data as the container

```bash
docker volume create mongo_vol
```

### Running a mongo instance in docker container and mount the directory `/data/db` inside container to the path managed by docker in host filesystem. Fin this path by run `docker inspect mongo_vol`.

```bash
docker run --name mongodb -d \
    -e MONGO_INITDB_ROOT_USERNAME="mongouser" \
    -e MONGO_INITDB_ROOT_PASSWORD="mongopwd" \
    -v mongo_vol:/data/db \
    -p 27017:27017 \
    mongo:4.4.3
```

### Connect the container with MongoDB Compass

![setup connection](mongo_compass.png)

### Add random sample data to database

![test database](mongo_db_com.png)
![test database data](mongo_compass_2.png)

### Check that container has the same data as mongo compass. Find in `/data/db` path

```bash
docker exec -it mongodb /bin/bash
```

### NOTE

Remove docker container:

```bash
docker rm -f mongodb
```

Refresh MongoDB Compass:

![connection not found](mongo_compass3.png)

Restart mongodb container with the same volume:

```bash
docker run --name mongodb -d \
    -e MONGO_INITDB_ROOT_USERNAME="mongouser" \
    -e MONGO_INITDB_ROOT_PASSWORD="mongopwd" \
    -v mongo_vol:/data/db \
    -p 27017:27017 \
    mongo:4.4.3
```

Refresh MongoDB Compass - Check that we have the same data although the container has been deleted and restarted:

![restore data](mongo_compass_2.png)