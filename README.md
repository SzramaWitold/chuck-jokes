# Chuck Jokes API

This app is designed for manage Chuck Norris Jokes

#Availabale commands for develop:

  `database:create`  Create new database

  `database:migrate` migrate database

  `database:seed`    seed database with fake data

  `help`             Help about any command

  `schedule:run`     schedule all job inside crone

  `test`             ...


#First start

1. create .env file base on .env_example
2. run `database:create` create new database
3. run `database:migrate` migrate models
4. run `database:seed` seed database with fake data

# run app
```
go run main.go
```

# Get swagger docs
```
http://{domain}:{port}/swagger/index.html#/
```

# Update swagger docs
```
swag init
```