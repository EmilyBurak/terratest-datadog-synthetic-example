# Terratest with Datadog Synthetic Test Terraform

## Introduction & Goals

In the course of my work recently, I've created Datadog Synthetics for testing API endpoints, such as heartbeats for health checks. In moving these to Terraform in the future, I'd like to explore [Terratest](https://terratest.gruntwork.io/) for testing terraform code, to make it more robust and honestly to learn more Go. Here I create a Datadog Synthetic Test and introspect with terratest on the monitor using the Datadog API called via. Go to test that there is a synthetic alert there.

## Technologies Used

- Terraform
- Datadog
- Terratest
- Go

## Infrastructure

- Creates a Datadog Synthetic Test and accompanying Monitor for testing.

## How It Works (High-Level)

The `synthetics_test.go` file calls Terraform using Terratest, then uses Go libraries to call the Datadog Monitors API and confirm the synthetic alert's presence as a test using `go test`. Terratest tears down the infrastructure after testing.

## Pre-reqs

- Datadog API key
- Datadog Application key
- Go
- Terraform

## How To Use

- Populate commented out or sample values in `main.tf`
- Run the Go test with `DD_SITE="datadoghq.com" DD_API_KEY="<DD_API_KEY>" DD_APP_KEY="<DD_APP_KEY>" go test`

## TO-DOs

- Test for something better than the presence of the field I'm testing for
- Better handling of API and app keys, use a var file for the Terraform part of it at least.
- Just write better Go code, this is the first non-trivial Go (if it is non-trivial) that I've written.
- Reformat `main.tf` out into proper file structure for TF.

## Lessons Learned / Observations

- I'm struggling to figure out a way to better test the Synthetic's presence and that it's working, other than having Terratest wait for a minute via. Go and then check for data coming into the monitor and that it's in an OK state. That's probably the best way forward but I don't know enough Go to implement that...yet.
- Go's pretty fun to work with, it's very particular in a nice way, the compiler is a Certified Friend for sure.
- Terratest is intuitive, presuming you know how to write the tests and what to test for...which is a big presumption sometimes, I'd think, depending on your infrastructure use case.

## Resources

- Datadog API docs
- Terratest docs
- Terraform Datadog provider docs
- StackOverflow, of course
