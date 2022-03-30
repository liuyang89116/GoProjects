## Test

### Post

```shell
curl -X POST http://localhost:8080/book -H "Content-Type:application/json" -d "{\"id\":\"123\", \"name\":\"How to pl ay soccer\", \"authors\":[\"Messi\", \"Ronaldo\"], \"press\":\"Chinese Soccer Association\"}"
```

### GET

#### GetId

```shell
curl -X GET -H "Content-Type:application/json" http://localhost:8080/book/123
```
* Response

```
[{"id":"123","name":"How to play soccer","authors":["Messi","Ronaldo"],"press":"Chinese Soccer Association"}]

```

#### GetAll

```shell
 curl -X GET -H "Content-Type:application/json" http://localhost:8080/book
```

* Response

```
[{"id":"123","name":"How to play soccer","authors":["Messi","Ronaldo"],"press":"Chinese Soccer Association"},{"id":"125","name":"How to swim","authors":["Messi","Ronaldo"],"press":"Chinese Soccer Association"}]
```
