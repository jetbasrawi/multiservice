#multiservice

An simple example of linking a Go application to a Redis database running on Docker and provisioned with Docker Compose. 

The cononical Docker Compose demo app which records the number of times a page has been viewed in a Redis database. The web application from the cononical example is ported to Go.

##Setup

Clone the example into your GOPATH.

The application uses Gulp to install and run the Go web application. You will need to install Gulp and a numeber of plugins in the application directory on you workstation.

Navigate to the application directory on your workstation and run the following commands.

`$ npm install --save-dev gulp`

`$ npm install gulp-util --save-dev`

`$ npm install --save node-notifier`

Navigate to the application directory in the CLI of your docker enviroment. Run the docker-compose up command. A Redis server and an Alpine container hosting the web application will be started.

The Gulp task runner will install and run the webserver.

If you are working directly on Linux or running docker-machine on OSX or Windows you should be able to view the application in the browser on your workstation at http://localhost:8080. If you are using Vagrant you will need to forward port 8080 to your host for this to work.

If you make any changes to the Go code or HTML, the Gulp task runner will detect the changes and install the new iteration and restart the application automatically resulting in a very fast develop test feedback loop.

