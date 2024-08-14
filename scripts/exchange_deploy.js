const {ethers, upgrades} = require("hardhat")

async function main() {
    const contractName = "Exchange"
    console.log(`deploying ${contractName}`);

    const contractObj = await ethers.getContractFactory(contractName)

    let resp = await contractObj.deploy("0xe514013221F102684E67F8dA551463D193755fFb")

    console.log(`deployed to ${resp.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
