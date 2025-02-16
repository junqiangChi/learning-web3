// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

contract CrowdFunding {
    address public beneficiary;
    uint256 public fundingGoal; //筹资目标数量
    uint256 public fundingAmount; //当前募集数量
    mapping(address => uint256) public funders;
    mapping(address => bool) private fundersInserted;

    address[] public fundersKey;

    bool AVAILABLED;

    constructor(address _beneficiary, uint256 _fundingGoal) {
        beneficiary = _beneficiary;
        fundingGoal = _fundingGoal;
        AVAILABLED = true;
    }

    function contribute() public payable {
        require(AVAILABLED, "contributions is close");
        uint256 amount = msg.value;
        uint256 fundingSum = fundingAmount + amount;
        if (fundingSum > fundingGoal) {
            uint256 fundingSub = fundingSum - fundingGoal;
            uint256 fundingPay = amount - fundingSub;
            funders[msg.sender] += fundingPay;
            fundingAmount += fundingPay;
            payable(msg.sender).transfer(fundingSub);
        } else {
            funders[msg.sender] += amount;
            fundingAmount += amount;
        }

        if (!fundersInserted[msg.sender]) {
            fundersInserted[msg.sender] = true;
            fundersKey.push(msg.sender);
        }
    }

    function close() public returns (bool) {
        if (fundingGoal > fundingAmount) {
            return false;
        }
        uint256 amount = fundingAmount;
        fundingAmount = 0;
        AVAILABLED = false;
        payable(beneficiary).transfer(amount);

        return true;
    }

    function getfundersCount() public view returns (uint) {
        return fundersKey.length;
    }
}
