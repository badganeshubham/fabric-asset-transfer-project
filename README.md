# Fabric Asset Transfer Project

This is a basic blockchain application built using *Hyperledger Fabric* and *JavaScript*. The project demonstrates the fundamental concept of asset transfer on a permissioned blockchain network. It includes a JavaScript client application, smart contracts (chaincode), and optionally an API interface.

---

##  Project Description

This project showcases a sample implementation of asset management using Hyperledger Fabric. It allows users to perform blockchain operations like creating, reading, updating, and deleting assets on the ledger. It is ideal for learning how to:

- Interact with Hyperledger Fabric using the Node.js SDK.
- Create and invoke chaincode.
- Register and authenticate users via Fabric CA.
- Explore transaction flows in a permissioned blockchain.

---

## 📁 Folder Structure
fabric-asset-transfer-project/
├── application-javascript/ # Node.js client app
├── chaincode/ # Chaincode (smart contracts)
├── api/ # Optional REST API interface
└── README.md # Project overview
---

## 🛠 Technologies Used

- Hyperledger Fabric
- JavaScript
- Node.js
- Fabric Node SDK
- Docker
- Git & GitHub

---

##  Prerequisites

- Node.js & npm
- Docker & Docker Compose
- Git
- Fabric samples and binaries installed

---

## Setup & Usage

### 1. Install Dependencies

Navigate to the client app:

```bash
cd application-javascript
npm install
