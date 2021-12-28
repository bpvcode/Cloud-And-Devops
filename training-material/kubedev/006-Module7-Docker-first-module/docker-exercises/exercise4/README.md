# Exercise 4

4 - "Agora que você já afiou o seu conhecimento sobre criação de imagens Docker, tá na hora de fazer o deploy de uma aplicação 100% em containers Docker. A aplicação está no GitHub do KubeDev e um detalhe MUITO importante, a aplicação precisa ser toda criada com apenas uma linha de comando." - kubedev community exercises

# Rotten-potatoes app

NOTE: Need docker login first

[Dockerfile added](Dockerfile)

Build image:

```bash
docker build . -t bpvcode/potatoes-app:v1
```

Push image to personal docker hub repo:

```bash
docker push bpvcode/potatoes-app:v1
```

**Run application:**

```bash
docker-compose up -d
```

Docker compose explain:

- 1º Create volume and network
- 2º Create mongo db container with volume
- 3º Create app container from previous build image
  - will fetch the image from `bpvcode/potatoes-app` repo, version `:v1`
  - Port binding

[Go to App](<http://localhost:5000>)

![rotten potatoes app](./rotten-potatoes-app.png)

## Configuração

MONGODB_DB => Nome do database

MONGODB_HOST => Host do MongoDB

MONGODB_PORT => Porto de acesso ao MongoDB

MONGODB_USERNAME => User do MongoDB

MONGODB_PASSWORD => Pass do MongoDB
