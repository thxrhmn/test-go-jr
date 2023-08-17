## TEST GO JR

### PREQUISITES
- Go v.120 or latest
- Git / Git bash (windows)
- PostgreSQL

### FEATURES
- Pagination
- Upload file only `.txt` and `.pdf` files allowed
- Filter by `title` and `description`
- Swagger
- Migration DB
- And more

### TECH & LIB
- Go & Echo Framework
- GORM
- PostgreSQL

### HOW TO USE?
```
git clone https://github.com/thxrhmn/test-go-jr.git
cd test-go-jr
go get
```

```
cp .env_example .env
```
Open the `.env` file and you can change `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_HOST`, and `DB_PORT` or you can adjust to what I wrote below. `SECRET_KEY` and `PORT` do not need to be changed. Make sure everything matches your PostgreSQL database settings.

```
SECRET_KEY=rahman
PORT=5000
DB_NAME=todolist
DB_USER=postgres
DB_HOST=localhost
DB_PASSWORD=root
DB_PORT=5432
```

### RUN 
```
go run main.go
```

### ADD EXAMPLE DATA
example data will be automatically added to database
```
chmod +x example-data.sh
./example-data.sh
```
### SWAGGER
open this link in your browser http://localhost:5000/swagger/index.html#/
you can test all enpoints using swagger

![The San Juan Mountains are beautiful!](swagger.png)

### POSTMAN COLLECTION
or you can use my postman collection https://www.postman.com/docking-module-geologist-95497667/workspace/thxrhmn/collection/26563934-734ad859-4500-4ec5-a328-4a3740ed0f94?action=share&creator=26563934

![The San Juan Mountains are beautiful!](postman.png)