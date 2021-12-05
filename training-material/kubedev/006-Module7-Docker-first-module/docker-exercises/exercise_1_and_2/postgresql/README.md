# PostgreSQL

NOTE: Exercise 2 using Visual Studio Code SQLTools extension instead PgAdmin

1 - Create a docker volume. Docker will manage a folder in host filesystem where to store the same data as the container

```bash
docker volume create postgres_vol
```

2 - Running a postgres instance in docker container and mount the directory `/data/db` inside container to the path managed by docker in host filesystem. Fin this path by run `docker inspect postgres_vol`

```bash
docker run --name postgresdb -d \
    -e POSTGRES_USER="postgresuser" \
    -e POSTGRES_PASSWORD="postgrespwd" \
    -e POSTGRES_DB="test" \
    -v postgres_vol:/data/db \
    -p 5433:5432 \
    postgres:12.9
```

3 - Connect the container with Visual Studio Code SQLTools extension

Choose PostgreSQL:

![setup connection](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres.png)

Insert connection details:

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres1.png)

Test that connection is successfully established and save connections:

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres2.png)

Connected:

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres3.png)

Add random data:

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres4.png)

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres4.1.png)

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres4.2.png)

4 - Check that container has the same data as Visual Studio Code SQLTools extension. Find in `/data/db` path

```bash
docker exec -it postgresdb /bin/bash
```

**NOTE:**

5 - Remove docker container

```bash
docker rm -f postgresdb
```

6 - Refresh Visual Studio Code SQLTools extension

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres5.png)

7 - Restart postgresdb container with the same volume

```bash
docker run --name postgresdb -d \
    -e POSTGRES_USER="postgresuser" \
    -e POSTGRES_PASSWORD="postgrespwd" \
    -e POSTGRES_DB="test" \
    -v postgres_vol:/var/lib/postgresql/data \
    -p 5433:5432 \
    postgres:12.9
```

8 - Refresh Visual Studio Code SQLTools extension - Check that we have the same data although the container has been deleted and restarted

![setup connection1](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/postgresql/postgres4.2.png)

## Conclusion

We create a postgres instance, based on a docker container, and we mount the volume data to the host filesystem. We used a Refresh Visual Studio Code SQLTools extension to interact with our database and we persist all data through docker volumes.
