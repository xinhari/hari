# Hari [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/xinhari.com/hari?status.svg)](https://godoc.org/xinhari.com/hari) [![Travis CI](https://travis-ci.org/xinhari/hari.svg?branch=main)](https://travis-ci.org/xinhari/hari) [![Go Report Card](https://goreportcard.com/badge/xinhari.com/hari)](https://goreportcard.com/report/xinhari.com/hari)

Hari is a distributed systems runtime for the Cloud and beyond.

## Overview

Hari addresses the key requirements for building distributed systems. It leverages the microservices
architecture pattern and provides a set of services which act as the building blocks of a platform. Xinhari deals
with the complexity of distributed systems and provides simpler programmable abstractions to build on.

Technology is constantly evolving. The infrastructure stack is always changing. Hari is a platform which
addresses these issues with a pluggable foundation and strongly defined apis to build on. Plug into any stack or cloud.

## Features

The runtime is composed of the following features:

### Services

Services are the core services that makeup the runtime. They provide a programmable abstraction layer for distributed systems infrastructure.

- **auth:** Authentication and authorization is a core requirement for any production ready platform. Micro builds in an auth service 
for managing service to service and user to service authentication.

- **broker:** A message broker allowing for async messaging. Microservices are event driven architectures and should provide messaging as a first
class citizen. Notify other services of events without needing to worry about a response.

- **config:** Manage dynamic config in a centralised location for your services to access. Has the ability to load config from multiple 
sources and enables you to update config without needing to restart services.

- **debug:** Built in aggregation of stats, logs and tracing info for debugging. The debug service scrapes all services for their info to 
help understand the overall scope of the system from one location. 

- **network:** A drop in service to service networking solution. Offload service discovery, load balancing and fault tolerance to the network.
The micro network dynamically builds a latency based routing table based on the local registry. It includes support for multi-cloud networking.

- **registry:** The registry provides service discovery to locate other services, store feature rich metadata and endpoint information. It's a
service explorer which lets you centrally and dynamically store this info at runtime.

- **runtime:** A service runtime which manages the lifecycle of your service, from source to running. The runtime service can run natively locally 
or on kubernetes, providing a seamless abstraction across both.

- **store:** State is a fundamental requirement of any system. We provide a key-value store to provide simple storage of state which can be shared
between services or offload long term to keep microservices stateless and horizontally scalable.

### Clients

Clients are entrypoints into the system. They enable access to your services through well known entrypoints.

- **api:** An api gateway which acts as a single entry point for the frontend with dynamic request routing using service discovery. 

- **bot:** A slack bot which enables you to query and interact with Hari directly from within slack. It's great for ChatOps.

- **cli:** Access services via the terminal. Every good developer tool needs a CLI as a defacto standard for operating a system. 

- **proxy:** An identity away proxy which allows you to access remote environments without painful configuration or vpn.

- **web:** A dashboard to explore services, describe their endpoints, the request and response formats and
query them directly.

### Framework

To write applications which run on Hari you can use the framework Xinhari.

- **xinhari:** Leverage the powerful [Xinhari](https://xinhari.com/xinhari) framework to develop microservices easily and quickly.
Xinhari abstracts away the complexity of distributed systems and provides simpler abstractions to build highly scalable microservices.

## Install

From source

```
go get xinhari.com/hari
```

Docker image

```
docker pull xinhari/hari
```

Linux install

```
wget -q  https://raw.githubusercontent.com/xinhari/hari/main/scripts/install.sh -O - | /bin/bash
```

Mac OS install
```
curl -fsSL https://raw.githubusercontent.com/xinhari/hari/main/scripts/install.sh | /bin/bash
```

Windows install
```
powershell -Command "iwr -useb https://raw.githubusercontent.com/xinhari/hari/main/scripts/install.ps1 | iex"
```

## Getting Started

Boot the entire runtime environment locally

```
hari server
```

### Create a service

```
# generate a service (follow instructions in output)
hari new example

# set to use server
hari env set server

# run the service
hari run example

# list services
hari list services

# call a service
hari call go.micro.service.example Example.Call '{"name": "John"}'
```

## Usage

See all the options

```
hari --help
```


See the [docs](https://xinhari.com/docs) for detailed information on the architecture, installation and use of the platform.

