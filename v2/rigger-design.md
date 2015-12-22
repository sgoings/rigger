rigger
======

What's New?
-----------

rigger is now a general purpose, pluggable workflow tool.

How does it work?
-----------------

rigger is a tool that enables you to build reusable clis by living under one
of the most important design patterns:

> Program to an interface, not an implementation

How does rigger do this?
------------------------

Consider it this way... in object oriented programming, you usually create:

  - a class
  - an instantiation (or multiple) of that class

In rigger terminology, we have:

  - a *workflow* which determines a cli tool's set of actions (a cli user interface)
  - a *provider* which is a set of actions that implements the workflow

What does this enable?
----------------------

Say we want a way to have a user expect to be able to create/destroy a Kubernetes
cluster via the actions:

- up
- down

But we want this 3 step process to be used no matter what type of infrastructure
we want to use to provision a Kubernetes cluster.

With rigger, you can define different providers (AWS, DigitalOcean, GCE) that
will operate under this single workflow's interface so that no matter what
cloud provider you're using you get an experience that's as consistent as
possible.

Workflows
---------

Consider a workflow to be an interface (or more accurately: an abstract class).
Workflows are created by people who care the most about the end user experience.
They also define what each step in the cli tool should generally provide
(inputs, outputs, etc.)

Provders
--------

Consider a provider to be a concrete class. These concrete classes are provided
by people that are intimately aware of the underlying functionality/logic. In
the above example, this would be cloud providers like AWS, GCE, etc.

Initialization
------------

`rigger init` is the process of instantiating the workflow using a specific provider
into an "object." The way that rigger defines the instantiation can come from a
preexisting answers file, a human, or a web service...
