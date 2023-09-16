# Chapter 1: Getting the grips with test driven development

This chapter introduce the principles and benefits of Test Driven Development.

The main resoin to care about the testing is to verify and prove that the code you write behaves and performs to expectation and requiremnets of the project.

This chapter cover the following main topics:

- The world or fundamentals of TDD
- The benefits of TDD
- Alternatives to TDD
- Test metrics

## Exploring the world of TDD

TDD is a technique that allows us to write automated tests with short feedback loops.
Was created as an Agile working practice, this allows us:

- Deliver code in an iterative process
- Consisting of writing functional code
- Verifying new code with tests
- Iteratively refactor new code if required

### Introduction to agile methodology

This precursos to the Agile movement was the waterfall methodology.
The 5 stages if waterfal methodology are:

1. Requirement
2. Design
3. Implementation
4. Verification
5. Maintenance

There are some difficulties when waterfall its used:

- Changing the course of the project or requirements is difficut
- Customer need to decide all of their requirements in detail at the beginning
- The process requires detailed documentation, which soecifies both requirements and the softare approach.

In practice, the waterfall model lacks the flexibility and iterative approach required for delivering complex software projects.

A better way of working named Agile emerged, which could address the challenges of the waterfall methodology.

The four core Agile values highlight the spirit of the movement:

1.  Individuals and interactions over the processes and tools
2.  Working software over comprehensive documentation
3.  Customer collaboration over contract negotiation
4.  Responding to change over following a plan

Unlike the waterfall model, the stages of the Agile software delivery methodology repeat, focusing on delivering software in small increments or iterations. In agile nomenclature, there iterations are called sprints

The sprint stages are:

- Plan:
  The product owner discusses project requirements that will be delivered in current sprint with key stakeholders
- Design
  This involves both, technical architecture and UI/UX design. This phase builds the requirements of Plan phase
- Implement
  The design are used as guide from which we implement the scoped functionality
- Test
  The test phase runs almost concurrently with implement phase, as test specifications can be written as soon as the design is completerd
- Release
  The release phase begins. This phase completes any client-facing documentation or release notes. At the end of this phase, the sprint is considered closed. A new sprint can beign, following the same cycle.

We will focus on learning how to leverage its process and technoques in our Go projects.

### Types of automated tests

Automates tests that involve tools and frameworks to verify the behavior of software systems.

All tests define their inputs and expected outputs according to the requirements of system under test. These tests are divided into several types of tests according to three criteria:

- The amount of knowledge they have of the system -> System knowledge type
- The type of requirement thy verify -> Requirement types
- The scope of funcionality they cover ->

**System knowledge**:
The tests can be divided into three categories according to how much internal knowledge of the system:

- User -> Black box tests: Implementation details are not known -> Output
  - Outputs are formulated according to requirements they verify
  - Tests tend not to be brittle to internal changes
- Developer -> White box tests: Implementation details are completely known -> Output
  - Tests are more detiled and uncover hidden errors that black box testing cannot deliver
  - Tests are often brittle if the internals of the system change
- Specialist user -> Gray box tests: Implementation details are partially known -> Output
  - Mix of black and white tests
  - Tests can verify more advanced use cases and requirements
  - Usually more time-consuming to write and run as well

**Requirements types**:
In general, we shold provide tests verify both, functionality and usability of system. We divide the tests of requirement that verify:

- Functional tests:
  This tests a cover funcionality. Whith functional tests from prior sprints ensuring that there are no regressions in functionality in later sprint. Usually these tests are black box tests.

- Non-functional tests:
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

The principle of TDD is: Write a unit test for a peace of functionality before implementing it.

The phases of TDD are:

1. Red phase: We consider what we want to test and translate the requirements into a test
2. Green phase: We swap from test code to implementation, writing just enoght code to make the failing test pass
3. Refactor phase: This is about cleaning up both, implementation code and test
4. We repeat this process until all requirements are tested and implemented and all tests pass

### TDD best practices

