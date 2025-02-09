// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MappingExample {
    mapping(address => uint) balances;
    mapping(address => int[]) records;
    address[] users;

    function addBalance() public payable {
        balances[msg.sender] += msg.value;
        records[msg.sender].push(int(msg.value));
    }

    function checkBalance() public view returns (uint) {
        return balances[msg.sender];
    }

    function addUser(address _user) public {
        require(balances[_user] == 0, "User already exist,");
        users.push(_user);
        balances[_user] = 100;
    }

    function transfer(address to, uint amount) public payable {
        require(balances[msg.sender] >= amount, "user address not enough");
        balances[msg.sender] -= amount;
        records[msg.sender].push(int(amount) * - 1);
        balances[to] += amount;
        records[msg.sender].push(int(amount));
    }

    function getUsers() public view returns (address[] memory){
        return users;
    }

    function getUserRecord() public view returns (int[] memory) {
        return records[msg.sender];
    }

}