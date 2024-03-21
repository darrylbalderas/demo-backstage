#!/bin/bash

# Terraform

terraform fmt -recursive

# Backstage

yarn fix
yarn lint:all
yarn test:all


# Security

gitleaks detect --source . -v
