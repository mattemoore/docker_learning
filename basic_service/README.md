# basic_service
Basic load balanced web site using Kubernets and Docker containers

## Requirements
Host w/ Docker and Kubernetes installed

## Usage
```
$ chmod u+x delete-basic-service.sh 
$ chmod u+x deploy-basic-service.sh 
$ chmod u+x stress-test.sh
```

Note that different hosts provided the responses

```
$ ./deploy-basic-service.sh
$ ./stress-test.sh
$ ./delete-basic-service.sh
```