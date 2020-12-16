# Link Rewriter

Simple HTTP server to do selective redirects, based on the hostname and an optional Path. 

## Get started

```
$ go build -o rewriter
$ # create config.yaml
$ ./rewriter
```

## Overview

This service rewrites Deep Links so that static web sites can handle URI parameters without breaking.

A URL like `https://vocdoni.link/entities/0x1234` can be easily handled by a mobile app. However, a static web site could load the `/entities/` path but there is no straighforward way to tell static web servers to ignore the `0x1234` parameter. So the `/entities/0x1234` would be attempted to load, and users would get a 404.

In addition, deep links like `https://vocdoni.link/entities/0x1234` may need to be redirected to `app.vocdoni.net` whereas deep links like `https://vocdoni.link/poll/...` may need to point to a different server.

This server allows to fulfill these scenarios. 

## Configuration

Create a `config.yaml` file on the same path as the server.

```yaml
cmd:
  port: 8080
  verbose: true
replacements:
  - source: localhost:8080
    target: target.domain.net
    path: /
  - source: source2.domain
    target: target.domain2.net
    path: /validation
```

For each domain and path that you want to capture, add an item to the `replacements` array.

In the example: 
- Calling `http://localhost:8080` would redirect to `https://target.domain.net`
- Calling `https://source2.domain/validation/0x1234` would redirect to `https://target.domain2.net/validation#/0x1234`

Make sure that the trailing slashes are correctly defined in the config file, and that the source URL's match the expected path.

If a replacement defined contains an empty `path` or it is set to `/` then, the entire URI path will be rewritten in the hash. 

