# BPHR: Blockchain-based Personal Health Records

**BPHR** is a blockchain-based solution that enables users to securely store and manage their personal health records on a distributed ledger. The application leverages Hyperledger Fabric for the blockchain infrastructure and provides a user-friendly interface to interact with the stored records.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)

## Features

- **User Management**: Register new users and manage their credentials.
- **Health Records**: Store, update, and retrieve personal health records.
- **Transactions**: Keep track of every change made to the records.
- **Security**: Built upon the secure foundations of Hyperledger Fabric.

## Prerequisites

- Docker & Docker Compose
- Node.js (v18.18.0 or later)
- Go (For chaincode development)
- Hyperledger Fabric binaries & samples
- Access to the terminal or command prompt

## Installation

1. **Clone the Repository**:
   bash
   git clone <repository-url>
   cd bphr
   

2. **Setup the Network**:
   Follow the setup instructions in the Hyperledger Fabric documentation, ensuring that you have your Fabric network up and running.

3. **Dependencies**:
   All required scripts will automatically handle the dependencies. Just ensure you have the prerequisites installed.

## Usage

1. **Start the Application**:
   Use the provided `run.sh` script to bootstrap the entire process.
   bash
   cd scripts
   chmod +x run.sh
   ./run.sh
   
   This script will handle:
   - Packaging the chaincode
   - Installing the chaincode on the peers
   - Approving the chaincode
   - Committing the chaincode to the channel
   - Installing application dependencies
   - Starting the application server
   - Testing the application functionalities
   - Cleaning up the network

2. **Web Interface**:
   After starting the application using `run.sh`, open your web browser and navigate to `http://localhost:3000` to access the application's interface.

3. **Shutdown**:
   The `run.sh` script will automatically shut down the network and the application at the end. If you need to manually shut down, you can use the `networkDown.sh` script.

## Support

For any questions or issues, please refer to the FAQ section or open an issue in this repository.

## Contributing

If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

---