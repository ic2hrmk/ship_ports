# Ship ports
Example of simple micro-service architecture

## Try it out
~~~bash
curl --request POST \
  --url http://localhost:8081/api/ports/import \
  --form file="@somefile.dat"
~~~


### Points of improvement
 - Separate 'City, Country etc from Port domain to other service' (task says only about 2 services)
 - Move to simplified ID, not connected with real port ID
 - Injection graph initialization  (overhead for those services)
 - Slice copy instead of setting as a value