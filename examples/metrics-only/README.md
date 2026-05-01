# Metrics Only Example

Generate a metrics remote_write starter:

```sh
alloygen generate --logs none --metrics mimir --traces none --remote-write-url http://mimir-nginx.monitoring.svc:80/api/v1/push --output ./metrics-only
```

