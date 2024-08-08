require("@nomicfoundation/hardhat-toolbox");
require('@openzeppelin/hardhat-upgrades')

const PRIVATE_KEY = process.env.META_ACCOUNT

module.exports = {
    // solidity: "0.4.18",
    // solidity: "0.5.16",
    solidity: "0.6.6",
    // solidity: "0.8.21",
    settings: {
        optimizer: {
            enabled: true,
            runs: 200
        },
        maxCodeSize: 50 * 1024 * 1024,
    },
    networks: {
        match: {
            url: "https://match-rpc-2.intoverse.co",
            accounts: [`${PRIVATE_KEY}`],
        },
        match_test: {
            url: "http://35.78.192.225:8545/",
            accounts: [`${PRIVATE_KEY}`],
        }
    }
};