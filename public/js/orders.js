var loginInfo = {};
var productsMap = {};
var orders = [];

$(document).ready(function() {
  $.ajax({
    type: 'GET',
    url: '/getLoginInfo',
    success: function(res) {
      loginInfo = res;

      $.ajax({
        type: 'GET',
        url: '/getProducts',
        success: function(res) {
          let rows = JSON.parse(res);
          if(rows[0] == null) {
            rows = products;
          }

          for(let i=0; i<rows.length; i++) {
            product = rows[i];
            productsMap[product.Product_id] = product;
          }

          showOrders();
        }
      });
    }
  });
});

function showOrders() {
  $.ajax({
    type: 'GET',
    url: '/getOrders',
    data: {
      'id': loginInfo.user_id
    },
    success: function(res) {
      orders = JSON.parse(res);
      let tbl = $("#tblOrders").empty();
      let tr,td;
      for(let i=0; i<orders.length; i++) {
        var order = orders[i];
        if(order.Status == "Placed") {
          tr = $("<tr>").append(buildOrderTd(order));
          tbl.append(tr);
        }
      }
    }
  });
}

// takes an order object and constructs a <td> element
function buildOrderTd(order) {
  let td = $("<td>").addClass("majorTd");

  // order ID, timestamp, and status
  let div1 = $("<div>").addClass("col3-left")
    .append($("<span>").html("Order ID: " + order.Order_id))
    .append($("<br>"))
    .append($("<span>").html("Order placed on " + order.Timestamp))
    .append($("<br>"))
    .append($("<span>").html("Order status: " + order.Status));

  // order items
  let div2 = $("<div>").addClass("col3-center");
  let items = order.Items;
  for(let i=0; i<items.length; i++) {
    let item = items[i];
    let product = productsMap[item];
    let span = $("<span>").addClass("itemSpan")
      .html(product.Name + "&nbsp;")
      .append($("<i class='fas fa-times-circle'></i>").click(removeHandler(order, item)));
    div2.append(span);
  }
  let btn = $("<button>").addClass("btn btn-md orderButton")
    .html("Update Order")
    .click(updateHandler(order)).appendTo(div2);

  // update or cancel the order
  let div3 = $("<div>").addClass("col3-right divOrderButtons")
    .append($("<button>").attr("type","button").click(function() {
      $.ajax({
        type: 'GET',
        url: '/payOrder',
        data: {
          id: order.Order_id
        },
        success: function(res) {
          showOrders();
        }
      });
    }).html("Pay").addClass("orderButton"))
    .append($("<br>"))
    .append($("<button>").attr("type","button").click(function() {
      $.ajax({
        type: 'GET',
        url: '/cancelOrder',
        data: {
          id: order.Order_id
        },
        success: function(res) {
          showOrders();
        }
      });
    }).html("Cancel").addClass("orderButton"));

  td.append(div1).append(div2).append(div3);

  return td;
}

function updateHandler(order) {
  return function() {
    console.log(order.Items);
    $.ajax({
      type: 'PUT',
      url: '/updateOrder',
      data: {
        'Order_id': order.Order_id,
        'User_id': order.User_id,
        'Timestamp': '',
        'Status': order.Status,
        'Items': order.Items
      },
      success: function(res) {
        showOrders();
      }
    });
  }
}

function removeHandler(order, item) {
  return function() {
    console.log(order.Items);
    let idx = order.Items.indexOf(item);
    if(idx > -1) {
      order.Items.splice(idx, 1);
    }
    console.log(order.Items);
    $(this).parent().remove();
  }
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























