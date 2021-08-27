FROM ubuntu:latest

#RUN - instruções
#Installar o curl no filesystem do container
#Fazer update do repostiótio
# RUN apt-get update
# RUN apt-get install curl --yes

#NOTA:
# Utilizar cache com camadas em instalação de pacotes junto com update, não e uma boa opção, mais vale desabilitar
# Como usamos cache ao construir as imagens, o apt-get pode ser um problema, porque se mais tarde nada mudar para cima dessa intrução,
# nunca vai fazer essa instrução, vai sempre a cache, ou seja nunca vai fazer update
#   Solução:
RUN apt-get update && \
    apt-get install curl --yes