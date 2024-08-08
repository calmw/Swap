const {ethers, upgrades} = require("hardhat")

async function main() {
    const contractName = "UniswapV2Factory"
    console.log(`deploying ${contractName}`);

    const contractObj = await ethers.getContractFactory(contractName)

    let resp = await contractObj.deploy('0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E')

    console.log(`deployed to ${resp.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
