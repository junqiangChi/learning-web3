// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

contract WETH {

    mapping (address => uint256) balances;
    mapping  (address => mapping(address=>uint256)) allowances;

    event Approval(address owner, address spender, uint256 value);
    event Deposit(address _address, uint256 amount);
    event Withdraw(address _address, uint256 amount);
    event Transfer(address srcAds, address toAds, uint256 amount);
    modifier isEnough(uint256 amount){
        require(balances[msg.sender] >= amount,"balance not enough");
        _;
    }
    
    function approve(address account,uint256 amount) public {
        allowances[msg.sender][account] = amount;
        emit Approval (msg.sender, account, amount);
    }

    function deposit() public payable {
        balances[msg.sender]+= msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint256 amount) public payable isEnough(amount) {
        balances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
        emit Withdraw(msg.sender, amount);
    }

    function totalSupply() public view  returns (uint256) {
        return balances[msg.sender];
    }

    function transfer(address _toAddress, uint256 amount) public returns (bool) {
        return transferFrom(msg.sender, _toAddress, amount);

    }

    function transferFrom(address _srcAdress, address _toAddress, uint256 amount) public payable returns  (bool) {
        require(balances[_srcAdress] >= amount, "balance not enough");
        if (_srcAdress != msg.sender){
            require(allowances[_srcAdress][msg.sender] >= amount,"");
            allowances[_srcAdress][msg.sender] -= amount;
        }
        balances[_srcAdress] -=amount;
        balances[_toAddress] += amount;
        emit Transfer(_srcAdress, _toAddress, amount);
        return true;
    }
    function receive() external payable { 
        deposit();
    }

    function fallback() external payable { 
        deposit();
    }


}