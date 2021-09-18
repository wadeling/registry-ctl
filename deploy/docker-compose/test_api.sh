#!/bin/bash

#curl -v -utestuser:testpassword  -H "Accept: application/vnd.docker.distribution.manifest.v2+json" "http://localhost:5000/v2/nginx/manifests/1.15"
#curl  -utestuser:testpassword   "http://localhost:5000/v2/nginx/manifests/1.15" > manifest_v1.txt
#curl -v -utestuser:testpassword   "http://localhost:5000/v2/_catalog?n=2" 
#curl -v -utestuser:testpassword   "http://localhost:5000/v2/_catalog" 
#curl -v -utestuser:testpassword   "http://localhost:5000/v2/_catalog" 

image_digest=sha256:9dc6c65944f0970ffd6ffc3c80e4c4d5e01e683e8acf75a2b4ee0350fb62fc76
requrl="http://localhost:5000/v2/library/nginx/manifests/$image_digest"
curl -v -u testuser:testpassword -X GET $requrl
