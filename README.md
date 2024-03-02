# K2 Native Delegation Module for MEV Plus

Plus module for performing K2 native delegation registrations and restaking cloud participation.

The K2 native delegation module for MEV Plus gives node runners the ability to automate their Proposer Registry and K2 validator native delegation registrations without the need for external parties or exposing validator keys outside of the node. Additionally their keys can easily be changed on their node as and when they wish and the module would facilitate this automatic registration without the need for extra steps or further action from the user.
Exiting K2 and Rewards claiming can also be programmatically configured by the user, turning contract interactions for validators to simple and better managed calls through MEV Plus.

## Tutorial

To follow a tutorial for the installation of the K2 native delegation module click [here](https://docs.restaking.cloud/middleware/install-native-delegation).

## Prerequisites

- [MEV Plus with K2 Native Delegation](https://github.com/pon-network/mev-plus) module plugged in.

- Eth1 Wallet: You should have an Ethereum wallet with sufficient funds to pay for gas fees. This wallet is used to pay for gas fees when interacting with the Proposer Registry and K2 contracts.

- MEV Plus connected to node by Builder API: The module requires the node to be connected to MEV Plus through the Builder API. This is required for the module to receive the presigned messages from the node and register validators on-chain.

## Supported Networks

- Mainnet
- Goerli
- Holesky

PoN deployment information can be found here:

```
https://github.com/pon-network/contract-deployments
```

## Configuration

You can configure the K2 Native Delegation module using the following flags:

### Required Flags

- `k2.eth1-private-key`: The private key of the proposer(s) representative wallet. This key is essential for signing registration messages and interacting with all the smart contracts.
NOTE: This key is the only wallet key that would be granted permission to manage validators on-chain and act as their representative. Keep this wallet safe.
This module also supports the use of multiple representative wallets, in which case provide a comma separated list of private keys. [eg. `k2.eth1-private-key 1234567890abcdef1234567890abcdef12345f,1234567890abcdef1234567890abcdef12345f`
]. The keys should be in order of priority, with the first key being the primary key. The module will use the first key to run contract calls and reward claim transactions. The additional keys are used to natively delegate or exit validators of varied payout recipients per key.

- `k2.beacon-node-url`: The URL of the beacon node. This URL is required for syncing with the Ethereum Consensus Layer.

- `k2.execution-node-url`: The URL of the execution node to connect to for on-chain execution.

### Optional Flags

- `k2.registration-only`: This flag is used to register validators on-chain in the Proposer Registry without natively delegating them to the K2 contract pool.

- `k2.web3-signer-url`: The module supports the use of a [Web3Signer](https://docs.web3signer.consensys.net/) to sign custom registration messages with a modified payout recipient address from the one configured on the node. This flag is optional and can be used to configure the Web3Signer URL. The validator keys to which their registration messages wish to be signed should be configured on the Web3Signer.

- `k2.payout-recipient`: The address of an alternative globally configured payout recipient. This address will be used for all validators if specified. If not specified, the payout recipient address configured on the node for each validator key will be used. To use this flag, the `k2.web3-signer-url` flag must also be specified in order to sign the registration messages with the alternative payout recipient address.

- `k2.max-gas-price`: The maximum gas price (denominated in WEI) to be used for on-chain transactions. This flag is optional and defaults to using the current netowrk gas price for execution. If set, any registration in which the gas price exceeds the maximum gas price, the registration/delegation will be skipped for that epoch.

- `k2.listen-address`: The address on which the module will listen for incoming requests. This flag is optional and defaults to `localhost:10000` if not specified. The API specifications can be found [here](#api).

- `k2.claim-threshold`: The threshold for claiming rewards from the K2 contract. This flag is optional and defaults to 0.0 KETH if not specified (claims any available rewards). If the rewards for a validator exceed the threshold, the rewards will be claimed from the K2 contract upon any request to the API.

- `k2.k2-lending-contract-address`: The address of the K2 lending contract you wish to provide to override the default contract address for a supported network, or to provide a contract address for an unsupported network.

- `k2.k2-node-operator-contract-address`: The address of the K2 node operator contract you wish to provide to override the default contract address for a supported network, or to provide a contract address for an unsupported network.

- `k2.proposer-registry-contract-address`: The address of the proposer registry contract you wish to provide to override the default contract address for a supported network, or to provide a contract address for an unsupported network. Especially if in registration-only mode. If not in registration-only mode, the module will obtain the contract address from the K2 contracts and does not need this flag.

- `k2.signature-swapper-url`: The URL of the signature swapper service. This flag is optional and defaults to the network-specific signature swapper URL if not specified (overridden), or can be used to provide a custom signature swapper URL for unsupported networks. The signature swapper is used to generate and manage ECDSA signatures as proof of ownership of the BLS keys.

- `k2.balance-verifier-url`: The URL of the balance verifier service. This flag is optional and defaults to the network-specific balance verifier URL if not specified (overridden), or can be used to provide a custom balance verifier URL for unsupported networks. The balance verifier is used to verify the effective balance of the proposer wallet for balance report to the contracts for reward claiming or exiting the protocol.

- `k2.strict-inclusion-list`: This flag is used to specify a list of validator public keys to solely process. The flag accepts a filepath to a JSON file containing the list of validator public keys/fee recipients to strictly process. The file is continuously monitored by the software and would pick up any changes immediately, allowing you to manage your registrations without restarting MEV Plus. The JSON file should be in the following format:

```json
[
    {
        "publicKey": string,
        "allowProposerRegistration": bool,
        "allowNativeDelegation": bool
    },
    ...,
    {
        "feeRecipientAddress": string,
        "allowProposerRegistration": bool,
        "allowNativeDelegation": bool
    }
]
```

You may either strictly include by a specific validator BLS Public Key in the entry list or a group of validators that share a Fee Recipient Address on your node.

eg. `k2.strict-inclusion-list ./inclusion-list.json`

```json inclusion-list.json
[
  {
    "publicKey": "0x93e2de67f75817c101c637b16efc4ba1de8374ed563a4cdcf2d6cc5ea6c1de4ab5abcefdb3bd2baa96a1a2ddb1847d08",
    "allowProposerRegistration": false,
    "allowNativeDelegation": true
  },
  {
    "feeRecipientAddress": "0x22A3864baaE65a9e8E5C163F80F850ADFe40Ed90",
    "allowProposerRegistration": true,
    "allowNativeDelegation": true
  },
]
```

**NOTE**: Either or both the `allowProposerRegistration` and `allowNativeDelegation` fields must be set as if not specified, and defaults to `false`. If both feilds are `false` due to not specified or being set as false, then effectively this entry should not be included in the inclusion list as it means you do not intend to include it in any process (proposer registration nor native delgation).
If `allowProposerRegistration` is set to `false`, the validator will not be registered on-chain in the Proposer Registry. And if this validator is intended to be natively delegated, and is not already in the Proposer Registry, the registration will fail, as native delegation requires registration in the Proposer Registry. However if this validator / group of validators are already registered in the proposer registry, their processing of netive delegation if set to true would be unaffected.

- `k2.exclusion-list`: This flag is used to specify a list of validator public keys to exclude from either on-chain registration or native delegation. The flag accepts a filepath to a JSON file containing the list of validator public keys/fee recipients to exclude. The file is continuously monitored by the software and would pick up any changes immediately, allowing you to manage your registrations without restarting MEV Plus. The JSON file should be in the following format:

```json
[
    {
        "publicKey": string,
        "allowProposerRegistration": bool,
        "allowNativeDelegation": bool
    },
    ...,
    {
        "feeRecipientAddress": string,
        "allowProposerRegistration": bool,
        "allowNativeDelegation": bool
    }
]
```

You may either exclude by a specific validator BLS Public Key in the entry list or a group of validators that share a Fee Recipient Address on your node.

eg. `k2.exclusion-list ./exclusion-list.json`

```json exclusion-list.json
[
  {
    "publicKey": "0x93e2de67f75817c101c637b16efc4ba1de8374ed563a4cdcf2d6cc5ea6c1de4ab5abcefdb3bd2baa96a1a2ddb1847d08",
    "allowProposerRegistration": false,
    "allowNativeDelegation": true
  },
  {
    "feeRecipientAddress": "0x22A3864baaE65a9e8E5C163F80F850ADFe40Ed90",
    "allowProposerRegistration": true,
    "allowNativeDelegation": true
  },
]
```

**NOTE**: The `allowProposerRegistration` and `allowNativeDelegation` fields are optional and default to `false` if not specified and thus completely excludes the entry from all processing during proposer registration and/or native delegation. You cannot set both fields to `true` as that essentially means you do not intend to exclude the entry from any process.
If `allowProposerRegistration` is set to `false`, the validator will not be registered on-chain in the Proposer Registry. And if this validator is intended to be natively delegated, and is not already in the Proposer Registry, the registration will fail, as native delegation requires registration in the Proposer Registry.

- `k2.representative-mapping`: This flag is used to optionally specify a mapping of representative wallets that should be used to process validators to specific validators by BLS Key or specific payout/feeRecipient addresses. The flag accepts a filepath to a JSON file containing the mapping of representative addresses designated to specific validators or that would handle validators that pay to different k2 fee/payout recipients (any ECDSA address). The file is continuously monitored by the software and would pick up any changes immediately, allowing you to manage your registrations without restarting MEV Plus. The JSON file should be in the following format:

```json
[
    {
        "representativeAddress": string,
        "feeRecipientAddress": string
    },
    ...,
    {
        "representativeAddress": string,
        "publicKey": string
    }
]
```

eg. `k2.representative-mapping ./representative-mapping.json`

```json representative-mapping.json
[
  {
    "representativeAddress": "0x22A3864baaE65a9e8E5C163F80F850ADFe40Ed90",
    "feeRecipientAddress": "0x22A3864baaE65a9e8E5C163F80F850ADFe40Ed90"
  },
  {
    "representativeAddress": "0x22A3864baaE65a9e8E5C163F80F850ADFe40Ed90",
    "publicKey": "0x93e2de67f75817c101c637b16efc4ba1de8374ed563a4cdcf2d6cc5ea6c1de4ab5abcefdb3bd2baa96a1a2ddb1847d08"
  }
]
```

**NOTE**: Cannot provide more than one representative-feeRecipient pair with the same representative. Cannot provide more than one representative-PublicKey pair. Ensure that the representative addresses are the wallets available in the configured `k2.eth1-private-key` flag. This file is optional and is used to strictly inform the module to use the representative address to process specific validators or set of validators with a common fee recipient address on the node. If the representative address is not found in the `k2.eth1-private-key` flag, the module will not process the validators to the specified payout recipient address. If the node registration has validators and/or a validators with a common fee recipient not strictly specified in this file, the module would use the next available representative address in the `k2.eth1-private-key` flag to process the registration if possible.

- `k2.logger-level`: The log level for the K2 Native Delegation module. This flag is optional and defaults to `info` if not specified. The available log levels are `debug`, `info`, `warn`, `error`, and `fatal`.

## How It Works

Validator Registration: The K2-Native-Delegation module enables node runners to register as validators on-chain by securely registering their BLS keys with the Proposer Registry contract. The module utilises the presigned messages broadcasted by the node through the Builder API of the consensus client to register validators on-chain.

Signature Swapper: The module uses the signature swapper to generate and manage ECDSA signatures as proof of ownership of the BLS keys. This ensures the security of the registration process and avoids spoofing.

Balance Verification: The module verifies the effective balance of the proposer wallet before registering validators on-chain. If the balance is insufficient (<32 ETH), the registration is skipped for that epoch. This verifiaction is also available as a remote designated verifier for each network that is used to balance report to the contracts for reward claiming or exiting the protocol.


Payout Management: Node runners can configure the payout recipient address. If not specified, payouts go to the fee recipients configured in your consensus client for each validator key.

## Benefits for Node Runners

**Security**: Validators can securely sign registration messages and interact with the K2 contract without revealing their private keys to external parties.

**Local Machine Control**: Node runners have full control over the registration, claiming and exiting processes on their local machines, reducing reliance on third-party services or DApps.

**Restaking Cloud Participation**: This module enables validators to participate in Restaking Cloud; k2 activities, potentially increasing their rewards in a secure and controlled manner.

**Custom Payouts**: Validators can configure payout recipient addresses, giving them flexibility in managing their rewards, or use their set fee recipients on the node, to receive rewards through both block proposer rewards and k2 rewards through restaking (exposure to multiple income streams to your configured node fee recipient or any ECDSA address set for payouts), including but not limited to rewards for natively delegating validators.

## Installation

The K2 Native Delegation module is a plugin for MEV Plus. To install the module, import the module into your MEV Plus `moduleList.go` file and instantiate the `NewK2Service()` and `NewCommand()` from [service.go](service.go) functions in the respective service and command lists variables within the [moduleList.go](https://github.com/pon-network/mev-plus/blob/main/moduleList/moduleList.go) file.

Build your modified MEV Plus binary and run it with the required flags.

For more information on building MEV Plus, check out
[How to build MEV Plus](https://github.com/pon-network/mev-plus#building-mev-plus)

## Usage

The module functionality can be enabled/disabled by running mev-plus with the k2 flags available.

```bash

# Enable K2 Native Delegation Module
mevPlus ............ -k2.eth1-private 1234567890abcdef1234567890abcdef12345f -k2.beacon-node-url http://localhost:5052 -k2.execution-node-url http://localhost:8545

```

## API

The module runs a REST API on `localhost` port `10000` by default. The API exposes the following important endpoints:

### POST `/eth/v1/exit`

This endpoint is used to exit the protocol. It accepts a JSON body with the BLS Public Key of the validator to exit.

```json
{
  "validator": string
}
```

### POST `/eth/v1/claim`

This endpoint is used to claim rewards from the K2 contract. It accepts a JSON body with a list of node operator representative addresss to check for rewards and claim.

```json
{
  "nodeOperators": [
    Representative Address {string},
    Representative Address {string},
    ...
  ]
}
```
*NOTE*: Payload is optional. If no payload is parsed `{}`, the module checks for rewards for the representative wallets configured under `k2.eth1-private-key` and claims any available rewards for those representatives (node operators).

### GET `/eth/v1/delegated-validators`

This endpoint is used to get the list of validators that are natively delegated to the K2 contract. It by default returns the list of all validators for the representative wallets configured under `k2.eth1-private-key` and their respective fee recipients. It optionally accepts a query parameter `representativeAddresses` to specify any representative wallets to check for their natively delegated validators. It also optionally accepts a query parameter `includeBalance` as (`true` string) to specify if the claimable rewards of representative node operators should be included in the response, as well as the effective balances of each node operator's delegated validator.

```
GET /eth/v1/delegated-validators?representativeAddresses=<representativeAddress>&includeBalance=true
```

You can pass multiple representative addresses as a comma-seperated string to `representativeAddresses` to get the list of validators for each representative address.

```
GET /eth/v1/delegated-validators?representativeAddresses=<representativeAddress1>,<representativeAddress2>,<representativeAddress3>&includeBalance=true
```

Response schema:
```json response schema
[
  {
    "representativeAddress": string,
    "claimableRewards": uint64, *only if includeBalance is true
    "delegatedValidators": [
      {
        "validatorPubKey": string,
        "representativeAddress": string,
        "effectiveBalance": uint64, (in wei) *only if includeBalance is true
      },
      ...
    ]
  },
  ...
]
```

### POST `/eth/v1/register`

This endpoint is used to register validators on-chain. It offers an alternative endpoint dedicated for signed registration messages to which can be processed immediately without having to wait for epoch changes from the node. It accepts a JSON body with a list of registration messages to register on-chain.

```json
[
  {
    "message": {
      "pubkey": string,
      "fee_recipient": string,
      "gas_limit": uint64,
      "timestamp": uint64,
    },
      "signature": string
    }
]
```



## License
[MIT](LICENSE.md)
