
# go-getting-started

A barebones Go app, which can easily be deployed to Heroku.

This application supports the [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) article - check it out.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/heroku/go-getting-started
$ cd $GOPATH/src/github.com/heroku/go-getting-started
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

You should also install [govendor](https://github.com/kardianos/govendor) if you are going to add any dependencies to the sample app.

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)


## Go .bash_file configuration

    export PATH=$GOPATH/bin:$PATH
    export PATH=$PATH:/Applications/Postgres.app/Contents/Versions/latest/bin

# Graphql Layer Service

The graphql layer running on [localhost:5000/graphql](http://localhost:5000/graphql), you could use a POST method or a GET method and add the query on the body.

The testing heroku application will be on the [heroku app](https//camaleon-reservations-api.herokuapp.com/graphql) and the same local dynamic could be used.

# If the database is new

## Set all the Reservations Status

    mutation {
        Pending : reservationStatus( description : "Pending", value : 1 ){ ... StatusParams }
        Approved : reservationStatus( description : "Approved", value : 2 ){ ... StatusParams }
        Canceled : reservationStatus( description : "Canceled", value : 3 ){ ... StatusParams }
        Available : reservationStatus( description : "Available", value : 4 ){ ... StatusParams }
        Completed : reservationStatus( description : "Completed", value : 5 ){ ... StatusParams }
    }

    fragment StatusParams on ReservationStatus {
        ID
        Description
        Value
    }


## Get all reservation status

    query {
        reservationStatuses {
            ID
            Description
            Value
        }
    }

## Create Client Info

    mutation{
        clientInfo(
            firstname : "",
            lastname : "",
            location : 0,
            email : "user@gmail.com",
            phone : "1234567890"
        ) {
            ID
            FirstName
            Phone
        }
    }

## Create Reservation

    mutation {
        reservation( 
            location : 0,
            table_id : 0,
            client_info_id : 0,
            status_id : 0,
            time_limit : 0,
            date : "02/28/2018 22:15:00",
            guests : 0
        ) {
            ID
            UID
            Table { ... TableParams }
            Status { ... ReservationStatusParams }
            ClientInfo { ... ClientInfoParams }
            Location { ... LocationParams }
            Date
            TimeLimit
            Guests
            Timestamp
            Updated
        }
    }

## Get Reservations by dates

If you don't add the status_id or set it as '0' it will just bring back all the reservations.

    query {
        reservations ( 
            location : 1, 
            start_date : "01/01/2018 00:00:00",
            end_date : "03/01/2018 23:59:59",
            status_id : 1
        ) {
            ID
            UID
            Table { ID, Name }
            Status { ID, Description, Value }
            ClientInfo { ID, FirstName, LastName }
            Location { ID, Name }
            Date
            TimeLimit
            Guests
            Timestamp
            Updated
        }
    }