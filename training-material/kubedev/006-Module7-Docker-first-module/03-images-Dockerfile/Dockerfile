FROM ubuntu

# INCORRECT - Se instalar algo depois de instalar o curl, nao vai fazer update antes de fazer install
# RUN apt-get update
# RUN apt-get install --yes curl

# CORRECT - Pode ser feito em varias linhas tambem com "\"
RUN apt-get update && apt-get install --yes curl

# Run this in container terminal, after create and execute the image:
# curl https://www.google.com
# Ex: root@3b55c84f0aaf:/# curl https://www.google.com