# MultiSendChain (MSC)

MultiSendChain (MSC) is a simple and secure tool for sending the same amount of cryptocurrency to multiple addresses in one operation. The application supports multiple blockchain networks and provides a guided, user-friendly command-line interface. MSC is designed to speed up sending crypto to multiple addresses simultaneously while maintaining safety and confirmation steps.

https://github.com/user-attachments/assets/a11b5772-d942-4e1d-88fc-c69691ad287a

## Available Chains

The project currently includes configuration for the following networks:

- Sepolia (Ethereum testnet)
- Ethereum Mainnet
- BNB Smart Chain Mainnet
- ETH Base
- ETh Arbitrum
- ETH Optimism
- Polygon
- Etc.

Other chains will be added in the future. You can also add your own chain configurations by creating RPC configuration files in the chains/ directory.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/<your-username>/msc.git
   cd msc
   ```

2. Prerequisites:
   - A recent version of Go (if you want to run from source).
   - A funded wallet with the private key (see Security & Wallet below).

3. Prepare a recipients file named `recipients.txt` in the project root. Each line should contain a single recipient address (for example, `0x...`).

## How to Run (user-friendly)

- Quick Start (recommended for most users):
  1. Create a `.env` file in the project root containing a single entry for your wallet private key:
     - `PRIVATE_KEY=your_private_key_here`
  2. Open a terminal, navigate to the project root, and run the application:
     - If you have Go installed: `go run ./cmd/msc`
     - Alternatively, build a binary first with `go build ./cmd/msc` and run the produced executable.

- Guided flow:
  - When the application starts it will prompt you to select a blockchain network from the available options.
  - You will be asked to confirm or change the recipients file path (default `recipients.txt`).
  - Enter the amount to send to each address and confirm the summary before the tool proceeds.

## Important Safety Notes

- Always double-check recipient addresses in `recipients.txt` before running any transfers. Mistakes cannot be undone.
- Test first on Sepolia (Ethereum testnet) with small amounts before using mainnets.
- Ensure your wallet has enough funds to cover both the total transfer amounts and network fees (gas).
- Keep your private key secure. Do not share it or store it in public places.

## Troubleshooting

- If the application cannot read your wallet nor connect to the network, check that your `.env` file is present and that `PRIVATE_KEY` is set correctly.
- Ensure your `recipients.txt` file uses one address per line and that addresses are valid for the selected network.

## Example files

Below are two example files you can use as starting points.

### recipients.txt (template)

Edit the `recipients.txt` file in the project root with one recipient address per line. Example:

```
0x1111111111111111111111111111111111111111
0x2222222222222222222222222222222222222222
0x3333333333333333333333333333333333333333
```

### .env.example

Copy `.env.example` to `.env` and replace the placeholder with your private key:

```
PRIVATE_KEY=0xYOUR_PRIVATE_KEY_HERE
```
