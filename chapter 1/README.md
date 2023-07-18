# Chapter 1: Getting the grips with test driven development

This chapter introduce the principles and benefits of Test Driven Development.

The main resoin to care about the testing is to verify and prove that the code you write behaves and performs to expectation and requiremnets of the project.

This chapter cover the following main topics:

-   The world or fundamentals of TDD
-   The benefits of TDD
-   Alternatives to TDD
-   Test metrics

## Exploring the world of TDD

TDD is a technique that allows us to write automated tests with short feedback loops.
Was created as an Agile working practice, this allows us:

-   Deliver code in an iterative process
-   Consisting of writing functional code
-   Verifying new code with tests
-   Iteratively refactor new code if required

### Introduction to agile methodology

This precursos to the Agile movement was the waterfall methodology.
The 5 stages if waterfal methodology are:

1. Requirement
2. Design
3. Implementation
4. Verification
5. Maintenance

There are some difficulties when waterfall its used:

-   Changing the course of the project or requirements is difficut
-   Customer need to decide all of their requirements in detail at the beginning
-   The process requires detailed documentation, which soecifies both requirements and the softare approach.

In practice, the waterfall model lacks the flexibility and iterative approach required for delivering complex software projects.

A better way of working named Agile emerged, which could address the challenges of the waterfall methodology.

The four core Agile values highlight the spirit of the movement:

1.  Individuals and interactions over the processes and tools
2.  Working software over comprehensive documentation
3.  Customer collaboration over contract negotiation
4.  Responding to change over following a plan

Unlike the waterfall model, the stages of the Agile software delivery methodology repeat, focusing on delivering software in small increments or iterations. In agile nomenclature, there iterations are called sprints

The sprint stages are:

-   Plan:
    The product owner discusses project requirements that will be delivered in current sprint with key stakeholders
-   Design
    This involves both, technical architecture and UI/UX design. This phase builds the requirements of Plan phase
-   Implement
    The design are used as guide from which we implement the scoped functionality
-   Test
    The test phase runs almost concurrently with implement phase, as test specifications can be written as soon as the design is completerd
-   Release
    The release phase begins. This phase completes any client-facing documentation or release notes. At the end of this phase, the sprint is considered closed. A new sprint can beign, following the same cycle.

We will focus on learning how to leverage its process and technoques in our Go projects.

### Types of automated tests

Automates tests that involve tools and frameworks to verify the behavior of software systems.

All tests define their inputs and expected outputs according to the requirements of system under test. These tests are divided into several types of tests according to three criteria:

-   The amount of knowledge they have of the system -> System knowledge type
-   The type of requirement thy verify -> Requirement types
-   The scope of funcionality they cover ->

**System knowledge**:
The tests can be divided into three categories according to how much internal knowledge of the system:

-   User -> Black box tests: Implementation details are not known -> Output
    -   Outputs are formulated according to requirements they verify
    -   Tests tend not to be brittle to internal changes
-   Developer -> White box tests: Implementation details are completely known -> Output
    -   Tests are more detiled and uncover hidden errors that black box testing cannot deliver
    -   Tests are often brittle if the internals of the system change
-   Specialist user -> Gray box tests: Implementation details are partially known -> Output
    -   Mix of black and white tests
    -   Tests can verify more advanced use cases and requirements
    -   Usually more time-consuming to write and run as well

**Requirements types**:
In general, we shold provide tests verify both, functionality and usability of system. We divide the tests of requirement that verify:

-   Functional tests:
    This tests a cover funcionality. Whith functional tests from prior sprints ensuring that there are no regressions in functionality in later sprint. Usually these tests are black box tests.

-   Non-functional tests:
    These tests cover all the aspects of the sysem that are not covered by functional. Cover aspects such as performance, usability and security. Usualy these tests are usually white box

**The testing pyramid**:
/|\
End to end
/ | \
Integration tests
/ | \
Unit tests

Unit tests:
They have small testing scope, covering the functionality of individual components.
Tratditionally thought of as white-box tests since thy are written by developers who knows all the impliementation detalis of the component

Integrations tests:
They test extend the scope and test the communication between multiple components. These tests canbe black-box or gray-box.

End to end:
They test the entire functionality of the application, ensuring the project deliverables are working according to requirements and can petentially be shepped

### The iterative approach of TDD

### TDD best practices

## Understanding the benefits and use of TDD

## Alternatives to TDD

## Understandint test metrics
