// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract TinyTest{
    uint[] x;
    function f(uint[] memory memoryArr) public returns  (uint[] memory){
        x = memoryArr;
        uint[] storage y = x;
        y[0] = 1;
        delete x; //删除x后，y也不能用了
        return y;
    }
    function getX() public view returns (uint[] memory){
        return x;
    }
}