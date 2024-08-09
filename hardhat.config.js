require("@nomicfoundation/hardhat-toolbox");
require('@openzeppelin/hardhat-upgrades')

const PRIVATE_KEY = process.env.Swap

module.exports = {
    solidity: "0.8.21",
    settings: {
        optimizer: {
            enabled: true,
            runs: 2000
        },
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