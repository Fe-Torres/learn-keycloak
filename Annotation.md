# ** Autenticação com Keycloak **

** Meios de autenticação **

- Auth2
> Autenticação para depois acontecer uma autorização/login.

- IdConect

- JWT
---
** Single sing-on **
> Autenticação de forma centralizada, ou seja, teremos vários micro services e todos eles autenticados no keycloak.

** Realm Keycloak **
> Separação lógica dos contextos do processo de autenticação. O realm "Master" serve para configurar o próprio keycloak.

---

** Clients Keycloak **
> São os sistemas que irão ter acessar o sitema de autenticação.

---
** Observação **

> Todas as vezes que você for criar um usuário no seu sistema, você vai criar no keycloak!

---

** Scope Keycloak **
> Para o clint ter acesso ao atributo do usuário no keycloak, precisamos criar um escope (scope). Ou seja, serve para mapear os atributos do usuário.

---

** Roles Users **
> Tem como criar roles no keycloak, para diferenciar os acessos do usuário (administrador, visitante, membro, etc.)

---

** Groups Users **
> Servem para detalhar as permissões/características de determinado grupo para que haja as especificações gerais de acesso. Feito isso, iremos apenas adicionar os usuários nesses grupos criados, sem precisar configurar cada um desses usuários.
> Obs: Caso seja necessário, tem como você retornar os grupos, os mesmos que o usuário está inserido.

---

** Keycloak com outros Providers **
> Google,facebook