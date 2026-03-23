# Project Structure

```
your-project/
│
├── internal/               # Internal application logic, not exposed to external use
│   ├── config/             # Application configuration (e.g., environment variables)
│   ├── controllers/        # API logic (Controllers)
│   ├── db/                 # Database access and config (Database)
│   ├── models/             # Database models and structs
│   ├── repository/         # Database access (Repositories)
│   └── service/            # Business logic (Services)
│
├── pkg/                    # Public reusable code (shared libraries)
│   ├── dotend/             # Dotenv functions
│   └── utils/              # Utility functions
│
├── migrations/             # Database migration files
│
├── test/                   # Unit and integration tests
│
├── main.go                 # Main function to start the app
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies checksum
└── README.md               # Documentation for your project
```

# Database

### GORM - Facilitate database interactions

Edit your database name in the following file `ms-{name}/internal/config/db.go`

```
var dbName = GetEnv("POSTGRES_DB") <- Change POSTGRES_DB for the right env name
```

### pressly/goose -> Database migration tool

Need to install goose to run the migrations https://pressly.github.io/goose/installation/

To create new migrations, run the following command
```
nx run $projectDir:create-migration --name=add-new-field //Specify the name of the migration
```

To run migrations, run the following command
```
nx run $projectDir:migrate
// Available args: --user, --host, --port, --password, --db, --sslmode
```

Other commands available [here](https://pressly.github.io/goose/documentation/cli-commands/)


### Keep GORM and the DB in sync

Even though we don't rely on GORM auto migration tool to manage our database, we should use GORM annotations to ensure that we reflect our database.
This means declaring indexes, primary key, table/column names and more. 

All this should reflect our database to avoid errors when inserting/updating/deleting information through GORM.

# Swagger

Swagger available at http://localhost:8087/swagger/index.html

How to define the API doc? Check [Declarative Comments Format](https://github.com/swaggo/swag?tab=readme-ov-file#declarative-comments-forma).


To generate the API doc, run the following command

```
nx run ms-tagpeak:openapi
```


# Libraries

- [GORM](https://gorm.io/docs/migration.html) -> ORM
- [pressly/goose](https://pressly.github.io/goose/) -> Database migration tool
- [Echo](https://echo.labstack.com/docs) -> Server
- [samber/lo](https://github.com/samber/lo) -> lo - Iterate over slices, maps, channels...
- [Swaggo | Echo-Swagger](https://github.com/swaggo/echo-swagger) -> OpenAPI 2.0 generation based on controllers

