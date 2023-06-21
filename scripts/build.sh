#!/bin/bash

VERSION="1.0.0"
echo "Building terraform provider for SAP BTP_${VERSION}"
go build -o terraform-provider-sap-btp_${VERSION}