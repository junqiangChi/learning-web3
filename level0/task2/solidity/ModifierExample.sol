pragma solidity >=0.8.0;

contract ModifierExample {
    address public owner;
    constructor() public {
        owner = msg.sender;
    }
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }
    function changeOwner(address newOwner) public onlyOwner {
        owner = newOwner;
    }
}