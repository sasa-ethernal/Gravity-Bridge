package types

// The go-ethereum ABI encoder *only* encodes function calls and then it only encodes
// function calls for which you provide an ABI json just like you would get out of the
// solidity compiler with your compiled contract.
// You are supposed to compile your contract, use abigen to generate an ABI , import
// this generated go module and then use for that for all testing and development.
// This abstraction layer is more trouble than it's worth, because we don't want to
// encode a function call at all, but instead we want to emulate a Solidity encode operation
// which has no equal available from go-ethereum.
//
// In order to work around this absurd series of problems we have to manually write the below
// 'function specification' that will encode the same arguments into a function call. We can then
// truncate the first several bytes where the call name is encoded to finally get the equal of the

const (
	// OutgoingBatchTxCheckpointABIJSON checks the ETH ABI for compatability of the OutgoingBatchTx message
	OutgoingBatchTxCheckpointABIJSON = `[{
	  "inputs": [
	    { "internalType": "bytes32", "name": "_peggyId", "type": "bytes32" },
	    { "internalType": "bytes32", "name": "_methodName", "type": "bytes32" },
		{ "internalType": "bytes32", "name": "_checkPoint", "type": "bytes32" },
	    { "internalType": "uint256[]", "name": "_amounts", "type": "uint256[]" },
	    { "internalType": "address[]", "name": "_destinations", "type": "address[]" },
	    { "internalType": "uint256[]", "name": "_fees", "type": "uint256[]" },
		{ "internalType": "uint256", "name": "_batchNonce", "type": "uint256" },
		{ "internalType": "address", "name": "_tokenContract", "type": "address" }
	  ],
	  "name": "updateValsetAndSubmitBatch",
	  "outputs": [
	    { "internalType": "bytes32", "name": "", "type": "bytes32" }
	  ],
	  "stateMutability": "pure",
	  "type": "function"
	}]`

	// ValsetCheckpointABIJSON checks the ETH ABI for compatability of the Valset update message
	ValsetCheckpointABIJSON = `[{
	  "inputs": [
	    { "internalType": "bytes32", "name": "_peggyId", "type": "bytes32" },
	    { "internalType": "bytes32", "name": "_checkpoint", "type": "bytes32" },
	    { "internalType": "uint256", "name": "_valsetNonce", "type": "uint256" },
	    { "internalType": "address[]", "name": "_validators", "type": "address[]" },
	    { "internalType": "uint256[]", "name": "_powers", "type": "uint256[]" }
	  ],
	  "name": "checkpoint",
	  "outputs": [
	    { "internalType": "bytes32", "name": "", "type": "bytes32" }
	  ],
	  "stateMutability": "pure",
	  "type": "function"
	}]`
)
