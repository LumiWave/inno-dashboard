set -x

sh ./prebuild.sh $1

rm -rf bin/inno-dashboard

go build -o bin/inno-dashboard.exe main.go

cd bin
./inno-dashboard.exe -c=config.yml