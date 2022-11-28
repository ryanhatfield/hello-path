# Hello Path

A simple Hello World server for a custom root path, useful for testing path-based ingress routes.

## Help Info

```
Usage: hello-path [--port PORT] PATHROOT

Positional arguments:
  PATHROOT

Options:
  --port PORT, -p PORT [default: 8081, env: HELLO_PATH_PORT]
  --help, -h             display this help and exit
```

## Container spec

Example container spec, passing required positional argument:

```yaml
spec:
  containers:
  - name: hello-path
    image: ryanhatfield/hello-path:latest
    args: ['/hello-world']
    livenessProbe:
      httpGet:
        path: /hello-world/live
        port: http
        scheme: HTTP
    readinessProbe:
      httpGet:
        path: /hello-world/ready
        port: http
        scheme: HTTP
    ports:
    - name: http
      containerPort: 8081
      protocol: TCP
```