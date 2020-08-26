# Hello World Example

This folder contains a hello-world sample code that executes a print from both the untrusted and trusted domains.

## FPU poisoning PoC

Example output showcasing an example run with and without FPU poisoning in an
untrusted domain and enclave, respectively. Notice the faulty result for
`VADDPD` when poisoning the denormals-as-zero MXCSR flag:

```
jo@breuer:~/Documents/gotee/example/hello-world$ ./main 
From an untrusted domain:
Hello World!
[ENTRY] Control Word: 0x37f
[ENTRY] MXCSR: 0x1fa0

Add(1,2) is 3
[EXIT] Control Word: 0x37f
[EXIT] MXCSR: 0x1fa2
--
From a trusted domain:
Hello World!
[ENTRY] Control Word: 0x1f7f
[ENTRY] MXCSR: 0x1fe0

Add(1,2) is 0
[EXIT] Control Word: 0x1f7f
[EXIT] MXCSR: 0x1fe0
```

## Compiling

Just type `make` in `example/hello-world/`.
This will generate a `main` executable.
If you have a look at the `Makefile`, you'll see that the command used to compile is `gotee build src/main.go`.
We set the `GOPATH` variable before that to allow importing nested folders (pre-go-module way of doing things).

## Running

Execute the generate program, i.e., `./main`. 
The expected output is: 

```
From an untrusted domain:
Hello World!
From a trusted domain:
Hello World!
```

Executing the script will generate a `enclavebin` binary, i.e., the code and data loaded inside the enclave.
This is just for your convinience, to allow you to inspect what code is loaded inside the enclave.

Optionally, you can run the code in simulation mode like this `SIM=1 ./main`. This allows to run Gotee programs without SGX.

## Comments

As you can see in `src/main.go`, you need to import explicitly the `gosec` package.
As mentionned in the general README, we had to disable `gosecure` calls that target functions defined in main.
We also discourage using the `gosecure` keyword outside of the main package for the moment (see README for more information).
