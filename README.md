<kbd>
  <img src="diagram-demo.png">
</kbd>


* Run Go on EC2
* Replace the IP on rds-regist.json

```shell script
consul agent -config-dir=config.json
```

```shell script
consul connect envoy -gateway=terminating \
-register -service rds-mysql-services-gateway \
-address '127.0.0.1:8443'
```

```shell script
consul connect envoy -sidecar-for ui-java -admin-bind 127.0.0.1:19100
```

```shell script
curl --request PUT --data @rds-regist.json localhost:8500/v1/catalog/register
```

```shell script
java -jar ui-java/target/nomad-shapshots-0.0.1-SNAPSHOT.jar
```