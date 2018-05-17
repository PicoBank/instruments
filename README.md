# Prerequisites

You need to make sure your environment is setup up correctly:
 * [go](https://golang.org/) -- The Go Programming Language
 * [dep](https://golang.github.io/dep/) -- Dependency management for Go

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

Running in development mode:

    DB_USER=pi DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=picobank bee run

Connect to `http://localhost:8080/v1/instrument`

# Build and deploy on Raspberry PI

In the source directory:

    GOOS=linux GOARCH=arm GOARM=5 bee pack

Copy the instruments.tar.gz file to the target Raspberry, untar and run:

    sed -i 's/runmode\s*=\s*dev/runmode = prod/' conf/app.conf
    DB_USER=pi DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=go_test ./instruments

Connect to `http://localhost:8080/v1/instrument`

# Installation notes

Initializing the database:

    psql postgres -f migrations/drop.sql
    psql postgres -f migrations/create.sql
    psql postgres -f migrations/sample-data.sql

# Setup notes

Installing BeeGo:

    go get github.com/lib/pq
    go get github.com/astaxie/beego
    go get github.com/beego/bee

Code generation:

    cd $GOPATH/src/github.com/picobank
    bee api instruments -driver=postgresql

