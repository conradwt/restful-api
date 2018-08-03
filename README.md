# Zero to Restful API using Buffalo

The purpose of this example is to provide details as to how one would go about creating a Restful API with the Buffalo Web Framework.

## Getting Started

## Software requirements

- [Go 1.10.3 or higher](https://golang.org/dl)

- [Buffalo 0.12.4 or higher](https://gobuffalo.io/en/docs/installation)

- [PostgreSQL 10.4.0 or higher](https://www.postgresql.org/download)

- [Postman 6.1.4 or higher](https://www.getpostman.com)

## Communication

- If you **need help**, use [Stack Overflow](http://stackoverflow.com/questions/tagged/go). (Tag 'go')
- If you'd like to **ask a general question**, use [Stack Overflow](http://stackoverflow.com/questions/tagged/go).
- If you **found a bug**, open an issue.
- If you **have a feature request**, open an issue.
- If you **want to contribute**, submit a pull request.

## Quick Installation

1.  clone this repository

    ```text
    $ git clone git@github.com:conradwt/zero-to-restful-api-using-buffalo.git
    ```

2.  change directory location

    ```text
    $ cp /path/to/zero-to-restful-api-using-buffalo restful_api
    $ cd /path/to/restful_api
    ```

3.  create, migrate, and seed the database

    ```text
    $ buffalo db create
    $ buffalo db migrate
    $ buffalo task db:seed
    ```

4.  start the server

    ```text
    $ buffalo dev
    ```

5.  navigate to the Postman application

    a. set the HTTP request type to

    ```text
    GET
    ```

    b. set the request URL to

    ```text
    http://localhost:3000/players
    ```

    c. click `Headers`, add the following key and value pair:

    Key

    ```text
    Content-Type
    ```

    Value

    ```text
    application/json
    ```

    d. click `Send`, and you should see something like the following:

    ```json
    [
      {
        "id": "39bac424-8734-4705-a06b-8a20c1bfd55c",
        "created_at": "2018-08-03T00:14:51.941791Z",
        "updated_at": "2018-08-03T00:14:51.941792Z",
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com",
        "age": 25,
        "position": "forward"
      },
      {
        "id": "7e46a404-b361-4e3c-b16d-73ebd779c28f",
        "created_at": "2018-08-03T00:14:51.953796Z",
        "updated_at": "2018-08-03T00:14:51.953797Z",
        "first_name": "Jane",
        "last_name": "Doe",
        "email": "jane.doe@example.com",
        "age": 23,
        "position": "goalkeeper"
      }
    ]
    ```

# Tutorial Installation

```text
WIP
```

## Production Setup

Ready to run in production? Please [check our deployment guides](https://gobuffalo.io/en/docs/building).

## Buffalo References

- Official website: https://gobuffalo.io
- Guides: https://gobuffalo.io/en/docs/goth
- Docs: https://godoc.org/github.com/gobuffalo/buffalo
- Mailing list: WIP
- Source: https://github.com/gobuffalo

## Support

Bug reports and feature requests can be filed with the rest for the Zero to Restful API using Buffalo project here:

- [File Bug Reports and Features](https://github.com/conradwt/zero-to-restful-api-using-buffalo/issues)

## License

Zero to Restful API using Buffalo is released under the [MIT license](https://mit-license.org).

## Copyright

copyright:: (c) Copyright 2018 Conrad Taylor. All Rights Reserved.
