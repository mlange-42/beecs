# beecs

Work-in-progress re-implementation of the [BEEHAVE](https://beehave-model.net) model
in [Go](https://go.dev) using the [Arche](https://github.com/mlange-42/arche) Entity Component System (ECS).

All the hard work to develop, parameterize and validate the original BEEHAVE model was done by Dr. Matthias Becher and co-workers.
I was not involved in that development in any way, and just re-implement the model following its ODD Protocol and the NetLogo code.

Beecs is currently at a state where it implements BEEHAVE's basic colony and foraging models.
Colony dynamics already mimic BEEHAVE quite well, but there are still differences that need to be addressed. See the [tests](https://github.com/mlange-42/beecs/tree/main/tests).

## Usage

### As a library

To add beecs as a dependency to an existing Go project, run this in the project's root folder:

```
go get github.com/mlange-42/beecs
```

Usage examples can be found in folder [_examples](https://github.com/mlange-42/beecs/blob/main/_examples).

### As a command line app

[TODO]

### With a graphical user interface

[TODO]

## Exploring the model

To explore the model code, start with reading files [`model/default.go`](https://github.com/mlange-42/beecs/blob/main/model/default.go)
and [`params/default.go`](https://github.com/mlange-42/beecs/blob/main/params/default.go).
Given that the model is developed with ECS, the structure should be quite obvious.

- [`comp`](https://github.com/mlange-42/beecs/blob/main/comp) contains all components.
- [`sys`](https://github.com/mlange-42/beecs/blob/main/sys) contains all systems.
- [`params`](https://github.com/mlange-42/beecs/blob/main/params) contains model parameters (as ECS resources).
- [`globals`](https://github.com/mlange-42/beecs/blob/main/globals) contains global variables (as ECS resources).

Also check out the [examples](https://github.com/mlange-42/beecs/blob/main/_examples).
