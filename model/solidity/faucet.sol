pragma solidity ^0.8.0;
// SPDX-License-Identifier: GPL-3.0
contract Faucet {

	//每天转账次数
	uint transferNum=0;

	//每次转账金额
	uint eachAmount =1 ether;

	function withdraw(address addr) payable public {

		require(transferNum < 10,"hightest transfer num");

		payable(addr).transfer(eachAmount);

		transferNum++;
	}

	// function pay() public payable {}
    fallback() payable external {}
    receive() payable external {}


}
