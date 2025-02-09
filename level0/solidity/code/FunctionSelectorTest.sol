// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract FunctionSelectorTest {
    constructor(){

    }
    bytes4 selector1 = bytes4(keccak256("functionName(uint256)"));


    function select(bytes4 selector ,uint x) external returns (uint z){
        (bool success, bytes memory data) =  address(this).call(abi.encodeWithSelector(selector,x));
        require(success, "Function call failed");
        z = abi.decode(data, (uint));
    }

    function getSelector() external pure returns (bytes4) {
        return this.select.selector;
    }

    //使用函数选择器与 `call` 时，需要确保调用的安全性，防止恶意代码执行。
    function execute(bytes4 selector, uint x) external returns (uint z) {
        (bool success, bytes memory data) = address(this).call(abi.encodeWithSelector(selector, x));
        require(success, "Call failed");
        z = abi.decode(data, (uint));
    }

    bytes4 func;
    function setFunction(bytes4 selector)external {
        func = selector;
    }

}
