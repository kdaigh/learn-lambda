# Learn Lambda

This is a very simple Go-based Lambda project created to become more familiar with Go and AWS Lambda functionality.

## Goals

* *Establish basic project functionality.* Write code that illustrates some basic functionality of Lambda with Go. (In this case, that means generating a random number between 0 and "max".)
* *Add Makefile functionality.* Use Makefile to make the project easier to use and interact with.

## Usage

### Prerequisites

Go can be downloaded [here](https://golang.org/doc/install).

AWS CLI can be downloaded [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html). Note: Installing the AWS Command Line Interface requires Python 2 version 2.6.5+ or Python 3 version 3.3+.

### Usage

Create a zip file of the project with the following command:

```
make zip
```

This will create a zip file titled *handler.zip* which can be uploaded directly to the Lambda function console once you've created a Lambda function.

An example test configuration for this project might be:

```
{
  "id": 45678,
  "max": 100
}
```

## Tech Stack

* [Go](https://golang.org/) - Language
* [AWS Lambda](https://aws.amazon.com/lambda/) - Serverless Platform

## Helpful Links

* [Go Based AWS Lambda Tutorial](https://www.youtube.com/watch?v=x_yCX4kSchY) - *Tutorial Edge* - Video tutorial covering the basics of project setup, configuration, and testing.
