# PURL - Parallel Client for URL

**purl** is a lightweight, parallelized alternative to the popular command-line tool curl. This program is designed for quick and small-scale testing, allowing users to assess how a service responds to multiple parallel requests, observe response times, and monitor resource consumption under pressure.

**curl** is widely used and well-known, everyone is familiar with its syntax and functionality. **purl** is designed to be used in a similar fashion to **curl**, with the addition of a `--parallel` flag to specify the number of parallel requests to make.

## What does the program do?

**purl** enables you to make multiple HTTP requests in parallel, providing a simple way to test and monitor the performance of a web service. The program utilizes goroutines to execute concurrent requests, allowing you to simulate a load on the target service and gather insights into its behavior under stress.

## How to Install

### Prerequisites

- Go installed on your machine

### Installation Steps

1. Install and update **purl** using the following command:

   ```bash
   GOPROXY=direct go install github.com/anonychun/purl@latest
   ```

   This command will automatically download, build, and install the **purl** executable in your `$GOBIN` directory.

2. Ensure that your `$GOBIN` directory is added to your system's PATH.

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   If you haven't set up your `$GOBIN` directory, you can use the default `$GOPATH/bin`:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

3. Verify the installation:

   ```bash
   purl --help
   ```

   This should display the help menu for **purl**.

## Usage

To run **purl**, use it in a similar fashion to curl:

```bash
purl [OPTIONS]
```

### Options

- `--parallel <number>` Specify the number of parallel requests (default is 1).

### Example

```bash
purl --parallel 100 https://example.com
```

## Disclaimer

**purl** is developed for experimental and fun purposes. If you want to test the performance of a service, there are many other tools that are more suitable for the job. Please be cautious when using this tool and ensure you have the necessary permissions before testing against any services.

I've been playing with Portainer and other monitoring tools. I use this program for quick, small-scale testing to determine how much load a service can handle, the duration of its response, and, most importantly, to monitor the resource consumption of a service under the pressure of many parallel requests.

## Extra

This is what I mean by quick and small-scale testing. I can simply copy the curl command from the browser, replace it with purl, and add the `--parallel` flag. Afterward, I run the command to get a quick overview of the service's performance.

You can also obtain curl commands from Postman and other tools, then use purl to test the service, as demonstrated in the example below.

https://github.com/anonychun/purl/assets/41361100/1593786b-6c55-4701-b78c-1002072dffd6
