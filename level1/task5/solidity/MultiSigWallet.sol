// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MultiSigWallet {
    event Log(address _ads);
    event Withdraw(address owner, uint256 amount);
    address[] owners;
    uint256 amount;
    uint256 lastSigNum;
    mapping(address => bool) isOwner;
    mapping(uint256 => mapping(address => bool)) public approved;
    struct Transaction {
        address toAds;
        uint256 value;
        bytes data;
        bool exected;
    }
    Transaction[] public transactions;
    modifier onlyOwner() {
        require(isOwner[msg.sender]);
        _;
    }
    modifier outOfTransaction(uint256 _txId) {
        require(_txId >= 0 && _txId < transactions.length);
        _;
    }

    modifier txExists(uint256 _txId) {
        require(_txId < transactions.length, "tx doesn't exist");
        _;
    }
    modifier notApproved(uint256 _txId) {
        require(!approved[_txId][msg.sender], "tx already approved");
        _;
    }
    modifier notExecuted(uint256 _txId) {
        require(!transactions[_txId].exected, "tx is exected");
        _;
    }

    event Submit(address _toAds, uint256 _amount, bytes data);
    event Approve(address indexed owner, uint256 indexed txId);
    event Revoke(address indexed owner, uint256 indexed txId);
    event Execute(uint256 indexed txId);

    constructor(address[] memory _owners, uint256 _lastSigNum) {
        require(_owners.length > 0, "tx not exist!");
        if (_lastSigNum <= 0 || _owners.length < _lastSigNum) {
            revert("invalid required number of owners");
        }
        for (uint256 i = 0; i < _owners.length; i++) {
            require(_owners[i] != address(0), "invalid owner");
            require(isOwner[_owners[i]], "owner is not unique");
            isOwner[_owners[i]] = true;
            owners.push(_owners[i]);
        }
        lastSigNum = _lastSigNum;
    }

    function deposit() external payable {
        amount += msg.value;
    }

    function submit(
        address _toAds,
        uint256 _amount,
        bytes calldata data
    ) external onlyOwner returns (uint256) {
        transactions.push(Transaction(_toAds, _amount, data, false));
        emit Submit(_toAds, _amount, data);
        return transactions.length - 1;
    }

    function approve(uint256 _txId) external outOfTransaction(_txId) {
        approved[_txId][msg.sender] = true;
        emit Approve(msg.sender, _txId);
    }

    function execute(
        uint256 _txId
    ) external onlyOwner txExists(_txId) notExecuted(_txId) {
        require(getApprovalCount(_txId) >= amount, "approvals < required");
        Transaction storage transaction = transactions[_txId];
        transaction.exected = true;
        (bool sucess, ) = transaction.toAds.call{value: transaction.value}(
            transaction.data
        );
        require(sucess, "tx failed");
        emit Execute(_txId);
    }

    function getApprovalCount(
        uint256 _txId
    ) public view returns (uint256 count) {
        for (uint256 index = 0; index < owners.length; index++) {
            if (approved[_txId][owners[index]]) {
                count += 1;
            }
        }
    }

    function revoke(
        uint256 _txId
    ) external onlyOwner txExists(_txId) notExecuted(_txId) {
        require(approved[_txId][msg.sender], "tx not approved");
        approved[_txId][msg.sender] = false;
        emit Revoke(msg.sender, _txId);
    }

    function getBalance() external view returns (uint256) {
        return amount;
    }

    receive() external payable {}
}
