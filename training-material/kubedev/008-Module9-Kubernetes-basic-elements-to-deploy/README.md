# Módulo 9 - Elementos básicos de um deploy no kubernetes

## POD

Menor objecto do cluster kubernetes.

Onde são executados os containers.

Dentro de um pod podemos ter mais de um container a correr - **side car pattern**, onde no container principal está a rodar a aplicação e em containers auxiliares executam-se rotinas nao funcionais, como recolha de logs, metricas... basicamente funcionalidades que não estão envolvidas no negócio da nossa aplicação.

Se não estivermos a utilizar esse pattern, **o mais provavel é termos só um container por cada pod**, para que consigamos garantir a escalabilidade, sem termos de escalar toda a aplicação em simultâneo.

Quando o POD é criado, **recebe um IP**, **tem o próprio filesystem**, como se fosse uma máquina.

NOTAS:

- Sempre que temos mais de um container no mesmo POD, eles compartilham a mesma rede e o mesmo filesystem
- Nunca devemos utilizar só o objecto POD para fazer o deploy da aplicação (Naked Pod), porque POD sózinho não garante escalabilidade nem resiliência.

Para criar um POD, criamos um ficheiro yaml (manifesto) com toda a especificação do objecto pod a criar no cluster kubernetes.

**Correr este comando no directorio corrente para criar o objecto POD com o nome `podname`:**

[POD manifesto](./pod.yml)

```bash
kubectl apply -f pod.yml
```

> NOTA:
>
> - Ver se imagem existe no docker hub, namespace `bpvcode` repositorio `stresss-test`, versão `v1`;
> - Neste momento ainda nao conseguimos aceder à aplicação via browser;

**Correr este comando para fazer o bind do porto da máquina local, com o porto do POD:**

```bash
kubectl port-forward pod/podname 8080:8080
```

8080 - Porto da máquina local

:

8080 - Porto do pod (que a aplicação está a expor, neste caso 8080)

> NOTAS FINAIS:
>
> - O objecto POD por si só não garante resiliência nem escalabilidade;
> - Se o node cair, não cria nova réplica/instância desse pod;
> - Se fizer delete ao POD agora, também não cria uma nova instância;

## LABELS and SELECTORS

### LABELS

Elementos `key:value` armazenados nos objectos do kubernetes, declarados na secção de `metadata`

Muito utilizado para organizar objectos, exemplo de labels:

- Tipo de aplicação
- Autor da aplicação
- Versão
- ...

**Criar 2 pods distintos, com labels distintas:**

[POD1 manifesto](label-latest.yml)

[POD2 manifesto](label-v1.yml)

```bash
kubectl apply -f ./label-latest.yml -f ./label-v1.yml

# Then:
kubectl port-forward pod/stress-test-pod-v1 9090:8080
kubectl port-forward pod/stress-test-pod-latest 8080:8080
```

Se acessarmos a `localhost:9090` e `localhost:8080` vemos que "são duas máquinas distintas a responder".

### SELECTORS

Forma de selecionar determinados objectos, mediante as LABELS que lhe foram atribuidas.

Exemplo de utilização de SELECTORS:

- REPLICASET para selecionar os pods;
- SERVICES para saber quais os serviços que serão expostos;
- ...

> NOTA: Uso prático de SELECTORS nas próximas duas secções;

## REPLICASET

Responsável por gerenciar o objecto POD.

Objecto que garante a quantidade de réplicas desejada e garante que o estado da aplicação está conforme o esperado.


**Criar um objecto do tipo RéplicaSet:**

```bash
kubectl apply -f ./replicaset.yml
```

**Ver pod replicaSet:**

```bash
kubectl get pods
```

```bash
kubectl describe pod <replicaset_pod_name>
```

**Fazer port binding da máquina local com o container:**

```bash
kubectl port-forward pod/<replicaset_pod_name> 8080:8080
```

**Ver ReplicaSet object:**

```bash
kubectl get replicaSet
```

```bash
kubectl describe replicaSet <replicaset_name>
```

**Fazer delete do pod:**

```bash
kubectl delete pod/<replicaset_pod_name>
```

**Neste caso é sempre gerado um novo POD, porque o objecto ReplicaSet está encarregue de se certificar que a quantidade de réplicas e o estado da aplicação está conforme desejado e estipulado do manifesto [`replicaset.yml`](replicaset.yml)**

**Teste de escalabilidade:**

```bash
kubectl scale replicaset <replicaSet_name> --replicas=10
```

Vamos pedir ao ReplicaSet object para aumentar o número de réplicas (10) e verificar que serão gerados mais pods.

Neste caso, o ReplicaSet object vai tentar ter sempre 10 intâncias do pod em STATUS running.

NOTAS FINAIS:

- ReplicaSet object garante:
  - Escalabilidade - pois permite ter várias réplicas do mesmo container a rodar e permite assim um downtime 0;
  - Resiliência - pois tenta ter sempre o numero de réplicas e o estado da aplicação como o desejado e gere todo esse processo;

- ReplicaSet object não gere a troca de versões de uma aplicação, ou seja, se mudarmos a imagem que estamos a usar no manifesto do replicaSet, não vai afetar os pods que já estão criados... esses pods são na mesma baseados na imagem antiga.
Para que haja uma atualização, esses pods que ja estao baseados na imagem antiga, teem de ser apagados, e quando forem gerados novos pods, já serão baseado na imagem nova.

## DEPLOYMENT

Responsável por gerenciar o ReplicaSet object.

[DEPLOYMENT manifesto](./deployment.yml)

Este objecto gere o ReplicaSet de forma a conseguirmos ter o processo de troca entre pods quando existem alterações no manifesto.

Garante por exemplo a gestão de versões, quando alterada a versão da imagem, de forma a que os PODS que ja estavam criados na versao antiga sejam reiniciados para a versão nova...

Quando o deployment manifest é alterado, **é criado um novo ReplicaSet**, mas só um fica ativo (com os novos pods em execução).
Isto porque o **Deployment object guarda os ReplicaSet objects para poder ter uma gestao das versões**, e assim garante tambem que **conseguimos fazer rollback para versões anteriores**

**Para ver o histórico de versões:**

```bash
kubectl rollout history deployment <deployment_Name>
```

**Para fazer undo e mudar o replicaSet que está ativo para um anterior:**

```bash
kubectl rollout undo deployment <deployment_Name>
```
