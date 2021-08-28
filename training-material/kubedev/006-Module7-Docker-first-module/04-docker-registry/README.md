# Módulo 7 - Docker first module

## Images - Docker Registry

Repositórios de imagens

### Docker HUB

No Docker HUB tem de criar primeiro o **namespace** e **repositório**(no user interface) antes de poder fazer push.

Exemplo:

> bpvcode/app-name:v1

    * `bpvcode` - **namespace**
    * `app-name` - **repositório**


NOTA:

Para o seguinte workflow utilizei este [Dockerfile](../03-images-Dockerfile/api-conversão/Dockerfile) como base;

Para correr o comando de build da imagem deverá ser no [path](../03-images-Dockerfile/api-conversão)  `../03-images-Dockerfile/api-conversão`.

**1º - LOGIN**

```bash
docker login
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

**4º - USE IMAGE FROM A REGISTRY TO RUN A CONTAINER**

```bash
docker container run -d -p 8080:8080  bpvcode/api-conversao:v5
```

A aplicação estará a rodar através do container e pode ser acedida localmente em `http://localhost:8080`