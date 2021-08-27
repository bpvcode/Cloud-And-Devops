# Kube Talk
**Live-Ingress-Controller-Cert-Manager**


## Certificados HTTPS:

É preciso uma entidade que certifique o dominio.

cert-manager

>https://cert-manager.io/docs/

Gestor de certificados
    * Issuers - Serviço que gera o certificado
      * Issuers - especifico para um namespace
      * Cluster Issuer - para todos os namespaces
    * Certificates -