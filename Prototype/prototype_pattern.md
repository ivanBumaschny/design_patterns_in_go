**Motivation
    - Complicated objects arent designed from scratch every time. You usually iterate an existing design and try to improve it
    - We make a copy of the prototype and customize it
        - Requires 'deep copy' support (all pointers to same memory slot)
    - We make the cloning convenient (e.g., via a Factory)
        - A factory that serves and, maybe, customizes the prototype at the point of creation

The Prototype Design Pattern is a partially or fully initialized object that you copy (clone) and make use of.

**Summary
    - To implement a prototype, partially construct an object and store it somewhere
    - Deep copy the prototype
    - Customize the resulting instance
    - A prototype factory provides a convenient API for using prototypes