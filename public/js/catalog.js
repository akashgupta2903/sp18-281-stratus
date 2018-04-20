$(document).ready(function() {
  // TODO: get actual product information
  /* 
  $.ajax({
    type: 'GET',
    url: '/users/getUserInfo',
    data: {
      username: username
    },
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
  let categories = {};
  let pd,cat,tabDiv,btn,contentDiv,tbl,lastTr;
  for(let i=0; i<productDataRaw.length; i++) {
    pd = productDataRaw[i];
    cat = pd.Category;
    if(cat in categories) {
      tbl = $("#"+cat).children().first()
      lastTr = tbl.children().last();
      if(lastTr.children().length === 4) {
        lastTr = $("<tr>").append(buildTd(pd)).appendTo(tbl);
      } else {
        lastTr.append(buildTd(pd));
      }
    } else {
      tabDiv = $("div.tabDiv");
      btn = $("<btn>")
        .addClass("tablinks")
        .html(toTitleCase(cat)); // TODO: add click handler
      if(Object.keys(categories === 0)) {
        btn.addClass("active");
      }
      tabDiv.append(btn);

      contentDiv = $("#contentDiv");
      let tab = buildTab(pd);
      categories[cat] = tab;
      contentDiv.append(tab);
    }
  }
});

function buildTab(product) {
  let div = $("<div>").addClass("tabcontent").attr("id", product.Category);
  let tbl = $("<table>").append($("<tr>").append(buildTd(product)));
  div.append(tbl);
  return div;
}

function buildTd(product) {
  ;
}

function toTitleCase(str)
{
  return str.replace(/\w\S*/g, function(txt) {
    return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
  });
}

// Switch to a tab on the user profile
function openTab(evt, tabName) {
  var i, tabcontent, tablinks;
  tabcontent = document.getElementsByClassName("tabcontent");
  for (i = 0; i < tabcontent.length; i++) {
    tabcontent[i].style.display = "none";
  }
  tablinks = document.getElementsByClassName("tablinks");
  for (i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace(" active", "");
  }
  document.getElementById(tabName).style.display = "block";
  evt.currentTarget.className += " active";
}









