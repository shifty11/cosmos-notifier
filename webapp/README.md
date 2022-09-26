# DAO DAO Notifier Frontend


## Run envoy
Envoy has to run to proxy grpc-web requests to the grpc server. 
```
envoy -c envoy_dev.yaml     # or use the docker-compose file
```

## Run frontend
Add the following line to `/etc/hosts` (optional):
```
127.0.0.1     test.mydomain.com
```

Then run the frontend:
```
flutter run --web-port 40001 --web-hostname test.mydomain.com
```

## Build freezed classes

```bash
flutter pub run build_runner build
```

