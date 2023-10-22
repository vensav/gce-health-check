# gce-health-check

A simple webserver that can be used for providing HTTP health check on GCE instances. Currently tested only on mac and linux


## Installation as sudo user

### Create service 

#### Example:- deploy github release 1.0.5 for linux_amd and let health-check run on port 1000
```
wget -O - https://raw.githubusercontent.com/vensav/gce-health-check/main/scripts/create_svc.sh linux_amd64  1.0.5 1000 | sudo bash
```

### Remove service
```
wget -O - https://raw.githubusercontent.com/vensav/gce-health-check/main/scripts/remove_svc.sh | sudo bash
```


## Build and test locally

### Build Examples
```
go tool dist list   # Get list of OS/Arch supported
env GOOS=linux GOARCH=amd64 go build -o ./build/gce-health-check-linux-amd64 main.go        # Linux x86-64
env GOOS=darwin GOARCH=arm64 go build -o ./build/gce-health-check-darwin-arm64 main.go      # Mac M1/M2
```

### Run on a different port (default to port 1000)
```
./gce-health-check-darwin-arm64 2000    
./gce-health-check-linux-amd64 2000    
```

### Response format
```
curl localhost:1000/health                  # Returns OK as text if service is up
curl localhost:1000/instance/name           # Returns instance name as text
curl localhost:1000/instance/external-ip    # Returns external ip attached to first nic card as text
```
Errors/Exceptions return non 200/201 response with error message in plain text


## CI Instructions
- Provide release tag manually while doing deploy
