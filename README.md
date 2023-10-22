# gce-health-check

A simple webserver that can be used for providing HTTP health check on GCE instances

## Build and compile
```
go tool dist list   # Get list of OS/Arch supported
env GOOS=linux GOARCH=amd64 go build -o ./build/gce-health-check-linux-amd64 main.go        # Linux x86-64
env GOOS=darwin GOARCH=arm64 go build -o ./build/gce-health-check-darwin-arm64 main.go      # Mac M1/M2
```

## Run on a different port (default to port 1000)
```
./gce-health-check-darwin-arm64 2000    
./gce-health-check-linux-amd64 2000    
```

## Response format
```
curl localhost:1000/health                  # Returns OK as text if service is up
curl localhost:1000/instance/name           # Returns instance name as text
curl localhost:1000/instance/external-ip    # Returns external ip attached to first nic card as text
```
Errors/Exceptions return non 200/201 response with error message in plain text

