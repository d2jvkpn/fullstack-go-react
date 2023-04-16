#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


#### backend
go mod init fullstack-go-react

#### frontend
create-react-app frontend
mv frontend/{node_modules,public,src,package.json} ./
mv frontend/README.md README.react.md
rm -r frontend

sed -i 's/frontend/fullstack-go-react/' package.json

# npm install --save-dev env-cmd
# rm package-lock.json
yarn add env-cmd --dev
echo "PUBLIC_URL=/site" > .env

#### ignores of tokie and docker
cat > .tokeignore <<EOF
# react
public/
package.json
package-lock.json
EOF

cat > .dockerignore <<EOF
wk*/
.git
target/
/logs
EOF
