**Motivation
    - Some objects are simple and can be created in a single constructor call
    - Other objects require a lot of ceremony to be created
    - Having a factory function with 10 argoments is not productive
    - Instead opt for piecewise, piece-by-piece, construction
    - Builder provides an API for construction an object step-by-step


**Summary
    - A builder is a separate component used for building an object
    - To make a builder fluent, return a reciever - allows chaining
    - Different facets of an object can be built with different builders working in tandem via a common struct 