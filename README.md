GoRestApp
=========

A golang REST API application, based on https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94 by Mohamed Labouardy.

This is part of my [#100DaysOfCode]{https://github.com/clcollins/100-days-of-code} to learn golang!

The App
-------

Rather than movies, in this app I'm CAT-aloguing kittens, in the grand tradition of the old Internet.

### API Endpoints

I will eventually make this into a webapp for browser consumption, as well, so I am going to deviate from Mr. Labouardy's spec by making my API accessbile at /api/v1/gatos.

```
GET     /api/v1/gatos       list the cats
GET     /api/v1/gatos/:id   get a particular cat
POST    /api/v1/gatos       add a cat
PUT     /api/v1/gatos/:id   add or replace a cat  (Replace?  no...)
PATCH   /api/v1/gatos/:id   update a cat
DELETE  /api/v1/gatos       delete a cat (NEVER!)
```
