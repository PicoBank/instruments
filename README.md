# Running

Running in development mode:

    DB_USER=pi DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=go_test bee run

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

    CREATE TABLE public.instruments (
        code character varying,
        value numeric
    );

    GRANT SELECT ON TABLE public.instruments TO pi;

    INSERT INTO public.instruments VALUES ('NYC', 1.34);
    INSERT INTO public.instruments VALUES ('NZU', 3.14);

# Setup notes

Installing BeeGo:

    go get github.com/astaxie/beego
    go get github.com/beego/bee

Code generation:

    cd $GOPATH/src/github.com/picobank
    bee api instruments -driver=postgresql

