This module contains APIs written in golang to handle the Product catalog part of Starbucks application(which our team is developing as part of Team Hackathon Project). Redis is used as the backend datastore for managing products.

Below is the list of APIs handling the catalogue part of the application:

| API Route           | Method           | Description                                                    |
| --------------------|------------------| ---------------------------------------------------------------|
| /getdetail?id={id}  | GET              | Responsible for getting a product based on it's id             |
| /like?id={id}       | POST             | Responsible for liking a product based on it's id              |
| /getallproducts     | GET              | Responsible for loading all the products in catalog            |
| /popular            | GET              | Responsible for getting top 3 products based on user likes     |


<img width="1237" alt="screen shot 2018-04-18 at 3 02 47 pm" src="https://user-images.githubusercontent.com/32351699/38961367-2950ce7c-431d-11e8-8df3-3af64e787257.png">


<img width="1227" alt="screen shot 2018-04-18 at 3 04 18 pm" src="https://user-images.githubusercontent.com/32351699/38961373-2ed4952c-431d-11e8-8c88-528c2ae1b8c4.png">


<img width="1230" alt="screen shot 2018-04-18 at 3 05 24 pm" src="https://user-images.githubusercontent.com/32351699/38961375-31a995a4-431d-11e8-9235-32537ef93681.png">


<img width="1231" alt="screen shot 2018-04-18 at 3 05 47 pm" src="https://user-images.githubusercontent.com/32351699/38961376-348e7dac-431d-11e8-88f7-b2eff2357daf.png">
