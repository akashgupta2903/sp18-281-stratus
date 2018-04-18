This module contains APIs written in golang to handle the Product catalog part of Starbucks application(which our team is developing as part of Team Hackathon Project). Redis is used as the backend datastore for managing products.

Below is the list of APIs handling the catalogue part of the application:

| API Route           | Method           | Description                                                    |
| --------------------|------------------| ---------------------------------------------------------------|
| /getdetail?id={id}  | GET              | Responsible for getting a product based on it's id             |
| /like?id={id}       | POST             | Responsible for liking a product based on it's id              |
| /getallproducts     | GET              | Responsible for loading all the products in catalog            |
| /popular            | GET              | Responsible for getting top 3 products based on user likes     |