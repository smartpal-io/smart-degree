# Smart Degree

[DECRIPTION]...

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

```
...
```

### Prerequisites

What things you need to install the software and how to install them

```
...
```

### Installing

A step by step series of examples that tell you have to get a development env running

```
npm init -y
npm install -E zeppelin-solidity
truffle compile
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

...

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

...
