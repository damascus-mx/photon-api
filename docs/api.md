# API
**Photon** API v1 contains various resources exposed for clients and other microservices within the **Photon** environment.

Currently, **Photon** API is using nomenclatures and standards as _Google_ dictates in their [article](https://cloud.google.com/apis/design/design_patterns).

At the current stage, **Photon** is meant to use the following URL.

_**https://api.DOMAINNAME.com/photon/v1**_

The following sections will show you all the exposed URIs for every microservice.

## User Microservice

**HTTP**

| Action        | Resource        | URI          | Description     | Needs authorization?     |
|:------------:|-----------------|:------------:|:---------------:|:---------------:|
| **POST**      | Create user      | /user         | Creates a new user             | _NO_ |
| **GET**      | Get all users | /user  | Obtain all users by pagination      | _YES_ |
| **GET**      | Get user by ID     | /user/:id          | Obtain a user with an ID     | _YES_ |
| **PUT**      | Update user by ID     | /user/:id          | Updates a user with an ID     | _YES_ |
| **DELETE**      | Delete user by ID     | /user/:id          | Deletes a user with an ID     | _YES_ |


## Task Microservice

**HTTP**

| Action        | Resource        | URI          | Description     | Needs authorization?     |
|:------------:|-----------------|:------------:|:---------------:|:---------------:|
| **POST**      | Create task      | /task         | Issues a new task             | _YES_ |
| **GET**      | Get all tasks | /task  | Obtain all tasks by pagination      | _YES_ |
| **GET**      | Get task by ID     | /task/:id          | Obtain a task with an ID     | _YES_ |
| **PUT**      | Update task by ID     | /task/:id          | Updates a task with an ID     | _YES_ |
| **DELETE**      | Delete task by ID     | /task/:id          | Deletes a task with an ID     | _YES_ |
