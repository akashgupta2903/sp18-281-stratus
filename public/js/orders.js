$(document).ready(function() {
  // TODO: get actual list of ordersi & products
  /*
  $.ajax({
    type: 'GET',
    url: '/getOrders',
    success: function(res) {
      //
    }
  });*/
  let productDataRaw = [
    {
      Product_id: 1,
      Name: "Cafe au Lait",
      Description: "Half coffee with half steaming milk",
      Ingredients: "milk, coffee beans, sugar",
      Category: "coffee",
      Price: 6.95,
      Likes: 5
    },
    {
      Product_id: 2,
      Name: "Croissant",
      Description: "Buttery flaky pastry originating from France",
      Ingredients: "wheat, eggs, powdered sugar",
      Category: "pastry",
      Price: 4.95,
      Likes: 10
    },
    {
      Product_id: 3,
      Name: "Mocha Frappucino",
      Description: "A coffee that tastes of chocolate",
      Ingredients: "milk, coffee beans, cocoa",
      Category: "coffee",
      Price: 5.95,
      Likes: 2
    },
    {
      Product_id: 4,
      Name: "Bagel",
      Description: "Goes great with lox or cream cheese",
      Ingredients: "wheat, eggs, yeast",
      Category: "pastry",
      Price: 2.95,
      Likes: 7
    },
    {
      Product_id: 5,
      Name: "Green Tea Latte",
      Description: "Green tea latte for those with a refined taste",
      Ingredients: "milk, matcha powder, sugar",
      Category: "coffee",
      Price: 3.95,
      Likes: 4
    },
  ];
  let productData = {};
  for(let i=0; i<productDataRaw.length; i++) {
    productData[productDataRaw[i].Product_id] = productDataRaw[i];
  }

  let orderData = [
    {
      order_id: 1,
      user_id: 1,
      timestamp: "2018-04-18 01:01:03",
      status: "Placed",
      items: [1,2,3]
    },
    {
      order_id: 2,
      user_id: 1,
      timestamp: "2018-04-18 01:01:02",
      status: "Placed",
      items: [2,4,5]
    },
    {
      order_id: 3,
      user_id: 1,
      timestamp: "2018-04-18 01:01:01",
      status: "Placed",
      items: [1,3,4,5]
    },
  ];

  let tbl = $("#tblOrders");
  let tr,td;
  for(let i=0; i<orderData.length; i++) {
    tr = $("<tr>").append(buildOrderTd(orderData[i]));
    tbl.append(tr);
  }
});

// takes an order object and constructs a <td> element
function buildOrderTd(order) {
  let td = $("<td>").addClass("majorTd");

  // order ID, timestamp, and status
  let div1 = $("<div>").addClass("col3-left")
    .append($("<span>").html("Order ID: " + order.order_id))
    .append($("<br>"))
    .append($("<span>").html("Order placed on " + order.timestamp))
    .append($("<br>"))
    .append($("<span>").html("Order status: " + order.status));

  // order items
  let div2 = $("<div>").addClass("col3-center").html("Items here");
  // TODO: append add-able and remove-able order items here

  // update or cancel the order
  let div3 = $("<div>").addClass("col3-right divOrderButtons")
    .append($("<button>").attr("type","button").click(function() {
      // TODO: ajax request to update the order
      window.location.reload(true);
    }).html("Update").addClass("orderButton"))
    .append($("<br>"))
    .append($("<button>").attr("type","button").click(function() {
      // TODO: ajax request to cancel the order
      window.location.reload(true);
    }).html("Cancel").addClass("orderButton"));

  td.append(div1).append(div2).append(div3);

  return td;
}

























