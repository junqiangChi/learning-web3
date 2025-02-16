// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

contract PiggyBank {
    address owner;

    constructor() {
        owner = msg.sender;
    }
    modifier isOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    event Deposit(address _address, uint _amount);
    event Withdraw(address _address, uint _amount);

    function deposit() public payable {
        emit Deposit(msg.sender, msg.value);
    }

    function withdraw() public payable isOwner {
        emit Withdraw(msg.sender, address(this).balance);
        selfdestruct(payable(msg.sender));
    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}
