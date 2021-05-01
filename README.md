# GO Gin-MVC template
## Get started
### clone project
```bash
$ git clone https://github.com/Nanshin-Seisakujo/gin-mvc-template
```

### install
```bash
$ go get
$ yarn install
```

### dev command
```bash
# defailt
$ go run main.go

# with air
$ air

# client develop
$ yarn dev
```

### build command
```bash
# server
$ go build

# client
$ yarn build
```

## Directory Structure
```
├── app
│    ├── config
│    │    ├── config.go
│    │    ├── development.yaml
│    │    └── production.yaml
│    ├── controllers
│    │    └── <name>.go
│    ├── db
│    │    └── db.go
│    ├── forms
│    │    └── <name>.go
│    ├── middlewares
│    │    └── <name>.go
│    ├── models
│    │    └── <name>.go
│    ├── modules
│    │    └── <name>.go
│    ├── server
│    │    ├── router.go
│    │    └── server.go
│    └── views
│         ├── index
│         │    └── index.tmpl
│         └── templates
│              └── <name>.tmpl
├── src
│    ├── sass
│    │    └── main.scss
│    └── typescript
│         └── main.ts
├── static
│    ├── css
│    │    └── main.css
│    ├── images
│    │    └── <somefile>
│    └── js
│         └── main.bundle.js
├── .air.toml
├── .browserslistrc
├── .env
├── .eslintignore
├── .eslintrc.js
├── .prettierignore
├── .prettierrc
├── go.mod
├── main.go
├── package.json
├── postcss.config.js
├── tsconfig.json
└── webpack.config.js
```

### /app/config
Describe the settings such as port and schema name.

### /app/controllers
Write functions to pour the data defined in model into view and insert data from view into the database.

### /app/db
Define functions for database configuration and connection.

### /app/forms
Describes schema validation of data.

### /app/middlewares
Define the middleware, for example, if authentication is required.

### /app/models
Define the schema to be used for the database tables.

### /app/modules
Create a module for common use.

### /app/server
This is files for start the server and create routes.

### /app/views
Save the html template.

### /src
Used as an entry point for pre-compiled sass and typescript files.

### /static
Used to serve as css and javascript files after compilation, and static files such as images.

### main.go
The executable file of the web server.