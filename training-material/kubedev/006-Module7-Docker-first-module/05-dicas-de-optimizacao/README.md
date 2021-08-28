# Módulo 7 - Docker first module

## Dicas de optimização

* Começar sempre com imagens oficiais ou feitas por nós
* Tentar utilizar imagens que ja tenham o que precisamos para nao ter de instalar tudo
* Utilizar sempre imagens com tags para garantir a idenpotência, ou seja, que sempre que o arquivo será executado, é executado da mesma forma, e tem sempre o mesmo comportamento, pois a versão da imagem não muda, enquanto que com o latest, pode sair alguma atualização que altera o comportamento desejado
  *  Exemplo:
        ```bash
        # CORRETO
        FROM node:14.15-alpine3.12
        ```
        ```bash
        # INCORRETO
        FROM node
        ```
* Utilizar layers de forma inteligente para garantir o uso de cache da melhor forma [exemplo 1](../03-images-Dockerfile/Dockerfile) [exemplo 2](../03-images-Dockerfile/api-conversão/Dockerfile);
* Utilizar `.dockerignore` faz com que não copie alguns arquivos não desejados para o container (exemplo node_modules)
