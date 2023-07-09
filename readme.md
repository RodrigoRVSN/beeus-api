# Beeus API

<!--- https://shields.io --->

![GitHub repo size](https://img.shields.io/github/repo-size/rodrigorvsn/beeus-api?style=for-the-badge)
![GitHub language count](https://img.shields.io/github/languages/count/rodrigorvsn/beeus-api?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/rodrigorvsn/beeus-api?style=for-the-badge)

___
<!--- #################### mudar badges #################### --->

- Database ERD

![image](https://github.com/RodrigoRVSN/beeus-api/assets/75763403/7ef0ec01-4cba-4fb8-8e1e-27144ac4c410)


<!--- #################### mudar imagem exemplo #################### --->
___

This is the API being used in [Beeus](https://github.com/william-james-pj/beeus).

![image](https://github.com/RodrigoRVSN/beeus-api/assets/75763403/47781665-2dfa-4e2d-8587-d07ad25b8e56)

___
## 💻 Prerequisites

- Have [docker](https://www.docker.com/) and [docker compose](https://docs.docker.com/compose/) installed or just the [Go cli](https://go.dev/doc/install)

<!--- #################### mudar pré-requisitos  ####################--->
___
## 🛠 Tools used

- [Golang](https://go.dev/): Programming language
- [GORM](https://gorm.io/): ORM
- [Fiber](https://gofiber.io/): Go framework
- [Docker](https://www.docker.com/): Run application in container
- [PostgreSQL](https://www.postgresql.org/): Relational database
- [Render](https://render.com/): Deploy server
- [ElephantSQL](https://www.elephantsql.com/): Deploy PostgreSQL database

<!--- #################### mudar ferramentas #################### --->
___
## 🚀 Running beeus-api

To run this application, follow the steps:

1. Create and fill a `.env` fill with the information provided in `.env.example`.

2. Build and run the database and application configured in `docker-compose.yaml`, using:

```bash
docker compose up
```

3. Done! You can see the documentation of routes in the Insomnia in `./docs/routes-beeus-api-insomnia.json`

## 🧪 Testing
Run all tests of the project recursively 

```bash
docker compose up test
```
_or if you have the Go CLI installed_
```bash
go test ./...
```

## ✍️Annotations

- [x]  **SOLID**
    - [x]  **S (SRP - Single Responsibility Principle)**
        
        > Uma classe deveria fazer apenas uma coisa e deveria ter apenas um motivo para ser modificada
        > 
        - [x]  Camada `repository` tem apenas uma responsabilidade: acesso ao banco de dados.
    - [x]  **O (OCP - Open–Closed Principle)**
        
        > Deve ser possível extender  o comportamento de uma classe sem modificá-la.
        > 
        - [x]  `DTOs` aplicando o `DTO` do usuário sem senha. Isso permite adicionar informações sobre o autor sem necessariamente mudar o `DTO` que está utilizando a implementação desse `DTO` de usuário sem senha
    - [x]  **L (LSP - Liskov Substitution Principle)**
        
        > Subclasses deveriam ser substituíveis pelas classes base.
        > 
        - [x]  `DocumentationControllerInterface` define os métodos implementados pelo `controller` de documentações. Novas subclasses seguiriam a mesma assinatura.
    - [x]  **I (ISP - Interface Segregation Principle)**
        
        > A classe não deveria ser forçada a implementar interfaces e métodos que não vão ser usados.
        > 
        - [x]  Interface de `gateway` utilizada no `repository` e no `use case`
    - [x]  **D (DIP - Dependency Inversion Principle)**
        
        > Dependa de abstrações, não implementações
        > 
        - [x]  Ao implementar a interface `gateway` , deixamos expostos métodos abstratos providos pela interface ao invés de uma implementação direta. Caso algo da camada repository mudasse, alteraríamos a implementação dela apenas.
- [x]  Design Patterns
    - [x]  Strategy
        - [x]  A interface **`DocumentationUseCaseInterface` declara os métodos do `use case` e contém a lógica de cada implementação. Com isso, os métodos da camada `controller` podem chamar os métodos do `use case` em tempo de execução, sem conhecer a implementação específica dele.**
    - [x]  Adapters
        - [x]  Utilização de `DTOs` para retornar os dados específicos esperados daquele `repository`
    - [x]  Factory
        - [x]  Implementação da criação de uma entidade `Tag`

---

# Arquitetura de Software

- [x]  Arquitetura utilizada
    - [x]  **Clean Architecture** + um pouco de Hexagonal
- [x]  Apresentação da abordagem de Clean Architecture
    - [x]  Acoplamento
        - [x]  Main
        - [x]  Controller
        - [x]  Use case
        - [x]  Repository
        - [x]  DTOs
    - [x]  Reuso e manutenibilidade
        - [x]  Utilização de SOLID e implementação de interfaces
- [x]  DDD
    - [x]  Domínio: entity e gateway
    - [x]  Application: DTOs, useCases e controllers
    - [x]  Infra: database, rotas
    - [x]  Linguagem ubíqua (específico por entidades)
    - [x]  Repositórios (acesso direto a camada de infra via injeção de dependência)

## 📫 Contribuiting with beeus-api

To contribute to beeus-api, follow the steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <name_branch>`.
3. Make the changes and confirm: `git commit -m '<message_commit>'`
4. Send to original repository: `git push origin <main_branch> / <local>`
5. Send the pull request.

Alternatively, see the GitHub documentation at [how to create a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).
