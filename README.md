# docker-compose-healthcheck

docker-compose depends_on: condition

## How can I wait for container X before starting Y?

By configuring a service `healthcheck` on X that is a (hard) dependency for Y together with a `depends_on` condition.

## Waiting for PostgreSQL to be "healthy"

A particularly common use case is a service that depends on a database, such as PostgreSQL.
We can configure docker-compose to wait for the PostgreSQL container to startup and be ready to accept requests before continuing.

```yaml
services:
  x:
    healthcheck:
      test:
      - CMD-SHELL
      - pg_isready -U postgres
      interval: 10s
      retries: 5
      timeout: 5s
    image: postgres:13.2-alpine
```

If the check is successful the container will be marked as `healthy`. Until then it will remain in an `unhealthy` state.

Services that depend on PostgreSQL can then be configured with the `depends_on` condition:

```yaml
services:
  y:
    depends_on:
      x:
        condition: service_healthy
```
