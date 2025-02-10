// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract Parent {
    uint public a;
    constructor(uint _a){
        a = _a;
    }
    function get() virtual public returns (uint);
}

abstract contract Parent1 {
    function get() virtual public returns (uint);
}

contract InheritTest is Parent(2) {
    uint public  b;
    constructor(){
        b = 1;
    }
    function get() public override view returns (uint) {
        return b;
    }
}

contract InheritTest1 is Parent {
    uint public  b;
    constructor() Parent(3){
        b = 3;
    }
    function get() public view override returns (uint){
        return b;
    }
}

contract InheritTest2 is Parent, Parent1 {
    uint public  b;
    constructor() Parent(3){
        b = 3;
    }
    function get() public override(Parent, Parent1) view returns (uint){
        return b;
    }
}

contract Car {
    uint speed;

    function driver() virtual public{}
}

contract ElectricCar is Car {
    uint batteryLevel;
    function driver()  public override {

    }
}

contract Person {
    function name() public virtual{

    }
}

contract Employee {
    function getId() public virtual  pure returns (uint){
        return 1;
    }
}

contract Manager is Person, Employee {
    function name() public override   {
    }
    function getId() public override pure returns (uint){
        return 1;
    }
}

abstract contract Shape{
    function area() public virtual returns (string memory);
}

contract Square is Shape {
    function area() public override pure returns (string memory){
        return "zhengfangxing";
    }
}
contract Circle is Shape {
    function area() public override pure returns (string memory){
        return "yuanxing";
    }
}