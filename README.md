# screwdriver
![](https://raw.githubusercontent.com/takasing/screwdriver/master/data/warsman.jpg)

The tools to manage EC2 Container Service(ECS)

### Caution
This repository depends on AWS, ECS specification.

### Installing
```go
go get github.com/takasing/screwdriver
```

### Configuring
The configurations of `screwdriver` command consists of environment variables and `yml`

#### AWS environment variable
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_REGION

#### Yaml file
Enable to use Golang template.
```yml
nginx:
  image: nginx:{{ .NGINX_IMAGE_TAG }}
  cpu: 128
  memory: 64
  portMappings:
    - containerPort: 80
    - hostPort: 80
  essential: true

api:
  image: golang:{{ .GOLANG_IMAGE_TAG }}
  cpu: 128
  memory: 64
  portMappings:
    - containerPort: 3000
    - hostPort: 3000
  essential: true
```

#### Environment variables for yaml
`screwdriver` picks up environment variables start with `SCREW_` prefix.  
If you define `SCREW_NGINX_IMAGE_TAG` environment variable, you can use `NGINX_IMAGE_TAG` in yml.

### Usage
```
usage: screwdriver [--version] [--help] <command> [<args>]

Available commands are:
    task    Operate ECS task
```

#### task
```
Usage: screw task <subcommand> [options]
Subcommands:
        defs          show the list of task definitions
        register      register task definition from configration file
```
