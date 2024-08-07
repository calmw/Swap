require("@nomicfoundation/hardhat-toolbox");
require('@openzeppelin/hardhat-upgrades')
require("@nomiclabs/hardhat-waffle");

const PRIVATE_KEY = process.env.Swap

module.exports = {
    solidity: "0.8.19",
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
            url: "https://lisbon-testnet-rpc.matchtest.co",
            accounts: [`${PRIVATE_KEY}`],
        }
    }
};