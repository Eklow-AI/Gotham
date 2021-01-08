# Gotham
[![Eklow AI](https://circleci.com/gh/Eklow-AI/Gotham.svg?style=shield&circle-token=ac7fe26ee034d9f334b586904d69c877633b2f54)](https://app.circleci.com/pipelines/github/Eklow-AI/Gotham)
The Gotham API for compatability scores

## Infra
The API runs as a containerized application on an AWS Elastic Beanstalk (EBS) environment

To set-up an EBS Enviroment run the following in your local terminal in the root of your project directory
* `eb init` All of the default values are okay. This will create an .elasticbeanstalk/ directory at the root of your project
* `eb create` will then create the instance
* Once created, your app is most likely going to fail. To fix this go to Elastic Beanstalk -> Enviroments -> [Name of enviroment] -> Config -> Software -> Scroll to the bottom and the PORT environment variable

To make sure your app passes the health AWS health check make sure your `/` route is defined and serves something. AWS uses this route's response to check if your app is doing okay

## Testing

CircleCI builds
