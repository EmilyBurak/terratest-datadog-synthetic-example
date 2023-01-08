
terraform {
  required_providers {
    datadog = {
      source = "DataDog/datadog"
    }
  }
}

variable "datadog_api_key" {
  type        = string
  description = "Datadog API Key"
  # default = "API key here"
}

variable "datadog_app_key" {
  type        = string
  description = "Datadog Application Key"
  # default = "API key here"
}

provider "datadog" {
  api_key = var.datadog_api_key
  app_key = var.datadog_app_key
}

resource "datadog_synthetics_test" "example" {
  type    = "api"
  subtype = "http"
  request_definition {
    method = "GET"
    # url    = "your-url-here"
  }
  request_headers = {
    Content-Type = "application/json"
  }
  assertion {
    type     = "statusCode"
    operator = "is"
    target   = "200"
  }
  locations = ["aws:eu-central-1"]
  options_list {
    tick_every = 900

    retry {
      count    = 2
      interval = 300
    }

    monitor_options {
      renotify_interval = 120
    }
  }
  name    = "An API test on your url"
  message = "Example message"
  tags    = ["foo:bar", "foo", "env:test"]

  status = "live"
}