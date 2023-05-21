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

3. Run all tests of the project recursively 

```bash
docker compose up test
```
_or if you have the Go CLI installed_
```bash
go test ./...
```

## 📫 Contribuiting with beeus-api

To contribue with beeus-api, follow the steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <name_branch>`.
3. Make the changes and confirm: `git commit -m '<message_commit>'`
4. Send to original repository: `git push origin <main_branch> / <local>`
5. Send the pull request.

Alternatively, see the GitHub documentation at [how to create a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).