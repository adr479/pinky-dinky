# Pinky-Dinky

Small service with endpoint to generate CPU and memory resource consumption.

## Build
To build the service, run:
```
docker build -t pinky-dinky .
```

## Run
Run the service:
```
docker run -p 8000:8000 pinky-dinky
```

## Use

### Load CPU
Try using up to 10 CPUs for 5 minutes:
```
curl -v localhost:8000/cpu
``` 

### Load Memory
Try using up to 2GB of memory:
```
curl -v localhost:8000/cpu
``` 

### Checking consumption
To track resource usage, run:

```
docker stats
```