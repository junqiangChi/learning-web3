// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TinyExample {
    uint[] data;

    function updateData(uint[] memory newData) public {
        data = newData;
    }

    function getData() public view returns (uint[] memory) {
        return data;
    }

    function modifyStorageData(uint index, uint value) public {
        require(index > data.length, "index out of bounds");
        data[index] = value;

    }

    function modifyMemoryData(uint[] memory memData) public pure returns (uint[] memory){
        require(memData.length <= 0, "This array is empty");
        memData[0] = 100;
        return memData;
    }
}