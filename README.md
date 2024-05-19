# user-service

# download the user service
git clone https://github.com/ABHISHEK-SUMAN-repo/user-service.git

# Download the PostgreSQL database

Download the PostgreSQL database : [postgresSQL](https://www.postgresql.org/download/)

# Download DBeaver

Watch how to download DBeaver
Download Link : [dbeaverYT](https://youtu.be/0BOjD6H9Uos?si=2NrD1rD0z8qJlL5z/)

# install and setup go

Download and install  : [go](https://go.dev/doc/install)

Go to the go directory and clone the repository on the same file.

**For windows set-up do the following:**

```
Environment variable edit:

| Variable  :  Value |
---------------------
| GOPATH    :  C:\go |  

```
System variable edit: Selected path and then add :

```
| Variable  :  Value |
---------------------
| GOPATH    :  C:directory_where_go_is_installed\Go\bin |
```

run command : 
```bash
go mod init user-service

# Run the command
**On Linux/MacOS:**

```bash
export APP_ENV=dev
go run main.go

```
**On Windows:**

```bash
set APP_ENV=dev
go run main.go
