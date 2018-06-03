# Smart Degree

Smart degree aims to provide a decentralized solution to manage degree. The main idea is for a school, or any entity managing diplomas, to store diplomas validation proof on a blockchain system. Smart Degree currently uses Ethereum as backing blockchain. Once a degree is committed on the blockchain, anyone is able to verify if a degree has been validated by the entity. For example, an employer can verify qualification of a candidate during a job interview.

## Getting Started

Smart Degree repository contains :

* truffle project for ethereum smart contract management
* decentralized web application to interact with the smart contract
* a backend written in Golang to interact with the smart contract

### Prerequisites

Node JS installed

### Installing

Download dependencies

```
npm install
```

### Compiling the smart contract

```
truffle compile
```

### Deploying the smart contract

#### Using local RPC

```
# assuming that a local rpc is running on port 8545
truffle migrate --network development # --reset optional
```

#### Using docker

```
docker-compose up -d
truffle migrate --network docker --reset
```

## Running the tests

```
truffle test
```

## Running the decentralized web application


```
npm run dev
```

Then visit :  http://localhost:8080/


## Running android test application

The android application call a ethereum smart contract to verify degree's hash 
see this repository : https://github.com/SmartContractCompanion/smart-degree-dapp-android

## Authors

Adbelhamid Bakhta

Karim Taam

Ludovic Maréchal

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

