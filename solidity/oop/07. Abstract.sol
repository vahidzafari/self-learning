// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/*
In an abstract contract, there are only a few functions that don't have the function 
body defined. You cannot deploy abstract contracts alone. However, you can inherit 
the contract and provide the definition to each function that's declared in the 
abstract contract.

If a contract inherits from an abstract contract and doesn't provide the implementation 
of the function, the inheriting contract would also be considered an abstract contract. 
Hence, it is the developer's responsibility to ensure that all the functions of the 
abstract contract are defined. Also, developer tools such as Remix and Truffle would 
not allow your inherited contract to be deployed.
*/

contract AbstractDeposit {
    function depositEther() public payable returns (bool);
}

contract DepositHolderImpl is AbstractDeposit {
    mapping(address => uint) deposits;

    function depositEther() public payable returns (bool) {
        deposits[msg.sender] += msg.value;
    }
}
