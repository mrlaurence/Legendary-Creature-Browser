<?php
  require_once "simple_html_dom.php";

  $baseUrl = "https://en.wikipedia.org";

  $scrapeInfo = array();
  $path = "";

  // $alphabeticalLists = getAlphabeticalList();
  // foreach($alphabeticalLists as $link){
  //   scrapeList($link->href);
  // }

  init();
  scrape();

  //scrapeList("/wiki/List_of_legendary_creatures_(K)");

  finish();

  function init(){
    global $path;
    set_time_limit(0);
    $path = "scrapes/".date("Y-m-d-H-i-s");
    mkdir($path);
    file_put_contents($path."/scrape.json","");
    file_put_contents($path."/logfile.log","");
  }

  function finish(){
    global $scrapeInfo, $path;
    file_put_contents($path."/scrape.json",json_encode($scrapeInfo));
  }

  function scrape(){
    logMessage("Beginning scrape...\n");
    foreach(getAlphabeticalList() as $list){
      scrapeList($list->href);
    }
  }

  function getAlphabeticalList(){
    logMessage("Accessing https://en.wikipedia.org/wiki/Lists_of_legendary_creatures...");
    $dom = file_get_html("https://en.wikipedia.org/wiki/Lists_of_legendary_creatures");
    if(!empty($dom)){
      logMessage("Returning links to alphabetical lists...\n");
      return $dom->find("h2+ul li a");
    }
    return NULL;
  }

  function scrapeList($url){
    global $baseUrl;
    logMessage("Accessing list ".$baseUrl.$url."...");
    $dom = file_get_html($baseUrl.$url);
    logMessage("Scraping items ...");
      foreach($dom->find('.mw-parser-output ul li') as $link){
        if($link->first_child() !== $link->last_child()){
          // echo $link->first_child()->getAttribute("href")."<br/>";
          scrapeItem($link->first_child()->getAttribute("href"));
        }
      }
  }

  function scrapeItem($url){
    global $baseUrl, $scrapeInfo;
    $source = $baseUrl.$url;
    if(strpos($url, "List_of")!==false || strpos($url, "mythology")!==false ||  strpos($url, "redlink=1")!==false || strpos($url, "wiktionary")!==false || strpos($url, "#")!==false){
      return;
    }
    logMessage("Scraping item at ".$source."...");
    try {
      $dom = file_get_html($source);
    } catch (\Exception $e) {
      return;
    }
    logMessage("Scraping name...");
    $name = $dom->find("#firstHeading")[0]->innertext;
    logMessage("name = ".$name);
    logMessage("Searching for image...");
    $imgList = $dom->find("ul.gallery img");
    $img = "";
    if(count($imgList)!==0){
      $img = "https:".$imgList[0]->src;
      logMessage("Gallery image found");
    }else{
      $imgList = $dom->find("img");
      if(count($imgList)!==0){
        $img = "https:".$imgList[0]->src;
        logMessage("Image found");
      }
    }
    $description = ["Shrek", "Fucked", "Donkey"];
    logMessage("Pushing to scrapeInfo...\n");
    array_push($scrapeInfo,array($name => array("Description"=>$description, "Img"=>$img, "Link"=>$source)));

  }

  function logMessage($msg){
    global $path;
    file_put_contents($path."/logfile.log", $msg."\n", FILE_APPEND);
  }

?>
