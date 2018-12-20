# Ship ports
Example of simple two layers micro-service architecture

## Try it out
Upload file:
~~~bash
curl --request POST \
  --url http://localhost:8081/api/ports/import \
  --form file="@somefile.dat"
~~~

Get existing data from DB:
~~~
curl --request GET \
  --url 'http://localhost:8081/api/ports?limit=200&offset=0'
~~~

## How to run

### Via Docker Compose
Easiest way to run whole infrastructure is to use _docker compose_:
~~~bash
docker-compose up
~~~

### Manual
If you need to run services separately:

**Gateway:**
~~~bash
go run ./entry/entry.go --kind=port-gtw --address=:8080
~~~

**Service:**
~~~bash
go run ./entry/entry.go --kind=port-srv --address=:10001
~~~

#### Note: if running manually, it's need to pass env. vars and provide Mongo DB
Example of _.env_ file:
~~~bash
MONGO_URL=:27017

GATEWAY_PORT_ADDRESS=:8080
SERVICE_PORT_ADDRESS=:10001

# Max. file size in bytes
PORT_MAX_IMPORT_FILE_SIZE=209715200
~~~

## Test
To run all available tests:
~~~bash
make run-tests
~~~

## Linting
You need to install linter at first (**ONLY ONCE**):
~~~bash
make install-linter
~~~

To run linter, use next command:
~~~bash
make run-linter
~~~

## Points of improvement
 - Separate 'City, Country etc from Port domain to other service' (task says only about 2 services)
 - Move to simplified ID, not connected with real port ID
 - Injection graph initialization  (overhead for those services)
 - Slice copy instead of setting as a value