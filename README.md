gae-ng-seed
===========
<a href="https://developers.google.com/appengine/">Google App Engine</a> / <a href="http://golang.org/">Go Lang</a> / <a href="http://angularjs.org/">Angular JS</a>

## Included

The server uses templating from Go's "html/template" package.

## Install & Setup

### Get Google App Engine SDK for Go 

Download available at ...
https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go

### Clone the gae-ng-seed project 

<code>git clone https://github.com/campbel/gae-ng-seed.git</code>

### Bower Install

<code>cd client</code>
<br/>
<code>bower install</code>

### Run the Go Dev Server from the project root 

<code>goapp serve</code>

Typically the GAE dev server will run on port 8080. Browse to http://localhost:8080 to see the app running.

## Notes

Before uploading your app make sure to change the application name in the app.yaml.

This project is just a start. For a more info on Angular, Go and Google's App Engine, visit their respective docs.
