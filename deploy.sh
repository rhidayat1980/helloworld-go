#!/usr/bin/env bash
project=glassy-operand-281303
region=asia-east1
service=helloworld

gcloud builds submit --tag gcr.io/${project}/${service}
gcloud run deploy ${service} --image gcr.io/${project}/${service} \
    --platform managed --region=${region}
