// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
// 使用 SafeMath 库的合约

contract AddTest {
    function add(uint x, uint y) public pure returns (uint) {
        return SafeMath.add(x, y);
    }
}

contract TestLib {
    using SafeMath for uint;
    function add(uint x, uint y) public pure returns (uint){
        return x.add(y);
    }
}

library Search {
    function indexOf(uint[] storage self, uint value) public view returns (int) {
        for (uint i = 0; i < self.length; i++) {
            if (self[i] == value) return int(i);
        }
        return - 1;
    }
}

contract A {
    using Search for uint[];
    uint[] data;

    constructor(){
        data = [1, 2, 3, 4];
    }

    function findIndex(uint n) public view returns (int) {
        return data.indexOf(n);
    }
}

library MathLib {
    function plus(uint x, uint y) public pure returns (uint) {
        return x + y;
    }

    function sub(uint x, uint y) public pure returns (uint) {
        require(x < y, "x must >= y");
        return x - y;
    }

    function mul(uint x, uint y) public pure returns (uint) {
        return x * y;
    }

    function divid(uint x, uint y) public pure returns (uint) {
        require(y != 0, "The divisor cannot be zero");
        return x / y;
    }
}

contract LibraryTest {
    using MathLib for uint;
    function plus(uint x, uint y) public pure returns (uint){
        return MathLib.plus(x, y);
    }

    function sub(uint x, uint y) public pure returns (uint){
        return MathLib.sub(x, y);
    }

    function mul(uint x, uint y) public pure returns (uint){
        return MathLib.mul(x, y);
    }

    function divide(uint x, uint y) public pure returns (uint){
        return MathLib.divid(x, y);
    }
}