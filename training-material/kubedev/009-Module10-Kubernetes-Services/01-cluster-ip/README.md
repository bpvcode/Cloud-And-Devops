# Módulo 10 - Kubernetes Services

## Services - Cluster IP

Serve para gerar conexão entre POD's só dentro do cluster Kubernetes, nunca para conexões fora do cluster.

**Deploy da aplicação api-conversao**

```bash
kubectl apply -f ./deployment.yml
```

**Find the IP of the POD**

Podemos comunicar através desse IP (colocando o IP no URL)

```bash
kubectl get pods -o wide
```

**Correr um container com ubuntu e curl, em modo iterativo, para fazer um pedido ao IP do POD a cima:**

```bash
kubectl run -i --tty --image kubedevio/ubuntu-curl ping-test --restart=Never --rm -- /bin/bash

root@ping-test: curl http://172.17.0.9:8080/temperatura/fahrenheitparacelsius/100
```

**RESULTADO FINAL**

**Aplicar o Service**

```bash
kubectl apply -f ./service.yml
```

**Fazer pedido, mas a um "domino" definido e não ao IP**

```bash
root@ping-test: curl http://app-api-service/temperatura/fahrenheitparacelsius/100
```
Service gere qual o endpoint que tem que acessar, mascarando-o com um dominio `app-api-service`.
