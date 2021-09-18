#!/bin/bash

username=testuser
password=testpassword
url="http://localhost:5000"
repo="library/nginx"
tag="1.20"

# get image digest which will be used for delete
#echo "get image digest"

digest=`curl -u$username:$password -X HEAD -I -H "Accept: application/vnd.docker.distribution.manifest.v2+json"  "${url}/v2/${repo}/manifests/${tag}" | grep Docker-Content-Digest | awk -F ':' '{print "sha256:"$3}'`
echo -e "$digest" > image_digest

echo "digest: $digest"
