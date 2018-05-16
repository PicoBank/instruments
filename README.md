# Get the source

    mkdir $GOPATH/src/github.com/picobank
    cd $GOPATH/src/github.com/picobank
    git clone git@github.com:picobank/instruments.git

# Running

Running in development mode:

    DB_USER=pi DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=picobank bee run

Connect to `http://localhost:8080/v1/instrument`

# Build and deploy on Rasbpberry PI

In the source directory:

    GOOS=linux GOARCH=arm GOARM=5 bee pack

Copy the instruments.tar.gz file to the target Raspberry, untar and run:

    sed -i 's/runmode\s*=\s*dev/runmode = prod/' conf/app.conf
    DB_USER=pi DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=go_test ./instruments

Connect to `http://localhost:8080/v1/instrument`

# Installation notes

Initializing the database:
    CREATE USER pi PASSWORD 'raspberry';
    ALTER ROLE pi SUPERUSER;
    CREATE DATABASE picobank OWNER pi;
    // connect as user pi: psql -U pi picobank
    CREATE SCHEMA instruments;

    CREATE TABLE instruments.instruments (
        code character varying,
        value numeric
    );

    GRANT SELECT ON TABLE instruments.instruments TO pi;

    INSERT INTO instruments.instruments VALUES ('NYC', 1.34);
    INSERT INTO instruments.instruments VALUES ('NZU', 3.14);

# Setup notes

Installing BeeGo:

    go get github.com/lib/pq
    go get github.com/astaxie/beego
    go get github.com/beego/bee

Code generation:

    cd $GOPATH/src/github.com/picobank
    bee api instruments -driver=postgresql

