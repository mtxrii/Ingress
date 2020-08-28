# Ingress
A sample product hub backend designed for use with a PostgreSQL database and a visual product hub frontend static site

# API
This app handles a list of products, each with three fields: `id`, `name` & `price`.

The HTTP requests currently available are as follows...

`/products/{key}` - GET: returns a list of saved products

`/product/{key}` - POST: saves new product from json in database

`/product/{id}/{key}` - PUT: updates a product of given id

`/product/{id}/{key}` - GET: returns a product of given id

`/product/{id}/{key}` - DELETE: removes a product of given id

`{key}` refers to whatever API key you specify in your `auth.env`

# Usage
### Create an `auth.env`

Follow this template:
```
APP_DB_USERNAME=database username
APP_DB_PASSWORD=database user password
APP_DB_SERVER=link to database
APP_DB_NAME=database name (probably same as username)

APP_KEY=key for verifying frontend interaction (you make this up)
```

### If running locally...

Download [executable](https://github.com/mtxrii/Ingress/releases) and run in the same directory as your `auth.env`

### If deploying to Heroku...

Create simple `Procfile` and `.godir` files.

Procfile:
```
web: ingress
```

.godir:
```
ingress
```

Put these files, as well as your `auth.env`, in your working directory, then follow instructions on [Heroku](https://devcenter.heroku.com/articles/git)
