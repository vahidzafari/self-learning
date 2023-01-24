// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/*
Libraries are similar to contracts, but you can't declare any state variable and you can't send ether.

A library is embedded into the contract if all library functions are internal.

Otherwise the library must be deployed and then linked before the contract is deployed.
*/

library Math {
    struct Data {
        uint num;
    }

    function set(Data storage self, uint _num) public {
        self.num = _num;
    }

    function add(Data storage self, uint y) public view returns (uint) {
        return self.num + y;
    }

    function sqrt(uint y) internal pure returns (uint z) {
        if (y > 3) {
            z = y;
            uint x = y / 2 + 1;
            while (x < z) {
                z = x;
                x = (y / x + x) / 2;
            }
        } else if (y != 0) {
            z = 1;
        }
        // else z = 0 (default value)
    }
}

contract TestMath {
    function testSquareRoot(uint x) public pure returns (uint) {
        return Math.sqrt(x);
    }
}

/*
To attach a library to the contract, we use a special directive called using X for Y;, 
where X is the library and Y is the data type. By defining this directive, we are saying 
that X is a library, we take all the functions from the X library, and allow those 
functions to be called on the Y data type. For example, the X library has a function 
called funcName(); in that case, you would be able to call Y.funcName() on the Y datatype. 
When funcName() is called, the first parameter to the funcName() function is passed as 
the Y datatype.
*/

contract TestUsing {
    using Math for Math.Data;
    Math.Data data;

    uint public result;

    function testSquareRoot(uint _x, uint _y) public {
        data.set(_x);
        result = data.add(_y);
    }
}