AAA(Arrange, Act, Assert) Pattern:

1. Arrange: We setup part of the test. We set up the UUT(unit under test) and all dependencies that it requires
2. Act: We perform the actions specified by the tests scenario. This steps uses the preconditions and inputs defined in Arrange
3. Assert: We confirm that UUT behaves according to requirements
4. If there is no more steps, the tests is considered passed. If there is more steps then we go back to Act phase
5. The Act and Assert steps can be repeated as many times as necessary for the test scenario. Just you should avoid to write lengthy or complicated tests.

Control scope:
Keep things as simple as posible to get some advantages

- Easy to debug
- Easy to mantain
- Faster execution of tests

Test outputs, not implementation:
It is important to focus on testing externally visible outputs, not implementation details.

Keep tests independent

- The tests should run independently of each other
- Starting point and end state of each test is as expected
- Best tests create their own UUT(Unit Under Test)

## Understanding the benefits and use of TDD

Pros:

- Allows development and testing processes to be done at once
- Developers analyse requirements in the beginning
- Increase the cofidence of code changes
- Developers have ownershig of their code quality

Cons:

- Requires witing more code
- Requirements need to be elaborated and more detail
- Test maintance can be time consuming

## Alternatives to TDD

Waterfall testing:
The entire project is delivered and all requirements are implemented by this point

Pros:

- The project is well structured and documented.
- Developers and testers acan rely on the project documentation to work independently.

Cons:

- The bugs can become complex
- A lot of implementation and testing might be wasted if the requirements change
- If there are delays in development, it can be easy to cut corners in the verification process, delivering an unstable product

Acceptance Test Driven Development:

- The customer requierements into a list of requirements that can be understood by all the stakeholders. These requirements are then converted in automated acceptance tests.

Pros:

- A complete suite of automated acceptance tests can be run after each commit or incremental code delivery
- The project will supported by a wide variety of stakeholders inside the business

Cons:

- Significant communication and synchronization effort is required
- It can be particularly challenging to write acceptance tests
- It can be challenging to get sample payloads or datasets from outsed of a project.

## Understanding test metrics

Important metrics:

- Requirement coverage
- Defect acount and distribution
- Defect resolution time
- Code coverage
- Burndown rates and chart

Code coverage:
Since the code coverage is important, tests shold be cover the following:

- The funcions you implemented
- The staments that your functions are composed of
- The different execution path of your functions
- The different conditions of your Boolean variables
- The different parameter values that can be passed to your functions

# Questions

1. What is the testing pyramid? What are its components?
2. What is the difference between functional and non-functional tests?
3. Explain what the red, green, and refactor TDD approach is.
4. What is ATDD?
5. What is code coverage?

# Answers

1. The testing pyramid specifies how automated test suites should be structured.
   At the bottom of the pyramid are unit tests, which test a single isolated component.
   Next up in the middle of the pyramid are integration tests, which test that multiple components are able to work together as specified.
   Finally, at the top of the test pyramid are end-to-end tests that test the behavior of the entire application.
2. Functional tests cover the correctness of a system, while non-functional tests cover the usability and performance of a system. Both types of tests are required to ensure that the system satisfies the customers’ needs.
3. The red, green, and refactor TDD approach refers to the three phases of the process. The red phase involves writing a new failing test for the functionality we intend to implement. The green phase involves writing enough implementation code to make all tests pass. Finally, the refactor phase involves optimizing both implementation and testing code to remove duplication and come up with better solutions.
4. Acceptance test-driven development. Just like TDD, ATDD puts tests first. ATDD is related to TDD, but it involves writing a suite of acceptance tests before the implementation begins. It involves multiple stakeholders to ensure that the acceptance test captures the customer’s requirements.
5. Code coverage is the percentage of your lines of code that are exercised by your unit test. This is calculated by considering the function statements, parameter values, and execution paths of your code. The Go test runner outputs the calculated code coverage. We should aim for a good value, but optimizing for 100% is normally not appropriate.
