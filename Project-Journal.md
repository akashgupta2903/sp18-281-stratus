## Weekly Progress

| Week  | Progress  | Challenge  |
| ------------ | ------------ | ------------ |
| 4/1 - 4/07 | <ol><li>Met in order to decide the application for our project; finally agreed on a Starbucks-like web application supporting user authentication and order placement</li><li>Discussed the backend architecture and the API endpoints needed to provide the basic application functionalities</li><li>Discussed within the team to decide the databases to be used for storing different types of data</li></ol> | <ol><li>Agreeing on an application topic for our project</li><li>Balancing the number of backend APIs so that a smooth user experience could be provided, while also maintaining a reasonable project scope</li><li>Deciding the backend microservice architecture for our application</li></ol>|
| 4/8 - 4/14 | <ol><li>After reading about the various NoSQL databases, the team agreed on using MongoDB to store orders and users data &amp; Redis to store the product catalog data</li><li>Assigned the backend APIs to individuals so that development could be performed in parallel</li></ol> | <ol><li>Agreeing on using Redis as persistent database for storing product data</li><li>Splitting the application into microservices in a logical way</li></ol>  |
| 4/15 - 4/21 | <ol><li>Team started writing the backend Go APIs as per the individual assignment</li><li>Team started setting up the different database clusters to be used</li></ol>  | <ol><li>Writing backend APIs in Go due to lack of familiarity with the language</li><li>Setting up the Mongo replica sets and the Redis master-slave cluster</li></ol>  |
| 4/22 - 4/28 | <ol><li>Started developing a basic user interface for the application</li><li>Started deploying the backend services on ECS as per our architecture</li><li>Finished writing the Go APIs</li></ol>  | <ol><li>Integrating the backend with the frontend</li><li>Setting up the ECS clusters for various services</li></ol>  |
| 4/29 - 5/04 | <ol><li>Further end-to-end integration &amp; testing of the application</li><li>Implemented several UI features to further improve the user experience</li></ol>  | <ol><li>Fixing bugs due to integration issues</li><li>Settling on a process to demonstrate how the application still functions correctly in the presence of a network partition</li></ol>  |


## Architectural Overview

<img width="702" alt="screen shot 2018-05-02 at 7 57 03 pm" src="https://user-images.githubusercontent.com/32351699/39558394-12ba0b2a-4e43-11e8-9384-4a9a037f43ff.png">

## Youtube Demonstration

[Video](https://youtu.be/yLtQAJmPgJo)