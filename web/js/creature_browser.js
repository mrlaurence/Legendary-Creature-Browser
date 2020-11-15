$('.carousel').carousel({
  interval: false
})

var apiLinkBase = "http://127.0.0.1";

function runSearch() {

  var searchTerms = document.getElementById("inputSearch").value;
  var resultsNo = document.getElementById("inputResultsNo").value;
  var sensitivity = document.getElementById("inputSensitivity").value;

  var xmlHttp = new XMLHttpRequest();

  var url = apiLinkBase + "/search?n=" + resultsNo + "&q=" + escape(searchTerms) + "&s=" + sensitivity * 100;

  xmlHttp.open("GET", url, false);
  xmlHttp.send(null);
  creatureData = xmlHttp.responseText;

  var creaturesObject = JSON.parse(creatureData);

  var caroInner = document.getElementById("carousel-inner");

  caroInner.innerHTML = "";

  var firstItem = true;

  if (Object.keys(creaturesObject).length === 0 && creaturesObject.constructor === Object) {
    caroInner.innerHTML = "<div style='text-align:center'><h2>No results :(</h2></div>";
  }
  else {
    Object.keys(creaturesObject).forEach(function(key, index) {

        var creatureName = key;
        var creatureTagString = creaturesObject[key].Description.join(', ');
        var creatureWikiLink = creaturesObject[key].Link;
        var creatureImageLink = creaturesObject[key].Img;

        if (firstItem) {
          firstItem = false;

          caroInner.innerHTML += '<div class="carousel-item active"><div class="creature-div"><h1>' + creatureName + '</h1><div class="creature-col"><p>Tags: ' + creatureTagString + '</p><p>Wikipedia page: <a href="' + creatureWikiLink + '">' + creatureName + '</a></p></div><div class="creature-col"><img src="' + creatureImageLink + '" /></div></div></div>'
        } else {
          caroInner.innerHTML += '<div class="carousel-item"><div class="creature-div"><h1>' + creatureName + '</h1><div class="creature-col"><p>Tags: ' + creatureTagString + '</p><p>Wikipedia page: <a href="' + creatureWikiLink + '">' + creatureName + '</a></p></div><div class="creature-col"><img src="' + creatureImageLink + '" /></div></div></div>'
        }
      })
}};

$('#search-form').submit(function() {
  runSearch();
  return false;
});

function getRandom() {
  var resultsNo = document.getElementById("inputResultsNo").value;
  var sensitivity = document.getElementById("inputSensitivity").value;

  var xmlHttp = new XMLHttpRequest();

  var url = apiLinkBase + "/random?n=" + resultsNo;

  xmlHttp.open("GET", url, false);
  xmlHttp.send(null);
  creatureData = xmlHttp.responseText;

  var creaturesObject = JSON.parse(creatureData);

  var caroInner = document.getElementById("carousel-inner");

  caroInner.innerHTML = "";

  var firstItem = true;

  if (Object.keys(creaturesObject).length === 0 && creaturesObject.constructor === Object) {
    caroInner.innerHTML = "<div style='text-align:center'><h2>No results :(</h2></div>";
  }
  else {
    Object.keys(creaturesObject).forEach(function(key, index) {

        var creatureName = key;
        var creatureTagString = creaturesObject[key].Description.join(', ');
        var creatureWikiLink = creaturesObject[key].Link;
        var creatureImageLink = creaturesObject[key].Img;

        if (firstItem) {
          firstItem = false;

          caroInner.innerHTML += '<div class="carousel-item active"><div class="creature-div"><h1>' + creatureName + '</h1><div class="creature-col"><p>Tags: ' + creatureTagString + '</p><p>Wikipedia page: <a href="' + creatureWikiLink + '">' + creatureName + '</a></p></div><div class="creature-col"><img src="' + creatureImageLink + '" /></div></div></div>'
        } else {
          caroInner.innerHTML += '<div class="carousel-item"><div class="creature-div"><h1>' + creatureName + '</h1><div class="creature-col"><p>Tags: ' + creatureTagString + '</p><p>Wikipedia page: <a href="' + creatureWikiLink + '">' + creatureName + '</a></p></div><div class="creature-col"><img src="' + creatureImageLink + '" /></div></div></div>'
        }
      })
}};
