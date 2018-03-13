# go-todos

A simple todo list with user accounts. 

- Golang with [Gin](https://github.com/gin-gonic/gin) and [Xorm](https://github.com/go-xorm/xorm) for the backend
- jQuery for a simple frontend without [JavaScript Fatigue](http://lucasfcosta.com/2017/07/17/The-Ultimate-Guide-to-JavaScript-Fatigue.html)
- Postgres for persistence
- Redis for user sessions
- Fluentd for log aggregation 
- Docker for building

# Setup

- Run `docker-compose -f docker-compose.yml -f docker-compose.prod.yml up`
- Go to [http://localhost:8080/](http://localhost:8080/).
- Register, create some todos, and be productive. 
