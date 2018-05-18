# Smart Degree

Smart degree aims to provide a decentralized solution to manage degree. The main idea is for a school, or any entity managing diplomas, to store diplomas validation on a blockchain system. Smart Degree currently uses Ethereum as backing blockchain. Once a degree is committed on the blockchain, anyone is able to verify if a degree has been validated by the entity. For example, an employer can verify qualification of a candidate during a job interview.

## Getting Started

Smart Degree repository contains :

* truffle project for ethereum smart contract management
* android application to interact with the smart contract
* decentralized web application to interact with the smart contract
* a backend written in Golang to interact with the smart contract

### Prerequisites

Node JS installed

### Installing

Download dependencies

```
npm install
```

Compile the smart contracts

```
truffle compile
```

Deploy the smart contracts

```
truffle migrate # --reset optional
```

## Running the tests

```
truffle test
```

## Running android test application

This android application call a ethereum smart contract to verify degree's hash 

```
truffle compile
truffle deploy
web3j truffle generate D:/github/smart-degree/build/contracts/SmartDegree.json -o D:/github/smart-degree/app-android/app/src/main/java/ -p com.degree.application.contract
gradle assemble -PETH_NETWORK_URL="http://[IP]:[PORT]/" -PETH_NETWORK_ID="5776"
```

## Authors

Adbelhamid Bakhta

Karim Taam

Ludovic Maréchal

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

