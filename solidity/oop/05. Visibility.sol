// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/*
Functions and state variables have to declare whether they are accessible by other contracts.

Functions can be declared as

public:
A function defined as public is accessible from both inside and outside of the contract. 
Solidity generates a getter function for public state variables, and these are also 
accessible from outside of the contract via these getter functions.


external: 
A function defined as external is only accessible from outside of the contract. 
Internally, the function will not be accessible directly. To access an external 
function within the same contract, you would have to call this.functionName(). 
For state variables, you cannot use the external keyword.

internal: 
A function or variable defined as internal is only accessible internally within 
the contract. Also, if Contract X inherits from Contract Y, which has internal functions 
or variables, Contract X can access all of Contract Y's internal functions and variables. 
When no visibility is specified, the variable would, by default, take internal visibility.

private: 
A function or variable defined as private is only accessible internally within the same 
contract. A private variable or function is not accessible in derived contracts.
*/

contract Base {
    // Private function can only be called
    // - inside this contract
    // Contracts that inherit this contract cannot call this function.
    function privateFunc() private pure returns (string memory) {
        return "private function called";
    }

    function testPrivateFunc() public pure returns (string memory) {
        return privateFunc();
    }

    // Internal function can be called
    // - inside this contract
    // - inside contracts that inherit this contract
    function internalFunc() internal pure returns (string memory) {
        return "internal function called";
    }

    function testInternalFunc() public pure virtual returns (string memory) {
        return internalFunc();
    }

    // Public functions can be called
    // - inside this contract
    // - inside contracts that inherit this contract
    // - by other contracts and accounts
    function publicFunc() public pure returns (string memory) {
        return "public function called";
    }

    // External functions can only be called
    // - by other contracts and accounts
    function externalFunc() external pure returns (string memory) {
        return "external function called";
    }

    // This function will not compile since we're trying to call
    // an external function here.
    // function testExternalFunc() public pure returns (string memory) {
    //     return externalFunc();
    // }

    // State variables
    string private privateVar = "my private variable";
    string internal internalVar = "my internal variable";
    string public publicVar = "my public variable";
    // State variables cannot be external so this code won't compile.
    // string external externalVar = "my external variable";
}

contract Child is Base {
    // Inherited contracts do not have access to private functions
    // and state variables.
    // function testPrivateFunc() public pure returns (string memory) {
    //     return privateFunc();
    // }

    // Internal function call be called inside child contracts.
    function testInternalFunc() public pure override returns (string memory) {
        return internalFunc();
    }
}
