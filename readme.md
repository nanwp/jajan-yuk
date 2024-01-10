## Backen for app Jajan yuk
Monorepo for micorserivce jajan-yuk app.<br>



<h3>Services</h3>
<ul>
  <li>Auth services</li>
    Microservices for authentification and authorization <br> Include endpoint : 
    <ul>
        <li>Login</li>
            Login can generate JWT access token and refresh token, and stored it to redis as a session login
        <li>Current-User</li>
            Current user, for authorization. get info curret user login
        <li>Refresh-Token</li>
            Refresh token use for generate new token if access token is expired 
        <li>Logout</li>
            Logout use for remove access token from redis
    </ul>
  <li>User servuces</li>
    Microservices for manage user<br>
    Include endpoint
    <ul>
        <li>Registration</li>
            For registration of new user
        <li>Update Profile</li>
            For update user profile
    </ul>
  <li>Email Services</li>
    Microservices for send email<br> This service is subscriber from pubsub topic, When there is a new message in the topic, this service will send an email to the destination<br>
    flow : Email service waiting for message -> another service sending a message to email-topic -> Email services can process this message
    <br>this is a implement of <b>SSE (Server Send Event)</b>
</ul>


<h3>Tech Stack</h3>
<ul>
  <li>Go</li>
  <li>Postgres</li>
  <li>Redis</li>
  <li>Docker</li>
  <li>Pubsub</li>
</ul>

<h3>Example Usage</h3>

<li>Login</li>

```
curl --location 'https://auth.jajan-yuk.pegelinux.my.id/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"nanda",
    "password":"password"
}'
```


