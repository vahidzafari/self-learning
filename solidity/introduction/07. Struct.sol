// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/*
You can define your own type by creating a struct.
They are useful for grouping together related data.
Structs can be declared outside of a contract and imported in another contract.

Variables are declared as either storage, memory or calldata to explicitly specify the location of the data.
storage - variable is a state variable (store on blockchain)
memory - variable is in memory and it exists while a function is being called
calldata - special data location that contains function arguments
*/

contract Struct {
    struct Person {
        string first_name;
        string last_name;
    }
    // An array of 'Person' structs
    Person[] public persons;

    function create(
        string calldata first_name,
        string calldata last_name
    ) public {
        // key value mapping
        persons.push(Person({first_name: first_name, last_name: last_name}));

        // initialize an empty struct and then update it
        Person memory person;
        person.first_name = first_name;
        person.last_name = last_name;

        persons.push(person);
    }

    // Solidity automatically created a getter for 'persons' so
    // you don't actually need this function.
    function get(
        uint _index
    ) public view returns (string memory first_name, string memory last_name) {
        Person storage person = persons[_index];
        return (person.first_name, person.last_name);
    }

    function updatePerson(
        uint _index,
        string calldata first_name,
        string calldata last_name
    ) public {
        Person storage person = persons[_index];
        person.first_name = first_name;
        person.last_name = last_name;
    }
}
