// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/*
Gas is essential to the Ethereum network. It is the fuel that allows it to operate, in the same way that a car needs gasoline to run.
Gas refers to the unit that measures the amount of computational effort required to execute specific operations on the Ethereum network.

Gas limit refers to the maximum amount of gas you are willing to consume on a transaction. 
A standard ETH transfer requires a gas limit of 21,000 units of gas.

Gas Price - the amount you want to pay per unit of gas as a fee to the miner

- useful linkes:
https://ethereum.github.io/yellowpaper/paper.pdf
https://etherscan.io/gastracker
https://ethgasstation.info/

How much ether do you need to pay for a transaction?
You pay gas spent * gas price amount of ether, where
gas is a unit of computation
gas spent is the total amount of gas used in a transaction
gas price is how much ether you are willing to pay per gas
Transactions with higher gas price have higher priority to be included in a block.

Unspent gas will be refunded.

Gas Limit
There are 2 upper bounds to the amount of gas you can spend

gas limit (max amount of gas you're willing to use for your transaction, set by you)
block gas limit (max amount of gas allowed in a block, set by the network)
*/

contract Gas {
    // 1 wei is equal to 1
    uint public oneWei = 1 wei;

    // 1 gwei is equal to 10^9 wei
    uint public oneGwei = 1 gwei;

    // 1 ether is equal to 10^18 wei
    uint public oneEther = 1 ether;

    uint public i = 0;

    // Using up all of the gas that you send causes your transaction to fail.
    // State changes are undone.
    // Gas spent are not refunded.
    function forever() public {
        // Here we run a loop until all of the gas are spent
        // and the transaction fails
        while (true) {
            i += 1;
        }
    }
}
