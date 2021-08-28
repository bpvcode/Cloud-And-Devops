# Módulo 7 - Docker first module

## Azure Container Registry

Alternativa ao Docker Hub, serve como repositório de imagens e muitos mais serviços.

[**Portal Microsoft Azure**](https://portal.azure.com)

Existem mais (Nexus, ACR, ECR).

**Criar um registry - Portal Azure**

* Menu Button -> All services > Container registries > + CREATE
  * Criar registry
  * Escolher Resource group, podemos ter varios
  * Definir o registry name -> `bpvcode.azurecr.io`
  * Escolher localização do servidor
  * SKU - Plano (Escolher `Basic`) se o storage for inferior a 10GB - ver mais informaões sobre preços

**Para poder acessar o registry criado através do docker cli localmente**, é necessário:

1º - Habilitar o Admin user. Portal Azure -> `**Settings -> Access keys -> Admin User - enable**`
  * Vai gerar um `username` e duas `passwords` (podemos usar qualquer uma delas) para authenticar o docker através da máquina local.

NOTA:

Se a máquina já estiver authenticada no Azure, é possivel ver a informação em `~/.docker/config.json` na máquina local.

2º - Run command to apply connection to this registry (Pass username and password in terminal)

```bash
docker login bpvcode.azurecr.io
```

**Push image to Azure**

1º - Build the image

```bash
# Correr comando no [path](../06-multistagebuild/app-go-example)
docker build -t bpvcode.azurecr.io/app-go-example:v1 -t bpvcode/app-go-example:latest -f Dockerfile-good .
```

2º - Push the image

```bash
# Correr comando no [path](../06-multistagebuild/app-go-example)
docker push bpvcode.azurecr.io/app-go-example:v1
docker push bpvcode.azurecr.io/app-go-example:latest
```

3º - Check if the image is in Azure repository:

Portal Azure -> `Menu Button > All services > Container registries`

Namespaces:

[namespaces](./azure-containers-registry-1.png)


Portal Azure -> `Menu Button > All services > Container registries > bpvcode > Repositories `

Repositories:

[repositories](./azure-containers-registry-2.png)

Portal Azure -> `Menu Button > All services > Container registries > bpvcode > Repositóries > app-example `

Images and versions inside the repo:

[images](./azure-containers-registry-3.png)