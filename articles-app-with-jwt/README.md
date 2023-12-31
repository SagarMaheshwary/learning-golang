# ARTICLES APP

Articles CRUD with jwt authentication build using Fiber, GORM, and PostgreSQL database.

### TABLES

Create tables that are required for this application:

```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL
);
```

```sql
CREATE TABLE articles (
    id BIGSERIAL PRIMARY KEY,
    slug VARCHAR(120) NOT NULL UNIQUE,
    title VARCHAR(120) NOT NULL,
    body TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### APIS

| URI               | METHOD | AUTHORIZATION         | REQUEST BODY          |
| ----------------- | ------ | --------------------- | --------------------- |
| /api/login        | POST   | -                     | email, password       |
| /api/register     | POST   | -                     | name, email, password |
| /api/profile      | GET    | JWT as "bearer" Token | -                     |
| /api/articles     | GET    | -                     | -                     |
| /api/articles     | POST   | JWT as "bearer" Token | title, body           |
| /api/articles/:id | GET    | -                     | -                     |
| /api/articles/:id | PUT    | JWT as "bearer" Token | title, body           |
| /api/articles/:id | DELETE | JWT as "bearer" Token | -                     |

### DEPENDENCIES

All the dependencies can be found in **go.mod** file. (which will be automatically downloaded when running the application.)

### RUNNING THE APPLICATION

Create a **.env** from **.env.example** in the project directory. Add your database credentials and jwt secret then start the application with **go run** command:

```bash
    go run main.go
```
