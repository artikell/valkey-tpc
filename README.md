# Valkey-TPC

Valkey-TPC is a benchmarking tool inspired by the TPC-C benchmark. It is designed to help you evaluate the performance of your transaction processing systems.

      ┌───────────────────────────┐          
      │                           │          
      │           Engine          │          
      │                           │          
      └────┬─────────────────┬────┘          
           │                 │               
      ┌────▼────┐      ┌─────▼────┐          
      │         │      │          │          
      │ Storage │      │ Workload │          
      │         │      │          │          
      └─────────┘      └──────────┘          
       *Valkey          *Reddit - Website    
       *Redis                                
       *KeyDB                                

## Table of Contents

- [Getting Started](#getting-started)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installing

A step by step series of examples that tell you how to get a development environment running.

```
make build
```

## Testing

Explain how to run the automated tests for this system.

```
make test
```

## Contributing

Please read [CONTRIBUTING](CONTRIBUTING) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
