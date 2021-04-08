# MongoDB Atlas AWS CloudFormation Resources

***This file will be accurate post GA of the MongoDB Atlas Resource Provider for CloudFormation***

### Status: BETA (actively looking for [feedback](https://feedback.mongodb.com/forums/924145-atlas/category/392596-atlas-cloudformation-resources) and [comments](https://github.com/mongodb/mongodbatlas-cloudformation-resources/issues/new))

Use AWS CloudFormation to manage [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

The quickest way to get started:

* Launch the [Quickstart](https://aws.amazon.com/quickstart/architecture/mongodb-atlas/).


# Support, Bugs, Feature Requests

Feature requests can be submitted at [feedback.mongodb.com](https://feedback.mongodb.com/forums/924145-atlas/category/392596-atlas-cloudformation-resources) - just select "Atlas CloudFormation Resources" as the category or vote for an already suggested feature.

CURRENT BETA INFO: Support questions submitted under the Issues section of this repo will be handled on a "best effort" basis.
Bugs should be filed under the Issues section of this repo.

POST GA: Support for the MongoDB Atlas Resource Provider for CloudFormation is provided under MongoDB Atlas support plans, starting with Developer. Please submit support questions within the Atlas UI. 

# Securing your Quickstart

This repository contains utilities which use both the AWS and MongoDB Atlas APIs to progamatically create, destory, or otherwise manage resources in the AWS and MongoDB clouds. This can have costs or other impact so please use these tools with diligence and follow best practices. 

The automated workflows require the following secrets to be placed into your Github forks' Settings/Secrets:

```
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_REGION
ATLAS_PUBLIC_KEY
ATLAS_PRIVATE_KEY
ATLAS_ORG_ID
```

## MongoDB Atlas Programmatic API key
It's necessary to generate and configure an API key for your organization for the acceptance test to succeed. To grant programmatic access to an organization or project using only the API you need to know:

The programmatic API key has two parts: a Public Key and a Private Key. To see more details on how to create a programmatic API key visit https://docs.atlas.mongodb.com/configure-api-access/#programmatic-api-keys.

The programmatic API key must be granted roles sufficient for the acceptance test to succeed. The Organization Owner and Project Owner roles should be sufficient. You can see the available roles at https://docs.atlas.mongodb.com/reference/user-roles.

You must configure Atlas API Access for your programmatic API key. You should allow API access for the IP address from which the acceptance test runs.

## Security - Setup 

Step 1) Create and note your MongoDB Atlas API Key.
Step 2) Create and note your AWS Access Key and AWS Secret Key ID.
Step 3) Follow the Github docs on how to [create a Secret](https://docs.github.com/en/actions/configuring-and-managing-workflows/creating-and-storing-encrypted-secrets#creating-encrypted-secrets-for-a-repository) in your clone of this repository.

# Requirements

- [AWS CloudFormation CLI](https://github.com/aws-cloudformation/cloudformation-cli) 
- (Optional - only need if building from source) [AWS CloudFormation CLI Go Plugin](https://github.com/aws-cloudformation/cloudformation-cli-go-plugin/)
- (Optional - only need if building from source) [Go](https://golang.org/doc/install) 

## Registering resources 

1. Please check that you satisfy all the [requirements](#Requirements) before proceeding.
2. Clone this repo. 
3. A helper script `cfn-submit-helper.sh` will build and submit each resource for you. You can also run the `cfn submit` tool yourself. 

The following command builds and registers all the types used in the quickstart:

```
cd mongodbatlas-cloudformation-resources\cfn-resources
./cfn-submit-helper.sh project cluster database-user project-ip-access-list network-peering
```

# Testing the Provider

Please see README for each resource for details on unit and integrated AWS testing.

