ad-space-auction/
  ├── supply/
  │   ├── api/
  │   │   ├── handlers/
  │   │   │   └── adspace_handler.go
  │   │   └── models/
  │   │       └── ad_space.go
  │   ├── database/
  │   │   └── mysql.go
  │   ├── Dockerfile
  │   ├── main.go
  │   ├── go.mod
  │   └── go.sum
  ├── demand/
  │   ├── api/
  │   │   └── handlers/
  │   │       └── ...
  │   ├── database/
  │   │   └── ...
  │   ├── Dockerfile
  │   ├── main.go
  │   ├── go.mod
  │   └── go.sum
  ├── auction/
  │   ├── api/
  │   │   └── handlers/
  │   │       └── ...
  │   ├── database/
  │   │   └── ...
  │   ├── Dockerfile
  │   ├── main.go
  │   ├── go.mod
  │   └── go.sum
  ├── docker-compose.yml
  ├── go.mod
  └── go.sum
