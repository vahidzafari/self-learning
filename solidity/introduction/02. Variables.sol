// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Variables {
    // State variables are stored on the blockchain.
    string public text = "Hello";
    uint public num = 123;
    uint public timestamp;
    address public sender;

    // coding convention to uppercase constant variables
    address public constant MY_ADDRESS =
        0x777788889999AaAAbBbbCcccddDdeeeEfFFfCcCc;
    uint public immutable MY_UINT = 123;

    function execute() public {
        // Local variables are not saved to the blockchain.
        uint i = 456;
        num = i;

        // Here are some global variables
        timestamp = block.timestamp; // Current block timestamp
        sender = msg.sender; // address of the caller
    }
}
