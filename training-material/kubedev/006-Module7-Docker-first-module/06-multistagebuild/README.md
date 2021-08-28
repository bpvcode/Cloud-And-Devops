# Módulo 7 - Docker first module

## Multistage Build

Bom para otimizar as imagens docker, principalmente se a aplicação for compilada.

**[Mau Exemplo](app-go-example)**

* Criar imagem a partir do Dockerfile:

```bash
docker build -t bpvcode/app-go-example:10 .
```

* Ver tamanho da imagem:

```bash
docker image ls
```

* Resultado final:

```bash
  REPOSITORY                                                                    TAG                                                                IMAGE ID       CREATED          SIZE
bpvcode/app-go-example                                                        latest                                                             e26271529947   5 minutes ago    882MB
bpvcode/app-go-example                                                        v1                                                                 e26271529947   5 minutes ago    882MB
```

**NOTA: 882MB**

* Neste caso estamos a compilar e a executar no mesmo stage


**[Bom Exemplo](a)**

* Criar imagem a partir do Dockerfile:

```bash
docker build -t bpvcode/app-go-example:10 -f Dockerfile-good .
```

* Ver tamanho da imagem:

```bash
docker image ls
```

* Resultado final:

```bash
  REPOSITORY                                                                    TAG                                                                IMAGE ID       CREATED          SIZE
bpvcode/app-go-example                                                        latest                                                             3530524e1acf   12 minutes ago       5.6MB
bpvcode/app-go-example                                                        v10                                                                3530524e1acf   12 minutes ago       5.6MB
```

**NOTA: 5.6MB**  

* Neste caso a imagem conta só a partir do ultimo stage, e a imagem do `alpine` é menos pesada do que a do `golang` e ja só temos o ficheiro executável neste workspace, para poder executar, e não todo o conteudo como no exemplo antigo;
* No caso de **linguagens interpretadas**, é necessário que o ultimo stage tenha o código da aplicação para poder ser interpretado;
* Contudo, no caso de **linguagens compiladas**, só precisamos de ter o ficheiro executável, e reduz em muito o tamanho da imagem final.
