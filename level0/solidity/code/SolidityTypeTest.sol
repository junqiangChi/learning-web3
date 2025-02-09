// SPDX-License-Identifier: MIT
pragma solidity >0.5.0;

contract SolidityTypeTest {
    function getAddress() public view  returns (address){
        return address(this);
    }

    function destroyContract(address payable recipient) public {
        selfdestruct(recipient);
    }


    function getContractInfo() public pure returns (string memory) {
        return (type(SolidityTypeTest).name);
    }

    function isContract(address addr) internal view returns (bool) {
        uint256 size;
        assembly {size := extcodesize(addr)} // assembly 表示接下来的代码是内联汇编。
        return size > 0;
    }
}