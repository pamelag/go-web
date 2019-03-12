# webapp

Web application built using domain driven design principles to act as API Gateway and also serve static files

## objectives

- Setup a dockerized dev environment with hot reloading support for Golang and a Postgres DB with scripts
- Serve static files and handle RESTful API requests for Authoring, Transformation and Search services.

**Background** 

The RightPrism app is built as a Single Page Application. It needs a server-side program to serve the initial html file along with other static assets and handle REST API requests from the app.

**Action Points**

1. Serve static assets, handle REST API requests and act as an API Gateway to provide access to individual backend services for Commit Logs, Stream Processor and Image-to-Wireframe transformations.
2. Communicate with these services using gRPC and Protocol buffer and provide a circuit breaker.
   Handle the authentication and authorization, and other security-related features such as handling secure session cookies, setting secure headers, Cross-site scripting, Cross-site request forgery and input sanitization.
3. Use postgresql as its database.
4. Make use of Prepared statements for faster exceution and also prevent sql injection attacks.

**Structure** 

The code is based on the concepts defined in the Domain driven design approach. It has the following packages (inside web)
authoring
transformation
search
postgres
server
static
rde

authoring, transformation and search contains the application services corresponding to the three use-cases. authoring is used by requirement developers to author Feature documents and Wireframes. transformation is used in migration projects to convert existing forms to wireframes. Finally, search allows users to find any content in any feature or wireframe across the whole project.

rde is pure domain package for the requirement development environment. It consists of Project, Feature and Wireframe structs and ProjectRepository interface. The domain layer is used by the application services.

postgres contains a postgresql implementation of the repository interface. This repository interface is defined in the rde (domain) package.

server is where the code for routing, decoding requests, encoding responses, access control and other middleware required by an enterprise application reside.

static package contains the static assets html, js, css.

**Inversion of control** 

Each layer defines interfaces that allow the outer layers to inject concrete implementations, leaving the control of which implementation to use, to the caller. This is crucial, as it keeps unnecessary dependencies out of the domain model. An example would be the ProjectRepository, where the domain model exposes an interface but leaves the implementation to the infrastructure layer. So it decouples the choice of persistence technology from the domain model and application services. The database technology, postgresql in this case, can be easily replaced with another one, without making any changes to the domain layer. This is a huge win, as domain is where the logic that is central to the application is defined.

**Domain** 

The domain layer consists of entities like Project, Feature and Wireframe. These entities are defined by their identities ProjectID, FeatureID and WireframeID respectively. It also has the ProjectRepository interface. This layer is where the behaviours and the business rules are defined. The entire application must obey these rules. These entities in the domain layer are Go structs and have fields that are named carefully to reflect the terminology of the domain.

**Application Services** 

The application services define the operations available to the users like creating a project, a feature or a wireframe, or rename a wireframe's title. These service create new domain entities and contain input validation and sanitization before a Project or a Feature is created and stored in the domain repository.

**Infrastructure** 

The infrastructure layer is responsible for adapting request to our application layer, handling connection pools for database access, communicating with users using transport protocols, etc. This is where most of the third party packages are used.

The cool thing about structuring the project in this way is that, we can easily pick and choose these third-party packages in the infrastructure layer without really affecting the application layer or domain logic. They become interchangeable, as long as they adhere to an interface contract.

**packages** 

Domain driven design mentions the use of modules. While they're essentially used for organizing code, they actually tell the story of the domain at a larger scale. The name of the module actually becomes another tool in the conversations. In this project "rde" is the domain package. It stands for requirement development environment.

Each application service typically groups functionality related to a certain use-case. They are pretty self-contained, which lets them evolve individually. The dependencies that used in the infrastructure layer are not allowed to be a part of the application layer or the domain layer. They get their own package and these packages will typically contain implementations that can be injected into the application layer.

The postgres package is one such package. The name suggests that its probably going to deal with some SQL queries.

Finally the project is wired up in the main package where everything comes together to form the application. It basically becomes a matter of configuring the third-party packages and then injecting them into the application services.
