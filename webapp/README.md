# webapp

A new Flutter project.

## Run envoy
Envoy has to run to proxy grpc-web requests to the grpc server. 
```
envoy -c envoy_dev.yaml     # or use the docker-compose file
```

## Build freezed classes

```bash
flutter pub run build_runner build
```

