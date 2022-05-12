<!DOCTYPE html>
<html>
<head>
  <title>GO-SEC-CODE</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://cdn.staticfile.org/twitter-bootstrap/5.1.1/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.staticfile.org/twitter-bootstrap/5.1.1/js/bootstrap.bundle.min.js"></script>
</head>
<body>

<div class="container-fluid p-5 bg-primary text-white text-center">
  <h1>GO-SEC-CODE</h1>
  <p>Go-sec-code is a  project for learning Go vulnerability code.</p> 
</div>
  
<div class="container mt-5">
  <div class="row">
    <div class="col-sm-4">
      <h3>CommandInject</h3>
      <h4><a href="/commandInject/vuln?dir=%2F">vuln code 1</a></h4>
      <h4><a href=/commandInject/vuln/host>vuln code 2</a></h4>
      <h4><a href="/commandInject/safe?dir=.%3Bwhoami">safe code</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>Cors</h3>
      <h4><a href="/cors/vuln/reflect">vuln code</a></h4>
      <h4><a href="/cors/vuln/any-origin-with-credential">vuln code 2</a></h4>
      <h4><a href="/cors/safe">safe code</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>CRLF-Injection</h3>        
      <h4><a href=/crlfInjection/safe?header>safe code</a></h4>
    </div>
  </div>
  <div class="row">
    <div class="col-sm-4">
      <h3>FileUpload</h3>
      <h4><a href=/fileUpload/vuln>vuln code</a></h4>
      <h4><a href=/fileUpload/safe>safe code</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>JSONP</h3>        
      <h4><a href="/jsonp/vuln/noCheck?callback=jsonp">vuln code</a></h4>
      <h4><a href="/jsonp/vuln/emptyReferer?callback=jsonp">vuln code 2</a></h4>
      <h4><a href="/jsonp/safe?callback=jsonp">safe code</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>PathTraversal</h3>        
      <h4><a href="/pathTraversal/vuln?file=../../../../../../../etc/passwd">vuln code</a></h4>
      <h4><a href="/pathTraversal/safe?file=../../../../../../../etc/passwd">safe code</a></h4>
    </div>
  </div>
  <div class="row">
    <div class="col-sm-4">
      <h3>SQLInjection</h3>
      <h4><a href="/sqlInjection/native/vuln/integer?id=1">vuln code 1</a></h4>
      <h4><a href="/sqlInjection/native/vuln/string?username=admin">vuln code 2</a></h4>
      <h4><a href="/sqlInjection/native/safe/integer?id=1">safe code 1</a></h4>
      <h4><a href="/sqlInjection/native/safe/string?username=admin">safe code 2</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>SSRF</h3>
      <h4><a href="/ssrf/vuln?url">vuln code 1</a></h4>
      <h4><a href="/ssrf/vuln/302?url">vuln code 2</a></h4>
      <h4><a href="/ssrf/vuln/dnsrebinding?url">vuln code 3</a></h4>
      <h4><a href="/ssrf/safe/whitelists?url">safe code 1</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>SSTI</h3>
      <h4><a href="/ssti/vuln?template">vuln code 1</a></h4>
      <h4><a href="/ssti/safe?template">safe code 1</a></h4>
    </div>
  </div>
  <div class="row">
    <div class="col-sm-4">
      <h3>XSS</h3>
      <h4><a href="/xss/vuln?xss">vuln code 1</a></h4>
      <h4><a href="/xss/vuln/store?xss">vuln code 2</a></h4>
      <h4><a href="/xss/safe?xss">safe code 1</a></h4>
    </div>
    <div class="col-sm-4">
      <h3>XXE</h3>
      <h4><a href="/xxe/vuln">vuln code 1</a></h4>
      <!-- <h4><a href="/xxe/vuln/gokogiri">vuln code 2</a></h4> -->
      <h4><a href="/xxe/safe">safe code 1</a></h4>
    </div>
  </div>
</div>

</body>
</html>