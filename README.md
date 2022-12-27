# griffin-backend 

<p>
"griffin backend" v2. Escape from redis and json type database. Use ent-go(developed by Meta) and gin to create a griffin web2 backend server.
</p>

## Execute

First generate `.env.serve` file that contains server information.

```console
$ docker-compose up -d
```
This will generate docker container and run it detached.

## Execute locally

First rewrite `.env.local` file to fit your local machine. After writing your environment file

```console
$ go run .
```


## API Documents

<p>

API document is written in Swagger. To access swagger doc, use `/swagger` endpoint
> swagger/index.html#
> 
> Swagger Document Embedding Access

</p>
 