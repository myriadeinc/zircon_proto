const PROTO_PATH = '/usr/lib/zircon/zircon.proto';

const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const serverurl = 'zircon_proto:8088';
const blockReferenceService = require('block.reference.js');
let packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
let zircon_proto = grpc.loadPackageDefinition(packageDefinition).zircon;

async function main() {
  let server = new grpc.Server();
  server.addService(
    zircon_proto.Zircon.service, 
    { validateBlock: blockReferenceService.processBlock });

  server.bindAsync(
    serverurl, 
    grpc.ServerCredentials.createInsecure(),
    // Callback function, no params
    () => {
      server.start() 
      console.log(`Server started on ${serverurl}`)
    });
  
}



main().catch(err=>console.log(err));
