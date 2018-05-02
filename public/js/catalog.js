var loginInfo = {};
var cart = [];

$(document).ready(function() {
  $.ajax({
    type: 'GET',
    url: '/getLoginInfo',
    success: function(res) {
      loginInfo = res;
    }
  });

  $.ajax({
    type: 'GET',
    url: '/getProducts',
    success: function(res) {
      var rows = JSON.parse(res);
      if(rows[0] == null) {
        rows = products;
      }
      console.log(rows);
      var tbl = $("#productTable").empty();
      var row,tr,td,h3,p,button;
      for(var i=0; i<10; i++) {
        row = rows[i];
        if(i % 4 === 0) {
          if(i !== 0) {
            tr.appendTo(tbl);
          }
          tr = $("<tr/>");
        }
        td = $("<td/>");
        h3 = $("<h3/>").html(row.Name).appendTo(td);
        p = $("<p/>").html("Category: " + row.Category +
          "<br/>Description: " + row.Description +
          "<br/>Ingredients: " + row.Ingredients +
          "<br/>Price: " + row.Price +
          "<br/>Likes: " + row.Likes).appendTo(td);
        button = $("<button/>").attr('type','button')
          .addClass('btn btn-lg productBtn')
          .click(addToCartHandler(row))
          .appendTo(td)
          .html('Add To Cart');
        td.appendTo(tr);
      }
      tbl.append(tr);
    }
  });
});

function addToCartHandler(product) {
  return function() {
    $("#cartBtn").show();
    let productId = product.Product_id;
    if(!cart.includes(productId)) {
      cart.push(productId);

      let tbl = $("#orderTable").children().last();
      let tr,td1,td2;
      tr = $("<tr/>");
      td1 = $("<td/>").html(product.Name);
      td2 = $("<td/>").html('$&nbsp;<span>' + product.Price + '</span>');
      tr.append(td1);
      tr.append(td2);
      tbl.before(tr);

      // TODO: enable removing items from the cart

      let oldTotal = $("#total");
      let newTotal = parseFloat(oldTotal.html()) + parseFloat(product.Price);
      oldTotal.html(newTotal.toFixed(2));
    }
  }
}

function placeOrder() {
  let orderId = getOrderId();
  $.ajax({
    type: 'POST',
    url: '/placeOrder',
    data: {
      'Order_id': orderId,
      'User_id': loginInfo.user_id,
      'Timestamp': '',
      'Status': 'Placed',
      'Items': cart
    },
    success: function(res) {
      window.location.href = '/orders';
    }
  });
}

function getOrderId() {
  return Math.floor(Math.random() * (9999999 - 20 + 1)) + 20;
}

function toTitleCase(str)
{
  return str.replace(/\w\S*/g, function(txt) {
    return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
  });
}

function logout() {
  $.ajax({
    type: 'GET',
    url: '/logout',
    success: function(res) {
      window.location.href = '/'
    }
  });
}

var products = [
  {
    "Product_id": 1,
    "Name": "Americano",
    "Description": "A shot or two of espresso with hot water added",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 4.95,
    "Likes": 5
  },
  {
    "Product_id": 10,
    "Name": "Red Eye",
    "Description": "Coffee with a shot of espresso",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 2.95,
    "Likes": 5
  },
  {
    "Product_id": 2,
    "Name": "Cappuccino",
    "Description": "Espresso with a little bit of steamed milk topped with a lot of foam",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 3.95,
    "Likes": 5
  },
  {
    "Product_id": 3,
    "Name": "Mocha",
    "Description": "Espresso with steamed milk and chocolate (you can add whipped cream as well!)",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 2.95,
    "Likes": 5
  },
  {
    "Product_id": 4,
    "Name": "Café au Lait",
    "Description": "Half coffee with half steamed milk",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 6.95,
    "Likes": 5
  },
  {
    "Product_id": 5,
    "Name": "Espresso",
    "Description": "Coffee brewed by forcing a small amount of nearly boiling water under pressure through finely ground beans",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 4,
    "Likes": 5
  },
  {
    "Product_id": 6,
    "Name": "Latte",
    "Description": "Espresso with steamed milk and little to no foam",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 1.95,
    "Likes": 5
  },
  {
    "Product_id": 7,
    "Name": "Breve",
    "Description": "Espresso with steamed half and half",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 7.95,
    "Likes": 5
  },
  {
    "Product_id": 8,
    "Name": "Macchiato",
    "Description": "Double shot of espresso with foam on top",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 5.85,
    "Likes": 5
  },
  {
    "Product_id": 9,
    "Name": "Chai Latte",
    "Description": "Half steamed milk with half of our home brewed chai and a pump of vanilla syrup",
    "Ingredients": "milk, coffee beans, sugar, whipped cream",
    "Category": "coffee",
    "Price": 4.95,
    "Likes": 5
  }
];






