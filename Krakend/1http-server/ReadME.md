# Build a krakenD plugin
- `docker run -it -v "$PWD:/app" -w /app krakend/builder:2.2.1 go build -buildmode=plugin -o krakend-server-example.so .`