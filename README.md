# BeeCS

Work-in-progress re-implementation of the [BEEHAVE](https://beehave-model.net) model
in [Go](https://go.dev) using the [Arche](https://github.com/mlange-42/arche) Entity Component System (ECS).

All the hard work to develop, parameterize and validate the original BEEHAVE model was done by Dr. Matthias Becher and co-workers.
I was not involved in that development in any way, and just re-implement the model following its ODD Protocol and the NetLogo code.

Beecs is currently at a state where it implements BEEHAVE's basic colony and foraging models.
Colony dynamics already mimic BEEHAVE quite well, but there are still differences that need to be addressed.

## Usage

There are currently no precompiled binaries.
Therefore, [Go](https://go.dev) is required to run beecs.

Clone the repository, and navigate into it:

```
git clone https://github.com/mlange-42/beecs.git
cd beecs
```

Run the default model setup with Go:

```
go run .
```

## Exploring the model

To explore the model code, start with reading file [`main.go`](https://github.com/mlange-42/beecs/blob/main/main.go).
Given that the model is developed with ECS, the structure should be very obvious.

- [`model/comp`](https://github.com/mlange-42/beecs/blob/main/model/comp) contains all components.
- [`model/sys`](https://github.com/mlange-42/beecs/blob/main/model/sys) contains all systems.
- [`model/res`](https://github.com/mlange-42/beecs/blob/main/model/sys) contains global parameters and variables (as ECS resources).
