# Exercise 5

5 - "Chegou um cliente pra você que possui todas as suas aplicações em data centers e a gestão dessas aplicações está cada vez mais complexa então pra iniciar um plano de gestão unificada e migração pra um ambiente cloud, as aplicações serão migradas pra containers. E hoje você precisa iniciar esse processo com um projeto piloto, o portal de conteúdos da empresa construido em Wordpress. Então hoje sua missão é criar esse ambiente wordpress pronto para a equipe de publicidade começar a popular." - kubedev community exercises

# Wordpress docker environment

**Run application:**

```bash
docker-compose up -d
```

Docker compose explain:

- 1º Create volume and network
- 2º Create mysql container with volume
- 3º Create wordpress container

[Go to App](<http://localhost:8080>)

![wordpress dashboard](./wordpress-dashboard.png)
