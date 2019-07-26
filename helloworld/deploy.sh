#!/usr/bin/env bash
set -x
aws s3 --profile personal_s3 cp hello  s3://opentikva/ --acl public-read