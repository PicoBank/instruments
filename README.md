# Running

    DB_USER=pi DB_PASSWORD=raspberry DB_HOST=localhost DB_NAME=go_test bee run

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

