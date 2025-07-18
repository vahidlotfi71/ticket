# Golang Project

## Initialization

Clone the project from [github repository](https://github.com/vahidlotfi71/ticket) .

Navigate to the folder and download the dependencies :

```shell
cd ticket
go mod tidy
```

Now copy the `.env.example` file to `.env` file :

```shell
cp -iv .env{.example,}
```

fill the necessary variables in `.env` file, database information and connection settings .

Project is set up !

Now run the project :

```shell
go run main.go
```

Then you will see this :

```shell
Loaded .env file successfully ...
Connected to the database successfully ...
Migrating table users ...
Starting the server on port 8000

 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.52.8                   │ 
 │               http://127.0.0.1:8000               │ 
 │       (bound on host 0.0.0.0 and port 8000)       │ 
 │                                                   │ 
 │ Handlers ............ 11  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 34312 │ 
 └───────────────────────────────────────────────────┘ 

````
