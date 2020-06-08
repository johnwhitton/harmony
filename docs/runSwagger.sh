# git clone git@github.com:swagger-api/swagger-ui.git
swagger generate spec -w ../cmd/harmony/ -o harmony.yaml -m
rm -rf swaggerui
mkdir swaggerui
cp -rf ./swagger-ui/dist/* ./swaggerui/.
# rm -rf ./swagger-ui
cp harmony.yaml ./swaggerui/.
cp index.html ./swaggerui/.
statik -src=/home/ec2-user/go/src/github.com/harmony-one/harmony/docs/swaggerui -p docs -dest ../
mv statik.go swaggerui.go
rm -rf swaggerui
cd ..
make linux_static
cp ~/go/src/github.com/harmony-one/harmony/bin/harmony ~/explorer
cd docs
