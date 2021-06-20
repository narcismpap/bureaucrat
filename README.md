# Bureaucr.at Coding Challenge
Author: Narcis M. Pap - https://www.linkedin.com/in/narcismpap/

## Running
You can run a built-in example using:

```
make build
make example
```

Running via Docker
```
docker run docker.pkg.github.com/narcismpap/bureaucrat/bureaucrat:latest
```

Arguments
```
Usage of /bureaucrat:
  -l string
        Staff #1 ref (default "GoT-005")
  -r string
        Staff #2 ref (default "GoT-006")
  -s string
        .json staff directory file (default "/GoT.json")
```

### Task
Bureaucr.at is a typical hierarchical organization. Claire, its CEO, has a hierarchy of employees reporting to her and each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a Manager.

Your task is to implement a corporate directory for Bureaucr.at with an interface to find the closest common Manager (i.e. farthest from the CEO) between two employees. You may assume that all employees eventually report up to the CEO.

Here are some guidelines:

Resolve ambiguity with assumptions.
The directory should be an in-memory structure.
A Manager should link to Employees and not the other way around.
We prefer that you to use Go.
How the program takes its input and produces its output is up to you.
