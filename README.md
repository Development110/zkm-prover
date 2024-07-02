# zkm-prover
A parallel proving service for [ZKM](https://github.com/zkMIPS/zkm).

1. Go to git repo

    https://github.com/zkMIPS/zkm-prover

    click fork to get the copy of code to your git repo

2. -- Install Rust
	curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

3. download go binary with below command

	wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz

4. change permission to root 
    sudo su

5.  install go

	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz

6.  Set Go path

	export PATH=$PATH:/usr/local/go/bin

7.  install Hardhat

	sudo apt update
	sudo apt install curl git
	curl -fsSL https://deb.nodesource.com/setup_22.x | sudo -E bash -
	sudo apt-get install -y nodejs

8. create iseven.go file unnder  /zkm-prover folder as per 
 
	https://docs.google.com/document/d/1uQRwePAjT24WKblA43Za3VAs_psI_8hLnXCBjHAR-0Q/edit

9.  Execute below command to install protobuf compiler

	apt-get install protobuf-compiler

10. 	execute the below command

PRIVATE_KEY=YOUR_PRIVATE_KEY SEG_SIZE=YOUR_SEG_SIZE CA_CERT_PATH=./tools/certs/ca.pem ELF_PATH=YOUR_ELF_PATH ARGS='PUBLIC_ARG PRIVATE_ARG' ENDPOINT=https://152.32.186.45:20002 RUST_LOG=info OUTPUT_DIR=/tmp/examples cargo run --example stage

