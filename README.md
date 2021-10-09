# orchestrator-service

## Initializing Module

* Make sure you are in the github repository directory after unzipping it.
```
go mod init github.com/JeetKaria06/orchestrator-service
```

## Synchronize Dependencies

```
go mod tidy
```

## For first RPC service

### Running Server
```
go run server/main.go
```

> Expected Output ```2021/10/09 18:33:29 server listening at [::]:50001```

### Running Client
```
go run client/main.go
```

> Expected Output ```2021/10/09 18:33:48 could not get response because: rpc error: code = Unknown desc = not implemented yet. Jeet will implement me                                                                                                                exit status 1 ```

## For Dummy Data Service

### Running Server
```
go run datamock/server/main.go
```

> Expected Output ```2021/10/09 18:36:10 server listening at [::]:10000 ```

### Running Client

```
go run datamock/client/main.go
```

> Expected Output ```2021/10/09 18:36:42 name:"Jeet Karia"  class:"10"  roll:100```

## For Orchestrator Service

### Running Servers

```
go run logic/orchestrator1/main.go
go run logic/orchestrator2/main.go
go run datamock/server/main.go
```

> In above, orchestrator-1 instance will be listening on 9000, orchestrator-2 instance on 9001 and dummy data service on 10000

### Running Client

```
go run logic/client/main.go
```

> Expected Output in case of name having less than 6 characters ```2021/10/09 18:42:09 could not get response because: rpc error: code = Unknown desc = length of name is less than 6      exit status 1 ```

</br>

---
**NOTE**

Operating System Used for Developement = Windows
GOPATH = %USERPROFILE%\go --> C:\Users\i\go
Path of Repository in local = D:\Activities\orchestrator-service

---
