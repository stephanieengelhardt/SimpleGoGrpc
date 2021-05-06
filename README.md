# gRPC Bank
Author: Stephanie Hingtgen

## Introduction
This project creates a distributed banking system that spins up branches for each customer to communicate with over gRPC. It also ensures eventual consistency between the branches by propogating deposits and withdraws that occur in each branch. Note that the customers here are all withdrawing from the same pool of money. This program reads in the input.json file, which specifies the customers and events they will perform, as well as the branches and the beginning balance. It will then create an output.json file that shows the status (success or error) of the different actions per branch and the final balance of each branch.

### Examples
You can find example inputs and outputs in the samples folder. This includes a case where two customers attempt to withdraw a valid amount, but together the combined withdraw is greater than the balance. It also includes a case of invalid inputs, such as negative money on queries and money that contains more than two decimal points.

## Running the Program
### MacOS
`make run`

For more verbose logging:
`make run-dev`

### Linux (GOOS=linux GOARCH=arm)
`make run-linux`

For more verbose logging:
`make run-linux-dev`

Note: You can check your GOOS and GOARCH by running `go env` (with golang installed) and make sure GOOS=linux and GOARCH=arm.

### All Other Platforms
To run on any other platform, you will need to compile the program. Please follow the guide below to do so.

## Building the Program
### Technology Requirements
- Golang 1.15

### Building for Your Platform
If you want to compile for a different operating system than what is supported in the Makefile, you can run the following from within the main folder:
`env GOOS=<os> GOARCH=<platform> go build -o bank`

To determine the GOOS and GOARCH you should use, run `go env`. 

Once the binary `bank` has been created, run `./bank -log-level=0` in the main folder.

For more information, check out [tutorial](https://medium.com/@utranand/building-golang-package-for-linux-from-windows-22fa23764808).

## Development
When updating the .proto file, run `make clean` before running `make build-osx` or `make build-linux`, which will create the .pb.go file for you.

Make sure to run with the log level at 1 while developing. If you're on MacOS you can do this by running `make run-dev`.


