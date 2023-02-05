# ascii-art-web
## [Link to the project](https://01.alem.school/intra/alem/div-01/ascii-art-web?event=621)
<br>

## Objectives

<br>
Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, <u><a href="https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art">ascii-art</a></u>.

Your webpage must allow the use of the different banners:
<ul>
    <li>shadow</li>
    <li>standard</li>
    <li>thinkertoy</li>
</ul>

Implement the following HTTP endpoints:

<ol>
<li>GET <code>/</code>: Sends HTML response, the main page.<br>
1.1. GET Tip: <u><a href="https://pkg.go.dev/html/template">go templates</a></u> to receive and display data from the server.</li>
<br>
<li>POST <code>/ascii-art</code>: that sends data to Go server (text and a banner)<br>
2.1. POST Tip: use form and other types of <u><a href="https://developer.mozilla.org/en-US/docs/Web/HTML/Element">tags</a></u> to make the post request.</li>
</ol>

The way you display the result from the POST is up to you. What we recommend are one of the following :
<ul>
    <li>Display the result in the route <code>/ascii-art</code> after the POST is completed. So going from the home page to another page.</li>
    <li>Or display the result of the POST in the home page. This way appending the results in the home page.</li>
</ul>
<br>

## HTTP status code

<br>



Your endpoints must return appropriate HTTP status codes.
<ul>
    <li>OK (200), if everything went without errors.</li>
    <li>Not Found, if nothing is found, for example templates or banners.</li>
    <li>Bad Request, for incorrect requests.</li>
    <li>Internal Server Error, for unhandled errors.</li>
</ul>


In the root project directory create a <code>README.MD</code> file with the following sections and contents:
<ul>
    <li>Description</li>
    <li>Authors</li>
    <li>Usage: how to run</li>
    <li>Implementation details: algorithm</li>
</ul>
<br>

## Instructions

<br>

<ul>
<li>HTTP server must be written in Go.</li>
<li>HTML templates must be in the project root directory templates.</li>
<li>The code must respect the <u><a href="https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md">good practices</a></u>.</li>
</ul>

<br>

## Allowed packages

<br>

<ul>
<li>
Only the 
<u><a href="https://pkg.go.dev/std">standard go</a></u>
packages are allowed
</li>
</ul>

<br>
