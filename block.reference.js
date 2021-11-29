'use strict';
const xmr = require('./xmr.js');
const logger = (s) => console.log(s)
// Value is 16^64 or 2^256
const hex256 = BigInt('0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF')
const INVALID = 0;
const LOW_SCORE = 1;
const VALID = 2;
const WINNER = 3;

const BlockReference = {
  // Protobuf handler (unfortunately does not seem its easy to avoid callbacks)
  processBlock: (call, callback) => {
    // Check protobuf params
  const pblock = call.request;

  const result = BlockReference.checkDifficulty(
    pblock.local_diff,
    pblock.global_diff,
    pblock.hex_result
  );
  return callback(null, 
    {
        block_status : result,
    });
  },

  constructJob: (call, callback) => {
    // Check protobuf params
  const input = call.request.jobstub;

  return callback(null, {
    jobstub: input
  });
  },



  // Checks if block sent by miner corresponds to the job they were assigned
    verifyBlock: (blob, nonce, extraNonce, seed_hash, result) => {
    try {
        let block = BlockReference.buildBlock(blob, nonce, extraNonce);
      block = BlockReference.convertBlock(block);
      block = BlockReference.hashBlock(block, seed_hash);
      return block.toString('hex') == result;
    } catch (err) {
        logger(err);
        return false;
      }
    },
 
  buildBlock: (blob, nonce, extraNonce) => {
    let block = Buffer.from(blob, "hex");
    // Value of 8 is given because our default reserve_offset is set to 8
    block.writeUInt32BE(extraNonce, 8);
    const NonceBuffer = Buffer.from(nonce, 'hex');
    return xmr.construct_block_blob(block, NonceBuffer);
  },
  convertBlock: (constructedBlock) =>  xmr.convert_blob(constructedBlock),
  hashBlock: (block, seed_hash) => xmr.randomx(block, Buffer.from(seed_hash, 'hex')),
  



  checkDifficulty: (localDiff, globalDiff, block) => {

    let rawBlock = Buffer.from(block, 'hex');
    rawBlock = Array.prototype.slice.call(rawBlock, 0).reverse();
    let hexBlock = Buffer.from(rawBlock).toString('hex');
    if(!hexBlock.trim()) return INVALID;
    
    const blockInt = BigInt(`0x${hexBlock}`)
    let hashDiff = hex256 / blockInt;

    hashDiff = BigInt(hashDiff)
    globalDiff = BigInt(globalDiff);
    localDiff = BigInt(localDiff);

    if (hashDiff >= globalDiff) {

      // We won the block reward!
      // Submit to monero node
      return WINNER;
    }
    if (hashDiff >= localDiff) {
      // Grant share to miner
      return VALID;
    }
    // Block does not match difficulty requirements: low difficulty share
    return LOW_SCORE;
  },

   // Checks if block sent by miner corresponds to the job they were assigned
   verifyBlockNonProto: (minerData, job) => {
    try {
      let block = BlockReference.buildBlock(job.blob, minerData.nonce, job.extraNonce);
      block = BlockReference.convertBlock(block);
      block = BlockReference.hashBlock(block, job.seed_hash);
      return block.toString('hex') == minerData.result;
    } catch (err) {
      logger(err);
      return false;
    }
  },
    // Used when our nonce is represented as an integer value 
    buildIntNonce: (blob, nonce) => {
        let block = Buffer.from(blob, "hex");
        const buf = Buffer.allocUnsafe(4);
        buf.writeUInt32LE(nonce);
        const NonceBuffer = buf;
        return xmr.construct_block_blob(block, NonceBuffer);
      },
  // For fallback:
  // new Buffer(minerData.nonce, 'hex').copy(block, 39);
  // Writing the nonce in a specific position if util does not work in testing

};

module.exports = BlockReference;


