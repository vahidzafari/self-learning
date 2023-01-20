// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/*
Getter functions can be declared view or pure.
View - function declares that no state will be changed.
Pure - function declares that no state variable will be changed or read.
Internal - functions are the type of functions that can only be called internally within the contract or by the contracts inherited from the current one.
External - functions are the type of functions that are part of the contract but can only be used externally and called outside the contract by the other

Note: We can also assign an internal quantifier to the variables of the contract.
*/

contract ViewAndPure {
    uint public x = 1;

    // Promise not to modify the state.
    function addToX(uint y) public view returns (uint) {
        return x + y;
    }

    // Promise not to modify or read from the state.
    function add(uint i, uint j) public pure returns (uint) {
        return i + j;
    }
}
