gae-ng-seed
===========

<a href="https://developers.google.com/appengine/">Google App Engine</a> / <a href="http://golang.org/">Go Lang</a> / <a href="http://angularjs.org/">Angular JS Seed Project</a>

This project applies the <a href="https://github.com/angular/angular-seed">Angular-Seed</a> into the GAE with Go context.

##Install

###Google App Engine SDK for Go 

Download the GAE SDK for Go from Google at ...
https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go

###Clone the gae-ng-seed project 

<code>git clone https://github.com/campbel/gae-ng-seed.git</code>

###Run the Go Dev Server from the project root 

<code>dev_appserver.py ./</code>

Typically the GAE dev server will run on port 8080. Browse to http://localhost:8080 to see the app running.

##Optional

###<a href="http://karma-runner.github.io/0.8/index.html">Karma</a>

The karma test runner is used to automatically watch and run tests when your code changes. This is great for following test driven development methodology. The gae-ng-seed project comes with a karma config that will work with the setup. Once Karma is installed head to /client/scripts and execute test.bat OR head to /client/config on the command line and execute <code>karma start</code>
