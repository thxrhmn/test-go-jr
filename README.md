## TEST GO JR

### PREQUISITES:
- go v.120
- git 
- postgresql (pg admin)


### HOW TO USE?
```
git clone https://github.com/thxrhmn/test-go-jr.git
cd test-go-jr
go get
```

```
cp .env_example .env
```
Open the `.env` file and you can change `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_HOST`, and `DB_PORT` or you can adjust to what I wrote below. Make sure everything matches your PostgreSQL database settings.

```
SECRET_KEY=rahman
DB_NAME=todolist
DB_USER=postgres
DB_HOST=localhost
DB_PASSWORD=root
DB_PORT=5432
PORT=5000
```

### RUN 
```
go run main.go
```
open this link in your browser

http://localhost:5000/swagger/index.html#/

you can test all enpoints using swagger
![The San Juan Mountains are beautiful!](swagger.png)
