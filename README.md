# envecho

ENV Echo is a go webserver to serve static response, configurable via environment variables.

- `ENVECHO_HOST`: host to accept requests, optional, default to `0.0.0.0`
- `ENVECHO_PORT`: port to accept requests, optional, default to `80`
- `ENVECHO_MESSAGE`: response message as text/plain, optional

```
docker run -p 8080:80 tatchnicolas/envecho:latest
```
