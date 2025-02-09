// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

contract ErrorHandlingExample {
    uint public balance;

    error Unauthorized(address caller);

    function sendHalf() public payable {
        require(msg.value % 2 == 0, "Even value required."); // 输入检查
        uint balanceBeforeTransfer = address(this).balance;
        // addr.transfer(msg.value / 2);
        assert(address(this).balance == balanceBeforeTransfer - msg.value / 2); // 内部错误检查
        revert Unauthorized(msg.sender);
    }

    function checkValue(uint value) public pure {
        if (value > 10) {
            revert("Value cannot exceed 10"); // 返回自定义错误信息
        }
    }
}