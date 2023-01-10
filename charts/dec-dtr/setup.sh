#!/bin/sh

cd ../..
git submodule init
git submodule update

cd ./charts/dec-dtr/product-edc/edc-tests/src/main/resources/deployment/helm/all-in-one
helm dep up
cd ../../../../../../../../catenax-at-home/deployment/helm/api-wrapper
helm dep up

cd ../../../../
helm dep up