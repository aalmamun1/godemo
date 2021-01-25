package main

const messageWithHomeLink = `
<a href="/" >Home</a><br><br>
`

// index page
const indexPage = `
<style>
body {background-color: Azure;}
h1   {color: RoyalBlue;}
label    {color: Black; }
input    {color: Black;}
button    {color: DarkGreen;}
</style>

<h1>Login</h1>
<form method="post" action="/login" name="loginForm">
    <label for="name"><b>User name</b></label><br>
    <input type="text" placeholder="Enter User Name" id="name" name="name"size="25"><br><br>
    <label for="password"><b>Password</b></label><br>
    <input type="password" placeholder="Enter Password" id="password" name="password"size="25"><br><br>
    <button type="submit"><b>Login</b></button><br><br>
</form>

<h1>Signup</h1>
<form method="post" action="/signup" name="SignupForm">
    <label for="name"><b>User name</b></label>
    <input type="text" placeholder="Enter User Name" id="signupName" name="signupName"size="25">
    <label for="password"><b>Password</b></label>
    <input type="password" placeholder="Enter Password" id="signupPassword" name="signupPassword"size="25"><br><br>
    <label for="firstName"><b>First name</b></label>
    <input type="text" placeholder="Enter First Name" id="firstName" name="firstName"size="25"> 
    <label for="lastName"><b>Last name</b></label>
    <input type="text" placeholder="Enter Last Name" id="lastName" name="lastName"size="25"><br><br>
    <label for="email"><b>Email</b></label>
    <input type="text" placeholder="Enter Email" name="email" size="68"><br><br>
    <button type="submit"><b>Signup</b></button>
</form>
`

// internal page
const internalPageTop = `
<style>
table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
}
td {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}

<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
* {box-sizing: border-box;}

body {
  margin: 0;
  font-family: Arial, Helvetica, sans-serif;
}

.topnav {
  overflow: hidden;
  background-color: #e9e9e9;
}

.topnav a {
  float: left;
  display: block;
  color: black;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

.topnav a:hover {
  background-color: #ddd;
  color: black;
}

.topnav a.active {
  background-color: #2196F3;
  color: white;
}

.topnav .search-container {
  float: right;
}

.topnav input[type=text] {
  padding: 6px;
  margin-top: 8px;
  font-size: 17px;
  border: none;
}

.topnav .search-container button {
  float: right;
  padding: 6px;
  margin-top: 8px;
  margin-right: 16px;
  background: #ddd;
  font-size: 17px;
  border: none;
  cursor: pointer;
}

.topnav .search-container button:hover {
  background: #ccc;
}

@media screen and (max-width: 600px) {
  .topnav .search-container {
    float: none;
  }
  .topnav a, .topnav input[type=text], .topnav .search-container button {
    float: none;
    display: block;
    text-align: left;
    width: 100%;
    margin: 0;
    padding: 14px;
  }
  .topnav input[type=text] {
    border: 1px solid #ccc;  
  }
}
</style>
<div class="topnav">
  <a class="active" href="/internalUser">Home</a>
  <a href="/internal" >List of all users</a>
  <div class="search-container">
    <form action="/searchByUserName">
      <input type="text" placeholder="Search by username.." name="search">
      <button type="submit">Submit</button>
    </form>
  </div>
</div>
<hr>
<table border="1">
<tr>
	<td>User ID</td>
	<td>User Name</td>
	<td>First Name</td>
	<td>Last Name</td>
	<td>Email ID</td>
	<td>Last Update</td>
	<td>Delete</td>
	<td>Update</td>
	
</tr>
`
const internalPageDown = `
</table>
<form method="post" action="/logout"><br>
    <button type="submit">Logout</button>
</form>
`
const updatePage = `
	<style>
	body {background-color: Azure;}
	h1   {color: RoyalBlue;}
	label    {color: Black; }
	input    {color: Black;}
	button    {color: DarkGreen;}
	</style>

	<h1>User Update</h1>
	<form method="post" action="/updateUser/%d" name="updateForm">
		<label for="name"><b>User name</b></label>
		<input type="text" placeholder="Enter User Name" id="signupName" name="signupName"size="25"  value=%s readonly  style="background-color: grey">
		<label for="password"><b>Password</b></label>
		<input type="password" placeholder="Enter Password" id="signupPassword" name="signupPassword"size="25"  value=%s ><br><br>
		<label for="firstName"><b>First name</b></label>
		<input type="text" placeholder="Enter First Name" id="firstName" name="firstName"size="25" value=%s > 
		<label for="lastName"><b>Last name</b></label>
		<input type="text" placeholder="Enter Last Name" id="lastName" name="lastName"size="25" value=%s ><br><br>
		<label for="email"><b>Email</b></label>
		<input type="text" placeholder="Enter Email" name="email" size="68" value=%s ><br><br>
		<button type="submit"><b>Update</b></button>
	</form>
	`
