# Módulo 7 - Docker first module

## Exercicio 2

>   A prova de certificação Certified Kubernetes Application Developer (CKAD) utiliza um ambiente Linux e o vim como editor de texto. Como alguns utilizam o ambiente Windows e nem sempre temos o WSL2 instalado, sua tarefa agora é criar uma imagem docker que possua os requisitos de ambiente:
> - Kubectl (https://kubernetes.io/docs/tasks/tools/install-kubectl)
> - Vim (https://www.vim.org)
> Lembrando que é preciso ter o arquivo config na pasta .kube pra acessar o cluster Kubernetes


**Explain [Dockerfile](Dockerfile)**

* Começamos com a imagem do `ubuntu` para utilizarmos um ambiente linux
* Definimos o workspace
* Fazemos update dos packages
* Instalamos o `curl`
* Baixamos o executável de instalação do `kubectl`
* Instalamos o `kubectl`
* Instalamos o `vim`
* Copiamos o [ficheiro de configuração](config) para o path `./root/.kube/config`
  * NOTA: colocamos root no path porque a imagem `ubuntu` roda em sudo por default. Se nao estiver em modo sudo, seri só `./.kube/config`


**Build da imagem:**

```bash
docker build -t bpvcode/ubuntu-kubectl-vim:v1 .
```

**Run do container baseado na imagem:**

```bash
docker run -it -p 8080:80 bpvcode/ubuntu-kubectl-vim:v1 /bin/bash
```

- `-it` - interactive terminal (bash terminal - `/bin/bash` at the end)
- `-p` - port binding - mapear porto 8080 da máquina local com o porto 80 do container

Output:

```bash
root@40ff559e566d:/#
```

**Conclusão**

Temos um container a correr, com o sistema operativo ubunto versão 20.4, com o curl, o vim e o kubectl instalado.

Por fim, as configurações ao cluster são atualizadas pelo `.kube/config` file.

Podemos, através deste cluster criar ficheiros de código com o vim e fazer deploy das nossas aplicações no kubernetes através das configurações do ficheiro `.kube/config`.

---

NOTA:

Comand to delete all container - **WARNING!!!!!**

```bash
docker rm -f $(docker ps -a -q)
```
