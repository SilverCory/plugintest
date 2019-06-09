#!/bin/bash
rm -R "plugins" &> /dev/null && mkdir "plugins" &> /dev/null
echo -e "\n\n\n"
go build -buildmode=plugin -o ./plugins/plug1.so ./modules/plug1 || (echo "Unable to build plugin 1." && exit 1);
go build -buildmode=plugin -o ./plugins/plug2.so ./modules/plug2 || (echo "Unable to build plugin 2." && exit 1);
go run ./cmd/app
