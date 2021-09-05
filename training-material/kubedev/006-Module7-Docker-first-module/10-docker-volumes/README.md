# Módulo 7 - Docker first module

## How to save data - Volumes

"Container vai "morrer" a determinada altura" - Nasce mas tem um periodo, e novo container sera gerado.

Por isso, precisamos de volumes, para nao perdermos dados, por exemplo uma base de dados.

- **Bind Mount** - Mapear um directório no container com um directório na máquina local;
- **Volume** - Gerenciado pelo docker daemon (é gerado como um objecto no docker). Também é mapeado um directório do container com a máquina local, mas que gere esse mapeamento é o docker;
- **tmpfs** - Armazenar o directório do container em memória

## Bind Mount

**Executar container com bind mount:**

```bash
docker container run -it -v "$(pwd)/vol:/app" ubuntu /bin/bash
```

1º - No directorio local onde for corrido este comando, é criado um directório `/vol`, que está mapeado ao directório `/app` a correr dentro do container.

2º - Tudo o que for adicionado ao directório `/vol` será mapeado dentro do container, neste caso o ficheiro `container.txt`, e vice-versa, tudo o que está no directório `/app` do container, será mapeado ao directorio `/vol` localmente.

3ª - Podem ser feitas alterações pelo container ou localmente que serão sempre visiveis para ambas as partes

> Podemos ter de dar permissões aos ficheiros criados

**NOTA:**

Outra forma de executar container com bind mount:

```bash
docker container run -it --mount type=bind,src="$(pwd)",dst=/app ubuntu /bin/bash
```

> Estamos a fazer um bind do directório local em '$(pwd)', com o directório `/app` do container. (Em ambos os paths será criada um directório `/vol`, que representa o directório mapeado entre ambas as máquinas)

## Docker Volumes

Desta vez a gestão do mapeamento está a cargo do docker daemon, não somos nós que especificamos qual o directório a mapear localmente.

**Criar docker volumes:**

```bash
docker volume create <volume_Name>
```

**Listar docker volumes:**

```bash
docker volume ls
```

**Executar container usando docker volume:**

```bash
docker container run -it -v "<volume_Name>:/app" ubuntu /bin/bash
```

**Ver qual o directório local a ser mapeado:**

```bash
docker volume inspect <volume_Name>

# Output:
[
    {
        "CreatedAt": "2021-09-05T02:43:05+01:00",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/test_volumes/_data",
        "Name": "test_volumes",
        "Options": {},
        "Scope": "local"
    }
]

```
O Mountpoint é o directório local que está a ser mapeado para o directório `/app` do container.

> NOTA: Cria um ficheiro no directório `/app`, dentro do teu container, e verifica no path `/var/lib/docker/volumes/test_volumes/_data` localmente se esses dados estão a ser lá armazenados

**NOTA:**

Outra forma de executar container com bind mount:

```bash
docker container run -it --mount type=volume,src="<volume_Name>",dst=/app ubuntu /bin/bash
```