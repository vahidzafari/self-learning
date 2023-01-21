// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Name {
    string public name;

    // A constructor is an optional function that is executed upon contract creation.
    constructor(string memory _name) {
        name = _name;
    }
}

contract Text {
    string public text;

    constructor(string memory _text) {
        text = _text;
    }
}

// Pass the parameters here in the inheritance list.
contract Child is Name("Input to Name"), Text("Input to Text") {

}

contract SecondChild is Name, Text {
    // Pass the parameters here in the constructor,
    // similar to function modifiers.
    constructor(
        string memory _name,
        string memory _text
    ) Name(_name) Text(_text) {}
}

// Parent constructors are always called in the order of inheritance
// regardless of the order of parent contracts listed in the
// constructor of the child contract.

// Order of constructors called:
// 1. Name
// 2. Text
// 3. ThirdChild
contract ThirdChild is Name, Text {
    constructor() Name("Name was called") Text("Text was called") {}
}

// Order of constructors called:
// 1. Name
// 2. Text
// 3. ForthChild
contract ForthChild is Name, Text {
    constructor() Name("Name was called") Text("Text was called") {}
}
