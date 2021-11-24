This project provides a template to start a rest application in Go. It provides following integration:

##### Rest Api server

1. Graceful shutdown
2. Recovery

##### Dependency Injection

1. Uber Fx integration for dependency injection

##### Messaging integration

1. Support of Kafka
2. Support of SQS

##### HTTP integration using Gox-HTTP

1. Gox Http provides utility to call a http endpoint. It provides following:
2. Define all endpoint and api config in configuration file
3. Circuit breaker using Hystrix
    1. Set concurrency for each api - this ensures that if we go beyond "concurrency" no of parallel requests then
       hystrix will reject the requests
    2. Set timeout for each api - the call will timeout if this request takes time > timeout defined
    3. acceptable_codes - list of "," separated status codes which are acceptable. These status codes will not be
       counted as errors and will not open hystrix circuit

---

## Code Breakup

##### Build and config

Build folder has command to build and run the code
1. sh build/build.sh -> build the code 
2. sh build/run.sh -> run the code

Config folder has a sample config:
1. Gox-Http config to configure a client
2. Messaging config to configure kafka integration

##### Entry point

```cmd/server/main.go``` is the main entry point for this app. This calls the command ```cmd/server/app/main.go```,
where we bind all the modules using Uber Fx.

```NewApplicationEntryPoint``` method is the main lifecycle method which will start eh rest app
server. ```cmd/server/app/server.go```
has the code where you can put all your API endpoints. This framework use Gin but you can replace it if you need.

##### Rest API handlers

As mentioned above, the rest apis are registered at ```cmd/server/app/server.go```. However, the full implementation of
the Rest apis are created in ```interna/handler/*.go```. You can add any new handlers here, and register them to
in ```cmd``` to keep ```cmd``` lightweight.

We have provided 2 sample handler:
1. user handler - this shows an approach where handlers are encapsulated in a struct.
2. random handler - this shows how to can add any other handler (approach 1 is recommended)

##### Pkg - Clients
This contains the code to talk to any other client. You should keep generic code here. Anyone else should be able to import 
```.../pkg/clients``` to re-use the code.

We have given a example of ```json place holder```, and provided how to use gox-http to make http calls.

##### Pkg - core
bootstrap - all code which needs at the time of booting your application.

##### Pkg - server
A re-usable implementation for a rest server