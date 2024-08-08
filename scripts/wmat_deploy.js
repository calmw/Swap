const {ethers, upgrades} = require("hardhat")

async function main() {
    const contractName = "WETH9"
    console.log(`deploying ${contractName}`);

    // const contractObj = await hre.ethers.deployContract(contractName, [], {});
    const contractObj = await ethers.getContractFactory(contractName)

    // await contractObj.deployed();
    let resp = await contractObj.deploy([])

    console.log(`deployed to ${resp.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
