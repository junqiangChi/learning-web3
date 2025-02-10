pragma solidity ^0.8.0;

contract Registry {
    mapping(uint256 => address) public documents;

    function register(uint256 hash) public {
        documents[hash] = msg.sender;
    }
}

// 用事件消耗gas更少
contract DocumentRegistry {
    event Registered(uint256 hash, address sender);
    function register(uint256 hash) public {
        emit Registered(hash, msg.sender);
    }
}