const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');
const requestify = require('requestify');
const request = require('request');
const app = express();
const url = require('url');
const fs = require('fs');

// Heroku environment variables
const PORT = process.env.PORT || 5555;
const API_PORT = process.env.API_PORT || 4000;
const USERS_PORT = process.env.USERS_PORT || 8000;
//const PRODUCTS_ENDPOINT = process.env.PRODUCTS_ENDPOINT || "ECS-first-run-alb-201447009.us-west-1.elb.amazonaws.com";
const PRODUCTS_ENDPOINT = process.env.PRODUCTS_ENDPOINT || "ECS-first-run-alb-1212423008.us-west-1.elb.amazonaws.com";
const ORDERS_ENDPOINT = process.env.ORDERS_ENDPOINT || "ECS-first-run-alb-454373483.us-west-1.elb.amazonaws.com";
const USERS_ENDPOINT = process.env.USERS_ENDPOINT || "ECS-first-run-alb-431339320.us-west-1.elb.amazonaws.com";

app.use(express.static(path.join(__dirname, 'public')));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

var server = app.listen(PORT, () => console.log(`Listening on ${ PORT }`));
var loginInfo = {};

app.get('/', (req, res) => res.render('index'));
app.get('/orders', (req, res) => res.render('orders'));
app.get('/catalog', (req, res) => res.render('catalog'));

app.get('/getEndpoints', function(req, res) {
  var endpoints = {
    prd: PRODUCTS_ENDPOINT,
    ord: ORDERS_ENDPOINT,
    usr: USERS_ENDPOINT,
    prt: API_PORT
  };
  res.json(endpoints);
});

app.get('/getLoginInfo', function(req, res) {
  res.json(loginInfo);
});

// USER APIS

app.get('/login', (req, res) => {
  let url = 'http://' + USERS_ENDPOINT + ':' + USERS_PORT + '/login';
  let options = {
    url: url,
    method: 'POST',
    json: true,
    body: {
      'id': '',
      'username': req.query.username,
      'email': '',
      'password': req.query.password
    }
  };
  request(options, function (error, response, body) {
    if (!error && !body.error) {
      loginInfo = body;
      res.sendStatus(200);
    } else {
      console.log("ERROR in /login: " + body.error);
      res.sendStatus(500);
    }
  });
});

app.get('/register', function(req, res) {
  let url = 'http://' + USERS_ENDPOINT + ':' + USERS_PORT + '/signup';
  let options = {
    url: url,
    method: 'POST',
    json: true,
    body: {
      'id': '',
      'username': req.query.username,
      'email': req.query.email,
      'password': req.query.password
    }
  };
  request(options, function (error, response, body) {
    if (!error && !body.error) {
      loginInfo = body;
      res.sendStatus(200);
    } else {
      console.log("ERROR in /register: " + body.error);
      res.sendStatus(500);
    }
  });
});

app.get('/logout', function(req, res) {
  let url = 'http://' + USERS_ENDPOINT + ':' + USERS_PORT + '/logout';
  let options = {
    url: url,
    method: 'POST'
  };
  request(options, function (error, response, body) {
    if (!error) {
      res.sendStatus(200);
    } else {
      console.log("ERROR in /logout");
      res.sendStatus(500);
    }
  });
});


// ORDER APIS

app.get('/getOrders', function(req, res) {
  let id = req.query.id;
  let url = 'http://' + ORDERS_ENDPOINT + ':' + API_PORT + '/getorders';
  //let url = 'http://ec2-52-53-181-58.us-west-1.compute.amazonaws.com:4000/getorders';
  let options = {
    url: url,
    method: 'GET',
    qs: {'id':id}
  };
  request(options, function (error, response, body) {
    if (!error && response.statusCode == 200) {
      res.send(body);
    } else {
      console.log("ERROR in /getOrders");
      res.sendStatus(500);
    }
  });
});

app.post('/placeOrder', function(req, res) {
  let newItems = req.body.Items.map(Number);
  let url = 'http://' + ORDERS_ENDPOINT + ':' + API_PORT + '/placeorder';
  let options = {
    url: url,
    method: 'POST',
    json: true,
    body: {
      'Order_id': parseInt(req.body.Order_id),
      'User_id': req.body.User_id,
      'Timestamp': req.body.Timestamp,
      'Status': req.body.Status,
      'Items': newItems
    }
  };
  request(options, function (error, response, body) {
    if (!error) {
      res.send(body);
    } else {
      console.log("ERROR in /placeOrder");
      res.sendStatus(500);
    }
  });
});

app.put('/updateOrder', function(req, res) {
  let newItems = req.body.Items.map(Number);
  let url = 'http://' + ORDERS_ENDPOINT + ':' + API_PORT + '/updateorder';
  let options = {
    url: url,
    method: 'PUT',
    json: true,
    body: {
      'Order_id': parseInt(req.body.Order_id),
      'User_id': req.body.User_id,
      'Timestamp': req.body.Timestamp,
      'Status': req.body.Status,
      'Items': newItems
    }
  };
  request(options, function (error, response, body) {
    if (!error) {
      console.log("success");
      res.sendStatus(200);
    } else {
      console.log("ERROR in /updateOrder");
      res.sendStatus(500);
    }
  });
});

app.get('/payOrder', function(req, res) {
  let id = parseInt(req.query.id);
  let url = 'http://' + ORDERS_ENDPOINT + ':' + API_PORT + '/payorder';
  let options = {
    url: url,
    method: 'POST',
    qs: {'id':id}
  };
  request(options, function (error, response, body) {
    if (!error) {
      res.sendStatus(200);
    } else {
      console.log("ERROR in /payOrder");
      res.sendStatus(500);
    }
  });
});

app.get('/cancelOrder', function(req, res) {
  let id = parseInt(req.query.id);
  let url = 'http://' + ORDERS_ENDPOINT + ':' + API_PORT + '/cancelorder';
  let options = {
    url: url,
    method: 'DELETE',
    qs: {'id':id}
  };
  request(options, function (error, response, body) {
    if (!error) {
      res.sendStatus(200);
    } else {
      console.log("ERROR in /cancelOrder");
      res.sendStatus(500);
    }
  });
});

// PRODUCT APIS

app.get('/getProducts', function(req, res) {
  let url = 'http://' + PRODUCTS_ENDPOINT + ':' + API_PORT + '/getallproducts';
  let options = {
    url: url,
    method: 'GET'
  };
  request(options, function (error, response, body) {
    if (!error && response.statusCode == 200) {
      res.send(body);
    } else {
      console.log("ERROR in /getProducts");
      res.sendStatus(500);
    }
  });
});

app.post('/likeProduct', function(req, res) {
  let id = parseInt(req.body.id);
  let url = 'http://' + PRODUCTS_ENDPOINT + ':' + API_PORT + '/like';
  let options = {
    url: url,
    method: 'POST',
    qs: {'id':id}
  };
  request(options, function (error, response, body) {
    if (!error) {
      res.send(body);
    } else {
      console.log("ERROR in /likeProduct");
      res.sendStatus(500);
    }
  });
});

function stop() {
  server.close();
}

module.exports = app;
module.exports.stop = stop;