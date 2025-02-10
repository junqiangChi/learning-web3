// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FunctionExample{
    bytes4 public storedSelector ;

    function square(uint x) public pure  returns (uint){
        return  x * x;
    }

    function double(uint x) public pure returns (uint){
        return 2 * x;
    }

    function storeSelector(bytes4 selector) public {
        storedSelector = selector;
    }

    function executeStoredFunction(uint x) public returns (uint){
        require(storedSelector != bytes4(0), "Selector not set");
        (bool success, bytes memory data) = address(this).call(abi.encodeWithSelector(storedSelector, x));
        require(success, "Function call failed");
        return abi.decode(data, (uint));
    }
}