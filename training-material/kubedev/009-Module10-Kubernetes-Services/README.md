# Módulo 10 - Kubernetes Services

## Services

Objectos Services gerenciam os endpoints dos POD's.

Funcionam como o ponto unico onde os POD's são registados e sempre que alguem queira acessar a um POD, **é da responsabilidade do Service atribuir essa conexão(IP) e fazer o balanceamento das conexões** para que não haja POD's com elevadas cargas de trabalho e outros sem nada...

## Services types

- [Cluster IP](./01-cluster-ip)
- [NodePort]()
- [LoadBalancer]()
- [ExternalName]()