# JOB PORTAL APP GO BACKEND
[zipkin](http://localhost:9411/)

# TECH STACK
1. elasticsearch for searching needs with lots of data and also lots of user requests
2. redis for caching data that is frequently retrieved and easy to cache, such as no search feature etc. in this case it is used for caching company data
3. otel, zipkin. for distributed tracing needs, easy to track request flow
4. Kafka for ETL processes, so the raw data which was originally relational is combined into one document and saved to elastic
5. using proto for contract event driven

# HOW TO RUN
1. init for install dependency, create topic, generated mock, generated api.yml, generated proto buf, and migrations up
```shell
make init
```
2. to view the api contract, run
```shell
make preview_open_api
```
and open [preview open api](http://127.0.0.1:8080/)

4. to run rest api, run
```shell
make run_rest_api
```
and the [base](http://127.0.0.1:8282/) url is running at http://localhost:8282

5. to run consumer etl, run
```shell
make run_consumer
```