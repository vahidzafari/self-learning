// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract ThisExample {
    address payable public owner;
    ThisExample public instance;
    ThisExample public instanceConverted;
    address public currentContractAddressConverted;

    constructor() {
        // this is a keyword, used in the contract to get the current contract's type.
        // this is also explicitly convertible to the address type.
        owner = payable(msg.sender);
        instance = this;
        instanceConverted = ThisExample(this);
        currentContractAddressConverted = address(this);
    }

    // selfdestruct(address recipient) is a globally available function. This is used to
    // destroy an already deployed contract instance. Once the contract is destroyed on
    // blockchain, all of its storage is released from blockchain.
    // If the recipient address is a contract address with a payable fallback function
    // defined, then the ether transfer triggered by the selfdestruct function does not
    // initiate a call to the payable fallback function. Hence, the recipient contract
    // will receive the ether balance of the destroyed contract silently, without calling
    // the fallback function.
    function kill() public {
        require(msg.sender == owner);
        selfdestruct(owner);
    }
}
