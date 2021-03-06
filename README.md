# Prerequisites

You need to make sure your environment is setup up correctly:
 * [go](https://golang.org/) -- The Go Programming Language
 * [dep](https://golang.github.io/dep/) -- Dependency management for Go
 * [docker](https://www.docker.com/) -- Build, Ship, and Run Any App, Anywhere


# Get the sources

Clone the Git repository:

```bash
git clone https://github.com/picobank/instruments.git
cd instruments
```

# Preparing the development environment


## MacOSX
    
    go get -u github.com/golang/dep/cmd/dep

## Docker

Start the development environment:

```bash
docker-compose up -d
```

The `app` container is bundled with:
 * dep -- dependency manager
 * beego -- application framework
 * goose -- database migration tool

# Running the development environment 

Run `dep ensure` to make sure vendor/ is in the correct state.

To regenerate SQLBoiler models: 

```bash
go generate
```

To test the generated code:

```bash
go test -v ./models/ -args -test.config=../database/sqlboiler.toml
```

## MacOSX

TODO

## Docker 

```bash
docker-compose exec app dep ensure
```

Run `goose` to initialize the `instruments` schema:

```bash
docker-compose exec app goose -dir database/migrations postgres "postgres://instruments:raspberry@db:5432/picobank?sslmode=disable" up
```

Run the application in development mode:

```bash
docker-compose exec app go run main.go
```

Connect to `http://localhost:8080/v1/instruments`

Stop the development environment:

```bash
docker-compose down
```

# Prepare the Raspberry PI

Copy SQL scripts to the Raspberry PI:

```bash
scp database/initdb/*.sql pi@raspberrypi.local:~
```

Connect to the Raspberry PI:

```bash
ssh pi@raspberrypi.local
```

Install Postgres on Rasbian:

```bash
sudo apt install -y postgresql libpq-dev postgresql-client postgresql-client-common
```

Create the database and the schema:

```bash
sudo -u postgres psql -f database.sql
sudo -u postgres psql -f schema.sql
```

Enable remote access to the PostgreSQL instance:

```bash
sudo sed -i "s/#listen_addresses = 'localhost'/listen_addresses = '*'/" /etc/postgresql/9.6/main/postgresql.conf
cat << EOF | sudo tee -a /etc/postgresql/9.6/main/pg_hba.conf
host    picobank        instruments     0.0.0.0/0               md5
host    picobank        instruments     ::/0                    md5
EOF
sudo service postgresql restart
```

# Build for and deploy on Raspberry PI

Run `goose` to initialize the `instruments` schema:

```bash
docker-compose exec app goose -dir database/migrations postgres "postgres://instruments:raspberry@raspberrypi.local:5432/picobank?sslmode=disable" up
```

Package the application:

```bash
docker-compose exec -e GOOS=linux -e GOARCH=arm -e GOARM=5 app bee pack
```

Copy the instruments.tar.gz file to the Raspberry and untar:
   
```bash
scp instruments.tar.gz pi@raspberrypi.local:~
ssh pi@raspberrypi.local
sudo tar xzvf instruments.tar.gz
```

Run the application in production mode:

```bash
sudo BEEGO_RUNMODE=prod DB_USER=instruments DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=picobank ./instruments
```

Connect to `http://raspberrypi.local:8080/v1/instruments`

# Load testing (WIP)

    $GATLING_HOME/bin/gatling.sh  -sf loadtests/simulations/ -bdf loadtests/bodies/ -df loadtests/data -m