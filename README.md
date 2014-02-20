gogular
=======
Gogular is a small proof-of-concept of a Go web server serving up AngularJS in a single-page app.

## Installing
=======
You will need the go runtime installed (http://www.golang.org/).

For AngularJS development, we recommend getting Yeoman (http://yeoman.io/).

In order to install all client-side packages, run:
```
$ cd angapp
$ npm install
$ bower install
```

To generate a new route, use:
```
$ yo angular:route routename
```

So we used "yo angular:route about" to generate the about us page.

If you add a library or create a route, remember to add it to the 
dependency list and add to index.html scripts list to include it in the 
app.

## Running
=======
Boot up the server from inside the gogular root:
```
$ go run httpserver.go
```