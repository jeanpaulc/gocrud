# gocrud
A tiny http server built with golang, mux, and gorm

## schema
### User
|        User       | Type          | JSON-field |
| ------------- | ------------- | ----- |
| ID (PK; Auto Increment)     | uint | id |
| LastName      | string      |   first_name |
| Email | string      |    last_name |
| Age | string     |    age |
| CreatedAt | time.Time      |    created_at |
| UpdatedAt | time.Time     |    updated_at |
