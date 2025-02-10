// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ArrayExample {
    uint[] arr1 = [1, 2, 3];

    function sunArrayValue(uint[] memory arr) public pure returns (uint){
        uint sum;
        for (uint i = 0; i < arr.length; i++) {
            sum += arr[i];
        }
        return sum;
    }

    function removeIndexValue(uint index) public {
        require(arr1.length > index && index >= 0, "index out of bounds");
        if (index < arr1.length - 1) {
            for (uint i = index; i < arr1.length; i++) {
                arr1[i] = arr1[i + 1];
            }
        }
        arr1.pop();
    }

    function getArr() public view returns (uint[] memory){
        return arr1;
    }

    function getArrLength() public view returns (uint){
        return arr1.length;
    }

    function pop() public {
        arr1.pop();
    }

    function push(uint value) public {
        arr1.push(value);
    }
}