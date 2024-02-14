package contracts

const (
	MULTICALL3_CONTRACT_ABI = `[
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "aggregate",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bytes[]",
			  "name": "returnData",
			  "type": "bytes[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bool",
				  "name": "allowFailure",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call3[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "aggregate3",
		  "outputs": [
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bool",
				  "name": "allowFailure",
				  "type": "bool"
				},
				{
				  "internalType": "uint256",
				  "name": "value",
				  "type": "uint256"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call3Value[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "aggregate3Value",
		  "outputs": [
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "blockAndAggregate",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			},
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getBasefee",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "basefee",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			}
		  ],
		  "name": "getBlockHash",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getBlockNumber",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getChainId",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "chainid",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockCoinbase",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "coinbase",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockDifficulty",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "difficulty",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockGasLimit",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "gaslimit",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockTimestamp",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "timestamp",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "addr",
			  "type": "address"
			}
		  ],
		  "name": "getEthBalance",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "balance",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getLastBlockHash",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bool",
			  "name": "requireSuccess",
			  "type": "bool"
			},
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "tryAggregate",
		  "outputs": [
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bool",
			  "name": "requireSuccess",
			  "type": "bool"
			},
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "tryBlockAndAggregate",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			},
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		}
	  ]`
	K2_LENDING_CONTRACT_ABI       = `[{"type":"constructor","inputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"DEFAULT_ADMIN_ROLE","inputs":[],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"allowance","inputs":[{"name":"owner","type":"address","internalType":"address"},{"name":"spender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"approve","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"assumedLiquidity","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"balanceOf","inputs":[{"name":"account","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"banNodeOperator","inputs":[{"name":"bannedNO","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"batchNodeOperatorDeposit","inputs":[{"name":"_blsPubkeys","type":"bytes[]","internalType":"bytes[]"},{"name":"_payoutRecipients","type":"address[]","internalType":"address[]"},{"name":"_blsSignatures","type":"bytes[]","internalType":"bytes[]"},{"name":"_ecdsaSignatures","type":"tuple[]","internalType":"struct IProposerRegistry.SignatureECDSA[]","components":[{"name":"v","type":"uint8","internalType":"uint8"},{"name":"r","type":"bytes32","internalType":"bytes32"},{"name":"s","type":"bytes32","internalType":"bytes32"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"blsPublicKeyToKicked","inputs":[{"name":"","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"blsPublicKeyToNodeOperator","inputs":[{"name":"","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"borrow","inputs":[{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"rstConfigParams","type":"tuple","internalType":"struct IK2Lending.RSTConfigParams","components":[{"name":"mintRST","type":"bool","internalType":"bool"},{"name":"initialSupply","type":"uint256","internalType":"uint256"},{"name":"recipientOfRemainingSupply","type":"address","internalType":"address"},{"name":"percentageContributionToIncentives","type":"uint256","internalType":"uint256"},{"name":"symbol","type":"string","internalType":"string"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"borrowDuration","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"borrowFor","inputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"rstConfigParams","type":"tuple","internalType":"struct IK2Lending.RSTConfigParams","components":[{"name":"mintRST","type":"bool","internalType":"bool"},{"name":"initialSupply","type":"uint256","internalType":"uint256"},{"name":"recipientOfRemainingSupply","type":"address","internalType":"address"},{"name":"percentageContributionToIncentives","type":"uint256","internalType":"uint256"},{"name":"symbol","type":"string","internalType":"string"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"borrowedLiquidity","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"claimKETH","inputs":[{"name":"lender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"claimableKETH","inputs":[{"name":"lender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"claimableKETHForNodeOperator","inputs":[{"name":"nodeOperator","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"daoAddress","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"debtPositions","inputs":[{"name":"","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"hook","type":"address","internalType":"address"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"designatedVerifierURI","type":"string","internalType":"string"},{"name":"principalAmount","type":"uint256","internalType":"uint256"},{"name":"interestPerSec_RAY","type":"uint256","internalType":"uint256"},{"name":"endTimestamp","type":"uint256","internalType":"uint256"},{"name":"slashAmount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"debtors","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"decimals","inputs":[],"outputs":[{"name":"","type":"uint8","internalType":"uint8"}],"stateMutability":"view"},{"type":"function","name":"decreaseAllowance","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"subtractedValue","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"delegatedClaim","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"deposit","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"depositFor","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"recipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"getBorrowedLiquidity","inputs":[],"outputs":[{"name":"res","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getDebtor","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"tuple","internalType":"struct IK2Lending.DebtPosition","components":[{"name":"debtor","type":"address","internalType":"address"},{"name":"hook","type":"address","internalType":"address"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"designatedVerifierURI","type":"string","internalType":"string"},{"name":"principalAmount","type":"uint256","internalType":"uint256"},{"name":"interestPerSec_RAY","type":"uint256","internalType":"uint256"},{"name":"endTimestamp","type":"uint256","internalType":"uint256"},{"name":"slashAmount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"}]}],"stateMutability":"view"},{"type":"function","name":"getOutstandingInterest","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getRemainingDuration","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getRoleAdmin","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"}],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"getTotalBorrowableAmount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getTotalBorrowableAmountWithMaxBorrowRatio","inputs":[{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"grantRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"hasRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"increaseAllowance","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"addedValue","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"increaseDebt","inputs":[{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"resetDuration","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"increaseDebtFor","inputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"resetDuration","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"initialize","inputs":[{"name":"_configurator","type":"address","internalType":"address"},{"name":"_daoAddress","type":"address","internalType":"address"},{"name":"_k2Incentives","type":"address","internalType":"address"},{"name":"_name","type":"string","internalType":"string"},{"name":"_symbol","type":"string","internalType":"string"},{"name":"_proposerRegistry","type":"address","internalType":"address"},{"name":"_nodeOperatorInclusionList","type":"address","internalType":"address"},{"name":"_maxSlashableRatio_RAY","type":"uint256","internalType":"uint256"},{"name":"_terminatePenalty_RAY","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"interestRateModelByType","inputs":[{"name":"","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"isNodeOperatorBanned","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"k2Incentives","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"lenders","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"cumulativeKethPerShareLU_RAY","type":"uint256","internalType":"uint256"},{"name":"kethEarned","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"liquidate","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"maxSlashableRatio_RAY","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"name","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorClaim","inputs":[{"name":"_recipientOverride","type":"address","internalType":"address"},{"name":"_nodeOperator","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorInclusionList","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorKick","inputs":[{"name":"_reporter","type":"address","internalType":"address"},{"name":"_blsPubkey","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorToBlsPublicKeyCount","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorToLendPosition","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"cumulativeKethPerShareLU_RAY","type":"uint256","internalType":"uint256"},{"name":"kethEarned","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorToPayoutRecipient","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorWithdraw","inputs":[{"name":"_recipientOverride","type":"address","internalType":"address"},{"name":"_nodeOperator","type":"address","internalType":"address"},{"name":"_blsPubkey","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"proposerRegistry","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"renounceRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"revokeRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"rstModule","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract RSTModule"}],"stateMutability":"view"},{"type":"function","name":"setBorrowDuration","inputs":[{"name":"_duration","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setConfigurator","inputs":[{"name":"newConfigurator","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setDelegatedRecipient","inputs":[{"name":"_recipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setDesignatedVerifierURI","inputs":[{"name":"designatedVerifierURI","type":"string","internalType":"string"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setHookAsDebtorForSBP","inputs":[{"name":"_hook","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setIndexUpdateBeacon","inputs":[{"name":"_beacon","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setInterestRateModel","inputs":[{"name":"_type","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"_newInterestRateModel","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setKETHAddress","inputs":[{"name":"_kETH","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setMinParams","inputs":[{"name":"newMinDepositLimit","type":"uint256","internalType":"uint256"},{"name":"newMinTransferLimit","type":"uint256","internalType":"uint256"},{"name":"newMinLockUpPeriod","type":"uint256","internalType":"uint256"},{"name":"_maxSlashableRatio_RAY","type":"uint256","internalType":"uint256"},{"name":"_terminatePenalty_RAY","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setNodeOperatorInclusionList","inputs":[{"name":"_nodeOperatorInclusionList","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setRSTModule","inputs":[{"name":"_rstModule","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"slash","inputs":[{"name":"slashType","type":"uint8","internalType":"enum ReporterRegistry.SlashType"},{"name":"debtor","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"recipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"supportsInterface","inputs":[{"name":"interfaceId","type":"bytes4","internalType":"bytes4"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"symbol","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"terminate","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"terminatePenalty_RAY","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"topUpAndTerminate","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"topUpSlashAmount","inputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"totalSupply","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"transfer","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"transferFrom","inputs":[{"name":"from","type":"address","internalType":"address"},{"name":"to","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"updateInterest","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateLenderPosition","inputs":[{"name":"lender","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateNodeOperatorPosition","inputs":[{"name":"nodeOperator","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"withdraw","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"claim","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"event","name":"Approval","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"spender","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Borrowed","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"designatedVerifier","type":"address","indexed":true,"internalType":"address"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"interestPaid","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"DelegatedClaimRecipientSet","inputs":[{"name":"claimRecipient","type":"address","indexed":true,"internalType":"address"},{"name":"delegateClaimer","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"DesignatedVerifierURIUpdated","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"designatedVerifierURI","type":"string","indexed":false,"internalType":"string"}],"anonymous":false},{"type":"event","name":"IncreasedDebt","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"designatedVerifier","type":"address","indexed":true,"internalType":"address"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"resetDuration","type":"bool","indexed":false,"internalType":"bool"},{"name":"interestPaid","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Initialized","inputs":[{"name":"version","type":"uint8","indexed":false,"internalType":"uint8"}],"anonymous":false},{"type":"event","name":"KETHClaimed","inputs":[{"name":"depositor","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"KETHDeposited","inputs":[{"name":"depositor","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"KETHWithdrawn","inputs":[{"name":"depositor","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Liquidated","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"liquidator","type":"address","indexed":true,"internalType":"address"},{"name":"liquidationAmount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"NodeOperatorClaimed","inputs":[{"name":"nodeOperator","type":"address","indexed":true,"internalType":"address"},{"name":"payoutRecipient","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"NodeOperatorDeposited","inputs":[{"name":"operator","type":"address","indexed":true,"internalType":"address"},{"name":"blsPublicKey","type":"bytes","indexed":false,"internalType":"bytes"},{"name":"payoutRecipient","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"NodeOperatorWithdrawn","inputs":[{"name":"operator","type":"address","indexed":true,"internalType":"address"},{"name":"blsPublicKey","type":"bytes","indexed":false,"internalType":"bytes"},{"name":"kicked","type":"bool","indexed":false,"internalType":"bool"}],"anonymous":false},{"type":"event","name":"RepaidSlashAmount","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"topupAmount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"RoleAdminChanged","inputs":[{"name":"role","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"previousAdminRole","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"newAdminRole","type":"bytes32","indexed":true,"internalType":"bytes32"}],"anonymous":false},{"type":"event","name":"RoleGranted","inputs":[{"name":"role","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"account","type":"address","indexed":true,"internalType":"address"},{"name":"sender","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"RoleRevoked","inputs":[{"name":"role","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"account","type":"address","indexed":true,"internalType":"address"},{"name":"sender","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"SBPHookUpdated","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"hook","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Slashed","inputs":[{"name":"debtor","type":"address","indexed":false,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"recipient","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Terminated","inputs":[{"name":"debtor","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Transfer","inputs":[{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"error","name":"ComeBackLater","inputs":[]},{"type":"error","name":"ExceedMaxBorrowRatio","inputs":[]},{"type":"error","name":"ExceedMaxSlashableAmountPerCorruption","inputs":[]},{"type":"error","name":"ExceedMaxSlashableAmountPerLiveness","inputs":[]},{"type":"error","name":"HasDebt","inputs":[]},{"type":"error","name":"InvalidDebtPositionType","inputs":[]},{"type":"error","name":"InvalidLength","inputs":[]},{"type":"error","name":"MinDepositLimit","inputs":[]},{"type":"error","name":"MustPaySlashedAmount","inputs":[]},{"type":"error","name":"NoDebt","inputs":[]},{"type":"error","name":"NoElements","inputs":[]},{"type":"error","name":"NoSlashedAmount","inputs":[]},{"type":"error","name":"NoTerminationWithRST","inputs":[]},{"type":"error","name":"NodeOperatorAlreadyRegistered","inputs":[]},{"type":"error","name":"NodeOperatorBLSKeyNotPermitted","inputs":[]},{"type":"error","name":"NodeOperatorInvalid","inputs":[]},{"type":"error","name":"NodeOperatorInvalidRepresentative","inputs":[]},{"type":"error","name":"NodeOperatorInvalidStatus","inputs":[]},{"type":"error","name":"NodeOperatorKicked","inputs":[]},{"type":"error","name":"NodeOperatorNotRegistered","inputs":[]},{"type":"error","name":"NotAbleToLiquidate","inputs":[]},{"type":"error","name":"NotAllowedToDecreaseInterestRate","inputs":[]},{"type":"error","name":"NotEnoughAssumedLiquidity","inputs":[]},{"type":"error","name":"NotEnoughOutstandingInterestToSlash","inputs":[]},{"type":"error","name":"TooSmall","inputs":[]},{"type":"error","name":"Unauthorized","inputs":[]},{"type":"error","name":"ZeroAddress","inputs":[]}]`
	K2_NODE_OPERATOR_CONTRACT_ABI = `[
		{
			"inputs": [],
			"stateMutability": "nonpayable",
			"type": "constructor"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"internalType": "address",
					"name": "previousAdmin",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "address",
					"name": "newAdmin",
					"type": "address"
				}
			],
			"name": "AdminChanged",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "beacon",
					"type": "address"
				}
			],
			"name": "BeaconUpgraded",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [],
			"name": "EIP712DomainChanged",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"internalType": "uint8",
					"name": "version",
					"type": "uint8"
				}
			],
			"name": "Initialized",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "uint256",
					"name": "newValue",
					"type": "uint256"
				}
			],
			"name": "MaxNativeDelegationPerNodeOperatorUpdated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "uint256",
					"name": "newValue",
					"type": "uint256"
				}
			],
			"name": "MaxNumOfOpenNativeDelegationsUpdated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "previousOwner",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "newOwner",
					"type": "address"
				}
			],
			"name": "OwnershipTransferred",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "implementation",
					"type": "address"
				}
			],
			"name": "Upgraded",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "target",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "bool",
					"name": "enabled",
					"type": "bool"
				}
			],
			"name": "isPartOfInclusionListUpdated",
			"type": "event"
		},
		{
			"inputs": [],
			"name": "E_BALANCE_REPORT_TYPEHASH",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "MAX_NATIVE_DELEGATION_PER_NODE_OPERATOR",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "MAX_OPEN_NATIVE_DELEGATION_CAPACITY",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes[]",
					"name": "_blsPubkeys",
					"type": "bytes[]"
				},
				{
					"internalType": "uint256[]",
					"name": "_effectiveBalances",
					"type": "uint256[]"
				},
				{
					"components": [
						{
							"internalType": "uint8",
							"name": "v",
							"type": "uint8"
						},
						{
							"internalType": "bytes32",
							"name": "r",
							"type": "bytes32"
						},
						{
							"internalType": "bytes32",
							"name": "s",
							"type": "bytes32"
						}
					],
					"internalType": "struct EIP712Verifier.SignatureECDSA[]",
					"name": "_designatedVerifierSignatures",
					"type": "tuple[]"
				}
			],
			"name": "batchNodeOperatorKick",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "eip712Domain",
			"outputs": [
				{
					"internalType": "bytes1",
					"name": "fields",
					"type": "bytes1"
				},
				{
					"internalType": "string",
					"name": "name",
					"type": "string"
				},
				{
					"internalType": "string",
					"name": "version",
					"type": "string"
				},
				{
					"internalType": "uint256",
					"name": "chainId",
					"type": "uint256"
				},
				{
					"internalType": "address",
					"name": "verifyingContract",
					"type": "address"
				},
				{
					"internalType": "bytes32",
					"name": "salt",
					"type": "bytes32"
				},
				{
					"internalType": "uint256[]",
					"name": "extensions",
					"type": "uint256[]"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "getDomainSeparator",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "incentivesHook",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "_initialOwner",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_k2Lending",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_reporterRegistry",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "_maxNumOfKeysThatCanRegister",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "_maxNumOfKeysThatCanRegisterPerMember",
					"type": "uint256"
				},
				{
					"internalType": "address",
					"name": "_proposerRegistry",
					"type": "address"
				}
			],
			"name": "init",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubKey",
					"type": "bytes"
				}
			],
			"name": "isNewBLSKeyPermitted",
			"outputs": [
				{
					"internalType": "bool",
					"name": "",
					"type": "bool"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"name": "isPartOfInclusionList",
			"outputs": [
				{
					"internalType": "bool",
					"name": "",
					"type": "bool"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubkey",
					"type": "bytes"
				},
				{
					"internalType": "uint256",
					"name": "_effectiveBalance",
					"type": "uint256"
				},
				{
					"components": [
						{
							"internalType": "uint8",
							"name": "v",
							"type": "uint8"
						},
						{
							"internalType": "bytes32",
							"name": "r",
							"type": "bytes32"
						},
						{
							"internalType": "bytes32",
							"name": "s",
							"type": "bytes32"
						}
					],
					"internalType": "struct EIP712Verifier.SignatureECDSA",
					"name": "_designatedVerifierSignature",
					"type": "tuple"
				}
			],
			"name": "isValidEffectiveBalanceReport",
			"outputs": [
				{
					"internalType": "bool",
					"name": "",
					"type": "bool"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "lending",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "_verifier",
					"type": "address"
				},
				{
					"internalType": "bool",
					"name": "_isEnabled",
					"type": "bool"
				}
			],
			"name": "manageDesignatedVerifier",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes[]",
					"name": "_blsPubKeys",
					"type": "bytes[]"
				},
				{
					"internalType": "uint256[]",
					"name": "_effectiveBalances",
					"type": "uint256[]"
				},
				{
					"components": [
						{
							"internalType": "uint8",
							"name": "v",
							"type": "uint8"
						},
						{
							"internalType": "bytes32",
							"name": "r",
							"type": "bytes32"
						},
						{
							"internalType": "bytes32",
							"name": "s",
							"type": "bytes32"
						}
					],
					"internalType": "struct EIP712Verifier.SignatureECDSA[]",
					"name": "_designatedVerifierSignatures",
					"type": "tuple[]"
				}
			],
			"name": "nodeOperatorClaim",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubkey",
					"type": "bytes"
				},
				{
					"internalType": "uint256",
					"name": "_effectiveBalance",
					"type": "uint256"
				},
				{
					"components": [
						{
							"internalType": "uint8",
							"name": "v",
							"type": "uint8"
						},
						{
							"internalType": "bytes32",
							"name": "r",
							"type": "bytes32"
						},
						{
							"internalType": "bytes32",
							"name": "s",
							"type": "bytes32"
						}
					],
					"internalType": "struct EIP712Verifier.SignatureECDSA",
					"name": "_designatedVerifierSignature",
					"type": "tuple"
				}
			],
			"name": "nodeOperatorKick",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubKey",
					"type": "bytes"
				},
				{
					"internalType": "uint256",
					"name": "_effectiveBalance",
					"type": "uint256"
				},
				{
					"components": [
						{
							"internalType": "uint8",
							"name": "v",
							"type": "uint8"
						},
						{
							"internalType": "bytes32",
							"name": "r",
							"type": "bytes32"
						},
						{
							"internalType": "bytes32",
							"name": "s",
							"type": "bytes32"
						}
					],
					"internalType": "struct EIP712Verifier.SignatureECDSA",
					"name": "_designatedVerifierSignature",
					"type": "tuple"
				}
			],
			"name": "nodeOperatorWithdraw",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubKey",
					"type": "bytes"
				}
			],
			"name": "onBLSKeyRegisteredToK2",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubKey",
					"type": "bytes"
				}
			],
			"name": "onBLSKeyUnregisteredToK2",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "owner",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "proposerRegistry",
			"outputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "proxiableUUID",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "renounceOwnership",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubKey",
					"type": "bytes"
				}
			],
			"name": "reportExitFromProposerRegistry",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubkey",
					"type": "bytes"
				},
				{
					"internalType": "uint256",
					"name": "_effectiveBalance",
					"type": "uint256"
				}
			],
			"name": "reportTypedHash",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "typedHash",
					"type": "bytes32"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "reporterRegistry",
			"outputs": [
				{
					"internalType": "contract ReporterRegistry",
					"name": "",
					"type": "address"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "_hook",
					"type": "address"
				}
			],
			"name": "setIncentivesHook",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"name": "totalNativeDelegationsForRepresentative",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "totalNumberOfNativeDelegationKeys",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "",
					"type": "address"
				}
			],
			"name": "totalNumberOfRegisteredKeysForInclusionListMember",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "totalOpenNativeDelegationCapacityConsumed",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "newOwner",
					"type": "address"
				}
			],
			"name": "transferOwnership",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address[]",
					"name": "_targets",
					"type": "address[]"
				},
				{
					"internalType": "bool[]",
					"name": "_enabled",
					"type": "bool[]"
				}
			],
			"name": "updateInclusionList",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "_maxNumOfKeysThatCanRegister",
					"type": "uint256"
				}
			],
			"name": "updateMaxNativeDelegationPerNodeOperator",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "_maxOpenNativeDelegationCapacity",
					"type": "uint256"
				}
			],
			"name": "updateOpenNativeDelegationCapacity",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "newImplementation",
					"type": "address"
				}
			],
			"name": "upgradeTo",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "newImplementation",
					"type": "address"
				},
				{
					"internalType": "bytes",
					"name": "data",
					"type": "bytes"
				}
			],
			"name": "upgradeToAndCall",
			"outputs": [],
			"stateMutability": "payable",
			"type": "function"
		}
	]`
	PROPOSER_REGISTRY_CONTRACT_ABI = `[
		{
		  "inputs": [],
		  "stateMutability": "nonpayable",
		  "type": "constructor"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "payoutPool",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "reporterRegistry",
			  "type": "address"
			}
		  ],
		  "name": "ContractsSet",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "uint8",
			  "name": "version",
			  "type": "uint8"
			}
		  ],
		  "name": "Initialized",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "uint256",
			  "name": "newThreshold",
			  "type": "uint256"
			}
		  ],
		  "name": "KickThresholdChanged",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "previousOwner",
			  "type": "address"
			},
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "newOwner",
			  "type": "address"
			}
		  ],
		  "name": "OwnershipTransferred",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ProposerActivated",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ProposerExited",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ProposerKicked",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ProposerOptedIntoPayoutPool",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ProposerOptedOutOfPayoutPool",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			},
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "payoutRecipient",
			  "type": "address"
			}
		  ],
		  "name": "ProposerPayoutRecipientUpdated",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ProposerPositionedForRagequit",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			},
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "signature",
			  "type": "bytes"
			},
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "payoutRecipient",
			  "type": "address"
			}
		  ],
		  "name": "ProposerRegistered",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "bytes",
			  "name": "blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ProposerReported",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "signatureSwapper",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "bool",
			  "name": "enabled",
			  "type": "bool"
			}
		  ],
		  "name": "SignatureSwapperPermissionsChanged",
		  "type": "event"
		},
		{
		  "inputs": [],
		  "name": "KICK_THRESHOLD",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "PAYOUT_CYCLE_LENGTH",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "PAYOUT_POOL",
		  "outputs": [
			{
			  "internalType": "contract PayoutPool",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "REGISTRATION_TYPEHASH",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "REPORTER_REGISTRY",
		  "outputs": [
			{
			  "internalType": "contract ReporterRegistry",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "activateProposers",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "activeValidators",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "",
			  "type": "bytes"
			}
		  ],
		  "name": "alternativeFeeRecipient",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes[]",
			  "name": "_blsKeys",
			  "type": "bytes[]"
			},
			{
			  "internalType": "address[]",
			  "name": "_payoutRecipients",
			  "type": "address[]"
			},
			{
			  "internalType": "bytes[]",
			  "name": "_blsSignatures",
			  "type": "bytes[]"
			},
			{
			  "components": [
				{
				  "internalType": "uint8",
				  "name": "v",
				  "type": "uint8"
				},
				{
				  "internalType": "bytes32",
				  "name": "r",
				  "type": "bytes32"
				},
				{
				  "internalType": "bytes32",
				  "name": "s",
				  "type": "bytes32"
				}
			  ],
			  "internalType": "struct ProposerRegistry.SignatureECDSA[]",
			  "name": "_ecdsaSignatures",
			  "type": "tuple[]"
			},
			{
			  "internalType": "bool[]",
			  "name": "_openClaims",
			  "type": "bool[]"
			}
		  ],
		  "name": "batchRegisterProposer",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes[]",
			  "name": "_blsKeys",
			  "type": "bytes[]"
			},
			{
			  "internalType": "address[]",
			  "name": "_payoutRecipients",
			  "type": "address[]"
			},
			{
			  "internalType": "bytes[]",
			  "name": "_blsSignatures",
			  "type": "bytes[]"
			},
			{
			  "components": [
				{
				  "internalType": "uint8",
				  "name": "v",
				  "type": "uint8"
				},
				{
				  "internalType": "bytes32",
				  "name": "r",
				  "type": "bytes32"
				},
				{
				  "internalType": "bytes32",
				  "name": "s",
				  "type": "bytes32"
				}
			  ],
			  "internalType": "struct ProposerRegistry.SignatureECDSA[]",
			  "name": "_ecdsaSignatures",
			  "type": "tuple[]"
			},
			{
			  "internalType": "bool[]",
			  "name": "_openClaims",
			  "type": "bool[]"
			},
			{
			  "internalType": "address[]",
			  "name": "_alternativeFeeRecipients",
			  "type": "address[]"
			}
		  ],
		  "name": "batchRegisterProposerWithoutPayoutPoolRegistration",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "",
			  "type": "bytes"
			}
		  ],
		  "name": "blsPublicKeyToProposer",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "activationBlock",
			  "type": "uint256"
			},
			{
			  "internalType": "uint256",
			  "name": "exitClaimAmount",
			  "type": "uint256"
			},
			{
			  "internalType": "uint256",
			  "name": "exitBlock",
			  "type": "uint256"
			},
			{
			  "internalType": "address",
			  "name": "payoutRecipient",
			  "type": "address"
			},
			{
			  "internalType": "address",
			  "name": "representative",
			  "type": "address"
			},
			{
			  "internalType": "enum ProposerRegistry.ProposerStates",
			  "name": "status",
			  "type": "uint8"
			},
			{
			  "internalType": "uint8",
			  "name": "reportCount",
			  "type": "uint8"
			},
			{
			  "internalType": "bool",
			  "name": "openClaim",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes[]",
			  "name": "_blsKeys",
			  "type": "bytes[]"
			}
		  ],
		  "name": "checkBatchOperationalStatus",
		  "outputs": [
			{
			  "internalType": "bool[]",
			  "name": "",
			  "type": "bool[]"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			},
			{
			  "internalType": "address",
			  "name": "_payoutRecipient",
			  "type": "address"
			},
			{
			  "internalType": "bytes",
			  "name": "_blsSignature",
			  "type": "bytes"
			},
			{
			  "internalType": "address",
			  "name": "_representative",
			  "type": "address"
			}
		  ],
		  "name": "computeTypedStructHash",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getDomainSeparator",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "getExitClaimAmount",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "getProposerAccounts",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			},
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getProposerLength",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "getProposerStatus",
		  "outputs": [
			{
			  "internalType": "uint8",
			  "name": "",
			  "type": "uint8"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "_start",
			  "type": "uint256"
			},
			{
			  "internalType": "uint256",
			  "name": "_end",
			  "type": "uint256"
			}
		  ],
		  "name": "getProposers",
		  "outputs": [
			{
			  "components": [
				{
				  "internalType": "uint256",
				  "name": "activationBlock",
				  "type": "uint256"
				},
				{
				  "internalType": "uint256",
				  "name": "exitClaimAmount",
				  "type": "uint256"
				},
				{
				  "internalType": "uint256",
				  "name": "exitBlock",
				  "type": "uint256"
				},
				{
				  "internalType": "address",
				  "name": "payoutRecipient",
				  "type": "address"
				},
				{
				  "internalType": "address",
				  "name": "representative",
				  "type": "address"
				},
				{
				  "internalType": "enum ProposerRegistry.ProposerStates",
				  "name": "status",
				  "type": "uint8"
				},
				{
				  "internalType": "uint8",
				  "name": "reportCount",
				  "type": "uint8"
				},
				{
				  "internalType": "bool",
				  "name": "openClaim",
				  "type": "bool"
				}
			  ],
			  "internalType": "struct ProposerRegistry.Proposer[]",
			  "name": "",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "historicalActivatedProposers",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_signatureSwapper",
			  "type": "address"
			},
			{
			  "internalType": "address",
			  "name": "_owner",
			  "type": "address"
			},
			{
			  "internalType": "uint256",
			  "name": "_kickThreshold",
			  "type": "uint256"
			},
			{
			  "internalType": "string",
			  "name": "_eip712Name",
			  "type": "string"
			},
			{
			  "internalType": "string",
			  "name": "_eip712Version",
			  "type": "string"
			},
			{
			  "internalType": "uint256",
			  "name": "_payoutCycleLength",
			  "type": "uint256"
			}
		  ],
		  "name": "init",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "isProposerOperational",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "isProposerReportable",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			},
			{
			  "internalType": "address",
			  "name": "_claimer",
			  "type": "address"
			}
		  ],
		  "name": "isRewardClaimAuthorized",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "name": "isSignatureSwapper",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "kickProposer",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			}
		  ],
		  "name": "optIntoPayoutPool",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "owner",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "positionForRagequit",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "name": "proposers",
		  "outputs": [
			{
			  "internalType": "bytes",
			  "name": "",
			  "type": "bytes"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsPublicKey",
			  "type": "bytes"
			}
		  ],
		  "name": "ragequitProposer",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			},
			{
			  "internalType": "address",
			  "name": "_payoutRecipient",
			  "type": "address"
			},
			{
			  "internalType": "bytes",
			  "name": "_blsSignature",
			  "type": "bytes"
			},
			{
			  "components": [
				{
				  "internalType": "uint8",
				  "name": "v",
				  "type": "uint8"
				},
				{
				  "internalType": "bytes32",
				  "name": "r",
				  "type": "bytes32"
				},
				{
				  "internalType": "bytes32",
				  "name": "s",
				  "type": "bytes32"
				}
			  ],
			  "internalType": "struct ProposerRegistry.SignatureECDSA",
			  "name": "_ecdsaSignature",
			  "type": "tuple"
			},
			{
			  "internalType": "bool",
			  "name": "_openClaim",
			  "type": "bool"
			}
		  ],
		  "name": "registerProposer",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			},
			{
			  "internalType": "address",
			  "name": "_payoutRecipient",
			  "type": "address"
			},
			{
			  "internalType": "bytes",
			  "name": "_blsSignature",
			  "type": "bytes"
			},
			{
			  "components": [
				{
				  "internalType": "uint8",
				  "name": "v",
				  "type": "uint8"
				},
				{
				  "internalType": "bytes32",
				  "name": "r",
				  "type": "bytes32"
				},
				{
				  "internalType": "bytes32",
				  "name": "s",
				  "type": "bytes32"
				}
			  ],
			  "internalType": "struct ProposerRegistry.SignatureECDSA",
			  "name": "_ecdsaSignature",
			  "type": "tuple"
			},
			{
			  "internalType": "bool",
			  "name": "_openClaim",
			  "type": "bool"
			},
			{
			  "internalType": "address",
			  "name": "_alternativeFeeRecipient",
			  "type": "address"
			}
		  ],
		  "name": "registerProposerWithoutPayoutPoolRegistration",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "renounceOwnership",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			}
		  ],
		  "name": "reportProposer",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "contract PayoutPool",
			  "name": "_payoutPool",
			  "type": "address"
			},
			{
			  "internalType": "contract ReporterRegistry",
			  "name": "_reporterRegistry",
			  "type": "address"
			}
		  ],
		  "name": "setContracts",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_signatureSwapper",
			  "type": "address"
			},
			{
			  "internalType": "bool",
			  "name": "_enabled",
			  "type": "bool"
			}
		  ],
		  "name": "setSignatureSwapper",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "_newThreshold",
			  "type": "uint256"
			}
		  ],
		  "name": "setThreshold",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "newOwner",
			  "type": "address"
			}
		  ],
		  "name": "transferOwnership",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			},
			{
			  "internalType": "address",
			  "name": "_newPayoutRecipient",
			  "type": "address"
			}
		  ],
		  "name": "updatePayoutRecipient",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_blsKey",
			  "type": "bytes"
			},
			{
			  "internalType": "address",
			  "name": "_representative",
			  "type": "address"
			},
			{
			  "internalType": "address",
			  "name": "_payoutRecipient",
			  "type": "address"
			},
			{
			  "internalType": "bytes",
			  "name": "_blsSignature",
			  "type": "bytes"
			},
			{
			  "components": [
				{
				  "internalType": "uint8",
				  "name": "v",
				  "type": "uint8"
				},
				{
				  "internalType": "bytes32",
				  "name": "r",
				  "type": "bytes32"
				},
				{
				  "internalType": "bytes32",
				  "name": "s",
				  "type": "bytes32"
				}
			  ],
			  "internalType": "struct ProposerRegistry.SignatureECDSA",
			  "name": "_ecdsaSignature",
			  "type": "tuple"
			}
		  ],
		  "name": "validateRegistrationSignature",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		}
	  ]`
)
