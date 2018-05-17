# Prerequisites

You need to make sure your environment is setup up correctly:
 * [go](https://golang.org/) -- The Go Programming Language
 * [dep](https://golang.github.io/dep/) -- Dependency management for Go
 * [docker](https://www.docker.com/) -- Build, Ship, and Run Any App, Anywhere


# Get the source

Clone the Git repository:

```bash
git clone https://github.com/picobank/instruments.git $GOPATH/src/github.com/picobank
cd $GOPATH/src/github.com/picobank
```


# Installing dependencies

Run `dep ensure` to make sure vendor/ is in the correct state:

```bash
dep ensure
```


# Running

Start the postgres instance from a dedicated terminal:

```bash
docker-compose up
```

Run database migrations tool to initialize and populate the `instruments` schema:

```bash
bee migrate -driver="postgres" -conn="postgres://instruments:raspberry@localhost:5432/picobank?sslmode=disable"
```

Run the application in development mode:

```bash
DB_USER=instruments DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=picobank bee run
```

Connect to `http://localhost:8080/v1/instrument`


# Prepare the Raspberry PI

Copy SQL scripts to the Raspberry PI:

```bash
scp database/*.sql pi@raspberrypi.local:~
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

Enable remote access to the Postgres instance:

```bash
sudo sed -i "s/#listen_addresses = 'localhost'/listen_addresses = '*'/" /etc/postgresql/9.6/main/postgresql.conf
cat << EOF | sudo tee -a /etc/postgresql/9.6/main/pg_hba.conf
host    picobank        instruments     0.0.0.0/0               md5
host    picobank        instruments     ::/0                    md5
EOF
sudo service postgresql restart
```


# Build and deploy on Raspberry PI

Run database migrations tool to initialize and populate the `instruments` schema:

```bash
bee migrate -driver="postgres" -conn="postgres://instruments:raspberry@raspberrypi.local:5432/picobank?sslmode=disable"
```

Package the application:

```bash
GOOS=linux GOARCH=arm GOARM=5 bee pack
```

Copy the instruments.tar.gz file to the Raspberry and untar:
   
```bash
scp instruments.tar.gz pi@raspberrypi.local:~
ssh pi@raspberrypi.local
sudo tar xzvf instruments.tar.gz
```

Run the application in production mode:

```bash
sudo sed -i 's/runmode\s*=\s*dev/runmode = prod/' conf/app.conf
sudo DB_USER=instruments DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=picobank ./instruments
```

Connect to `http://localhost:8080/v1/instrument`


# Setup notes

Installing BeeGo:

    go get github.com/lib/pq
    go get github.com/astaxie/beego
    go get github.com/beego/bee

Code generation:

    cd $GOPATH/src/github.com/picobank
    bee api instruments -driver=postgresql

