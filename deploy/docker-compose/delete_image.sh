#!/bin/bash

username=testuser
password=testpassword
url="http://localhost:5000"
repo="library/nginx"
tag="1.20"

# get image digest which will be used for delete
echo "get image digest"
digest=`curl -u$username:$password -X HEAD -I -H "Accept: application/vnd.docker.distribution.manifest.v2+json"  "${url}/v2/${repo}/manifests/${tag}" | grep Docker-Content-Digest | awk -F ':' '{print "sha256:"$3}'`
## delete \r
digest=`echo $digest|tr -d '\r'`
echo "digest: $digest"
#echo ${#digest}

#delete image by digest
requrl="${url}/v2/${repo}/manifests/$digest"
echo "requrl:$requrl"
curl -v -X DELETE -u $username:$password ${requrl}


echo "end"
