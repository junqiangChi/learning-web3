// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SolidityTest {
    mapping(address => bool) private  isWriteAccount;
    mapping(address => uint256) private balances;

    constructor() {
        isWriteAccount[0xe2899bddFD890e320e643044c6b95B9B0b84157A] = true;
    }
    function isWriteAuth() view private {
        require(isWriteAccount[msg.sender], "This account is black");
    }

    function deposit() public payable {
        isWriteAuth();
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint256 amount) public payable {
        isWriteAuth();
        require(amount < balances[msg.sender], "balance is enough");
        balances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
    }

    function checkBalance() public view returns (uint256) {
        isWriteAuth();
        return balances[msg.sender];
    }
}