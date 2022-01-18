# DBoM Storage Provider for Grafeas

[Grafeas](https://github.com/grafeas/grafeas) with a dockerized DBoM backend. Running this project will create a grafeas API server at the specified port utilizing an existing [DBoM API Gateway](https://github.com/DBOMproject/chainsource-gateway/) instance as the storage.

Grafeas ("scribe" in Greek) is an open-source artifact metadata API that provides a uniform way to audit and govern your software supply chain. Grafeas defines an API spec for managing metadata about software resources, such as container images, Virtual Machine (VM) images, JAR files, and scripts. You can use Grafeas to define and aggregate information about your project's components. Grafeas provides organizations with a central source of truth for tracking and enforcing policies across an ever growing set of software development teams and pipelines. Build, auditing, and compliance tools can use the Grafeas API to store, query, and retrieve comprehensive metadata on software components of all kinds.

## Current Status

This repository currently contains an MVP implementation of the Storage Provider with the following features implemented:

- [x] Project Methods
  - [x] `CreateProject`
  - [x] `GetProject`
  - [x] `ListProjects`
  - [x] `DeleteProject` - **DBoM does not support delete operations**
- [x] Occurrence Methods
  - [x] `CreateOccurrence`
  - [x] `BatchCreateOccurrences`
  - [x] `GetOccurrence`
  - [x] `ListOccurrences`
  - [x] `UpdateOccurrence`
  - [x] `DeleteOccurrence`  **DBoM does not support delete operations**
- [x] Note Methods
  - [x] `CreateNote`
  - [x] `BatchCreateNotes`
  - [x] `GetNote`
  - [x] `ListNotes`
  - [x] `UpdateNote`
  - [x] `DeleteNote` **DBoM does not support delete operations**
- [ ] Misc Methods
  - [x] `GetOccurrenceNote`
  - [x] `ListNoteOccurrences`
  - [ ] `GetVulnerabilityOccurrencesSummary`
- [ ] Filtering Support (for `List` methods)
- [ ] Pagination


## Usage

If you don't already have a DBoM node running, you can set up a minimal node using the instructions in the [deployments repository](https://github.com/DBOMproject/deployments/tree/master/docker-compose-quickstart).

You will also need docker & docker-compose to run the quickstart. 

Configure the appropriate endpoint in [docker-config.yaml](./docker-config.yaml) and start the service using `docker-compose up`

## Getting help

If you have any queries on insert-project-name, feel free to reach us on any of our [communication channels](https://github.com/DBOMproject/community/blob/master/COMMUNICATION.md) 

## Getting involved

This project is currently in the MVP stage and we welcome contributions to drive it to completion. Please refer to [CONTRIBUTING](CONTRIBUTING.md) for more information.


