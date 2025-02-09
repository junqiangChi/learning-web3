// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract StructTest {
    struct ProduceInfo {
        uint produceId;
        string name;
        uint score;
        uint count;
    }

    struct OrderInfo {
        uint orderId;
        ProduceInfo produceInfo;
        address user;
    }

    constructor (){
        produces[1] = ProduceInfo(1, "pen", 20, 3200);
        produces[2] = ProduceInfo(2, "box", 100, 3100);
        produces[3] = ProduceInfo(3, "bag", 200, 3300);
        produces[4] = ProduceInfo(4, "pc", 20000, 300);
    }

    mapping(uint => ProduceInfo) produces;
    mapping(uint => OrderInfo) orders;


    function createOrder(uint _id, uint _produceId) public {
        orders[_id] = OrderInfo(_id, produces[_produceId], msg.sender);
    }

    function getOrderInfo(uint orderId) public view returns (OrderInfo memory){
        return orders[orderId];
    }z
}