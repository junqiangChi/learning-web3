// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


interface IVault {
    function deposit(uint amount) external payable;

    function withdraw(uint amount) external payable;
}

contract Bank is IVault {
    mapping(address => uint) balances;

    function deposit(uint amount) external payable {
        balances[msg.sender] += amount;
    }

    function withdraw(uint amount) external payable {
        require(balances[msg.sender] < amount, "Account balance not enough");
        balances[msg.sender] -= amount;
    }
}


interface IERC20 {
    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address recipient, uint256 amount) external returns (bool);

    function allowance(address owner, address spender) external view returns (uint256);

    function approve(address spender, uint256 amount) external returns (bool);

    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

interface Reward {
    function reward(address user, uint amount) external payable;
}

contract MyToken {
    Reward immutable reward;
    IERC20 immutable ier20;

    constructor(Reward _reward){
        reward = _reward;
    }

    function sendBonus(address user) public {
        reward.reward(user, 100);
    }
}