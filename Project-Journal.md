# CMPE 281 Stratus Team Project: Starbucks

## Weekly Progress

| Week  | Progress  | Challenge  |
| ------------ | ------------ | ------------ |
| 4/1 - 4/07 | <ol><li>Met in order to decide the application for our project; finally agreed on a Starbucks-like web application supporting user authentication and order placement</li><li>Discussed the backend architecture and the API endpoints needed to provide the basic application functionalities</li><li>Discussed within the team to decide the databases to be used for storing different types of data</li></ol> | <ol><li>Agreeing on an application topic for our project</li><li>Balancing the number of backend APIs so that a smooth user experience could be provided, while also maintaining a reasonable project scope</li><li>Deciding the backend microservice architecture for our application</li></ol>|
| 4/8 - 4/14 | <ol><li>After reading about the various NoSQL databases, the team agreed on using MongoDB to store orders and users data &amp; Redis to store the product catalog data</li><li>Assigned the backend APIs to individuals so that development could be performed in parallel</li></ol> | <ol><li>Agreeing on using Redis as persistent database for storing product data</li><li>Splitting the application into microservices in a logical way</li></ol>  |
| 4/15 - 4/21 | <ol><li>Team started writing the backend Go APIs as per the individual assignment</li><li>Team started setting up the different database clusters to be used</li></ol>  | <ol><li>Writing backend APIs in Go due to lack of familiarity with the language</li><li>Setting up the Mongo replica sets and the Redis master-slave cluster</li></ol>  |
| 4/22 - 4/28 | <ol><li>Started developing a basic user interface for the application</li><li>Started deploying the backend services on ECS as per our architecture</li><li>Finished writing the Go APIs</li></ol>  | <ol><li>Integrating the backend with the frontend</li><li>Setting up the ECS clusters for various services</li></ol>  |
| 4/29 - 5/04 | <ol><li>Further end-to-end integration &amp; testing of the application</li><li>Implemented several UI features to further improve the user experience</li></ol>  | <ol><li>Fixing bugs due to integration issues</li><li>Settling on a process to demonstrate how the application still functions correctly in the presence of a network partition</li></ol>  |


## Architectural Overview

### Diagram

<img width="739" alt="screen shot 2018-05-04 at 6 02 50 pm" src="https://user-images.githubusercontent.com/32351699/39658180-777b8858-4fc5-11e8-84a5-7de086138b7e.png">


### Description

1. Web Browsers
  * Written in HTML, CSS, and JavaScript, with jQuery
  * 3 basic pages: home page (supporting login &amp; register; catalog page (supporting viewing products &amp; placing an order); orders page (supporting updating, paying for, or canceling an order)
2. Web Server
  * Node.js server responsible for handling requests from the UI, and forwarding them to the correct service
3. API Layer
  * Go APIs written by the team members, responsible for processing the request and interacting with the database layer as needed
  * Each microservice (user APIs, order APIs, or product APIs) is deployed on an AWS ECS instance cluster of 2 containers, with a load balancer in front
  * Load balancer configured to use sticky sessions so that requests from a particular client are always routed to the same service
4. Persistence Layer
  * Distributed NoSQL databases, sharded on the type of data (user, order, or product data)
  * User &amp; order data stored in MongoDB clusters with one primary and four secondary nodes, with automated reelection of the primary node in the case of a network partition, to basically ensure availability
  * Product data stored in Redis cluster with one master node and four slaves, with a quorum approach employed to basically ensure consistency in the case of a network partition

## Youtube Demonstration

<<<<<<< HEAD
[Video](https://youtu.be/p_lyTTVW22M)
=======
[Video](https://youtu.be/yLtQAJmPgJo)
>>>>>>> 2972f3cb9698211b6bb2d69a86c4693e7dddcc15
