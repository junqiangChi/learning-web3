// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

contract EtherWallet {
    event Log(address _ads);
    event WithDraw(address owner);
    address payable owner;
    uint256 amount;

    modifier isOwner() {
        require(msg.sender == owner, "acount is not Owner");
        _;
    }

    constructor(address _owner) {
        owner = payable(_owner);
    }

    function deposit() public payable {
        amount += msg.value;
    }

    function withdraw() external isOwner {
        payable(owner).transfer(amount);
        amount = 0;
        emit WithDraw(owner);
    }

    function withdraw3() external {
        require(msg.sender == owner, "Not owner");
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, "Call Failed");
    }

    receive() external payable {
        emit Log(msg.sender);
    }
}
