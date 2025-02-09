// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ControllerTest {
    uint public a;
    uint internal  b;
    uint private c;

    function setB(uint _b) public {
        b = _b;
    }

    function setC(uint _c) public {
        c = _c;
    }

    function getC() public view returns (uint) {
        return c;
    }

    function execute(uint n) public returns (uint) {
        try new ExecutorTest().sumNum(n) returns (uint sum){
            return sum;
        }catch {
            return 0;
        }
    }

}

contract ExecutorTest {
    function sumNum(uint n) public pure returns (uint){
        uint sum;
        for (uint i = 1; i <= n; i++) {
            sum += i;
        }
        return sum;
    }
}