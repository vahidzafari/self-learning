// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract GlobalVariables {
    // the miner's EOA address of the current block
    address coinbase = block.coinbase;

    // This provides the current difficulty level of the network as a uint256 value,
    // which is used by the miners to solve the Proof of Work (PoW) puzzle.
    uint difficulty = block.difficulty;

    // This provides the total gas limit of the current block as a uint256 value.
    // At the time of writing this book, 8 million gas is allocated to each block
    //on Ethereum mainnet.
    uint gasLimit = block.gaslimit;

    // This provides the current block number, which is generated for the current
    //block as a uint256 value. Block numbers are generated consecutively.
    uint blockNumber = block.number;

    // This provides the timestamp when the block is generated, as a uint256 value.
    // The timestamp is in seconds since Unix epoch.
    uint timestamp = block.timestamp;

    // This provides the full call data, represented as bytes, which is sent by the
    // transaction initiator. The data includes the first 4 bytes of method-hash to be executed,
    // followed by the arguments of the method.
    bytes data = msg.data;

    // This provides the sender of the message's current call as the address. When an
    // external contract call is initiated from a contract, msg.sender will return the
    // contract address of the caller contract.
    address sender = msg.sender;

    // This provides the first 4 bytes of the function signature, to be executed as bytes4.
    bytes4 sig = msg.sig;

    // This provides the amount of wei sent along with the transaction, as a uint256 value.
    uint value = msg.value;

    // This provides the gas price of the transaction set by transaction initiator, as a uint256 value.
    uint gasprice = tx.gasprice;

    // This provides the original sender of the transaction from the full call chain as an address.
    // As all of the transactions are always initiated from an EOA, we can say that tx.origin will
    // always return an address of an EOA. Therefore, the use of tx.origin is not recommended.
    address origin = tx.origin;
}
