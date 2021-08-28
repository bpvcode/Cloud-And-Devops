# Módulo 7 - Docker first module

## Docker Registry

Repositórios de imagens

### Docker HUB

No Docker HUB tem de criar primeiro o **namespace** e **repositório**(no user interface) antes de poder fazer push.

Exemplo:

> bpvcode/app-name:v1

    * `bpvcode` - **namespace**
    * `app-name` - **repositório**
    * `v1` - **version**


NOTA:

* Para o seguinte workflow utilizei este [Dockerfile](../03-images-Dockerfile/api-conversão/Dockerfile) como base;
* Para correr o comando de build da imagem deverá ser no [path](../03-images-Dockerfile/api-conversão)  `../03-images-Dockerfile/api-conversão`.

**1º - LOGIN**

```bash
docker login -u bpvcode
```

**2º - BUILD IMAGE**

```bash
docker build -t bpvcode/api-conversao:v5 -t bpvcode/api-conversao:latest .
```

**3º - PUSH IMAGE TO REGISTRY**

```bash
docker push bpvcode/api-conversao:v5
docker push bpvcode/api-conversao:latest
```

To push the images to a repository

NOTA:

Boa prática, fazer push da versão com `tag` e `latest`

**4º - USE IMAGE FROM A REGISTRY TO RUN THE APLICATION IN A CONTAINER**

```bash
docker container run -d -p 8080:8080  bpvcode/api-conversao:v5
```

**º - USE THE APPLICATION**

A aplicação estará a disponivel, a ser servida através do container e pode ser acedida localmente em `http://localhost:8080/api-docs` [Swagger](http://localhost:8080/api-docs).