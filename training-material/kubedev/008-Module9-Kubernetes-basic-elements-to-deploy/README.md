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

[POD manifesto](./myfirstpod.yml)

Correr este comando no directorio corrente para criar o objecto POD com o nome `podname`:

```bash
kubectl apply -f myfirstpod.yml
```

> NOTA:
>
> - Ver se imagem existe no docker hub, namespace `bpvcode` repositorio `stresss-test`, versão `v1`;
> - Neste momento ainda nao conseguimos aceder à aplicação via browser;

Correr este comando para fazer o bind do porto da máquina local, com o porto do POD:

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

## REPLICASET

## DEPLOYMENT
