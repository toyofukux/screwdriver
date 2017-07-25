# screwdriver
**Deprecated** This repository is no more mentenanced.  

![](https://raw.githubusercontent.com/takasing/screwdriver/master/data/warsman.jpg)

The tools to manage [EC2 Container Service(ECS)](http://aws.amazon.com/ecs/details/)  

Today, we use [Docker](https://www.docker.com/) containers to develop Web Services, and we get start moving towards an Immutable Infrastructure with ECS and `screwdriver`!!  

### Caution
This repository depends on AWS, ECS specification.

### Future
- [ ] task
  - [x] show the list of task definitions
  - [x] register task new definition
  - [ ] describe task definition
- [ ] service
  - [x] show the list of services
  - [x] create service
  - [x] update service
  - [ ] describe service
- [x] cluster
  - [x] show the list of clusters
  - [x] create cluster
  - [x] delete cluster
- [ ] Auto Scaling
- [ ] ELB
- [ ] Blue-Green Deployment

And you have some suggestion, create new Issue :)

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
  portmappings:
    - containerport: 80
      hostPort: 80
  essential: true

api:
  image: golang:{{ .GOLANG_IMAGE_TAG }}
  cpu: 128
  memory: 64
  portmappings:
    - containerport: 3000
      hostport: 3000
  essential: true
```
Check [ecs.ContainerDefinition](https://github.com/aws/aws-sdk-go/blob/master/service/ecs/api.go#L1002-L1094), and **Lowercase** fields such as `PortMappings` to `portmappings`

#### Environment variables for yaml
`screwdriver` picks up environment variables start with `SCREW_` prefix.  
If you define `SCREW_NGINX_IMAGE_TAG` environment variable, you can use `NGINX_IMAGE_TAG` in yml.

### Usage
```
usage: screwdriver [--version] [--help] <command> [<args>]

Available commands are:
    cluster    Operate ECS cluster
    service    Operate ECS service
    task       Operate ECS task
```

#### task
```
Usage: screwdriver task <subcommand> [options]
Subcommands:
        defs          show the list of task definitions
        register      register task definition from configration file
```

#### service
```
Usage: screwdriver service <subcommand> [options]
Subcommands:
        list        show the list of ECS services
        create      create ECS service
        update      update ECS service
```

#### cluster
```
Usage: screwdriver cluster <subcommand> [options]
Subcommands:
        list        show the list of ECS cluster
        create      create ECS cluster
        delete      delete ECS service
```
