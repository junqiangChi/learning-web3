// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FuntionTest {
    uint data;

    function getData() public view returns (uint){
        return data;
    }

    function getData1() private view returns (uint){
        return data;
    }

    function getData2() internal view returns (uint){
        return data;
    }

    function getData3() external view returns (uint){
        return data;
    }

    function setData(uint _data) public {
        data = _data;
    }

    function setData1(uint _data) private {
        data = _data;
    }

    function setData2(uint _data) internal {
        data = _data;
    }

    function setData3(uint _data) external {
        data = _data;
    }

    function product(uint x, uint y) public pure returns (uint){
        return x * y;
    }

    function deposit() external payable {
        data = msg.value;
    }
}