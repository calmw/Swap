const {ethers, upgrades} = require("hardhat")
const {bigint} = require("hardhat/internal/core/params/argumentTypes");
const {BigNumber} = require("ethers");

async function main() {
    const contractName = "ToxB"
    console.log(`deploying ${contractName}`);

    const contractObj = await ethers.getContractFactory(contractName)

    let resp = await contractObj.deploy(BigNumber.from("1000000000000000000000000000000000000"))

    console.log(`deployed to ${resp.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
