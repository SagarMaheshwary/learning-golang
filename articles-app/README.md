# ARTICLES APP

CRUD api build using Fiber framework and PostgreSQL database.

### DEPENDENCIES

```bash
go get github.com/gofiber/fiber/v2
```

```bash
go get github.com/lib/pq
```

```bash
go get github.com/gofor-little/env
```

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

NOTE: we are not using the users table but it's built in a way that we can easily add user authentication in this app.

### RUNNING THE APPLICATION

Create a **.env** from **.env.example** in the project directory and add your database credentials.
Start the application with **go run** command:

```bash
    go run main.go
```
