# E-Ticketing API 🚍 

A simple API system for managing terminals, gates, transactions, and admin authentication using JWT. Built with Golang and Echo framework.

---

## Getting Started 🚀
- <h3>Clone this repository</h3>

  ```
  git clone https://github.com/dzarurizkyy/backend-development-test
  ```

- <h3>Configure Database and JWT in `config/config.yml`</h3>

  ```
  database:
  host:     localhost
  port:     5432
  user:     your_db_user
  password: your_db_password
  dbname:   your_db_name

  jwt:
  secret: your_jwt_secret
  ```

- <h3>Create Database and Tables</h3>

  ```
  psql -U your_db_user -d your_db_name -f database/table.sql
  ```

- <h3>Install Dependencies</h3>

  ```
  go get 
  go mod tidy
  ```

- <h3>Run Server</h3>

  ```
  go run main.go
  ```


## Documentation 📚
- [System Design](https://github.com/dzarurizkyy/backend-development-test/blob/main/graph/system-design.png)
- [Database Design](https://github.com/dzarurizkyy/backend-development-test/blob/main/graph/database-design.png)
- [Postman](https://documenter.getpostman.com/view/29935661/2sB34Zr4bA)

## Contributor 🤝
- [Dzaru Rizky Fathan Fortuna](https://www.linkedin.com/in/dzarurizky/)
