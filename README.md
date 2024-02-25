1. Build project using go build
2. Get checksum using SHA256 of the build produced
3. Add this checksum and path of binary to  /spire/conf/server/server.conf
4. Run sudo ./bin/spire-server run 
5. Run sudo ./bin/spire-server token generate
6. sudo ./bin/spire-agent run -joinToken=4d2462e7-31ec-4e4c-8626-2f3fc068e7e1 -logLevel=DEBUG