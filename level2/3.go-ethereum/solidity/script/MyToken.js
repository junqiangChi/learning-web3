const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("MyTokenModule", (m) => {
    const myToken = m.contract("MyToken", ["MyToken", "erc20", 0, 0]);
    return { myToken };
});
