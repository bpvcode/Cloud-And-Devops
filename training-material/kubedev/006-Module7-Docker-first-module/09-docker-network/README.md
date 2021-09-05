# Módulo 7 - Docker first module

## Docker Network

Os processos devem ser todos isolados, um por container, mas é necessário existir a comunicação entre containers se os processos comunicarem entre si. Ai entra o docker network, que gere exactamente a comunicação entre os container.

### Network types

- **bridge** - Forma padrão;
- **host** - Container passa a ter a mesma rede que a máquina local;
- **overlay** - Container comunica com vários docker hosts, uso de docker swarm (orquestrador de containers do docker);
- **macvlan** - Simula enedereço mac de red no container

Comando de info:

```bash
docker network
```

Criar rede:

```bash
docker network create <network_Name>
```

Listar redes:

```bash
docker network ls
```

Conectar um container já em execução a uma rede:

```bash
docker network connect <network_Name> <container_ID>
```

> NOTA: Se ambos os containers estiverem conectados à mesma rede, poderão comunicar entre si

Conectar um container já em execução a uma rede:

```bash
docker network connect <network_Name> <container_ID>
```

Desconectar um container já em execução a uma rede:

```bash
docker network disconnect <network_Name> <container_ID>
```

Conectar um container a uma rede no momento da criação do mesmo:

```bash
docker container run -d --network <network_Name> --name <container_Name> <image>:<version>
```

Conectar um container à rede local do host - localhost:

```bash
docker container run -d --network <network_Name> --name <container_Name> <image>:<version>
```

> NOTA: Container que está a ser executado, compartilha a mesma rede que a máquina local