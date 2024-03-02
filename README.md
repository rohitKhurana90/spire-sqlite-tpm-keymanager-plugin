1. Build project using go build
2. Get checksum using SHA256 of the build produced
3. Add this checksum and path of binary to  /spire/conf/server/server.conf
4. Run sudo ./bin/spire-server run 
5. Run sudo ./bin/spire-server token generate
6. sudo ./bin/spire-agent run -joinToken=1c63b5d2-713e-4883-b5e3-c01920da2b1f -logLevel=DEBUG


TPM important commands

Create a persistent handle
tpm2_evictcontrol -C o -c key.ctx 0x81007000

Remove a persistent handle
tpm2_evictcontrol -C o -c 0x81008000



