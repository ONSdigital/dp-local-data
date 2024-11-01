# dp-local-data

⚠️ This repository will be archived in October 2024 as it is no longer in development. ⚠️

`dp-local-data` is a cmd-cli tool enabling developers to easily:
 - Clean all CMD data out of their local environment.
 - Import the prerequisite hierarchy/codelist data required to import datasets.
 
The default config will load the hierarchy and code lists required for the `Suicides` dataset. 

## Prerequisites
`dp-local-data` uses Go Modules so requires a go version of **1.11** or later. 

`dp-local-data` requires:
- `dp-code-list-scripts` 
- `dp-hierarchy-builder`

to be one your `$GOPATH`:
```
go get github.com/ONSdigital/dp-code-list-scripts
go get github.com/ONSdigital/dp-hierarchy-builder
```

### Getting started
Clone the code
```
git clone git@github.com:ONSdigital/dp-local-data.git
```
:warning: `dp-local-data` uses Go Modules and **must** be cloned to a location outside of your `$GOPATH`.

Install the binary
```bash
go install github.com/ONSdigital/dp-local-data
```

### Usage
If the install was successful running
```
dp-local-data
```
Should present you with a help menu similar to:
```
dp-local-data is a tool for cleaning CMD data out of local dev env and/or importing the prerequisite hierarchy/codelist data required to successfully import datasets

Usage:
	dp-local-data [-commands]
Commands:
  -clean
    	Drop all local CMD data from Neo4j and MongoDB and deletes any Zededee collections
  -help
    	Display help info
  -import
    	Import the generic hierarchies and code lists specified in config.yml
```

### Config
TODO - coming soon.
