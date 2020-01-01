# Photon
**Photon** is a platform where students can fetch their lab materials easier than ever.

In addition, **Photon** contains more functionalities like:
* Lab management
* Virtual classrooms
* Statistics

## The platform
The **Photon** platform is currently subdivided into 4 modules:
* **Photon** for _Students_ is intended to help students on their everyday. Improves material fetching and class management.
* **Photon** for _Teachers_ helps to manage classrooms and tasks/homework.
* **Photon** for _Labs_ is a sub-system where qualified personnel is able to manage warehouses.
* **Photon** for _Managers_ is a little module that retrieves useful data for decision-makers.

## The API
**Photon** API v1 is using the _Microservice_ pattern working along with the HTTP Protocol, turning it into a REST Microservice.

The API is not completely attached to the HTTP Protocol, it also uses Websockets and MQ Brokers for every need.

See the GODOC reference to see all the exposed resources -aka URIs-.

## API Dependencies
Since **Photon** is using Go Modules for package management, the following dependencies are related to environment requirements.

**Photon** API v1 was developed with the following Programming Language(s) / Tool(s) versions:
* Go 1.13
* PostgreSQL 12
* Redis 5.0.7

## Architecture
**Photon** API's architecture is completely based on unclebob's clean architecture concept and is also using some concepts from bxcodec's Golang's clean architecture like Delivery layer.

If you want for information, check this [article](https://pusher.com/tutorials/clean-architecture-introduction) and this [GitHub Repo](https://github.com/bxcodec/go-clean-arch).

### Diagrams
![alt text](https://images.ctfassets.net/1es3ne0caaid/2zvDDUcdpuYqIM06WgU2sC/d706d509886f88be185fa007f6b43402/clean-architecture-ex-4.png "Clean architecture schema")
##### _unclebob's Clean Architecture_
![alt text](https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png "Clean architecture for Go")
##### _bxcodec's Golang's Clean Architecture_

### Code
**Photon** API v1 is currently using multiple architectrual / design patterns (repository, dependency injection, singleton, etc).

* **_Bin_**: Intended to start all the required services -persistence, cron jobs, routers, env variables, etc-.
* **_Core_**: Wraps the main application logic into one single package. It's meant to be reused during the application's lifecycle.
  * _Config_: Application's global configurations.
  * _Helper_: Wrapper for all 3-rd party services/libraries.
  * _Middleware_: Extra handlers as middlewares.
  * _Util_: Useful and completely reusable functions.

* **_Entity_**: Contains all the entities required by business logic.
* **_Infrastructure_**: Contains all volatile modules.
  * _Delivery_: Protocols for data delivery.
  * _Handler_: Functions triggered by delivery's layer patterns.
  * _Repository_: Handles all persistence's layer operations.
  * _Service_: Services required by the application.
* **_Usecase_**: Main Business logic's Use cases

## Deploy
**Photon** is inteded to use Travis CI, Docker, Docker Compose and Kubernetes for automated deploy & CI.

It is recommended to use AWS infrastructure using the following services:
* _Internet Gateway_.- NAT Pool and ACLs for VPC.
* _VPC_.- Virtual Private Cloud, isolates instances and services into a private cloud.
* _EC2_.- Server instance service on the cloud, runs **Photon**'s API binaries.
* _RDS_.- Main Persistence layer, Relational Database instance service using PostgreSQL.
* _ElastiCache_.- In-Memory persistence layer, uses Redis.
* _Lambda_.- Cloud functions, triggers before and after a request -HTTP requests mostly-.
* Other

## Maintanance
**Photon** is mainly backed by Damascus Mexico Group, however, this project is currently under the following users' responsabilities:
  * [@maestre3d](https://github.com/maestre3d)
