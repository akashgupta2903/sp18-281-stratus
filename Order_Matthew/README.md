This module contains the API written in golang to handle the update order functionality (adding new items to an existing order, or removing items from an existing order) of the Starbucks application (which our team is developing as part of Team Hackathon Project). MongoDB is used as the database for managing orders.

Below is the list of APIs handled in this part of the application:

| API Route           | Method           | Description                                                    |
| --------------------|------------------| ---------------------------------------------------------------|
| /updateorder  | PUT              | Responsible for updating an order's items. The order object is passed within the request body, with some fields being set by the server prior to insertion into the database.            |
