# Go Example VPS Locaweb
Exemplo de uma aplicação segura em Go no VPS da Locaweb


# Estrutura do Projeto baseado no Mercurius 
Este App foi construído usando o framework Web em Go chamado Mercurius: https://github.com/novatrixtech/mercurius

```
/conf 
Application configuration including environment-specific configs

/conf/app
Middlewares and routes configuration

/handler
HTTP handlers

/locale
Language specific content bundles

/lib
Common libraries to be used across your app

/model
Models

/public
Web resources that are publicly available

/public/templates
Jade templates

/repository
Database comunication following repository pattern

main.go
Application entry
```