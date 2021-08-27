# Módulo 2 - Introdução

## Container and kubernetes overview

### Como era antes dos containers:

* Montagem de ambiente localmente para aquela aplicação especifica (mysql, golang, etc...)
* Programar, testar etc...
* Resultado final será um pacote (ex: executável)
* Pacote final é entregue para o administrador de sistemas, que vai fazer o deploy da aplicação, para a equipa de qa fazer homolgação e testes do que foi desenvolvido.
* O pacote seria entregue junto com manuais para montagem de ambiente
* Operador configura ambiente no servidor (onde correm outras aplicações)
  * **É muito comum no servidor existir conflitos entre pacotes de aplicações diferentes**, pois os ambientes de cada aplicação variam entre si(variam as stacks), e geram conflitos entre as aplicações.

**Foi para resolver esse problema de conflitos entre ambientes que foram criados os containers**


### O que são containers?

Containers são "caixas" que isolam processos (mysql, app java, app go, mongo db, etc...) e conseguimos utilizar esse container em qualquer ambiente, e podem ser executados em qualquer serviço que faça execução de containers (posso executar o mesmo container na minha maquina local como num desses serviços, que o comportamento será o mesmo). Com isto ganhamos isolamentos de processos (isolados num ambiente) e idempotência (mesmo comportamento localmente ou remoto).

### Docker

Ferramenta de criação e execução de containers, facilita a gestão dos containers

Como se cria um container através de uma aplicação?

* Tenho de ter o código da aplicação;
* Tenho de ter um Dockerfile
    * Dockerfile funciona como "a receita" para a criação da imagem (blueprint para a execução do container);
* Através do Dockerfile e do Código da aplicaçao, vamos criar uma imagem docker;
* Devemos hospedar as imagens criadas em repositórios de imagens (DockerHub, Nexus);
* Podemos depois, aceder a essas imagens, em qualquer lugar atraves desse repositório (nao estão só guardadas localmente numa maquina);
*Essas imagens docker darão origem a containers;


**Saimos do cenário em que a aplicação é executada toda dentro de um servidor (antes dos containers). Passando a ser executado em vários containers com comunicação,ou não, entre si (Container para API, container para DB), trabalhamos então de forma mais isolada**

O que o Docker não nos dá?

**Escalabilidade**, por exemplo, a criação de réplicas de todos os processos a serem executados.

Precisamos de um orquestrador de containers...

### Kubernetes - open source project

Ferramenta ideal para o cenário de operação no mundo DEVOPS, no mesmo ambiente executa processo go, java, .net, databases...

* Escalabilidade - Cria replicas de containers;
* Resiliência - Mecanismos de monitoramento de containers, garante que o numero de replicas vai ser garantido no cluster;
* Estratégias de atualização da aplicação - garantindo downtime 0 (Tambem faz rollback);
* Execuções agendadas;
* Service Discovery - Aplicação segue como um serviço que tem um endereço do container que esta executando o processo;

**Arquitectura do Kubernetes**

* Cluster - conjunto de máquinas que podem ter 2 papeis distintos no cluster
  * Control Plane - orquestrar / gerenciar o cluster
    * Kube API - `kubectl` comunica com o control plane atraves do Kube API
    * ETCD - E onde são armazenados os dados referentes ao cluster. Nunca é acessado directamente, mas sim atraves do Kube API Server
    * Kube Scheduler - Onde os manifestos vão ser lidos e organizados. Qualquer processo no kubernetes passa pelo Kube Scheduler
    * Kube Controller Manager - Gere todos os controladores do cluster kubernetes (Replica Set, Admission Control,...)
  * Nodes - executar os containers (executar workloads) é normal ter mais nodes
    * Kubelet - Comunica com o control plane utilizando o Kube API Server, verifica o estado do node que esta a ser executado;
    * Kube Proxy - Responsavel pela comunicação entre os nodes
    * Container Runtime - Responsavel por executar os containers (Pode rodar Docker cri-o)