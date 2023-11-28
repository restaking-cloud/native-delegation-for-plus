# K2 Native Delegation Module for MEV Plus
Plus module for performing K2 native delegation registrations and restaking cloud participation.

The K2 native delegation module for MEV Plus gives node runners the ability to automate their Proposer Registry and K2 validator native delegation registrations without the need for external parties or exposing validator keys outside of the node. Additionally their keys can easily be changed on their node as and when they wish and the module would facilitate this automatic registration without the need for extra steps or further action from the user.
Exiting K2, Withdrawals and Rewards claiming can also be programmatically configured by the user, turning contract interactions for validators to simple and better managed calls through MEV Plus.

# Tutorial

To follow a tutorial for the installation of the K2 native delegation module click [here](https://docs.restaking.cloud/middleware/install-native-delegation).

## Prerequisites
- [MEV Plus with K2 Native Delegation](https://github.com/pon-network/mev-plus) module plugged in.

- Eth1 Wallet: You should have an Ethereum wallet with sufficient funds to pay for gas fees. This wallet is used to pay for gas fees when interacting with the Proposer Registry and K2 contracts.

- MEV Plus connected to node by Builder API: The module requires the node to be connected to MEV Plus through the Builder API. This is required for the module to receive the presigned messages from the node and register validators on-chain.

## Supported Networks

- Mainnet
- Goerli

Please note that currently only `Goerli` supports native delegation at this time. However, `Mainnet` supports native delegation `pre-registration` due to the proposer registry deployments already existing on mainnet.

PoN deployment information can be found here:
```
https://github.com/pon-network/contract-deployments
```

## Configuration
You can configure the K2 Native Delegation module using the following flags:

### Required Flags

- `k2.eth1-private-key`: The private key of the proposer(s) representative wallet. This key is essential for signing registration messages and interacting with all the smart contracts.
NOTE: This key is the only wallet key that would be granted permission to manage validators on-chain and act as their representative. Keep this wallet safe.

- `k2.beacon-node-url`: The URL of the beacon node. This URL is required for syncing with the Ethereum Consensus Layer.

- `k2.execution-node-url`: The URL of the execution node to connect to for on-chain execution.

### Optional Flags

- `k2.web3-signer-url`: The module supports the use of a [Web3Signer](https://docs.web3signer.consensys.net/) to sign custom registration messages with a modified payout recipient address from the one configured on the node. This flag is optional and can be used to configure the Web3Signer URL. The validator keys to which their registration messages wish to be signed should be configured on the Web3Signer.

- `k2.payout-recipient`: The address of an alternative globally configured payout recipient. This address will be used for all validators if specified. If not specified, the payout recipient address configured on the node for each validator key will be used. To use this flag, the `k2.web3-signer-url` flag must also be specified in order to sign the registration messages with the alternative payout recipient address.


## How It Works

Validator Registration: The K2-Native-Delegation module enables node runners to register as validators on-chain by securely registering their BLS keys with the Proposer Registry contract. The module utilises the presigned messages broadcasted by the node through the Builder API of the consensus client to register validators on-chain.

Signature Swapper: The module uses the signature swapper to generate and manage ECDSA signatures as proof of ownership of the BLS keys. This ensures the security of the registration process and avoids spoofing.

K2 Contract Interaction: The module communicates with the K2 contract on the blockchain to register validators and participate in restaking cloud activities. A requirement of K2 native delegation is registration with Proof of Neutrality single sign on proposer registry, which is also facilitated by this module.

Payout Management: Node runners can configure the payout recipient address. If not specified, payouts go to the fee recipients configured in your consensus client for each validator key.

## Benefits for Node Runners
**Security**: Validators can securely sign registration messages and interact with the K2 contract without revealing their private keys to external parties.

**Local Machine Control**: Node runners have full control over the registration process on their local machines, reducing reliance on third-party services or DApps.

**Restaking Cloud Participation**: This module enables validators to participate in Restaking Cloud; k2 activities, potentially increasing their rewards in a secure and controlled manner.

**Custom Payouts**: Validators can configure payout recipient addresses, giving them flexibility in managing their rewards, or use their set fee recipients on the node, to receive rewards through both block proposer rewards and rewards through restaking (exposure to multiple income streams to your configured node fee recipient).

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

## License
[MIT](LICENSE.md)
