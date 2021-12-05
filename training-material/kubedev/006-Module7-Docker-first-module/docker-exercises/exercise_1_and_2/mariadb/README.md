# MariaDB

NOTE: Exercise 2 using DBeaver instead phpmyadmin

1 - Create a docker volume. Docker will manage a folder in host filesystem where to store the same data as the container

```bash
docker volume create mariadb_vol
```

2 - Running a mariadb instance in docker container and mount the directory `/var/lib/mysql` inside container to the path managed by docker in host filesystem. Fin this path by run `docker inspect mariadb_vol`

```bash
docker run --name mariadb -d \
    -e MARIADB_ROOT_PASSWORD="mariadbrootpwd" \
    -v mariadb_vol:/var/lib/mysql \
    -p 3307:3306 \
    mariadb:10.6
```

3 - Connect the container with DBeaver

Choose MariaDB:

![setup connection](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/mariadb/mariadb.png)

Insert connection details:

![setup connection](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/mariadb/mariadb1.png)

Add random data:

![setup connection](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/mariadb/mariadb3.png)

![setup connection](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/mariadb/mariadb4.png)

4 - Check that container has the same data as DBeaver. Find in `/var/lib/mysql` path

```bash
docker exec -it mariadb /bin/bash
```

**NOTE:**

5 - Remove docker container

```bash
docker rm -f mariadb
```

6 - Refresh DBeaver

7 - Restart mariadb container with the same volume

```bash
docker run --name mariadb -d \
    -e MARIADB_ROOT_PASSWORD="mariadbrootpwd" \
    -v mariadb_vol:/var/lib/mysql \
    -p 3307:3306 \
    mariadb:10.6
```

8 - Refresh DBeaver - Check that we have the same data although the container has been deleted and restarted

![setup connection](/training-material/kubedev/006-Module7-Docker-first-module/docker-exercises/exercise_1_and_2/mariadb/mariadb4.png)

## Conclusion

We create a mariadb instance, based on a docker container, and we mount the volume data to the host filesystem. We used DBeaver to interact with our database and we persist all data through docker volumes.