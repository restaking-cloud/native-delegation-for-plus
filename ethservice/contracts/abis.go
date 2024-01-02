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
	K2_LENDING_CONTRACT_ABI = `[
		{
			"inputs": [],
			"stateMutability": "nonpayable",
			"type": "constructor"
		},
		{
			"inputs": [],
			"name": "ComeBackLater",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "ExceedMaxBorrowRatio",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "ExceedMaxSlashableAmountPerCorruption",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "ExceedMaxSlashableAmountPerLiveness",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "HasDebt",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "InvalidDebtPositionType",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "InvalidLength",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "MinDepositLimit",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "MustPaySlashedAmount",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NoDebt",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NoElements",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NoTerminationWithRST",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NodeOperatorAlreadyRegistered",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NodeOperatorBLSKeyNotPermitted",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NodeOperatorInvalid",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NodeOperatorInvalidRepresentative",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NodeOperatorInvalidStatus",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NodeOperatorKicked",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NodeOperatorNotRegistered",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NotAbleToLiquidate",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NotAllowedToDecreaseInterestRate",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NotEnoughAssumedLiquidity",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "NotEnoughOutstandingInterestToSlash",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "TooSmall",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "Unauthorized",
			"type": "error"
		},
		{
			"inputs": [],
			"name": "ZeroAddress",
			"type": "error"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "owner",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "spender",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "value",
					"type": "uint256"
				}
			],
			"name": "Approval",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "borrower",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "designatedVerifier",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "interestPaid",
					"type": "uint256"
				}
			],
			"name": "Borrowed",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "borrower",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "string",
					"name": "designatedVerifierURI",
					"type": "string"
				}
			],
			"name": "DesignatedVerifierURIUpdated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "borrower",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "designatedVerifier",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "bool",
					"name": "resetDuration",
					"type": "bool"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "interestPaid",
					"type": "uint256"
				}
			],
			"name": "IncreasedDebt",
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
					"internalType": "address",
					"name": "depositor",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "KETHClaimed",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "depositor",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "KETHDeposited",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "depositor",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "KETHWithdrawn",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "borrower",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "liquidator",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "liquidationAmount",
					"type": "uint256"
				}
			],
			"name": "Liquidated",
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
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "NodeOperatorClaimed",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "operator",
					"type": "address"
				},
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
			"name": "NodeOperatorDeposited",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "operator",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "bytes",
					"name": "blsPublicKey",
					"type": "bytes"
				},
				{
					"indexed": false,
					"internalType": "bool",
					"name": "kicked",
					"type": "bool"
				}
			],
			"name": "NodeOperatorWithdrawn",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "borrower",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "topupAmount",
					"type": "uint256"
				}
			],
			"name": "RepaidSlashAmount",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				},
				{
					"indexed": true,
					"internalType": "bytes32",
					"name": "previousAdminRole",
					"type": "bytes32"
				},
				{
					"indexed": true,
					"internalType": "bytes32",
					"name": "newAdminRole",
					"type": "bytes32"
				}
			],
			"name": "RoleAdminChanged",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "account",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "sender",
					"type": "address"
				}
			],
			"name": "RoleGranted",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "account",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "sender",
					"type": "address"
				}
			],
			"name": "RoleRevoked",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "borrower",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "hook",
					"type": "address"
				}
			],
			"name": "SBPHookUpdated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"internalType": "address",
					"name": "debtor",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "address",
					"name": "recipient",
					"type": "address"
				}
			],
			"name": "Slashed",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"internalType": "address",
					"name": "debtor",
					"type": "address"
				}
			],
			"name": "Terminated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "from",
					"type": "address"
				},
				{
					"indexed": true,
					"internalType": "address",
					"name": "to",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "value",
					"type": "uint256"
				}
			],
			"name": "Transfer",
			"type": "event"
		},
		{
			"inputs": [],
			"name": "DEFAULT_ADMIN_ROLE",
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
					"internalType": "address",
					"name": "owner",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "spender",
					"type": "address"
				}
			],
			"name": "allowance",
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
					"name": "spender",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "approve",
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
			"inputs": [],
			"name": "assumedLiquidity",
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
					"name": "account",
					"type": "address"
				}
			],
			"name": "balanceOf",
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
					"internalType": "struct IProposerRegistry.SignatureECDSA[]",
					"name": "_ecdsaSignatures",
					"type": "tuple[]"
				}
			],
			"name": "batchNodeOperatorDeposit",
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
			"name": "blsPublicKeyToKicked",
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
					"name": "",
					"type": "bytes"
				}
			],
			"name": "blsPublicKeyToLendPosition",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "cumulativeKethPerShareLU_RAY",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "kethEarned",
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
			"name": "blsPublicKeyToNodeOperator",
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
					"name": "",
					"type": "bytes"
				}
			],
			"name": "blsPublicKeyToPayoutRecipient",
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
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "debtPositionType",
					"type": "uint8"
				},
				{
					"internalType": "address",
					"name": "designatedVerifier",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"components": [
						{
							"internalType": "bool",
							"name": "mintRST",
							"type": "bool"
						},
						{
							"internalType": "uint256",
							"name": "initialSupply",
							"type": "uint256"
						},
						{
							"internalType": "address",
							"name": "recipientOfRemainingSupply",
							"type": "address"
						},
						{
							"internalType": "uint256",
							"name": "percentageContributionToIncentives",
							"type": "uint256"
						},
						{
							"internalType": "string",
							"name": "symbol",
							"type": "string"
						}
					],
					"internalType": "struct IK2Lending.RSTConfigParams",
					"name": "rstConfigParams",
					"type": "tuple"
				}
			],
			"name": "borrow",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "borrowDuration",
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
					"name": "debtor",
					"type": "address"
				},
				{
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "debtPositionType",
					"type": "uint8"
				},
				{
					"internalType": "address",
					"name": "designatedVerifier",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"components": [
						{
							"internalType": "bool",
							"name": "mintRST",
							"type": "bool"
						},
						{
							"internalType": "uint256",
							"name": "initialSupply",
							"type": "uint256"
						},
						{
							"internalType": "address",
							"name": "recipientOfRemainingSupply",
							"type": "address"
						},
						{
							"internalType": "uint256",
							"name": "percentageContributionToIncentives",
							"type": "uint256"
						},
						{
							"internalType": "string",
							"name": "symbol",
							"type": "string"
						}
					],
					"internalType": "struct IK2Lending.RSTConfigParams",
					"name": "rstConfigParams",
					"type": "tuple"
				}
			],
			"name": "borrowFor",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "borrowedLiquidity",
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
					"name": "lender",
					"type": "address"
				}
			],
			"name": "claimKETH",
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
			"inputs": [
				{
					"internalType": "address",
					"name": "lender",
					"type": "address"
				}
			],
			"name": "claimableKETH",
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
			"inputs": [
				{
					"internalType": "bytes",
					"name": "_blsPubkey",
					"type": "bytes"
				}
			],
			"name": "claimableKETHForNodeOperator",
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
			"name": "daoAddress",
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
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"name": "debtPositions",
			"outputs": [
				{
					"internalType": "address",
					"name": "debtor",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "hook",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "designatedVerifier",
					"type": "address"
				},
				{
					"internalType": "string",
					"name": "designatedVerifierURI",
					"type": "string"
				},
				{
					"internalType": "uint256",
					"name": "principalAmount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "interestPerSec_RAY",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "endTimestamp",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "slashAmount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
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
			"name": "debtors",
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
			"name": "decimals",
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
					"internalType": "address",
					"name": "spender",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "subtractedValue",
					"type": "uint256"
				}
			],
			"name": "decreaseAllowance",
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
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "deposit",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "address",
					"name": "recipient",
					"type": "address"
				}
			],
			"name": "depositFor",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "debtor",
					"type": "address"
				}
			],
			"name": "getDebtor",
			"outputs": [
				{
					"components": [
						{
							"internalType": "address",
							"name": "debtor",
							"type": "address"
						},
						{
							"internalType": "address",
							"name": "hook",
							"type": "address"
						},
						{
							"internalType": "address",
							"name": "designatedVerifier",
							"type": "address"
						},
						{
							"internalType": "string",
							"name": "designatedVerifierURI",
							"type": "string"
						},
						{
							"internalType": "uint256",
							"name": "principalAmount",
							"type": "uint256"
						},
						{
							"internalType": "uint256",
							"name": "interestPerSec_RAY",
							"type": "uint256"
						},
						{
							"internalType": "uint256",
							"name": "endTimestamp",
							"type": "uint256"
						},
						{
							"internalType": "uint256",
							"name": "slashAmount",
							"type": "uint256"
						},
						{
							"internalType": "uint256",
							"name": "maxSlashableAmountPerLiveness",
							"type": "uint256"
						},
						{
							"internalType": "uint256",
							"name": "maxSlashableAmountPerCorruption",
							"type": "uint256"
						}
					],
					"internalType": "struct IK2Lending.DebtPosition",
					"name": "",
					"type": "tuple"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "debtPositionType",
					"type": "uint8"
				},
				{
					"internalType": "uint256",
					"name": "newBorrowAmount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "currentBorrowAmount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "duration",
					"type": "uint256"
				}
			],
			"name": "getExpectedInterest",
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
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "debtPositionType",
					"type": "uint8"
				},
				{
					"internalType": "uint256",
					"name": "interestAmount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "currentBorrowAmount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "duration",
					"type": "uint256"
				}
			],
			"name": "getMaxBorrowableAmount",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "amount",
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
					"name": "debtor",
					"type": "address"
				}
			],
			"name": "getOutstandingInterest",
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
					"name": "debtor",
					"type": "address"
				}
			],
			"name": "getRemainingDuration",
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
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				}
			],
			"name": "getRoleAdmin",
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
			"name": "getTotalBorrowableAmount",
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
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				},
				{
					"internalType": "address",
					"name": "account",
					"type": "address"
				}
			],
			"name": "grantRole",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				},
				{
					"internalType": "address",
					"name": "account",
					"type": "address"
				}
			],
			"name": "hasRole",
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
					"name": "spender",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "addedValue",
					"type": "uint256"
				}
			],
			"name": "increaseAllowance",
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
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "debtPositionType",
					"type": "uint8"
				},
				{
					"internalType": "address",
					"name": "designatedVerifier",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"internalType": "bool",
					"name": "resetDuration",
					"type": "bool"
				}
			],
			"name": "increaseDebt",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "debtor",
					"type": "address"
				},
				{
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "debtPositionType",
					"type": "uint8"
				},
				{
					"internalType": "address",
					"name": "designatedVerifier",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerLiveness",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "maxSlashableAmountPerCorruption",
					"type": "uint256"
				},
				{
					"internalType": "bool",
					"name": "resetDuration",
					"type": "bool"
				}
			],
			"name": "increaseDebtFor",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "_keth",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_configurator",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_daoAddress",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_k2Incentives",
					"type": "address"
				},
				{
					"internalType": "string",
					"name": "_name",
					"type": "string"
				},
				{
					"internalType": "string",
					"name": "_symbol",
					"type": "string"
				},
				{
					"internalType": "address",
					"name": "_proposerRegistry",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "_nodeOperatorInclusionList",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "_maxBorrowRatio_RAY",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "_maxSlashableRatio_RAY",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "_terminatePenalty_RAY",
					"type": "uint256"
				},
				{
					"internalType": "contract RSTModule",
					"name": "_rstModule",
					"type": "address"
				}
			],
			"name": "initialize",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "interestPerSecLU_RAY",
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
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "",
					"type": "uint8"
				}
			],
			"name": "interestRateModelByType",
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
			"name": "k2Incentives",
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
			"name": "keth",
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
					"name": "",
					"type": "address"
				}
			],
			"name": "lenders",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "cumulativeKethPerShareLU_RAY",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "kethEarned",
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
					"name": "debtor",
					"type": "address"
				}
			],
			"name": "liquidate",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "maxBorrowRatio_RAY",
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
			"name": "maxSlashableRatio_RAY",
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
			"name": "minDepositLimit",
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
			"name": "minLockUpPeriod",
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
			"name": "minTransferLimit",
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
			"name": "name",
			"outputs": [
				{
					"internalType": "string",
					"name": "",
					"type": "string"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes[]",
					"name": "_blsPublicKeys",
					"type": "bytes[]"
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
					"internalType": "struct IProposerRegistry.SignatureECDSA",
					"name": "_ecdsaSignature",
					"type": "tuple"
				}
			],
			"name": "nodeOperatorDeposit",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "nodeOperatorInclusionList",
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
					"name": "_reporter",
					"type": "address"
				},
				{
					"internalType": "bytes",
					"name": "_blsPubkey",
					"type": "bytes"
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
					"internalType": "address",
					"name": "_nodeOperator",
					"type": "address"
				},
				{
					"internalType": "bytes",
					"name": "_blsPubkey",
					"type": "bytes"
				}
			],
			"name": "nodeOperatorWithdraw",
			"outputs": [],
			"stateMutability": "nonpayable",
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
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				},
				{
					"internalType": "address",
					"name": "account",
					"type": "address"
				}
			],
			"name": "renounceRole",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes32",
					"name": "role",
					"type": "bytes32"
				},
				{
					"internalType": "address",
					"name": "account",
					"type": "address"
				}
			],
			"name": "revokeRole",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "rstModule",
			"outputs": [
				{
					"internalType": "contract RSTModule",
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
					"name": "newConfigurator",
					"type": "address"
				}
			],
			"name": "setConfigurator",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "string",
					"name": "designatedVerifierURI",
					"type": "string"
				}
			],
			"name": "setDesignatedVerifierURI",
			"outputs": [],
			"stateMutability": "nonpayable",
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
			"name": "setHookAsDebtorForSBP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "enum IK2Lending.DebtPositionType",
					"name": "_type",
					"type": "uint8"
				},
				{
					"internalType": "address",
					"name": "_newInterestRateModel",
					"type": "address"
				}
			],
			"name": "setInterestRateModel",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "newMinDepositLimit",
					"type": "uint256"
				}
			],
			"name": "setMinDepositLimit",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "newMinLockUpPeriod",
					"type": "uint256"
				}
			],
			"name": "setMinLockupPeriod",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "newMinTransferLimit",
					"type": "uint256"
				}
			],
			"name": "setMinTransferLimit",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "_nodeOperatorInclusionList",
					"type": "address"
				}
			],
			"name": "setNodeOperatorInclusionList",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "enum ReporterRegistry.SlashType",
					"name": "slashType",
					"type": "uint8"
				},
				{
					"internalType": "address",
					"name": "debtor",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "address",
					"name": "recipient",
					"type": "address"
				}
			],
			"name": "slash",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "slashedLiquidity",
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
					"internalType": "bytes4",
					"name": "interfaceId",
					"type": "bytes4"
				}
			],
			"name": "supportsInterface",
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
			"name": "symbol",
			"outputs": [
				{
					"internalType": "string",
					"name": "",
					"type": "string"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "terminate",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "terminatePenalty_RAY",
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
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "topUpAndTerminate",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "debtor",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "topUpSlashAmount",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [],
			"name": "totalSupply",
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
					"name": "to",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "transfer",
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
					"internalType": "address",
					"name": "from",
					"type": "address"
				},
				{
					"internalType": "address",
					"name": "to",
					"type": "address"
				},
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "transferFrom",
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
			"inputs": [],
			"name": "updateInterest",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "lender",
					"type": "address"
				}
			],
			"name": "updateLenderPosition",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "bytes",
					"name": "blsPublicKey",
					"type": "bytes"
				}
			],
			"name": "updateNodeOperatorPosition",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "_maxBorrowRatio_RAY",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "_maxSlashableRatio_RAY",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "_terminatePenalty_RAY",
					"type": "uint256"
				}
			],
			"name": "updateRatioSettings",
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
			"name": "userLastInteractedTimestamp",
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
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "bool",
					"name": "claim",
					"type": "bool"
				}
			],
			"name": "withdraw",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`
	K2_NODE_OPERATOR_CONTRACT_ABI  = `[
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
