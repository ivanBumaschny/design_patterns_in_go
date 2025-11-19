**Motivation
    - Object creation logic becomes too complicated
        - You do not want to initialize element by element every time you start a complex struct. You usually want default values & business logic
    - Struct has too many fields, needs to initialize all correctly
    - Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to
        - Builder builds "slowly" (piecewise means little by little, not everything at once)
        - Factories aims to create the object in one go
        - A separate function (Factory Function, a.k.a Constructor)
        - That may exist in a separate struct (Factory). Just for the sake of organization. Can even be in a different package

A Factory is a Component responsible solely for the wholesale (not piecewise) creation of objects.

**Summary
    - A _factory function_ (a.k.a constructor) is a helper function for making struct instances
    - Gives additional flexibility at object creation
    - A factory is any entity that can take care of objkect creation
    - Can be a function or a dedicated structure, depending on use
        - If its a function, it should return a function that generates things (higher order function)
        - If its a structure, it should be accompained by an interface that implements & exposes the utilization methods (for example, the Create() method to initialize the objects)